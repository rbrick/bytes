package cmd

import (
	"log"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rbrick/bytes/api"
	"github.com/spf13/cobra"
)

// httpCmd represents the http command
var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "Starts HTTP server",
	Long:  `Starts the backend RESTful API service`,
	Run: func(cmd *cobra.Command, args []string) {
		rpcClient, err := ethclient.Dial(os.Getenv("RPC_HOST"))

		if err != nil {
			log.Fatalln(err)
		}

		server := api.NewServer(rpcClient)
		log.Fatalln(server.Run(os.Getenv("HOST")))
	},
}

func init() {
	rootCmd.AddCommand(httpCmd)
}
