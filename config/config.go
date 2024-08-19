package config

type DBConfig struct {
	Host                string `json:"host"                    validate:"required"`
	Port                int    `json:"port"                    validate:"required"`
	Name                string `json:"name"                    validate:"required"`
	User                string `json:"user"                    validate:"required"`
	Pass                string `json:"pass"                    validate:"required"`
	MaxIdleTimeInMinute int    `json:"max_idle_time_in_minute" validate:"required"`
	EnableSSLMode       bool   `json:"enable_ssl_mode"`
}

type DB struct {
	Read  DBConfig `json:"read"  validate:"required"`
	Write DBConfig `json:"write"  validate:"required"`
}

type Mode string

const DebugMode = Mode("debug")
const ReleaseMode = Mode("release")

type Config struct {
	Mode         Mode   `json:"mode,omitempty" validate:"required"`
	ServiceName  string `json:"service_name,omitempty" validate:"required"`
	HttpPort     int    `json:"http_port,omitempty"`
	JwtSecretKey string `json:"jwt_secret_key"`
	DB           DB     `json:"db,omitempty" validate:"required"`
}

var config *Config

func init() {
	config = &Config{}
}

func GetConfig() Config {
	return *config
}
