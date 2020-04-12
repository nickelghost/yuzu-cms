package config

// Config represents the configuration of the CMS instance
type Config struct {
	AppPort           int
	AppSecret         []byte
	AppTheme          string
	AppForwardWebpack bool
	AppUserName       string
	AppUserPassword   []byte
	DBHost            string
	DBPort            int
	DBUser            string
	DBPassword        string
	DBName            string
	DBRequireSSL      bool
}
