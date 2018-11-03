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
	"agenda/service"
	"fmt"

	"github.com/spf13/cobra"
)

// queryuserCmd represents the queryuser command
var queryuserCmd = &cobra.Command{
	Use:   "queryuser",
	Short: "Query user",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("queryuser called")

		_, flag := service.GetCurUser()
		if flag == false {
			fmt.Println("Please login")
			return
		}
		allUser := service.ListAllUser()
		for _, u := range allUser {
			fmt.Println("----------------")
			fmt.Println("Username: ", u.Name)
			fmt.Println("Phone: ", u.Phone)
			fmt.Println("Email: ", u.Email)
			fmt.Println("----------------")
		}

	},
}

func init() {
	RootCmd.AddCommand(queryuserCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// queryuserCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// queryuserCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
