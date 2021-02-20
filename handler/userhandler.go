package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"studygin/db"
)

//Client提交的数据
type SqlUser struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

//应答体（响应client的请求）
type SqlResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

var sqlResponse SqlResponse

func DeleteData(c *gin.Context) {
	name := c.Query("name")
	var count int
	//1、先查询
	sqlStr := "select count(*) from user where name=?"
	err := db.SqlDb.QueryRow(sqlStr, name).Scan(&count)
	if count <= 0 || err != nil {
		sqlResponse.Code = http.StatusBadRequest
		sqlResponse.Message = "删除的数据不存在"
		sqlResponse.Data = "error"
		c.JSON(http.StatusOK, sqlResponse)
		return
	}
	//2、再删除
	delStr := "delete from user where name=?"
	ret, err := db.SqlDb.Exec(delStr, name)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		sqlResponse.Code = http.StatusBadRequest
		sqlResponse.Message = "删除失败"
		sqlResponse.Data = "error"
		c.JSON(http.StatusOK, sqlResponse)
		return
	}
	sqlResponse.Code = http.StatusOK
	sqlResponse.Message = "删除成功"
	sqlResponse.Data = "OK"
	c.JSON(http.StatusOK, sqlResponse)
	fmt.Println(ret.LastInsertId()) //打印结果
}

func UpdateData(c *gin.Context) {
	var u SqlUser
	err := c.Bind(&u)
	if err != nil {
		sqlResponse.Code = http.StatusBadRequest
		sqlResponse.Message = "参数错误"
		sqlResponse.Data = "error"
		c.JSON(http.StatusOK, sqlResponse)
		return
	}
	sqlStr := "update user set age=? ,address=? where name=?"
	ret, err := db.SqlDb.Exec(sqlStr, u.Age, u.Address, u.Name)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		sqlResponse.Code = http.StatusBadRequest
		sqlResponse.Message = "更新失败"
		sqlResponse.Data = "error"
		c.JSON(http.StatusOK, sqlResponse)
		return
	}
	sqlResponse.Code = http.StatusOK
	sqlResponse.Message = "更新成功"
	sqlResponse.Data = "OK"
	c.JSON(http.StatusOK, sqlResponse)
	fmt.Println(ret.LastInsertId()) //打印结果
}

func GetMulData(c *gin.Context) {
	address := c.Query("address")
	sqlStr := "select name,age from user where address=?"
	rows, err := db.SqlDb.Query(sqlStr, address)
	if err != nil {
		sqlResponse.Code = http.StatusBadRequest
		sqlResponse.Message = "查询错误"
		sqlResponse.Data = "error"
		c.JSON(http.StatusOK, sqlResponse)
		return
	}
	defer rows.Close()
	resUser := make([]SqlUser, 0)
	for rows.Next() {
		var userTemp SqlUser
		rows.Scan(&userTemp.Name, &userTemp.Age)
		userTemp.Address = address
		resUser = append(resUser, userTemp)
	}
	sqlResponse.Code = http.StatusOK
	sqlResponse.Message = "读取成功"
	sqlResponse.Data = resUser
	c.JSON(http.StatusOK, sqlResponse)
}

func GetData(c *gin.Context) {
	name := c.Query("name")
	sqlStr := "select age,address from user where name=?"
	var u SqlUser
	err := db.SqlDb.QueryRow(sqlStr, name).Scan(&u.Age, &u.Address)
	if err != nil {
		sqlResponse.Code = http.StatusBadRequest
		sqlResponse.Message = "查询错误"
		sqlResponse.Data = "error"
		c.JSON(http.StatusOK, sqlResponse)
		return
	}
	u.Name = name
	sqlResponse.Code = http.StatusOK
	sqlResponse.Message = "读取成功"
	sqlResponse.Data = u
	c.JSON(http.StatusOK, sqlResponse)
}

func InsertData(c *gin.Context) {
	var u SqlUser
	err := c.Bind(&u)
	if err != nil {
		fmt.Println(err)
		sqlResponse.Code = http.StatusBadRequest
		sqlResponse.Message = "参数错误"
		sqlResponse.Data = "error"
		c.JSON(http.StatusOK, sqlResponse)
		return
	}
	sqlStr := "insert into user(name, age, address) values (?,?,?)"
	//ret, err := db.SqlDb.Exec(sqlStr, u.Name, u.Age, u.Address)
	_, err = db.SqlDb.Exec(sqlStr, u.Name, u.Age, u.Address)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		sqlResponse.Code = http.StatusBadRequest
		sqlResponse.Message = "写入失败"
		sqlResponse.Data = "error"
		c.JSON(http.StatusOK, sqlResponse)
		return
	}
	sqlResponse.Code = http.StatusOK
	sqlResponse.Message = "写入成功"
	sqlResponse.Data = "OK"
	c.JSON(http.StatusOK, sqlResponse)
	//fmt.Println(ret.LastInsertId()) //打印结果

}
