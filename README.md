# BTCD Applications (Experiments with BTCD)

[ConnectAndAddBlocksToServer.go](https://github.com/DARA-Project/BTCD-Applications/blob/master/ConnectAndAddBlockToServer.go) - This application used an example file from the main BTCD repo [here](https://github.com/btcsuite/btcd/blob/master/rpcclient/examples/btcdwebsockets/main.go). The application creates RPC notification handlers, connects to a local BTCD server, and tries to get the current block count in the blockchain. Subsequently, it tries to create a new transactions and sends the transaction to the server. Finally, it shuts down the connection with the server.<br />
**Goal:** Set up an RPC server and try to send it transactions that the CPU miner can mine on the server.<br />
**What Works:** Establishing notification handlers and getting the block count.<br />
**What Doesn't Work:** Creating and sending a transaction to the server to mine.<br />

![ConnectAndAddBlocksToServer img](https://github.com/DARA-Project/BTCD-Applications/blob/master/images/btcdRPCScript%20img.jpg)

[ConnectPeers.go](https://github.com/DARA-Project/BTCD-Applications/blob/master/ConnectPeers.go) - This application is basically a copy [of](https://github.com/btcsuite/btcd/blob/master/peer/example_test.go) in the main BTCD repo. Creates an inbound peer and outbound peer. Inbound peer listens for connection from outbound peer, establishes connection, and disconnects.<br />
**Goal:** Try to connect multiple peers using the BTCD API.<br />
**What Works:** Peers are able to connect and then disconnect.<br />
**What Doesn't Work:** Everything works.<br />

[CreateBlockChainAndPaymentAddress.go](https://github.com/DARA-Project/BTCD-Applications/blob/master/CreateBlockchainAndPaymentAddress.go) - This application creates a block chain and tries to instantiate a payment address. It was one of the earliest applications, so there isn't very much going on here. I was interesting in understanding how payment addresses work.<br />
**Goal:** Understanding how payment addresses work and setting up a chain work.<br />
**What Works:** Sets up a blockchain and creates a Bitcoin payment address.<br />
**What Doesn't Work:** Everything works.<br />

[CreateTransactionsForMempoolAndGenerateMultipleBlocksUsingCPUMiner.go](https://github.com/DARA-Project/BTCD-Applications/blob/master/CreateTransactionsForMempoolAndGenerateMultipleBlocksUsingCPUMiner.go) - This application tries to put a bunch of the results from above together to create an underlying database that stores the blockchain; the blockchain with simulated network parameters; generate a Bitcoin payment address to pay to; generate a block template; create a memory pool of transactions to mine; add transactions to the memory pool; and attempt to generate blocks by instantiating a CPU miner. <br />
**Goal:** Manually generate transactions to mine so that the CPU miner can attempt to create blocks.<br />
**What Works:** Able to generate blocks successfully from the transactions added to the memory pool.<br />
**What Doesn't Work:** Works fairly well but as a final step, we need to add the generated blocks to the blockchain.<br />

[CreateTransactionsForMempoolAndMineBlocks.go](https://github.com/DARA-Project/BTCD-Applications/blob/master/CreateTransactionsForMempoolAndMineBlocks.go) - This application does the same things [CreateTransactionsForMempoolAndGenerateMultipleBlocksUsingCPUMiner.go](https://github.com/DARA-Project/BTCD-Applications/blob/master/CreateTransactionsForMempoolAndGenerateMultipleBlocksUsingCPUMiner.go) with the exception that instead of attempting to generate multiple blocks through the CPU miner in the end, it attempts to generate one block and process it to see if it can be successfully added to the blockchain.<br />
**Goal:** Generate a block from an underlying source of transactions and see if it can be successfully added to the blockchain.<br />
**What Works:** Works when the issue mentioned in what doesn't work doesn't happen.<br />
**What Doesn't Work:** Issue with hash generation. Sometimes the application fails because the block hash generated exceeds 0x800...Max hash should be 0x7ff.....Tried to play around with it to fix it, but ended up abondoning this approach in favour of using the command line utility as recommended by the BTCD developers and on discussion with Stew.<br />

[RunBTCDMainLocally.go](https://github.com/DARA-Project/BTCD-Applications/blob/master/RunBTCDMainLocally.go) - This application is basically a copy of the [main](https://github.com/btcsuite/btcd/blob/master/btcd.go) BTCD file. The idea of using this was to try attempt to change line 27 in the original file to see if we could possibly create multiple instances of the underlying DB for each BTCD process. This was necessary because running the program on the command line required instantiating a DB object for the local storage of the blockchain, but it wasn't possible to run multiple BTCD processes due to the fact that all processes ended up using the same BTCD file path. No changes were actually made to the file, but this was the plan to address this issue and attempt to run multiple BTCD processes on the command line (so we could test distributed invariants using DARA).
<br />
**Goal:** Attempt to run multiple BTCD processes by modifying the blockDB prefix so there is no contention on the underlying database for each process we spawn.<br />
**What Works:** Unable to get this to work.<br />
**What Doesn't Work:** Need to work on getting this up and running.<br />

[rpc.cert](https://github.com/DARA-Project/BTCD-Applications/blob/master/rpc.cert) - This file is important to running the applications so please leave it here. It's not being used in terms of building the applications so you don't need to worry much about it.
