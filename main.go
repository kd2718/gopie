package main

import (
	"fmt"
	"github.com/kd2718/gopie/prepdb"
	"time"
	"github.com/jinzhu/gorm"
)

type Person struct{
	gorm.Model
	Name string `gorm:"type:varchar(100);unique_index"`
	Age int64
	Male bool
}

func main() {
	fmt.Println("Some go code...")
	fmt.Println("time now:", time.Now())

	conn, err := prepdb.GetGonn()
	defer conn.Close()
	if err != nil {
		fmt.Println(err)
	}
	conn.AutoMigrate(&Person{})
	new_person := Person{Name: "Cheri", Age: 33, Male: false}

	conn.Create(&new_person)
	fmt.Println("new person", new_person.Name, "new id", new_person.ID)

	var kory Person

	conn.Where(&Person{Male: true}).First(&kory)

	fmt.Println("Its a Me", kory.Name, kory.ID, kory.Age, kory.Male)
}