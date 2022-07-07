package config

type (
	Config struct {
		App   `yaml:"app"`
		Http  `yaml:"http"`
		Log   `yaml:"logger"`
		Mysql `yaml:"mysql"`
		Redis `yaml:"redis"`
		Wechat `yaml:"wechat"`
	}

	App struct {
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}

	Http struct {
		Addr string `env-required:"true" yaml:"addr" env:"HTTP_ADDR"`
		Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
		Domain string `env-required:"true" yaml:"domain" env:"HTTP_DOMAIN`
	}

	Log struct {
		Level string `env-required:"true" yaml:"log_level" env:"LOG_LEVEL"`
	}

	Mysql struct {
		Host     string `env-required:"true" yaml:"host" env:"MYSQL_HOST"`
		Username string `env-required:"true" yaml:"username" env:"MYSQL_USERNAME"`
		Password string `env-required:"true" yaml:"password" env:"MYSQL_PASSWORD"`
		Database string `env-required:"true" yaml:"database" env:"MYSQL_DATABASE"`
		Prefix   string `env-required:"true" yaml:"prefix" env:"MYSQL_PREFIX"`
		ShowSql  string `env-required:"true" yaml:"show_sql" env:"MYSQL_SHOW_SQL"`
	}

	Redis struct {
		RedisHost string `env-required:"true" yaml:"redis_host" env:"REDIS_HOST"`
		RedisPass string `yaml:"redis_pass" env:"REDIS_PASS"`
		RedisDatabase string `env-required:"true" yaml:"redis_database" env:"REDIS_HOST"`
	}

	Wechat struct {
		AppId	string 	`env-required:"true" yaml:"app_id" env:"APP_ID"`
		Secret	string 	`env-required:"true" yaml:"secret" env:"SECRET"`
		EncodingAESKey	string 	`yaml:"encoding_aes_key" env:"ENCODING_AES_KEY"`
		Token	string 	`yaml:"token" env:"TOKEN"`
	}
)
