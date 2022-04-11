package mysql

import "gorm.io/gorm"

type MysqlDB struct {
	Db *gorm.DB
}

func init() {

}
