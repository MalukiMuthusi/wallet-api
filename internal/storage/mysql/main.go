package mysql

import (
	"fmt"
	"os"

	"github.com/MalukiMuthusi/wallet-api/internal/utils"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlDB struct {
	Db *gorm.DB
}

func init() {

}

func SetUp() (*gorm.DB, error) {
	user := viper.GetString(utils.DbUser)
	pass := viper.GetString(utils.DbPwd)
	port := viper.GetString(utils.DbPort)
	dbName := viper.GetString(utils.DbName)

	var dsn string

	if viper.GetBool(utils.DbHostedOnCloud) {
		instanceConnectionName := viper.GetString(utils.DbConnectionName)
		socketDir, isSet := os.LookupEnv("DB_SOCKET_DIR")
		if !isSet {
			socketDir = "/cloudsql"
		}

		dsn = fmt.Sprintf("%s:%s@unix(/%s/%s)/%s?parseTime=true", user, pass, socketDir, instanceConnectionName, dbName)

	} else {

		dsn = fmt.Sprintf("%s:%s@tcp(127.0.0.1:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, port, dbName)
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	return db, err
}
