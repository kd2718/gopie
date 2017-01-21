package prepdb

import(
	"github.com/jinzhu/gorm"
        _ "github.com/jinzhu/gorm/dialects/mysql"
)

func GetGonn() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "gopie:aaaa@/gopie?charset=utf8&parseTime=True&loc=Local")
	return db, err
}