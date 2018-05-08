# Directed-Studies (BTCD)

### Table of Contents
1. [General Information](#General)
2. [Papers Read (Academic and Non-Academic)](#Papers)
3. [BTCD Information](#BTCD)
4. [Details About Scripts In the Repo](#Scripts)
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

<a name="Scripts" />

### 4. Details About Scripts in the Repo

<a name="Future" />

### 5. Progress, Issues, and Things to Do
