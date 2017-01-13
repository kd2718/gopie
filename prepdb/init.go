package prepdb

import(
	"github.com/jinzhu/gorm"
        _ "github.com/jinzhu/gorm/dialects/mysql"
)

func get_conn() (db gorm.db, err error) {
	db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
}