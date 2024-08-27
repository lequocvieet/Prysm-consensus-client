build-prsymc-ctl:
	rm -rf prysmctl && cd prysm/ &&  bazel build //cmd/prysmctl:prysmctl --config=release 

copy-bin-file:
	cp -f /home/lequocvieet/Documents/devnet/prysm/bazel-bin/cmd/prysmctl/prysmctl_/prysmctl  /home/lequocvieet/Documents/devnet/

run-prysm: 
	./prysmctl testnet generate-genesis --fork capella --num-validators 64 --genesis-time-delay 600 --chain-config-file config.yml --geth-genesis-json-in genesis.json  --geth-genesis-json-out genesis.json --output-ssz genesis.ssz
