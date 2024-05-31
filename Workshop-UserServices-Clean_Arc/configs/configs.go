package configs

type Configs struct {
	PostgreSQL PostgreSQL
	App        Fiber
	Sarama     Sarama
	Redis      Redis
}

type Fiber struct {
	Host string
	Port string
}

// Database
type PostgreSQL struct {
	Host     string
	Port     string
	Protocol string
	Username string
	Password string
	Database string
	SSLMode  string
}

type Sarama struct {
	Host  []string
	Group string
}

type Redis struct {
	Addr string
}
