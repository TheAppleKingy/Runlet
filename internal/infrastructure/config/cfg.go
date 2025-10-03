package config

// config is interface for structures representing any configs
type config interface {
	// parse defines how data will be parsed from config files into config structures
	parse() error
}

// LoadConfigs load data for config structures defined in package "config".
func LoadConfigs() error {
	configs := []config{
		&DBConfig,
		&RunnersConfig,
		&AuthConfig,
		&AppConfig,
	}
	for _, cfg := range configs {
		if err := cfg.parse(); err != nil {
			return err
		}
	}
	return nil
}
