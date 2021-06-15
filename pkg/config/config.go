package config

type Config struct {
	Host         string `env:"HOST" envDefault:"localhost"`
	Port         string `env:"PORT" envDefault:"8081"`
	YAMLFilePath string `env:"YAML_FILE_PATH" envDefault:"./assets/files/data.yaml"`
}
