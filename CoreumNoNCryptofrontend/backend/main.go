package main

import (
	"github.com/CoreumFoundation/coreum/pkg/client"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/gin-gonic/gin"  // Recommended for REST API
)

// Initialize Coreum client
func initCoreum() (*client.Client, error) {
	config := client.Config{
		ChainID:       "coreum-mainnet-1",  // or testnet
		RPCEndpoint:   "https://full-node.mainnet-1.coreum.dev:26657",
		AccountPrefix: "core",
	}
	
	return client.NewClient(config)
}

func main() {
	// Initialize Gin router
	r := gin.Default()
	
	// Initialize Coreum client
	coreumClient, err := initCoreum()
	if err != nil {
		panic(err)
	}
	
	// Setup routes
	setupRoutes(r, coreumClient)
 
	// Run server
	r.Run(":8080")
} 