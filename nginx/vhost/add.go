package vhost

import (
	"fmt"
	"nginxctl/helper"
	"nginxctl/nginx"
	"nginxctl/nginx/configuration"
)

type vhostConfig struct {
	path string
}

// AddVHost create a new vhost configuration based on
// given --server-name, --root, and --type (reverse_proxy or static_site)
func AddVHost(c *configuration.Configuration) {
	path := fmt.Sprintf("%s/%s", nginx.SitesAvailableDirPath, c.ServerName)

	vc := &vhostConfig{path}

	// We have to make sure that no duplicate configuration
	// check if config (filepath) is exist, return error if exist
	isPath, _ := helper.IsFileExist(path)
	if isPath {
		err := fmt.Errorf(
			"Error: unable to create new configuration for %s, file already exist",
			path,
		)
		helper.ShowError(err)
	}

	// --type can only reverse_proxy or static_site
	if c.Type != nginx.StaticSite && c.Type != nginx.ReverseProxy {
		err := fmt.Errorf(
			"Error: --type can only %s or %s",
			nginx.StaticSite,
			nginx.ReverseProxy,
		)
		helper.ShowError(err)
	}

	switch c.Type {
	case nginx.StaticSite:
		vc.staticSiteVHost(c)
		return

	case nginx.ReverseProxy:
		vc.reverseProxyVHost(c)
		return
	}
}

func (v *vhostConfig) staticSiteVHost(c *configuration.Configuration) {
	if c.Root == "" {
		err := fmt.Errorf("Error: --root is required when you choose static_site type")
		helper.ShowError(err)
	}
	fmt.Println("Generate static site configuration")

	config, err := c.StaticSiteTemplate()
	if err != nil {
		helper.ShowError(err)
	}
	v.write(config)
}

func (v *vhostConfig) reverseProxyVHost(c *configuration.Configuration) {
	if c.ProxyPass == "" {
		err := fmt.Errorf("Error: --proxy-pass is required when you choose reverse_proxy type")
		helper.ShowError(err)
	}
	fmt.Println("Generate reverse proxy configuration")

	config, err := c.ReverseProxyTemplate()
	if err != nil {
		helper.ShowError(err)
	}
	v.write(config)
}

func (v *vhostConfig) write(config []byte) {
	// Write the configuration to /etc/nginx/sites-available
	err := helper.WriteConfig(v.path, config)
	if err != nil {
		helper.ShowError(err)
	}
	fmt.Println("Write the configuration to", v.path)

	// Write the configuration link to /etc/nginx/sites-enabled
	fmt.Println("Write the configuration link to", nginx.SitesEnabledDirPath)
	link, err := helper.WriteLink(v.path, nginx.SitesEnabledDirPath)
	if err != nil {
		helper.ShowError(err)
	}
	fmt.Println(link)
	fmt.Println("Successfully added configuration!")
}
