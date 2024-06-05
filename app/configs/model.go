package configs

type AppConfig struct {
	Name         string
	Env          string
	Port         string
	ExpiredToken int
	SecretKey    string
	UseRabbitMQ  bool
}

type DbConfig struct {
	Host        string
	Port        string
	DbName      string
	Username    string
	Password    string
	DbIsMigrate bool
	DebugMode   bool
}

type RedisConfig struct {
	Host                string `json:"host"`
	Port                int    `json:"port"`
	Password            string `json:"password"`
	Database            int    `json:"database"`
	Master              string `json:"master"`
	PoolSize            int    `json:"pool_size"`
	PoolTimeout         int    `json:"pool_timeout"`
	MinIdleConn         int    `json:"min_idle_conn"`
	DefaultCacheTimeOut int    `json:"default_cache_time_out"`
}

type RabbitMQConfig struct {
	Username string `json:"username"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Password string `json:"password"`
}

type Configs struct {
	App      AppConfig
	DB       DbConfig
	Redis    RedisConfig
	RabbitMQ RabbitMQConfig
}
