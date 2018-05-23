# BTCD Applications (Experiments with BTCD)

### Table of Contents
1. [Approaches and General Information](#General)
2. [BTCD Applications](#Apps)
    1. [ConnectAndAddBlocksToServer.go](#App1)
    2. [ConnectPeers.go](#App2)
    3. [CreateBlockChainAndPaymentAddress.go](#App3)
    4. [CreateTransactionsForMempoolAndGenerateMultipleBlocksUsingCPUMiner.go](#App4)
    5. [CreateTransactionsForMempoolAndMineBlocks.go](#App5)
    6. [RunBTCDMainLocally.go](#App6)
    7. [rpc.cert](#App7)
3. [Config Files](#Config)

<a name="General" />

### 1. Approaches and General Information

The purpose of these applications was to set up a working application that could be used to test Bitcoin invariants that could be tested using DARA from dynamic execution of these applications. The primary interest from our perspective is the validation that happens before a block is added to the blockchain after being mined and the validation process for the longest blockchain across multiple peers. There were two approaches here:<br />
1) Write Applications in Go: This was the original approach that I used that involved modifying developer tests in an attempt to set up a blockchain and transactions that could be mined by peers to try to add blocks to the blockchain. The developers suggested using the command line utility for testing purposes in comparison to writing applications manually. The advantage of this approach is that it gives us more control over the kinds of things we can do in terms of exercising BTCD but it's harder to set up and will require a lot more time to come up with. In particular, there's a lot of parameters within the configs for the various structs that need to be created to create a working server, CPU miner, blockchain, etc.<br /><br />
2) Config Files: Use config confiles and a command line utility called [BTCCTL](https://github.com/btcsuite/btcd/tree/master/cmd/btcctl) within BTCD to set up an RPC server and blockchain instance to which blocks are added. The config files discussed below were used in this approach. The developers suggested using this approach for testing over writing applications manually in Go that utilized BTCD. The advantage of this approach is that it's easier to set up and the developers seem to think that this is the best approach to test the application. Unfortunately, it does lead to more of a blackbox based testing approach and may only test the program with certain predefined scenarios. It may be possible to exercise some control by deciding the transactions that are added to the memory pool but it's still less controllable than using Go applications.

<a name="Apps" />

### 2. BTCD Applications

<a name="App1" />

**2.1 ConnectAndAddBlocksToServer.go**

[ConnectAndAddBlocksToServer.go](https://github.com/DARA-Project/BTCD-Applications/blob/master/ConnectAndAddBlockToServer.go) - This application used an example file from the main BTCD repo [here](https://github.com/btcsuite/btcd/blob/master/rpcclient/examples/btcdwebsockets/main.go). The application creates RPC notification handlers, connects to a local BTCD server, and tries to get the current block count in the blockchain. Subsequently, it tries to create a new transactions and sends the transaction to the server. Finally, it shuts down the connection with the server.<br />
**Goal:** Set up an RPC server and try to send it transactions that the CPU miner can mine on the server.<br />
**What Works:** Establishing notification handlers and getting the block count.<br />
**What Doesn't Work:** Creating and sending a transaction to the server to mine.<br />

![ConnectAndAddBlocksToServer img](https://github.com/DARA-Project/BTCD-Applications/blob/master/images/btcdRPCScript%20img.jpg)

<a name="App2" />

**2.2 ConnectPeers.go**

[ConnectPeers.go](https://github.com/DARA-Project/BTCD-Applications/blob/master/ConnectPeers.go) - This application is basically a copy [of](https://github.com/btcsuite/btcd/blob/master/peer/example_test.go) in the main BTCD repo. Creates an inbound peer and outbound peer. Inbound peer listens for connection from outbound peer, establishes connection, and disconnects.<br />
**Goal:** Try to connect multiple peers using the BTCD API.<br />
**What Works:** Peers are able to connect and then disconnect.<br />
**What Doesn't Work:** Everything works.<br />

![ConnectPeers img](https://github.com/DARA-Project/BTCD-Applications/blob/master/images/ConnectPeers%20img.jpg)

<a name="App3" />

**2.3 CreateBlockChainAndPaymentAddress.go**

[CreateBlockChainAndPaymentAddress.go](https://github.com/DARA-Project/BTCD-Applications/blob/master/CreateBlockchainAndPaymentAddress.go) - This application creates a block chain and tries to instantiate a payment address. It was one of the earliest applications, so there isn't very much going on here. I was interesting in understanding how payment addresses work.<br />
**Goal:** Understanding how payment addresses work and setting up a chain work.<br />
**What Works:** Sets up a blockchain and creates a Bitcoin payment address.<br />
**What Doesn't Work:** Everything works.<br />

![CreateBlockChainAndPaymentAddress img](https://github.com/DARA-Project/BTCD-Applications/blob/master/images/CreateBlockChainAndPaymentAddress.jpg)

<a name="App4" />

**2.4 CreateTransactionsForMempoolAndGenerateMultipleBlocksUsingCPUMiner.go**

[CreateTransactionsForMempoolAndGenerateMultipleBlocksUsingCPUMiner.go](https://github.com/DARA-Project/BTCD-Applications/blob/master/CreateTransactionsForMempoolAndGenerateMultipleBlocksUsingCPUMiner.go) - This application tries to put a bunch of the results from above together to create an underlying database that stores the blockchain; the blockchain with simulated network parameters; generate a Bitcoin payment address to pay to; generate a block template; create a memory pool of transactions to mine; add transactions to the memory pool; and attempt to generate blocks by instantiating a CPU miner. <br />
**Goal:** Manually generate transactions to mine so that the CPU miner can attempt to create blocks.<br />
**What Works:** Able to generate blocks successfully from the transactions added to the memory pool.<br />
**What Doesn't Work:** Works fairly well but as a final step, we need to add the generated blocks to the blockchain.<br />

![CreateTransactionsForMempoolAndGenerateMultipleBlocksUsingCPUMine img](https://github.com/DARA-Project/BTCD-Applications/blob/master/images/CreateTransactionsForMempoolAndGenerateMultipleBlocksUsingCPUMiner%20img.jpg)

<a name="App5" />

**2.5 CreateTransactionsForMempoolAndMineBlocks.go**

[CreateTransactionsForMempoolAndMineBlocks.go](https://github.com/DARA-Project/BTCD-Applications/blob/master/CreateTransactionsForMempoolAndMineBlocks.go) - This application does the same thing as [CreateTransactionsForMempoolAndGenerateMultipleBlocksUsingCPUMiner.go](https://github.com/DARA-Project/BTCD-Applications/blob/master/CreateTransactionsForMempoolAndGenerateMultipleBlocksUsingCPUMiner.go) with the exception that instead of attempting to generate multiple blocks through the CPU miner in the end, it attempts to generate one block and process it to see if it can be successfully added to the blockchain.<br />
**Goal:** Generate a block from an underlying source of transactions and see if it can be successfully added to the blockchain.<br />
**What Works:** Works when the issue mentioned in what doesn't work doesn't happen.<br />
**What Doesn't Work:** Issue with hash generation. Sometimes the application fails because the block hash generated exceeds 0x800...Max hash should be 0x7ff.....Tried to play around with it to fix it, but ended up abondoning this approach in favour of using the command line utility as recommended by the BTCD developers and on discussion with Stew.<br />

![CreateTransactionsForMempoolAndMineBlocks img](https://github.com/DARA-Project/BTCD-Applications/blob/master/images/CreateTransactionsForMempoolAndGenerateMultipleBlocksUsingCPUMiner%20img.jpg)

<a name="App6" />

**2.6 RunBTCDMainLocally.go**

[RunBTCDMainLocally.go](https://github.com/DARA-Project/BTCD-Applications/blob/master/RunBTCDMainLocally.go) - This application is basically a copy of the [main](https://github.com/btcsuite/btcd/blob/master/btcd.go) BTCD file. The idea of using this was to try attempt to change line 27 in the original file to see if we could possibly create multiple instances of the underlying DB for each BTCD process. This was necessary because running the program on the command line required instantiating a DB object for the local storage of the blockchain, but it wasn't possible to run multiple BTCD processes due to the fact that all processes ended up using the same BTCD file path. No changes were actually made to the file, but this was the plan to address this issue and attempt to run multiple BTCD processes on the command line (so we could test distributed invariants using DARA).
<br />
**Goal:** Attempt to run multiple BTCD processes by modifying the blockDB prefix so there is no contention on the underlying database for each process we spawn.<br />
**What Works:** Unable to get this to work.<br />
**What Doesn't Work:** Need to work on getting this up and running.<br />

<a name="App7" />

**2.7 rpc.cert**

[rpc.cert](https://github.com/DARA-Project/BTCD-Applications/blob/master/rpc.cert) - This file is important to running the applications so please leave it here. It's not being used in terms of building the applications so you don't need to worry much about it.

<a name="Config" />

### 3. Config Files

These are my local copies of the config files that I used to test BTCD. The [docs](https://github.com/DARA-Project/BTCD-Applications/tree/master/docs) folder section 4.1 contains more information about scripts that can be used with these config files. Please note that config files are placed in the following folders: <br />
1) [btcd.conf](https://github.com/DARA-Project/BTCD-Applications/blob/master/config/btcd.conf): ~/.btcd/btcd.conf<br />
2) [btcwallet.conf](https://github.com/DARA-Project/BTCD-Applications/blob/master/config/btcwallet.conf): ~/.btcwallet/btcwallet.conf

See section 4.1 of the docs folder for more information on how they are used.
