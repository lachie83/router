package nginx

import (
	"bytes"
	"regexp"
	"testing"
	"text/template"

	"github.com/Masterminds/sprig"
	"github.com/deis/router/model"
)

func TestDisableServerTokens(t *testing.T) {
	routerConfig := &model.RouterConfig{
		WorkerProcesses:          "auto",
		MaxWorkerConnections:     "768",
		TrafficStatusZoneSize:    "1m",
		DefaultTimeout:           "1300s",
		ServerNameHashMaxSize:    "512",
		ServerNameHashBucketSize: "64",
		GzipConfig: &model.GzipConfig{
			Enabled:     true,
			CompLevel:   "5",
			Disable:     "msie6",
			HTTPVersion: "1.1",
			MinLength:   "256",
			Proxied:     "any",
			Types:       "application/atom+xml application/javascript application/json application/rss+xml application/vnd.ms-fontobject application/x-font-ttf application/x-web-app-manifest+json application/xhtml+xml application/xml font/opentype image/svg+xml image/x-icon text/css text/plain text/x-component",
			Vary:        "on",
		},
		BodySize:          "1m",
		ProxyRealIPCIDRs:  []string{"10.0.0.0/8"},
		ErrorLogLevel:     "error",
		UseProxyProtocol:  false,
		EnforceWhitelists: false,
		WhitelistMode:     "extend",
		SSLConfig: &model.SSLConfig{
			Enforce:           false,
			Protocols:         "TLSv1 TLSv1.1 TLSv1.2",
			SessionTimeout:    "10m",
			UseSessionTickets: true,
			BufferSize:        "4k",
			HSTSConfig: &model.HSTSConfig{
				Enabled:           false,
				MaxAge:            15552000, // 180 days
				IncludeSubDomains: false,
				Preload:           false,
			},
		},

		DisableServerTokens: true,
	}

	var b bytes.Buffer

	tmpl, err := template.New("nginx").Funcs(sprig.TxtFuncMap()).Parse(confTemplate)

	if err != nil {
		t.Fatalf("Encountered an error: %v", err)
	}

	err = tmpl.Execute(&b, routerConfig)

	if err != nil {
		t.Fatalf("Encountered an error: %v", err)
	}

	validDirective := regexp.MustCompile(`(?m)^(\s*)server_tokens off;$`)

	if !validDirective.Match(b.Bytes()) {
		t.Errorf("Expected: 'server_tokens off' in the configuration. Actual: no match")
	}

}
