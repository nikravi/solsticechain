# solsticechain

Proof of Authority blockchain, with predefined list of nodes

* miners are set in the configuration file and included in the genesis block, thus immutable
* only miners which are in the genesis block are allowed to include blocks
* address similar to bitcoin
* mining consists of signing the block signature by the miners
* mining each 1 minute, with 10 coins reward
* nodes with sync using network communication (commands)
* a simple RPC server is implemented 

## run locally
```
export NODE_ID=3005
go install
cd $GOPATH/bin
solsticechain createwallet #this will be the first miner's address, minerAddress1
solsticechain createblockchain -address <minerAddress1>
solsticechain getbalance -address <minerAddress1> # returns the balance of minerAddress1
solsticechain printchain # prints the blockchain
solsticechain startnode -miner <minerAddress1> # starts the node and the API server
solsticechain send -from <minerAddress1> -to <clientNode1> -amount 10 -data 'test send' -mine

# in a new terminal
export NODE_ID=3000
solsticechain createwallet #this will be a client node (non miner)

# in a new terminal
export NODE_ID=3010
solsticechain createwallet #this will be second miner's address, minerAddress2
solsticechain startnode -miner <minerAddress2> # starts the node and downloads the blockchain from minerNode1

```
## cli commands
```
Usage:
  createblockchain -address ADDRESS - Create a blockchain and send genesis block reward to ADDRESS
  createwallet - Generates a new key-pair and saves it into the wallet file
  getbalance -address ADDRESS - Get balance of ADDRESS
  listaddresses - Lists all addresses from the wallet file
  printchain - Print all the blocks of the blockchain
  reindexutxo - Rebuilds the UTXO set
  send -from FROM -to TO -amount AMOUNT -data DATA -mine - Send AMOUNT of coins from FROM address to TO, optional data, mine on the same node (if-mine)
  startnode -miner ADDRESS - Start a node with ID specified in NODE_ID env. var. -miner enables mining
```
## api 

* get block count
`http://localhost:30058/block-count`
* get block
`http://localhost:30058/blocks/1` (by block height)
`http://localhost:30058/blocks/66b5abab811574a13f35af36ed6750990a6cb18eac61adb66fee39e92d74ec4b` (by block hash)
