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
	"nginxctl/caller/vhost"
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
		onlyArgs := make(map[string]bool, 3)
		onlyArgs[nginx.SitesAvailableDir] = OnlySitesAvailable
		onlyArgs[nginx.SitesEnabledDir] = OnlySitesEnabled
		onlyArgs[nginx.ConfdDir] = OnlyConfd

		vhost.List(onlyArgs)
	},
}

var OnlySitesAvailable bool
var OnlySitesEnabled bool
var OnlyConfd bool

func init() {
	vhostCmd.AddCommand(listCmd)

	// --sites-available
	listCmd.Flags().BoolVar(
		&OnlySitesAvailable,
		"sites-available",
		false,
		"Only show configurations on /etc/nginx/sites-available",
	)

	// --sites-enabled
	listCmd.Flags().BoolVar(
		&OnlySitesEnabled,
		"sites-enabled",
		false,
		"Only show configurations on /etc/nginx/sites-enabled",
	)

	// --confd
	listCmd.Flags().BoolVar(
		&OnlyConfd,
		"confd",
		false,
		"Only show configurations on /etc/nginx/conf.d",
	)
}
