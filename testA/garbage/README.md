# Garbage Folder

General experiments in trying to get bootstrapped to BTCD but didn't really contribute much to the project.

[other.go](https://github.com/sumahmood/Directed-Studies/blob/master/testA/garbage/other.go) - This script tries to create an initial coinbase transaction and adds more transactions to try to mine them into a block. I tried to generate a block header manually, add transactions, and then process them into a block. It didn't work out too well, so I just left it and asked the BTCD people about a better way to do it. They suggested using RPC to mine blocks, so I abandoned this experiment.

[peer.go](https://github.com/sumahmood/Directed-Studies/blob/master/testA/garbage/peer.go) - This script is basically a copy of [this script](https://github.com/btcsuite/btcd/blob/master/peer/example_test.go) in the main BTCD repo. I tried to get it to work locally to try to see what was happening. Was able to get the basic script to work in trying to initialize and create an outbound peer following their script. Didn't really contribute much to testing but was useful in trying to understand the process.
