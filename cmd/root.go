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
	"os"
	"fmt"
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/config"

	"github.com/spf13/cobra"
)

type Flags struct {
	AutoApprove bool
	Status string
	LastUsed int
	Username string
}

const (
	GITHUB_URL = "https://github.com/tuladhar/cleanup-aws-access-keys"
)

type State struct {
	TableData [][]string
}

var (
	flags = Flags{}
	client *iam.Client
	state = &State{}
)

func AskApproval() bool {
	var answer string
	for {
		fmt.Printf("Proceed? (yes/no): ")
		fmt.Scanln(&answer)
		switch answer {
		case "no":
			return false
		case "yes":
			return true
		default:
			fmt.Println("Please type 'yes' or 'no'.")
		}
	}
}

func LoadConfigOrDie() (aws.Config) {
	config, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	return config
}

var rootCmd = &cobra.Command{
	Use:   "cleanup-aws-access-keys",
	Version: "1.0",
	Short: fmt.Sprintf("A cloud security tool to search and clean up unused AWS access keys (%s)\n", GITHUB_URL),
	Long: ``,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
