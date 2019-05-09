package common

type CfgPg struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
	Schema   string
}

type CfgRd struct {
}

type CfgMg struct {
}

type CfgDb struct {
	Postgre *CfgPg
	Redis   *CfgRd
	Mongo   *CfgMg
}

type Service struct {
	Host       string
	Port       int
	SSL        bool
	PrivateKey string
	PublicKey  string
}

type Config struct {
	Database *CfgDb
	Service  *Service
}

func LoadConfig() (*Config, error) {
	var config = &Config{
		Database: &CfgDb{
			Postgre: &CfgPg{
				Host:     "171.244.49.164",
				Port:     5432,
				Username: "postgres",
				Password: "xoacdi!@#",
				Database: "spin_app",
			},
		},
		Service: &Service{
			Host: "localhost",
			Port: 8080,
		},
	}
	return config, nil
}
