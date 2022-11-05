/*
Copyright © 2022 Anggit M Ginanjar <anggit.ginanjar.dev@gmail.com>

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
	"nginxctl/nginx"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Get virtualhost configurations",
	Long: `Get all nginx virtualhost configurations from "/etc/nginx/sites-available",
"/etc/nginx/sites-enabled" (for enabled vhost), and "/etc/nginx/conf.d" directory`,
	Run: func(cmd *cobra.Command, args []string) {
		vhosts, err := nginx.GetAllVHosts()
		if err != nil {
			fmt.Println(err)
			return
		}

		sitesAvailableNum := len(vhosts[nginx.SitesAvailableDir])
		confdNum := len(vhosts[nginx.ConfdDir])
		sitesEnabledNum := len(vhosts[nginx.SitesEnabledDir])

		// Show configurations on /etc/nginx/sites-available
		if sitesAvailableNum > 0 {
			fmt.Printf("\nVirtualhost configurations on %s (total %d):\n\n",
				nginx.SitesAvailableDirPath,
				sitesAvailableNum,
			)
			for _, v := range vhosts[nginx.SitesAvailableDir] {
				fmt.Println("*", v)
			}
		}

		// Show configurations on /etc/nginx/conf.d
		if confdNum > 0 {
			fmt.Printf("\nVirtualhost configurations on %s (total %d):\n\n",
				nginx.ConfdDirPath,
				confdNum,
			)
			for _, v := range vhosts[nginx.ConfdDir] {
				fmt.Println("*", v)
			}
		}

		// Show configurations on /etc/nginx/sites-enabled
		if sitesEnabledNum > 0 {
			fmt.Printf("\nEnabled virtualhost configurations on %s (total %d):\n\n",
				nginx.SitesEnabledDirPath,
				sitesEnabledNum,
			)
			for _, v := range vhosts[nginx.SitesEnabledDir] {
				fmt.Println("*", v)
			}
		}
	},
}

func init() {
	vhostCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	//listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	//listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
