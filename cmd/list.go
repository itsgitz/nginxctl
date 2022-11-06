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
THE SOFTWARE. */
package cmd

import (
	"fmt"
	"nginxctl/helper"
	"nginxctl/nginx"
	"nginxctl/nginx/vhost"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Get virtualhost configurations",
	Long: `Get all nginx virtualhost configurations from "/etc/nginx/sites-available",
"/etc/nginx/sites-enabled" (for enabled vhost), and "/etc/nginx/conf.d" directory`,
	Run: func(cmd *cobra.Command, args []string) {
		specific := &specificArgument{}
		specific.sitesAvailable = onlySitesAvailable
		specific.sitesEnabled = onlySitesEnabled
		specific.confd = onlyConfd
		list(specific)
	},
}

type specificArgument struct {
	sitesAvailable bool
	sitesEnabled   bool
	confd          bool
}

var (
	onlySitesAvailable bool
	onlySitesEnabled   bool
	onlyConfd          bool
)

func init() {
	vhostCmd.AddCommand(listCmd)

	// --sites-available
	listCmd.Flags().BoolVar(
		&onlySitesAvailable,
		"sites-available",
		false,
		"Only show configurations on /etc/nginx/sites-available",
	)

	// --sites-enabled
	listCmd.Flags().BoolVar(
		&onlySitesEnabled,
		"sites-enabled",
		false,
		"Only show configurations on /etc/nginx/sites-enabled",
	)

	// --confd
	listCmd.Flags().BoolVar(
		&onlyConfd,
		"confd",
		false,
		"Only show configurations on /etc/nginx/conf.d",
	)
}

func list(specific *specificArgument) {
	vhosts, err := vhost.GetAllVHosts()
	if err != nil {
		helper.ShowError(err)
	}

	if specific.sitesAvailable {
		// Only --sites-available
		show(vhosts[nginx.SitesAvailableDir], nginx.SitesAvailableDirPath)
	} else if specific.sitesEnabled {
		// Only --sites-enabled
		show(vhosts[nginx.SitesEnabledDir], nginx.SitesEnabledDirPath)
	} else if specific.confd {
		// Only --confd
		show(vhosts[nginx.ConfdDir], nginx.ConfdDirPath)
	} else {
		// All, means no argument specified
		show(vhosts[nginx.SitesAvailableDir], nginx.SitesAvailableDirPath)
		show(vhosts[nginx.SitesEnabledDir], nginx.SitesEnabledDirPath)
		show(vhosts[nginx.ConfdDir], nginx.ConfdDirPath)
	}
}

func show(vhosts []string, path string) {
	num := len(vhosts)
	fmt.Printf("Configurations on %s (total %d):\n", path, num)
	if num == 0 {
		fmt.Printf("* No config file\n\n")
		return
	}

	for _, v := range vhosts {
		fmt.Println("*", v)
	}
	fmt.Println()
	return
}
