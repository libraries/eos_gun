# Create development wallet

```sh
$ cleos wallet create -f password
$ cleos wallet open
$ cleos wallet unlock --password $(cat password)
$ cleos wallet list

$ cleos wallet create_key
$ echo -n EOS5AWb5oN3z8hyvMuxtGyGCufmz4znjTqJMANuyqwf2LNHA7D1gV > public_key

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
$ cleos create account eosio alice $(cat public_key)
$ cleos create account eosio bob $(cat public_key)
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

# Multiple

```sh
$ wget https://github.com/eosio/eos/releases/download/v2.0.4/eosio_2.0.4-1-ubuntu-18.04_amd64.deb
$ sudo apt install ./eosio_2.0.4-1-ubuntu-18.04_amd64.deb
$ wget https://github.com/EOSIO/eosio.cdt/releases/download/v1.7.0/eosio.cdt_1.7.0-1-ubuntu-18.04_amd64.deb
$ sudo apt install ./eosio.cdt_1.7.0-1-ubuntu-18.04_amd64.deb

$ git clone https://github.com/EOSIO/eosio.contracts
$ cd eosio.contracts
$ ./build.sh -c /usr/opt/eosio.cdt -e /usr/opt/eosio/2.0.4/
```

```sh
$ keosd --unlock-timeout=9999999
$ cleos wallet create -f password
$ cleos wallet open
$ cleos wallet unlock --password $(cat password)
$ cleos wallet list

$ cleos wallet create_key
$ echo -n EOS5AWb5oN3z8hyvMuxtGyGCufmz4znjTqJMANuyqwf2LNHA7D1gV >> public_key
$ cleos wallet import --private-key 5KQwrPbwdL6PhXujxW37FSSQZ1JiwsST4cqQzDeyXtP79zkvFD3
```

```sh
$ nodeos -e -p eosio --plugin eosio::chain_api_plugin --plugin eosio::net_api_plugin --plugin eosio::producer_api_plugin

# https://www.jianshu.com/p/99bdf3f908f6
# https://www.bcskill.com/index.php/archives/884.html

$ curl -X POST http://127.0.0.1:8888/v1/producer/schedule_protocol_feature_activations -d '{"protocol_features_to_activate": ["0ec7e080177b2c02b278d5088611686b49d739925a92d9bfcacd7fc6b74053bd"]}' | jq

$ cd eosio.contracts
$ cleos set contract eosio build/contracts/eosio.boot -p eosio@active
$ curl -X POST http://127.0.0.1:8888/v1/producer/schedule_protocol_feature_activations -d '{"protocol_features_to_activate": ["299dcb6af692324b899b39f16d5a530a33062804e41f09dc97e9f156b4476707"]}' | jq
$ cleos push transaction '{"delay_sec":0,"max_cpu_usage_ms":0,"actions":[{"account":"eosio","name":"activate","data":{"feature_digest":"299dcb6af692324b899b39f16d5a530a33062804e41f09dc97e9f156b4476707"},"authorization":[{"actor":"eosio","permission":"active"}]}]}'

$ cleos set contract eosio build/contracts/eosio.bios
```

```sh
$ cleos create key --to-console
# Private key: 5JpCkrddVm5DTCwNkZQhDSu7DWSiT75U4chg3gPCrAUgkftuQGg
# Public key: EOS5PAkKG9T4FMDq978apLduwKewH9vx5StipckZguf8YL3usKFYg
$ cleos wallet import --private-key 5JpCkrddVm5DTCwNkZQhDSu7DWSiT75U4chg3gPCrAUgkftuQGg
$ cleos create account eosio inita EOS5PAkKG9T4FMDq978apLduwKewH9vx5StipckZguf8YL3usKFYg EOS5PAkKG9T4FMDq978apLduwKewH9vx5StipckZguf8YL3usKFYg
```

```sh
$ nodeos --producer-name inita --plugin eosio::chain_api_plugin --plugin eosio::net_api_plugin --http-server-address 127.0.0.1:8889 --p2p-listen-endpoint 127.0.0.1:9877 --p2p-peer-address 127.0.0.1:9876 --config-dir node2 --data-dir node2 --private-key [\"EOS5PAkKG9T4FMDq978apLduwKewH9vx5StipckZguf8YL3usKFYg\",\"5JpCkrddVm5DTCwNkZQhDSu7DWSiT75U4chg3gPCrAUgkftuQGg\"]
```

```json
{
    "schedule": [
        {
            "producer_name": "inita",
            "authority": [
                "block_signing_authority_v0",
                {
                    "threshold": 1,
                    "keys": [
                        {
                            "key": "EOS5PAkKG9T4FMDq978apLduwKewH9vx5StipckZguf8YL3usKFYg",
                            "weight": 1
                        }
                    ]
                }
            ]
        }
    ]
}
```

```sh
$ cleos push action eosio setprods "data.json" -p eosio@active
```


Private key: 5HqNMyRKkdv1mT4sC5NPX8y9dSb7tEGuTXCm3Hg59S17YE3HFdR
Public key: EOS6gztebm2XNHzy8Z9dyE8C7MgjkAsvofuh1zu3r2pfqe4Zb6gpw


cleos create account eosio initc EOS6gztebm2XNHzy8Z9dyE8C7MgjkAsvofuh1zu3r2pfqe4Zb6gpw EOS6gztebm2XNHzy8Z9dyE8C7MgjkAsvofuh1zu3r2pfqe4Zb6gpw

nodeos --producer-name initc --plugin eosio::chain_api_plugin --plugin eosio::net_api_plugin --p2p-peer-address 3.0.115.46:9876 --config-dir node2 --data-dir node2 --private-key [\"EOS6gztebm2XNHzy8Z9dyE8C7MgjkAsvofuh1zu3r2pfqe4Zb6gpw\",\"5HqNMyRKkdv1mT4sC5NPX8y9dSb7tEGuTXCm3Hg59S17YE3HFdR\"]

```json
{
    "schedule": [
        {
            "producer_name": "inita",
            "authority": [
                "block_signing_authority_v0",
                {
                    "threshold": 1,
                    "keys": [
                        {
                            "key": "EOS5PAkKG9T4FMDq978apLduwKewH9vx5StipckZguf8YL3usKFYg",
                            "weight": 1
                        }
                    ]
                }
            ]
        },
        {
            "producer_name": "initc",
            "authority": [
                "block_signing_authority_v0",
                {
                    "threshold": 1,
                    "keys": [
                        {
                            "key": "EOS6gztebm2XNHzy8Z9dyE8C7MgjkAsvofuh1zu3r2pfqe4Zb6gpw",
                            "weight": 1
                        }
                    ]
                }
            ]
        }
    ]
}
```

### D

nodeos --producer-name initd --plugin eosio::chain_api_plugin --plugin eosio::net_api_plugin --p2p-peer-address 3.0.115.46:9876 --config-dir node2 --data-dir node2 --private-key [\"EOS5p5BWHxjRA1rxAaYS6x2UwxkJFyu2MRRCRsVm1vio93kXnD62v\",\"5Jot4yBKHugVqf9pdzpbbypSPah2gmvFuTgGnzBJXnu4su9w7gE\"]
