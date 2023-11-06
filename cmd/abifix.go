package cmd

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// ABI.go

type ABI struct {
	Inputs          []Put  `json:"inputs"`
	Name            string `json:"name"`
	Outputs         []Put  `json:"outputs"`
	StateMutability string `json:"stateMutability"`
	Type            string `json:"type"`
}

// Put.go

type Put struct {
	InternalType string `json:"internalType"`
	Name         string `json:"name"`
	Type         string `json:"type"`
}

// abifixCmd represents the abifix command
var abifixCmd = &cobra.Command{
	Use:   "abifix",
	Short: "Fixes an ABI json file for abigen",
	Long: `A command used to "fix" ABI files for the Solidity utility command: abigen. 
	For whatever reason, abigen, won't allow for type error, it just simply will not work. 
	This command simply transforms a valid ABI json file, removing the errors, to create an abigen valid file`,
	Run: func(cmd *cobra.Command, args []string) {
		readAbi, err := os.Open(args[0])

		if err != nil {
			log.Fatalln(err)
		}

		var allAbis []ABI

		err = json.NewDecoder(readAbi).Decode(&allAbis)

		if err != nil {
			log.Fatalln(err)
		}

		updatedAbis := []ABI{}

		for _, abi := range allAbis {
			if abi.Type != "error" {
				updatedAbis = append(updatedAbis, abi)
			}
		}

		fileName := filepath.Base(readAbi.Name())
		baseName := strings.Split(fileName, ".")[0] + "_fixed.abi"
		ext := filepath.Ext(fileName)
		directory := filepath.Dir(readAbi.Name())

		marshaled, err := json.Marshal(updatedAbis)

		if err != nil {
			log.Fatalln(err)
		}
		os.WriteFile(filepath.Join(directory, baseName+ext), marshaled, os.ModePerm)
	},
}

func init() {
	rootCmd.AddCommand(abifixCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// abifixCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// abifixCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
