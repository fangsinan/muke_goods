package config

type UserSrv struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
	Name string `mapstructure:"name"`
}

type JwtConfig struct {
	Sign string `mapstructure:"key"`
}

type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}

type ServerConfig struct {
	Name       string       `mapstructure:"name"`
	WebPort    int          `mapstructure:"webPort"`
	WebHost    string       `mapstructure:"webHost"`
	UserSrv    UserSrv      `mapstructure:"user_srv"`
	Jwtinfo    JwtConfig    `mapstructure:"jwt"`
	ConsulInfo ConsulConfig `mapstructure:"consul" json:"consul"`
}
