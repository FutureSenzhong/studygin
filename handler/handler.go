package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func GetReader(c *gin.Context)  {
	response, err := http.Get("https://raw.githubusercontent.com/gin-gonic/logo/master/color.png")
	if err != nil || response.StatusCode != http.StatusOK {
		c.Status(http.StatusServiceUnavailable)
		return
	}

	reader := response.Body
	contentLength := response.ContentLength
	contentType := response.Header.Get("Content-Type")

	extraHeaders := map[string]string{
		"Content-Disposition": `attachment; filename="gopher.png"`,
	}

	c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
}

func Ping(c *gin.Context)  {
	//fmt.Println(c)
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func SomeJson(c *gin.Context) {
	data := map[string]interface{}{
		"lang": "GO语言",
		"tag":  "<br>",
		"tags":  "<br>",
	}

	// 输出 : {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
	c.AsciiJSON(http.StatusOK, data)
}

func Index(c *gin.Context)  {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "这是网站主页",
	})
}

func Json(c *gin.Context)  {
	c.JSON(200, gin.H{
		"html": "<b>Hello, world!</b>",
	})
}

func PureJson(c *gin.Context)  {
	c.PureJSON(200, gin.H{
		"html": "<b>Hello, world!</b>",
	})
}

func Post(c *gin.Context)  {
	//获取url路径传值
	id := c.Query("id")
	page := c.DefaultQuery("page", "0")

	//获取form表单传值
	name := c.PostForm("name")
	message := c.PostForm("message")

	fmt.Printf("url 路径传参：id: %s; page: %s; \n", id, page)
	fmt.Printf("form 表单传参：name: %s; message: %s \n", name, message)
}

func SecureJSON(c *gin.Context)  {
	names := []string{"lena", "austin", "foo"}

	// 将输出：while(1);["lena","austin","foo"]
	c.SecureJSON(http.StatusOK, names)
}

func Upload(c *gin.Context)  {
	// 单文件
	file, _ := c.FormFile("file")
	log.Println(file.Filename)

	// 上传文件至指定目录
	_ = c.SaveUploadedFile(file, "update")

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}


type Person struct {
	Name    string `form:"name"`
	Address string `form:"address"`
}

func StartPage(c *gin.Context) {
	var person Person
	if c.ShouldBindQuery(&person) == nil {
		log.Println("====== Only Bind By Query String ======")
		log.Println(person.Name)
		log.Println(person.Address)
	}
	c.String(200, "Success")
}

func LongAsync(c *gin.Context) {
	// 创建在 goroutine 中使用的副本
	cCp := c.Copy()
	go func() {
		// 用 time.Sleep() 模拟一个长任务。
		time.Sleep(10 * time.Second)

		// 请注意您使用的是复制的上下文 "cCp"，这一点很重要
		log.Println("Done! in path " + cCp.Request.URL.Path)
	}()
}

func LongSync(c *gin.Context) {
	// 用 time.Sleep() 模拟一个长任务。
	time.Sleep(5 * time.Second)

	// 因为没有使用 goroutine，不需要拷贝上下文
	log.Println("Done! in path " + c.Request.URL.Path)
}

