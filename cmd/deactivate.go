/*
MIT License

Copyright (c) 2022 Puru Tuladhar (ptuladhar3@gmail.com)

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/
package cmd

import (
	"fmt"
	"context"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"

	"github.com/spf13/cobra"
)

func deactivateKeys() {
	var n int
	for _, d := range state.TableData {
		_, err := client.UpdateAccessKey(context.TODO(), &iam.UpdateAccessKeyInput{
			UserName: &d[0],
			AccessKeyId: &d[1],
			Status: types.StatusTypeInactive,
		})
		if err != nil {
			fmt.Printf("Unable to deactivate access key %s for username %s: %s\n", d[1], d[0], err)
			continue
		}
		n += 1
	}
	fmt.Printf("\nSuccessfully deactivated %d access key(s).\n", n)
}

// deactivateCmd represents the deactivate command
var deactivateCmd = &cobra.Command{
	Use:   "deactivate",
	Short: "Deactivate access key(s)",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		flags.Status = "active"
		SearchKeys()
		
		if !flags.AutoApprove {
			fmt.Println()
			fmt.Printf("Are you sure you want to DEACTIVATE %d access key(s)?\n", len(state.TableData))
			fmt.Printf("WARNING: You can't use a disabled key to make AWS API calls but you can activate it again later.\n")
			if AskApproval() {
				deactivateKeys()
			}
		} else {
			deactivateKeys()
		}
	},
}

func init() {
	rootCmd.AddCommand(deactivateCmd)

	deactivateCmd.Flags().IntVarP(&flags.LastUsed, "last-used", "", 0, "access key was last used n days.")
	deactivateCmd.Flags().StringVarP(&flags.Username, "username", "", "", "access key owned by username")
	deactivateCmd.Flags().StringVarP(&flags.Status, "status", "", "", "access key status: active or inactive")	
	deactivateCmd.Flags().BoolVarP(&flags.AutoApprove, "auto-approve", "", false, "automatic yes to prompts and run non-interactively.")
}
