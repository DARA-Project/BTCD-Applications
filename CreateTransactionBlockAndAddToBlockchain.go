var genesisCoinbaseTx = wire.MsgTx{
	Version: 1,
	TxIn: []*wire.TxIn{
		{
			PreviousOutPoint: wire.OutPoint{
				Hash:  chainhash.Hash{},
				Index: 0xffffffff,
			},
			SignatureScript: []byte{
				0x04, 0xff, 0xff, 0x00, 0x1d, 0x01, 0x04, 0x45, /* |.......E| */
				0x54, 0x68, 0x65, 0x20, 0x54, 0x69, 0x6d, 0x65, /* |The Time| */
				0x73, 0x20, 0x30, 0x33, 0x2f, 0x4a, 0x61, 0x6e, /* |s 03/Jan| */
				0x2f, 0x32, 0x30, 0x30, 0x39, 0x20, 0x43, 0x68, /* |/2009 Ch| */
				0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x6f, 0x72, /* |ancellor| */
				0x20, 0x6f, 0x6e, 0x20, 0x62, 0x72, 0x69, 0x6e, /* | on brin| */
				0x6b, 0x20, 0x6f, 0x66, 0x20, 0x73, 0x65, 0x63, /* |k of sec|*/
				0x6f, 0x6e, 0x64, 0x20, 0x62, 0x61, 0x69, 0x6c, /* |ond bail| */
				0x6f, 0x75, 0x74, 0x20, 0x66, 0x6f, 0x72, 0x20, /* |out for |*/
				0x62, 0x61, 0x6e, 0x6b, 0x73, /* |banks| */
			},
			Sequence: 0xffffffff,
		},
	},
	TxOut: []*wire.TxOut{
		{
			Value: 0x12a05f200,
			PkScript: []byte{
				0x41, 0x04, 0x67, 0x8a, 0xfd, 0xb0, 0xfe, 0x55, /* |A.g....U| */
				0x48, 0x27, 0x19, 0x67, 0xf1, 0xa6, 0x71, 0x30, /* |H'.g..q0| */
				0xb7, 0x10, 0x5c, 0xd6, 0xa8, 0x28, 0xe0, 0x39, /* |..\..(.9| */
				0x09, 0xa6, 0x79, 0x62, 0xe0, 0xea, 0x1f, 0x61, /* |..yb...a| */
				0xde, 0xb6, 0x49, 0xf6, 0xbc, 0x3f, 0x4c, 0xef, /* |..I..?L.| */
				0x38, 0xc4, 0xf3, 0x55, 0x04, 0xe5, 0x1e, 0xc1, /* |8..U....| */
				0x12, 0xde, 0x5c, 0x38, 0x4d, 0xf7, 0xba, 0x0b, /* |..\8M...| */
				0x8d, 0x57, 0x8a, 0x4c, 0x70, 0x2b, 0x6b, 0xf1, /* |.W.Lp+k.| */
				0x1d, 0x5f, 0xac, /* |._.| */
			},
		},
	},
	LockTime: 0,
	}

	origTrans := btcutil.NewTx(&genesisCoinbaseTx) // btcutil/tx.go

	var transactions []*btcutil.Tx
	transactions = append(transactions, origTrans)
	transactions = append(transactions, trans2)

	merkles := blockchain.BuildMerkleTreeStore(transactions, false)
	merkleHash := merkles[len(merkles) - 1]

	// Creating a block header. Adapted from wire/blockheader_test.go
	nonce64, err := wire.RandomUint64()
	if err != nil {
		fmt.Printf("RandomUint64: Error generating nonce: %v\n", err)
	}
	nonce := uint32(nonce64)
	bits := uint32(0x207fffff) // Using powlimit defined chaincfg/params.go for the simnet blockchain
	blockHead := wire.NewBlockHeader(1, chaincfg.SimNetParams.GenesisHash, merkleHash, bits, nonce)

	// Creating an actual block
	msgBlock := wire.NewMsgBlock(blockHead) // btcd/wire/msgblock.go

	// add transaction btcd/wire/msgblock.go
	errCB := msgBlock.AddTransaction(&genesisCoinbaseTx)
	if errCB != nil {
		fmt.Printf("Error adding transaction: %v\n", err)
	}

	errTrans := msgBlock.AddTransaction(str)

	if errTrans != nil {
		fmt.Printf("Error adding transaction: %v\n", err)
	}

	newBlock := btcutil.NewBlock(msgBlock) // btcutil/block.go
	fmt.Println(newBlock.Hash())
	isMainChain, isOrphan, err := chain.ProcessBlock(newBlock,
    	blockchain.BFNone) // In github.com/btcsuite/btcd/blob/master/blockchain/process.go. Would be good to test this further.
	                       // Flags defined in process.go. BF.None indicates no flags.
	if err != nil {
    	fmt.Printf("Failed to process block: %v\n", err)
   		return
	}