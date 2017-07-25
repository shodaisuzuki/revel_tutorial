package controllers

import (
    // "github.com/revel/revel"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "github.com/jinzhu/gorm"
    "myapp/app/models"
    "fmt"
)

var DB *gorm.DB

func InitDB() {
    // gormでMySQL接続
    // 失敗したらerrに格納される
    // Openの第二引数は {username}:{password}@/{dbname}?charset=utf8&parseTime=True&loc=Local
    db, err := gorm.Open("mysql", "root:root@/my_app?charset=utf8&parseTime=True&loc=Local")
    if err != nil {
        fmt.Println(err)
    }

    db.DB()
    // マイグレード
    db.AutoMigrate(&models.Article{})

    // dbをDB(*gorm.DB)として外からも使えるようにしてあげます
    // (関数の外でvarで宣言してるのでこうすれば他のコントローラーからもDB.hoge()って感じで使える)
    DB = db

}
