package domain

type SectionService struct {
	App   AppAccount   `json:",omitempty"`
	DB    DBAccount    `json:",omitempty"`
	Redis RedisAccount `json:",omitempty"`
}

type AppAccount struct {
	Name         string `json:",omitempty"`
	Environtment string `json:",omitempty"`
	URL          string `json:",omitempty"`
	Port         string `json:",omitempty"`
	SecretKey    string `json:",omitempty"`
}

type DBAccount struct {
	Username     string `json:",omitempty"`
	Password     string `json:",omitempty"`
	URL          string `json:",omitempty"`
	Port         string `json:",omitempty"`
	DBName       string `json:",omitempty"`
	MaxIdleConns int    `json:",omitempty"`
	MaxOpenConns int    `json:",omitempty"`
	MaxLifeTime  int    `json:",omitempty"`
	Timeout      string `json:",omitempty"`
}

type RedisAccount struct {
	Username     string `json:",omitempty"`
	Password     string `json:",omitempty"`
	URL          string `json:",omitempty"`
	Port         int    `json:",omitempty"`
	MinIdleConns int    `json:",omitempty"`
	Timeout      string `json:",omitempty"`
}
