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
	"context"
	"fmt"

	"github.com/apex/log"
	"github.com/drewstinnett/go-letterboxd"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

// batchCmd represents the batch command
var batchCmd = &cobra.Command{
	Use:   "batch",
	Short: "Retrieve a batch of movies from letterboxd.com",
	Run: func(cmd *cobra.Command, args []string) {
		userWatched, err := cmd.Flags().GetStringArray("watched")
		cobra.CheckErr(err)

		// Get lists
		listsA, err := cmd.Flags().GetStringArray("list")
		cobra.CheckErr(err)
		lists, err := letterboxd.ParseListArgs(listsA)
		cobra.CheckErr(err)

		// Get Watch lists
		watchLists, err := cmd.Flags().GetStringArray("watchlist")
		cobra.CheckErr(err)

		filmOpts := &letterboxd.FilmBatchOpts{
			Watched:   userWatched,
			List:      lists,
			WatchList: watchLists,
		}

		ctx := context.Background()
		filmC := make(chan *letterboxd.Film)
		done := make(chan error)
		go client.Film.StreamBatch(ctx, filmOpts, filmC, done)
		for {
			select {

			case film := <-filmC:
				d, err := yaml.Marshal([]letterboxd.Film{
					*film,
				})
				cobra.CheckErr(err)
				fmt.Println(string(d))
				stats.Total++
			case err := <-done:
				if err != nil {
					log.WithError(err).Error("Error batch streaming watched")
				} else {
					log.Debug("Finished batch streaming")
					return
				}
			default:
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(batchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// batchCmd.PersistentFlags().String("foo", "", "A help for foo")
	batchCmd.PersistentFlags().StringArray("watched", []string{}, "Watched films for a given user")
	batchCmd.PersistentFlags().StringArray("list", []string{}, "User list in the format of {username}/{list-slug}")
	batchCmd.PersistentFlags().StringArray("watchlist", []string{}, "Films on a given users Watch List")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// batchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
