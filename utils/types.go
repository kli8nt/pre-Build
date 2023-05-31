package utils

type Data struct {
	Technology           string   `env:"TECHNOLOGY"`
	Version              string   `env:"VERSION"`
	RepositoryURL        string   `env:"REPO_URL"`
	GithubToken          string   `env:"TOKEN"`
	ApplicationName      string   `env:"APP_NAME" `
	RunCommand           string   `env:"RUN_CMD"`
	InstallCommand       string   `env:"INSTALL_CMD"`
	BuildCommand         string   `env:"BUILD_CMD"`
	DependenciesFiles    []string `env:"DEPENDENCIES" envSeparator:":"`
	IsStatic             bool     `env:"STATIC"`
	OutputDirectory      string   `env:"OUTPUT"`
	EnvironmentVariables []string `env:"ENVARS" envSeparator:";"`
	Port                 int      `env:"PORT" `
}
