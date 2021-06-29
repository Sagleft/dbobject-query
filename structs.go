package dbquery

// DBCredentials - database connection data
type DBCredentials struct {
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"db"`
	Host     string `json:"host"` // by default: localhost
	Port     string `json:"port"`
}
