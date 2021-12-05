package proxies

type Proxies struct {
	Proxies []Proxy `yaml:"proxies"`
}

type Proxy struct {
	Endpoint    string            `yaml:"endpoint"`
	Api_version string            `yaml:"api_version"`
	Methods     []string          `yaml:"methods"`
	Headers     map[string]string `yaml:"headers"`
	Http_result int               `yaml:"http_result"`
	Body        string            `yaml:"body"`
}
