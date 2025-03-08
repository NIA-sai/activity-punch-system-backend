package config

type Mode string

const (
	ModeDebug   Mode = "debug"
	ModeRelease Mode = "release"
)

type Config struct {
	Host   string `envconfig:"HOST"`
	Port   string `envconfig:"PORT"`
	Prefix string `envconfig:"PREFIX"`
	Mode   Mode   `envconfig:"MODE"`
	OTel   OTel
	Mysql  Mysql
	JWT    JWT
}

type Mysql struct {
	Host     string `envconfig:"HOST"`
	Port     string `envconfig:"PORT"`
	Username string `envconfig:"USERNAME"`
	Password string `envconfig:"PASSWORD"`
	DBName   string `envconfig:"DB_NAME"`
}

type JWT struct {
	AccessSecret string `envconfig:"ACCESS_SECRET"`
	AccessExpire int64  `envconfig:"ACCESS_EXPIRE"`
}

type OTel struct {
	Enable      bool   `envconfig:"ENABLE"`
	ServiceName string `envconfig:"SERVICE_NAME"`
	Endpoint    string `envconfig:"ENDPOINT"`
	AgentHost   string `envconfig:"AGENT_HOST"`
	AgentPort   string `envconfig:"AGENT_PORT"`
}
