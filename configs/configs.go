package configs

type Configs struct {
	App Gin
}

type Gin struct {
	Host string
	Port string
}
