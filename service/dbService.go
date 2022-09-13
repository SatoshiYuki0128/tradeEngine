package service

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"time"
)

func DbConnect() *sql.DB {
	UserName := viper.GetString("DB.UserName")
	Password := viper.GetString("DB.Password")
	Addr := viper.GetString("DB.Addr")
	Port := viper.GetInt64("DB.Port")
	Database := viper.GetString("DB.Database")
	MaxLifetime := viper.GetInt("DB.MaxLifetime")
	MaxOpenConns := viper.GetInt("DB.MaxOpenConns")
	MaxIdleConns := viper.GetInt("DB.MaxIdleConns")

	connectString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", UserName, Password, Addr, Port, Database)
	//fmt.Println(connectString)

	DB, err := sql.Open("mysql", connectString)
	if err != nil {
		fmt.Println("db連線失敗!")
	}
	DB.SetConnMaxLifetime(time.Duration(MaxLifetime) * time.Second)
	DB.SetMaxOpenConns(MaxOpenConns)
	DB.SetMaxIdleConns(MaxIdleConns)
	return DB
}
