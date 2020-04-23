package config

// Config represents the configuration of the CMS instance
type Config struct {
	AppPort           int    `validate:"required,min=1,max=49151"`
	AppSecret         []byte `validate:"required"`
	AppTheme          string `validate:"required"`
	AppForwardWebpack bool
	AppUserName       string `validate:"required"`
	AppUserPassword   []byte `validate:"required"`
	DBHost            string `validate:"required"`
	DBPort            int    `validate:"required,gte=1,lte=49151"`
	DBUser            string `validate:"required"`
	DBPassword        string
	DBName            string `validate:"required"`
	DBRequireSSL      bool
}
