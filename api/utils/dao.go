package utils

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
}

func Engine_scan() (*xorm.Engine, error) {
	return xorm.NewEngine("sqlite3", "./starecms.db")
}
