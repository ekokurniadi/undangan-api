package config

type constantName struct {
	Host     string
	Database string
	Username string
	Password string
	Port     string
}

func DeclareConstant(Host string, Database string, Username string, Password string, Port string) constantName {
	return constantName{
		Host:     Host,
		Database: Database,
		Username: Username,
		Password: Password,
		Port:     Port,
	}
}
