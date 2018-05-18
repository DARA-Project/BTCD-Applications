package main

import (
	"fmt"
	"github.com/btcsuite/btcd/database"
	_ "github.com/btcsuite/btcd/database/ffldb"
	"github.com/btcsuite/btcd/blockchain"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
	"os"
	"path/filepath"
)


func main() {
	dbPath := filepath.Join(os.TempDir(), "exampleprocessblock")
	_ = os.RemoveAll(dbPath)
	db, err := database.Create("ffldb", dbPath, chaincfg.SimNetParams.Net)
	if err != nil {
		fmt.Printf("Failed to create database: %v\n", err)
		return
	}
	defer os.RemoveAll(dbPath)
	defer db.Close()

	// Create a new BlockChain instance for a simulation test bitcoin network
	chain, err := blockchain.New(&blockchain.Config{
		DB:          db,
		ChainParams: &chaincfg.SimNetParams,
		TimeSource:  blockchain.NewMedianTime(),
	})
	if err != nil {
		fmt.Printf("Failed to create chain instance: %v\n", err)
		return
	}

	// Generate address and pay to it
	addrHash := [20]byte{0x01}

	addr, err := btcutil.NewAddressPubKeyHash(addrHash[:],
		&chaincfg.SimNetParams)
	if err != nil {
		fmt.Printf("NewAddressPubKeyHash: unexpected error: %v\n", err)
	}

	fmt.Printf("Chain: %d\n", chain)
	fmt.Printf("Address: %d\n", addr)
}
