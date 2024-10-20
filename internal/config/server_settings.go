package config

import "time"

type ServerSettings struct {
	AppEnv     string             `mapstructure:"APP_ENV"`
	HTTPServer HTTPServerSettings `mapstructure:",squash"`
	Clients    ClientsSettings    `mapstructure:",squash"`
}

type HTTPServerSettings struct {
	Address          string        `mapstructure:"HTTP_SERVER_ADDRESS"`
	Timeout          time.Duration `mapstructure:"HTTP_SERVER_TIMEOUT" envDefault:"10s"`
	IdleTimeout      time.Duration `mapstructure:"HTTP_SERVER_IDLE_TIMEOUT" envDefault:"60s"`
	RequestLimitByIP int           `mapstructure:"HTTP_SERVER_REQUEST_LIMIT_BY_IP" envDefault:"100"`
}

type ClientsSettings struct {
	SSO Client `mapstructure:",squash"`
}

type Client struct {
	Address      string        `mapstructure:"SSO_CLIENT_ADDRESS"`
	Timeout      time.Duration `mapstructure:"SSO_CLIENT_TIMEOUT"`
	RetriesCount int           `mapstructure:"SSO_CLIENT_RETRIES_COUNT"`
	// TODO: implement secure transport
	// Insecure     bool          `mapstructure:"SSO_CLIENT_INSECURE"`
}
