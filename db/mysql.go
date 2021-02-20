package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"studygin/config"
)

var SqlDb *sql.DB //数据库连接db

func init() {
	//1、打开数据库
	//parseTime:时间格式转换(查询结果为时间时，是否自动解析为时间);
	// loc=Local：MySQL的时区设置
	user := config.EnvConfig.MySQL.Username
	pwd := config.EnvConfig.MySQL.Password
	host := config.EnvConfig.MySQL.Host
	port := config.EnvConfig.MySQL.Port
	name := config.EnvConfig.MySQL.Name
	sqlStr := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true&loc=Local",
		user, pwd, host, port, name)
	var err error
	SqlDb, err = sql.Open("mysql", sqlStr)
	if err != nil {
		fmt.Println("数据库打开出现了问题：", err)
		return
	}
	//2、 测试与数据库建立的连接（校验连接是否正确）
	err = SqlDb.Ping()
	if err != nil {
		fmt.Println("数据库连接出现了问题：", err)
		return
	}
	fmt.Println("数据库连接成功!")
}

//todo:  go-sql-driver地址：https://github.com/go-sql-driver/mysql
