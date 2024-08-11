package config

type DB struct {
	Host string `json:"host"`
	Port int    `json:"port"`
	User string `json:"user"`
	Pass string `json:"pass"`
	Name string `json:"name"`
}

type Config struct {
	Port int `json:"port"`
	Db   DB  `json:"db"`
}
