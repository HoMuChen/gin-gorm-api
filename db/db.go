package db
import (
  "fmt"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"

  "../configs"
)

var db *gorm.DB
var err error

func Init() {
  connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
    configs.PG_HOST,
    configs.PG_PORT,
    configs.PG_USER,
    configs.PG_DB,
    configs.PG_PWD,
  )
  db, err = gorm.Open("postgres", connStr)

  if err != nil {
    fmt.Println("[Error] failed to open db connection")
    fmt.Println(err)
  }

  fmt.Println("[Info] database connected @", configs.PG_HOST, configs.PG_PORT)
}

func GetDB() *gorm.DB {
  return db
}

func Close() {
  fmt.Println("DB connection is now closing...")
  db.Close()
}
