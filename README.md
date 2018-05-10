# Directed-Studies (BTCD)

### Table of Contents
1. [General Information](#General)
2. [Papers Read (Academic and Non-Academic)](#Papers)
3. [BTCD Information](#BTCD)
    1. [BTCD Insights](#Insights)
    2. [BTCD Helpful Links](#Links)
4. [Details About Scripts In the Repo](#Scripts)
    1. [peer](#peer)
    2. [testA](#testA)
5. [Progress, Issues, and Things to Do](#Future)

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
