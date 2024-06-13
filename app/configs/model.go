package configs

type Configs struct {
	App      AppConfig
	DB       DbConfig
	Redis    RedisConfig
	RabbitMQ RabbitMQConfig
	Email    EmailConfig
}

type AppConfig struct {
	Name         string `json:"name"`
	Env          string `json:"env"`
	Port         string `json:"port"`
	ExpiredToken int    `json:"expired_token"`
	SecretKey    string `json:"secret_key"`
	UseRabbitMQ  bool   `json:"use_rabbit_mq"`
}

type DbConfig struct {
	Host        string `json:"host"`
	Port        string `json:"port"`
	DbName      string `json:"db_name"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	DbIsMigrate bool   `json:"db_is_migrate"`
	DebugMode   bool   `json:"debug_mode"`
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

type EmailConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	From     string `json:"from"`
}
