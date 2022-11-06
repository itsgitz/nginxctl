package configuration

import (
	"bytes"
	"text/template"
)

func (c *Configuration) StaticSiteTemplate() ([]byte, error) {
	var result bytes.Buffer
	var configuration string = `server {
	server_name {{ .ServerName }};

	root {{ .Root }};
	index index.html;
	access_log off;

	location / {
		try_files $uri $uri/ =404;
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
