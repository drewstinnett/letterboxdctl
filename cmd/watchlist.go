/*
Copyright © 2022 Drew Stinnett <drew@drewlink.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"

	"github.com/apex/log"
	"github.com/drewstinnett/go-letterboxd"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

// watchlistCmd represents the watchlist command
var watchlistCmd = &cobra.Command{
	Use:   "watchlist <username>",
	Short: "Show all the watchlist films for a user",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filmC := make(chan *letterboxd.Film)
		errorC := make(chan error)
		fmt.Println("---")
		go client.User.StreamWatchList(ctx, args[0], filmC, errorC)
		for {
			select {

			case film := <-filmC:
				d, err := yaml.Marshal([]letterboxd.Film{
					*film,
				})
				cobra.CheckErr(err)
				fmt.Print(string(d))
				stats.Total++
			case err := <-errorC:
				if err != nil {
					log.WithError(err).Error("Error streaming watchlist")
				} else {
					return
				}
			default:
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(watchlistCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// watchlistCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// watchlistCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
