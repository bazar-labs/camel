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
- add requestid (Clients may provide a unique identifier for a request to perform at most once execution. When a requestid is resubmitted, it will not cause the work to be performed again; the response message will be the current state of items affected by the original successful execution.)
- input validation
- clean up contants and interfaces across modules and repos

- rename
- restructure
- add store enable/disable
- add endpoints
- test

master economy contracts

```yaml
- factory
- inventory_registry
- inventory_controller
- behaviors
  - purchase_item
  - open_lootbox
  - drop_item
```

game economy contracts

```yaml
- inventory_registry
  - address: 0x0...

- inventory_controller
  - address: 0x0...

- behaviors
  - purchase_item
    - address: 0x0...
    - enabled: true

  - open_lootbox
    - address: 0x0...
    - enabled: true

  - drop_item
    - address: 0x0...
    - enabled: true
```
