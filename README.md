bazel build //cmd/prysmctl:prysmctl --config=release


./prysmctl testnet generate-genesis --fork capella --num-validators 64 --genesis-time-delay 600 --chain-config-file config.yml --geth-genesis-json-in genesis.json  --geth-genesis-json-out genesis.json --output-ssz genesis.ssz