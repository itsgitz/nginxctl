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
	"nginxctl/nginx/vhost"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a virtualhost configuration",
	Long: `Warning: becareful, it will delete configuration on both
/etc/nginx/sites-available and /etc/nginx/sites-enabled directory`,
	Run: func(cmd *cobra.Command, args []string) {
		vhost.DeleteVHost(serverName)
	},
}

func init() {
	vhostCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().StringVar(&serverName, "server-name", "", "Select the server_name that will be deleted")
	deleteCmd.MarkFlagRequired("server-name")
}
