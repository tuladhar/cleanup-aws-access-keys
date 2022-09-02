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
	"log"
	"time"
	"context"
	"strings"

	"github.com/spf13/cobra"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/olekukonko/tablewriter"
)

func SearchKeys() {
	today := time.Now()
	fmt.Printf("Starting %s v%s (%s) at %s\n", rootCmd.Use, rootCmd.Version, GITHUB_URL, today.Format(time.UnixDate))

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"UserName", "Access Key ID", "Last Used", "Last Service Used", "Region", "Status"})

	// Load AWS config
	config := LoadConfigOrDie()
	
	// Create an Amazon IAM service client
	client = iam.NewFromConfig(config)

	// Use paginator in case there are more than 100 users
	paginator := iam.NewListUsersPaginator(client, &iam.ListUsersInput{})

	// Loop until paginator has no more data
	for {
		page, err := paginator.NextPage(context.TODO())
		if err != nil {
			log.Fatal(err)
		}

		usernameFound := false
		for _, user := range page.Users {
			// Skip if --username flag is set, and doesn't match the IAM user
			if flags.Username != "" && *user.UserName != flags.Username {
				continue
			}

			// Fetch access keys filter by username
			keys, err := client.ListAccessKeys(context.TODO(), &iam.ListAccessKeysInput{
				UserName: user.UserName,
			})
			if err != nil {
				log.Fatal(err)
			}

			// Loop through each access keys
			for _, m := range keys.AccessKeyMetadata {
				status := strings.ToLower(fmt.Sprintf("%s", m.Status))

				// Skip if --status flag is set, and doesn't match the key status
				if flags.Status != "" && status != flags.Status {
					continue
				}

				// Request when was the last used date of access keys
				lu, err := client.GetAccessKeyLastUsed(context.TODO(), &iam.GetAccessKeyLastUsedInput{
					AccessKeyId: m.AccessKeyId,
				})
				if err != nil {
					log.Print(err)
					continue
				}

				luDate := lu.AccessKeyLastUsed.LastUsedDate
				luServiceName := aws.ToString(lu.AccessKeyLastUsed.ServiceName)
				luRegion := aws.ToString(lu.AccessKeyLastUsed.Region)
				luInDays := 0
				luDateHuman := "N/A"

				// Calculate how many days ago the access keys was last accessed relatively to today
				if luDate != nil {
					luInDays = int(today.Sub(*luDate).Hours() / 24)
					if luInDays == 0 {
						luDateHuman = fmt.Sprintf("today", )
					} else {
						luDateHuman = fmt.Sprintf("%d days ago", luInDays)
					}
				}

				// Skip if --last-used N days is out of bound
				if luInDays < flags.LastUsed {
					continue
				}

				// If --last-used flag is set to -1 then skip used access keys
				if flags.LastUsed == -1 && luDate != nil {
					continue
				}

				// Populate the table to be later rendered as such:
				// +---------------+----------------------+---------------+-------------------+-----------+----------+
				// |   USERNAME    |    ACCESS KEY ID     |   LAST USED   | LAST SERVICE USED |  REGION   |  STATUS  |
				// +---------------+----------------------+---------------+-------------------+-----------+----------+
				// | devops        | AKIAYD7BEQCFXIKIPU49 | N/A           | N/A               | N/A       | inactive |
				// | tuladhar.puru | AKIAYR7BEQCFXIKYPU40 | 180 days ago  | sts               | us-east-1 | active   |
				// | puru.tuladhar | AKIAK7UAUHHZ29ZNEEHQ | today         | s3                | eu-west-1 | active   |
				// +---------------+----------------------+---------------+-------------------+-----------+----------+				
				data := []string{aws.ToString(m.UserName), aws.ToString(m.AccessKeyId), luDateHuman, luServiceName, luRegion, status}
				state.TableData = append(state.TableData, data)
				table.Append(data)
			}
			
			// Break out of loop if --username flag condition is met 
			if flags.Username != "" && *user.UserName == flags.Username {
				usernameFound = true
				break
			}
		}
		// Break out of loop if --username flag condition is met
		// or no more users left to paginate. 
		if usernameFound || !paginator.HasMorePages() {
			break
		}
	}

	// Render the table if there's table data
	fmt.Println()
	if len(state.TableData) != 0 {
		table.Render()
	} else {
		fmt.Printf("No access key(s) found.\n")
		os.Exit(0)
	}
}

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for access key(s)",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		SearchKeys()
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	searchCmd.Flags().IntVarP(&flags.LastUsed, "last-used", "", 0, "access key was last used n days.")
	searchCmd.Flags().StringVarP(&flags.Username, "username", "", "", "access key owned by username")
	searchCmd.Flags().StringVarP(&flags.Status, "status", "", "", "access key status: active or inactive")
}
