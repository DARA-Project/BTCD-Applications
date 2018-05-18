# BTCD Applications (Experiments with BTCD)

[ConnectAndAddBlocksToServer.go](https://github.com/DARA-Project/BTCD-Applications/blob/master/ConnectAndAddBlockToServer.go) - This application used an example file from the main BTCD repo [here](https://github.com/btcsuite/btcd/blob/master/rpcclient/examples/btcdwebsockets/main.go). The script creates RPC notification handlers, connects to a local BTCD server, and tries to get the current block count in the blockchain. Subsequently, it tries to create a new transactions and sends the transaction to the server. Finally, it shuts down the connection with the server.<br />
**Goal:** Set up an RPC server and try to send it transactions that the CPU miner can mine on the server.<br />
**What Works:** Establishing notification handlers and getting the block count.<br />
**What Doesn't Work:** Creating and sending a transaction to the server to mine.<br />
