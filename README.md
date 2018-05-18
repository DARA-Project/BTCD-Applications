# BTCD Applications (Experiments with BTCD)

[ConnectAndAddBlocksToServer.go](https://github.com/DARA-Project/BTCD-Applications/blob/master/ConnectAndAddBlockToServer.go) - This application used an example file from the main BTCD repo [here](https://github.com/btcsuite/btcd/blob/master/rpcclient/examples/btcdwebsockets/main.go). The script creates RPC notification handlers, connects to a local BTCD server, and tries to get the current block count in the blockchain. Subsequently, it tries to create a new transactions and sends the transaction to the server. Finally, it shuts down the connection with the server.<br />
**Goal:** Set up an RPC server and try to send it transactions that the CPU miner can mine on the server.<br />
**What Works:** Establishing notification handlers and getting the block count.<br />
**What Doesn't Work:** Creating and sending a transaction to the server to mine.<br />

[ConnectPeers.go](https://github.com/DARA-Project/BTCD-Applications/blob/master/ConnectPeers.go) - This application is basically a copy [of](https://github.com/btcsuite/btcd/blob/master/peer/example_test.go) in the main BTCD repo. Creates an inbound peer and outbound peer. Inbound peer listens for connection from outbound peer, establishes connection, and disconnects.
**Goal:** Try to connect multiple peers using the BTCD API.<br />
**What Works:** Peers are able to connect and then disconnect.<br />
**What Doesn't Work:** Everything works.<br />

[CreatBlockChainAndPaymentAddress.go](https://github.com/DARA-Project/BTCD-Applications/blob/master/CreateBlockchainAndPaymentAddress.go) -

[CreateTransactionBlockAndAddToBlockchain.go](https://github.com/DARA-Project/BTCD-Applications/blob/master/CreateTransactionBlockAndAddToBlockchain.go) -

[CreateTransactionsForMempoolAndGenerateMultipleBlocksUsingCPUMiner.go](https://github.com/DARA-Project/BTCD-Applications/blob/master/CreateTransactionsForMempoolAndGenerateMultipleBlocksUsingCPUMiner.go) -

[CreateTransactionsForMempoolAndMineBlocks.go](https://github.com/DARA-Project/BTCD-Applications/blob/master/CreateTransactionsForMempoolAndMineBlocks.go) -

[RunBTCDMainLocally.go](https://github.com/DARA-Project/BTCD-Applications/blob/master/RunBTCDMainLocally.go) -
