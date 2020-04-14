# Create development wallet

```sh
$ cleos wallet create -f password
$ cleos wallet open
$ cleos wallet unlock --password $(cat password)
$ cleos wallet list

$ cleos wallet create_key
$ echo -n EOS5AWb5oN3z8hyvMuxtGyGCufmz4znjTqJMANuyqwf2LNHA7D1gV >> public_key

$ cleos wallet import --private-key 5KQwrPbwdL6PhXujxW37FSSQZ1JiwsST4cqQzDeyXtP79zkvFD3
```

# Start keosd and nodeos

```sh
$ pkill keosd
$ keosd --unlock-timeout=9999999

$ nodeos -e -p eosio --plugin eosio::producer_plugin --plugin eosio::producer_api_plugin --plugin eosio::chain_api_plugin --plugin eosio::http_plugin --plugin eosio::history_plugin --plugin eosio::history_api_plugin --filter-on="*" --access-control-allow-origin='*' --contracts-console --http-validate-host=false --verbose-http-errors --max-transaction-time 31536000
$ curl http://localhost:8888/v1/chain/get_info
```

# Create test accounts

```sh
$ cleos create account eosio bob $(cat public_key)
$ cleos create account eosio alice $(cat public_key)
```

# Deploy EOS

```sh
$ git clone https://github.com/EOSIO/eosio.contracts --branch v1.7.0 --single-branch
$ cd eosio.contracts/contracts/eosio.token

$ cleos create account eosio eosio.token EOS6MRyAjQq8ud7hVNYcfnVPJqcVpscN5So8BhtHuGYqET5GDW5CV
$ eosio-cpp -I include -o eosio.token.wasm src/eosio.token.cpp --abigen
$ cleos set contract eosio.token ./ --abi eosio.token.abi -p eosio.token@active
$ cleos push action eosio.token create '[ "eosio", "100000000000.0000 EOS"]' -p eosio.token@active
$ cleos push action eosio.token issue '[ "eosio", "100000000000.0000 EOS", "memo" ]' -p eosio@active
$ cleos push action eosio.token transfer '[ "eosio", "alice", "0.0001 EOS", "memo" ]' -p eosio@active
```
