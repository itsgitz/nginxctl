package configuration

import (
	"bytes"
	"text/template"
)

func (c *Configuration) ReverseProxyTemplate() ([]byte, error) {
	var result bytes.Buffer
	var configuration string = `server {
	server_name {{ .ServerName }};

	access_log off;	

	location / {	
		proxy_pass		  {{ .ProxyPass }};
		proxy_redirect	  off;
		proxy_set_header  Host $host;
		proxy_set_header  X-Real-IP $remote_addr;
		proxy_set_header  X-Forwarded-For $proxy_add_x_forwarded_for;
		proxy_set_header  X-Forwarded-Host $server_name;		
		proxy_set_header  X-Forwarded-Proto $scheme;

		##
		# Proxy Cache Settings
		##
		proxy_cache default;
		proxy_cache_lock on;
		proxy_cache_use_stale updating;
		add_header X-Cache-Status $upstream_cache_status;
	}

    listen {{ .Listen }};
}`

	t, err := template.New("configuration").Parse(configuration)
	if err != nil {
		return nil, err
	}

	err = t.Execute(&result, c)
	if err != nil {
		return nil, err
	}

	return result.Bytes(), nil
}
