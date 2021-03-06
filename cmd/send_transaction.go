// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/johnnyeven/chain/blockchain"
	"github.com/johnnyeven/chain/services"
)

var (
	from string
	to string
	amount uint64
)

// sendTransactionCmd represents the sendTransaction command
var sendTransactionCmd = &cobra.Command{
	Use:   "sendTransaction",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		c := blockchain.NewBlockChain(blockchain.Config{
			NewGenesisBlockFunc: blockchain.NewGenesisBlock,
		})

		tran := blockchain.NewTransaction(from, to, amount, c)

		service := services.GetBlockChainService()
		if service != nil {
			service.GetTransChannel() <- &tran
		}
	},
}

func init() {
	RootCmd.AddCommand(sendTransactionCmd)

	sendTransactionCmd.Flags().StringVarP(&from, "from", "f", "", "")
	sendTransactionCmd.Flags().StringVarP(&to, "to", "t", "", "")
	sendTransactionCmd.Flags().Uint64VarP(&amount, "amount", "a", 0, "")
}
