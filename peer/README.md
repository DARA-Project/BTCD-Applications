# Peer Folder

This folder contains the following scripts:

[btcdRPCScript.go](https://github.com/sumahmood/Directed-Studies/blob/master/peer/btcdRPCScript.go) - This script basically uses [this script](https://github.com/btcsuite/btcd/blob/master/rpcclient/examples/btcwalletwebsockets/main.go) to try to connect to a running local RPC server using websockets. It gets a list of unspent transactions and shuts down the connection with the server after 10 seconds. [This image](https://github.com/DARA-Project/Directed-Studies/blob/master/images/btcdRPCScript%20img.jpg) illustrates the objects this string instantiates and interacts with.
