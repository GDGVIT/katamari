/*
Copyright © 2020 DSCVIT

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"context"
	"fmt"
	"katamari/internal/utils"
	"os"
	"sync"
	"time"

	"github.com/google/go-github/v32/github"
	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build your katamari project",
	Long:  `Fetch all repos from the specified organization, clone the READMEs and generate static pages ready for hosting!`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 1 {
			msg := fmt.Sprintf("Ignoring extra arguments after %s", chalk.Green.Color(fmt.Sprintf("'%s'", args[0])))
			utils.Warn("optional", msg)
		} else if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}

		utils.Info("sill", fmt.Sprintf("fetching repos for %s", chalk.Green.Color(fmt.Sprintf("'%s'", args[0]))))

		os.MkdirAll("./content/readmes/", 0755)

		client := github.NewClient(nil)
		repos, _, err := client.Repositories.ListByOrg(context.Background(), args[0], &github.RepositoryListByOrgOptions{Type: "public"})
		if err != nil {
			utils.Err("enoent", err.Error())
			os.Exit(1)
		}

		var wg sync.WaitGroup

		for _, repo := range repos {
			wg.Add(1)
			utils.Info("sill", fmt.Sprintf("Fetching readme for repo %s", *repo.Name))
			go func(client *github.Client, repo *github.Repository, wg *sync.WaitGroup) {
				defer wg.Done()
				readme, _, err := client.Repositories.GetReadme(context.Background(), args[0], *repo.Name, nil)
				if err != nil {
					utils.Err("enoent", err.Error())
					return
				}

				content, err := readme.GetContent()
				if err != nil {
					utils.Err("enoent", err.Error())
					return
				}

				f, err := os.Create(fmt.Sprintf("./content/readmes/%s.md", *repo.Name))
				if err != nil {
					utils.Err("enoent", err.Error())
				}

				f.WriteString("---\n")
				f.WriteString(fmt.Sprintf("title: %s\n", *repo.Name))
				f.WriteString(fmt.Sprintf("date: %s\n", time.Now().UTC().Format("2006-01-02T15:04:05-0700")))
				f.WriteString(fmt.Sprintf("draft: false\n"))
				f.WriteString("---\n")
				f.WriteString(content)
			}(client, repo, &wg)
		}

		wg.Wait()
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
}
