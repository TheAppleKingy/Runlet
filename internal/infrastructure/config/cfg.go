package config

// config is interface for structures representing any configs
type config interface {
	// parse defines how data will be parsed from config files or anywhere into config structures
	parse() error
}

// LoadConfigs load data for config structures defined in package "config".
func LoadConfigs() error {
	for _, cfg := range registry.cfgs {
		if err := cfg.parse(); err != nil {
			return err
		}
	}
	return nil
}

type configRegistry struct {
	cfgs []config
}

func (r *configRegistry) register(cfg config) {
	r.cfgs = append(r.cfgs, cfg)
}

var (
	DBConfig      = dbConfig{}
	AuthConfig    = authConfig{}
	AppConfig     = appConfig{}
	RunnersConfig = runnersConfig{}
	registry      = configRegistry{}
)

func init() {
	registry.register(&DBConfig)
	registry.register(&AuthConfig)
	registry.register(&AppConfig)
	registry.register(&RunnersConfig)
}
