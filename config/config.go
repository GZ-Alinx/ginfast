package config


// ps: 一定要注意 settings-dev.yaml每个字段名称和结构体的tag一一对应!!
type ServerConfig struct {
	Name string `mapstructure:"name"`
	Port int `mapstructure:"port"`
	MysqlInfo MysqlConfig `mapstructure:"mysql"`
	RedisInfo RedisConfig `mapstructure:"redis"`
	LogsAddress string `mapstructure:"logsAddress"`
	JWTKey JWTConfig `mapstructure:"jwt"`
}

type MysqlConfig struct {
	Host string `mapstructure:"host"`
	Port int `mapstructure:"port"`
	User string `mapstructure:"username"`
	PassWord string `mapstructure:"password"`
	DBName string `mapstructure:"dbname"`
}

type RedisConfig struct {
	Host string `mapstructure:"host"`
	Port int `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB int `mapstructure:"db"`
}

type JWTConfig struct {
	SigningKey string `mapstructure:"key" json:"key"`
}