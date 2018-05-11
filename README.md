# Directed-Studies (BTCD)

### Table of Contents
1. [General Information](#General)
2. [Papers Read (Academic and Non-Academic)](#Papers)
3. [BTCD Information](#BTCD)
    1. [BTCD Insights](#Insights)
    2. [BTCD Helpful Links](#Links)
    3. [Transcripts of Chats with Developers](#Chats)
4. [Details About Scripts In the Repo](#Scripts)
    1. [peer](#peer)
    2. [testA](#testA)
5. [Progress, Issues, and Things to Do](#Future)
    1. [Progress](#Progress)
    2. [Issues](#Issues)
    3. [Future Tasks](#Tasks)

<a name="General" />

### 1. General

The goal in this project was to use BTCD to validate invariants around validating the blockchain. The blockchain used was on a local simulated Bitcoin network (Simnet in BTCD) where we had multiple RPC servers running with different peers and underlying resources (e.g. databases) and to use executions from BTCD to determine where these invariants if any Bitcoin invariants were validated. The key thing to model is the communication of these servers in validating the longest blockchain to determine whether to keep or discard a transaction. This will involve creating dummy transactions and sending them to servers so they can be mined by peers.

<a name="Papers" />

### 2. Papers Read (Academic and Non-Academic)

These were the sources I read to try to get a grasp of Bitcoin and how the protocol operated at a high level. The articles presented some more concrete details about the implementation of the protocol and brought the things the paper discussed a bit more down to earth. The paper is still the key resource. It's fairly short and shouldn't take very long to read.

[Bitcoin Paper (Original)](https://bitcoin.org/bitcoin.pdf)

Other Articles on the Paper:

[Cracking the Bitcoin Whitepaper (what I used to get a better grasp of it)](https://medium.com/@FolusoOgunlana/cracking-the-bitcoin-white-paper-c5f479ce748d)

[Bitcoin White Paper in 4 minutes](https://hackernoon.com/dissecting-the-bitcoin-whitepaper-in-four-minutes-5c8c5e5f8010)

<a name="BTCD" />

### 3. BTCD Information

As indicated in the instructions, please make sure that BTCD is installed on your machine using glide to avoid future issues. The process to update BTCD and its dependencies is also described.

[BTCD Readme for Main Repo](https://github.com/btcsuite/btcd)

[Webchat (I used this frequently to interact with the developers)](https://webchat.freenode.net/?channels=btcd)

[BTCD Install Instructions](https://github.com/btcsuite/btcd/blob/master/docs/README.md)

[BTCD Documentation](https://github.com/btcsuite/btcd/tree/master/docs)

<a name="Insights" />

**3.1 BTCD Insights**

-Block chain constructed using btcd/blockchain/chain.go. Can use parameters defined in btcd/chaincfg/params.go to create a simulated
Bitcoin network.

-New transaction created by generating a new tx (transaction message) in btcd/wire/msgtx.go. Can construct an actual
transaction from this message using btcutil/tx.go.

-Constructing a new block first requires generating a new block header in btcd/wire/blockheader.go. There are numerous
parameters required here including version (set to 1), hash of the previous block, merkleRootHash (set as the hash of the transaction), 
difficulty bits, and nonce. Nonce was randomly generated. Difficulty bits is something that needs to be tuned so we can do a reasonable
amount of mining. We take the blockheader and change it to a msgblock using btcd/wire/msgblock.go. Finally, we take this msgblock and
convert it to a bitcoint block in btcutil/block.go.

<a name="Links" />

**3.2 BTCD Helpful Links**

[Useful example of how to set up a simple SimNet using the command line BTCD utilities](https://gist.github.com/davecgh/2992ed85d41307e794f6)

[A list of useful scripts for testing purposes](https://gist.github.com/davecgh). 

I highly recommend using these as a reference point to set up tests. I found them very useful.

<a name="Chats" />

**3.3 Transcripts of Chats with Developers**

<a name="Scripts" />

### 4. Details About Scripts in the Repo

<a name="peer" />

**4.1 peer**

[Readme File](https://github.com/sumahmood/Directed-Studies/blob/master/peer/README.md)

<a name="testA" />

**4.2 testA**

[Readme File - Garbage Folder](https://github.com/sumahmood/Directed-Studies/blob/master/testA/garbage/README.md)

[Readme File - tmp Folder](https://github.com/sumahmood/Directed-Studies/blob/master/testA/tmp/README.md)

<a name="Future" />

### 5. Progress, Issues, and Things to Do

<a name="Progress" />

**5.1 Progress**

This details all the progress I've made. The config files referred to below are placed in the config folder. The purpose here was to try to set up a local Simnet (a local simulated bitcoin peer-to-peer network) that can construct a chain and take in transactions which miners can use to mine blocks and validate new blocks before being added to the chain. The key goal was to create an architecture where multiple RPC servers, each with their own underlying blockchain versions (in the database). Miners mine new transactions into blocks, which are then validated and added to the blockchain based on the longest blockchain invariant (the network always accepts the longest valid blockchain as the true blockchain and discards any others).

This was valuable from the DARA point of view because setting up the system in this way allows us to change as little of the underlying code as possible and gives a lot of control over the configuration parameters without having to manually set them up. Our goal is to verify invariants, in particular around the new blocks that are accepted into the chain and verification of the longest chain. So, by setting up the network in this way, we can focus on setting up the logging in the underlying BTCD code instead of expending too much effort in setting up a local bitcoin test network.

1) To start up a simnet, the recommended process is detailed [here](https://gist.github.com/davecgh/2992ed85d41307e794f6).

2) Was able to get the config files set up to run btcd. The config file itself is placed in ~/.btcd/btcd.conf, which is essentially used to specify various configurations for the server. The config file includes information about how to set up the configuration parameters in the file. I've included a sample config file in the config folder for btcd.conf which should be helpful in getting started.

3) Set up a wallet with a new Simnet address using btcwallet. The wallet also uses a config file that has various settings. I've included my version of the config file for reference. The advantage of using a config file once again is that you don't have to manually specify the paramaters on the command line. The config file is in ~/.btcwallet/btcwallet.conf.

4) Followed the rest of the setup process detailed in 1. to get the RPC server running and was able to connect to a local network with a peer and miner ready to send transactions to mine blocks.

<a name="Issues" />

**5.2 Issues**

There were two issues that are presently holding me back. The architecture we want to set up is detailed here. 

One of the issues around [this script](https://gist.github.com/davecgh/2992ed85d41307e794f6) relates to the fact that the wallet doesn't get updated with Bitcoin as the gist indicates. My belief is that this problem arises from this [open issue](https://github.com/btcsuite/btcwallet/issues/496) in BTCD as this is the error message received by the wallet when it tries to connect to the RPC server. We need an open wallet while doing our experimentation to enable transactions to be processed in BTCD. 

The other issue relates to connecting peers. I still need to sort out the right way forward in connecting peers to the RPC servers, but the first problem is more pressing.

<a name="Tasks" />

**5.3 Future Tasks**

