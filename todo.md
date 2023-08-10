- add ctx for blockchain package functions
- add api endpoints
- master contract addresses needs to be stored in the db, not as config
- think about sync vs async writes to protocol
- think about error boundraries for txs
- think about gas limit for txs
- remove TODOs/FIXMEs
- lock down itemDefIDToURI based on the ERC-1155 standard on both sides
- use magiclink to authenticate
- handle errors like not found
- use nanoid instead of bigint id
- check that factory+master contracts exists on startup

1. abigen new contracts
2. deploy new contracts
3. add them to a table? master_contracts
4. update logic to deploy all
5. add enable/disable logic
