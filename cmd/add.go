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
	"nginxctl/nginx/configuration"
	"nginxctl/nginx/vhost"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a virtualhost configuration",
	Long: `Currently only support adding configuration
for static website and reverse proxy`,
	Run: func(cmd *cobra.Command, args []string) {
		conf := &configuration.Configuration{}
		conf.ServerName = serverName
		conf.Type = siteType
		conf.Listen = listen
		conf.Root = root
		conf.ProxyPass = proxyPass
		vhost.AddVHost(conf)
	},
}

var (
	serverName string
	siteType   string
	listen     int64

	// root, document root for static website
	root string

	// proxyPass, for proxy_pass address
	proxyPass string
)

func init() {
	vhostCmd.AddCommand(addCmd)

	// --server-name
	addCmd.Flags().StringVar(&serverName, "server-name", "", "Define the server_name")

	// --type
	addCmd.Flags().StringVar(
		&siteType,
		"type",
		"",
		`Define the vhost type, the value can be "static_site" and "reverse_proxy"`,
	)

	// --root (document root)
	addCmd.Flags().StringVar(&root, "root", "", "Define the document root directory")

	// --proxy-address (proxy_pass address)
	addCmd.Flags().StringVar(&proxyPass, "proxy-pass", "", "Define the proxy_pass address")

	// --listen port
	addCmd.Flags().Int64Var(&listen, "listen", 80, "Define the listen port")

	// all flags are required
	addCmd.MarkFlagRequired("server-name")
	addCmd.MarkFlagRequired("type")
	addCmd.MarkFlagsRequiredTogether("server-name", "type")
	addCmd.MarkFlagsMutuallyExclusive("root", "proxy-pass")
}
