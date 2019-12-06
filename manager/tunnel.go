package manager

// TunnelConfig represents all the configs collected from a user's config file
type TunnelConfig struct {
	Owner                  string   `json:"owner" mapstructure:"owner"`
	APIKey                 string   `json:"api_key" mapstructure:"api-key"`
	Cainfo                 string   `json:"cainfo" mapstructure:"cainfo"`
	Capath                 string   `json:"capath" mapstructure:"capath"`
	Certfile               string   `json:"certfile" mapstructure:"certificate"`
	DNS                    []string `json:"dns" mapstructure:"dns"`
	DirectDomains          []string `json:"direct_domains" mapstructure:"direct-domains"`
	FastFailRegexps        []string `json:"fast_fail_regexps" mapstructure:"fast-fail-regexps"`
	Logfile                string   `json:"logfile" mapstructure:"logfile"`
	MaxLogsize             int64    `json:"max_logsize" mapstructure:"max-logsize"`
	MetricsAddress         string   `json:"metrics_address" mapstructure:"metrics-address"`
	NoAutodetect           bool     `json:"no_autodetect" mapstructure:"no-autodetect"`
	NoCertVerifyFlag       bool     `json:"no_cert_verify" mapstructure:"no-cert-verify"`
	NoHTTPCertVerifyFlag   bool     `json:"no_http_cert_verify" mapstructure:"no-http-cert-verify"`
	NoProxyCachingFlag     bool     `json:"no_proxy_caching_flag" mapstructure:"no-proxy-caching"`
	NoSslBumpDomains       []string `json:"no_ssl_bump_domains" mapstructure:"no-ssl-bump-domains"`
	OutputConfigFlag       bool     `json:"output_config_flag" mapstructure:"output-config"`
	Pacurl                 string   `json:"pacurl" mapstructure:"pac"`
	Pidfile                string   `json:"pidfile" mapstructure:"pidfile"`
	Proxy                  string   `json:"proxy" mapstructure:"proxy"`
	ProxyTunnelFlag        bool     `json:"proxy_tunnel_flag" mapstructure:"proxy-tunnel"`
	ProxyUserpwd           string   `json:"proxy_userpwd" mapstructure:"proxy-userpwd"`
	Readyfile              string   `json:"readyfile" mapstructure:"readyfile"`
	RemoveCollidingTunnels bool     `json:"remove_colliding_tunnels"`
	RestURL                string   `json:"rest_url" mapstructure:"rest-url"`
	ScproxyPort            int      `json:"scproxy_port" mapstructure:"scproxy-port"`
	ScproxyReadLimit       uint     `json:"scproxy_read_limit" mapstructure:"scproxy-read-limit"`
	ScproxyWriteLimit      uint     `json:"scproxy_write_limit" mapstructure:"scproxy-write-limit"`
	SePort                 int      `json:"se_port" mapstructure:"se-port"`
	SharedTunnel           bool     `json:"shared_tunnel_flag" mapstructure:"shared-tunnel"`
	TunnelCainfo           string   `json:"tunnel_cainfo" mapstructure:"tunnel-cainfo"`
	TunnelCapath           string   `json:"tunnel_capath" mapstructure:"tunnel-capath"`
	TunnelCert             string   `json:"tunnel_cert" mapstructure:"tunnel-cert"`
	TunnelCertSuffix       string   `json:"tunnel_cert_suffix" mapstructure:"tunnel-cert-suffix"`
	TunnelDomains          []string `json:"tunnel_domains" mapstructure:"tunnel-domains"`
	TunnelIdentifier       string   `json:"tunnel_identifier" mapstructure:"tunnel-identifier"`
	PoolSize               int32    `json:"pool_size" mapstructure:"pool-size"`
	VeryVerbose            bool     `json:"very_verbose" mapstructure:"very-verbose"`
}
