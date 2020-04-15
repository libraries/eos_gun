package main

import (
	"encoding/hex"
	"flag"
	"log"
	"math/rand"
	"time"

	eos "github.com/eoscanada/eos-go"
	"github.com/eoscanada/eos-go/ecc"
	"github.com/eoscanada/eos-go/system"
	"github.com/eoscanada/eos-go/token"
)

var (
	flAPI       = flag.String("api", "http://3.0.115.46:28888", "API address")
	flGoroutine = flag.Int("goroutine", 8, "Number of goroutines")
)

const (
	cTokenAccount = "eosio.token"
	cFrom         = "eosio"
	cFromPrivKey  = "5KQwrPbwdL6PhXujxW37FSSQZ1JiwsST4cqQzDeyXtP79zkvFD3"
	cTo           = "alice"
	cSymbol       = "EOS"
	cQuantity     = "0.0001 EOS"
	cExpiration   = time.Hour
)

type eosgun struct {
	api *eos.API
	c   chan *eos.PackedTransaction
}

func (e *eosgun) create() {
	from := eos.AccountName(cFrom)
	to := eos.AccountName(cTo)
	quantity, err := eos.NewEOSAssetFromString(cQuantity)
	if err != nil {
		log.Panicln(err)
	}
	memo := ""

	for {
		txOpts := &eos.TxOptions{}
		if err := txOpts.FillFromChain(e.api); err != nil {
			log.Panicln(err)
		}
		buf := make([]byte, 8)
		rand.Read(buf)
		nonce := hex.EncodeToString(buf)
		tx := eos.NewTransaction([]*eos.Action{token.NewTransfer(from, to, quantity, memo), system.NewNonce(nonce)}, txOpts)
		tx.SetExpiration(cExpiration)
		_, packedTx, err := e.api.SignTransaction(tx, txOpts.ChainID, eos.CompressionNone)
		if err != nil {
			log.Panicln(err)
		}
		e.c <- packedTx
	}
}

func (e *eosgun) send() {
	for t := range e.c {
		r, err := e.api.PushTransaction(t)
		if err != nil {
			log.Panicln(err)
		}
		log.Println(r.Processed.ID)
	}
}

func jakobu() *eosgun {
	r := &eosgun{}
	a := func() *eos.API {
		a := eos.New(*flAPI)
		p, err := ecc.NewPrivateKey(cFromPrivKey)
		if err != nil {
			log.Panicln(err)
		}
		a.SetCustomGetRequiredKeys(func(tx *eos.Transaction) (keys []ecc.PublicKey, e error) {
			return []ecc.PublicKey{p.PublicKey()}, nil
		})
		keybag := &eos.KeyBag{}
		if err := keybag.ImportPrivateKey(cFromPrivKey); err != nil {
			log.Panicln(err)
		}
		a.SetSigner(keybag)
		return a
	}()
	r.api = a
	r.c = make(chan *eos.PackedTransaction, 1024)
	return r
}

func main() {
	flag.Parse()
	gun := jakobu()

	for i := 0; i < *flGoroutine; i++ {
		go gun.create()
	}
	for i := 0; i < *flGoroutine; i++ {
		go gun.send()
	}

	tic, err := gun.api.GetCurrencyBalance(cTo, cSymbol, cTokenAccount)
	if err != nil {
		log.Panicln(err)
	}
	time.Sleep(time.Minute)
	toc, err := gun.api.GetCurrencyBalance(cTo, cSymbol, cTokenAccount)
	if err != nil {
		log.Panicln(err)
	}

	log.Println(tic[0], toc[0])
}
