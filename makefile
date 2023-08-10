.PHONY: server, abigen, deploy

server:
	go run cmd/server/main.go

abigen:
	cd protocol && forge build && cd .. && go run cmd/abigen/main.go

deploy:
	cd protocol && forge script script/DeployMasterContracts.s.sol --rpc-url http://127.0.0.1:8545 --sender 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266

