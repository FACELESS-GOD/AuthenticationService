package Configurator

import (
	"os"

	"github.com/joho/godotenv"
)

type Cred struct {
	DBConn  string
	RDBAddr string
	RDBPass string
}

type Helper struct {
	Cred Cred
}

func NewHelper() Helper {
	Help := Helper{}
	return Help
}

func (Hel Helper) LoadConfig(EnvType string) error {
	err := godotenv.Load("")

	if err != nil {
		return nil
	}
	switch EnvType {
	case "Test":
		Hel.Cred.DBConn = os.Getenv("Dev_DBCONN")
		Hel.Cred.RDBAddr = os.Getenv("Dev_RDBADDR")
		Hel.Cred.RDBPass = os.Getenv("DEV_RDBPASSWORD")

	case "PROD":
		Hel.Cred.DBConn = os.Getenv("PROD_DBCONN")
		Hel.Cred.RDBAddr = os.Getenv("PROD_RDBADDR")
		Hel.Cred.RDBPass = os.Getenv("PROD_RDBPASSWORD")

	}
	return nil
}
