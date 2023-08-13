.PHONY: server, abigen, deploy

server:
	go run cmd/server/main.go

abigen:
	cd protocol && forge build && cd .. && go run cmd/abigen/main.go

deploy:
	cd protocol && forge script script/DeployMasterContracts.s.sol --rpc-url http://127.0.0.1:8545 --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 --broadcast -vvvv
