package config

type AppConfig struct {
	Model  Model  `mapstructure:"model"`
	Commit Commit `mapstructure:"commit"`
}

type Model struct {
	ModelType string `mapstructure:"modelType"`
	Url       string `mapstructure:"url"`
	ApiKey    string `mapstructure:"apiKey"`
	Name      string `mapstructure:"name"`
}

type Commit struct {
	Template     string `mapstructure:"template"`
	TemplateFile string `mapstructure:"templateFile"`
	Issue        string `mapstructure:"issue"`
	Type         string `mapstructure:"type"`
}
