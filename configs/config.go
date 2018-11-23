package configs

// Configs models the configuration required for connection string
type Config struct {
	DBAddress      []string `json:"dbAddress"`
	UserName       string   `json:"userName"`
	Password       string   `json:"password"`
	DBName         string   `json:"dbName"`
	CollectionName string   `json:"CollectionName"`
}

// DbConfigs A collection of authorized db connection strings
var DbConfigs = map[string]Config{
	"mongo": {[]string{"localhost:27017"}, "", "", "IceCreams", "BenJerrysAPI"},
}
