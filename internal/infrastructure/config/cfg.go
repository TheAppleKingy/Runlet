package config

type config interface {
	parse()
}

func LoadConfigs() {
	configs := []config{
		&DBConfig,
		&RunnersConfig,
	}
	for _, cfg := range configs {
		cfg.parse()
	}
}
