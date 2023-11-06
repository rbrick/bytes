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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// httpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// httpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
