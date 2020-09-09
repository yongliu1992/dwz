package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yongliu1992/mmh3"
	"gopkg.in/ini.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"strings"
)

type Wz struct {
	ID     uint `gorm:"primaryKey"`
	Sub    string
	Origin string
}

var dsn = "?charset=utf8mb4&parseTime=True&loc=Local"
var db *gorm.DB
var err error
var hostName string
func main() {
	cfg, err := ini.Load("app.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	dbs := cfg.Section("db")
	dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",dbs.Key("db_user").String(),dbs.Key("db_pwd").String(),dbs.Key("db_host").String(),dbs.Key("db_port").String(),dbs.Key("db_name").String())+dsn
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	r := gin.Default()
	r.GET("/:name", handlerRedirect)
	hostName = cfg.Section("server").Key("host").String()
	log.Fatal(r.Run(":"+cfg.Section("server").Key("port").String()))
}
func handlerNew(c *gin.Context) {
	ori := c.Query("url")
	if ori == "" {
		c.JSON(400, map[string]interface{}{"error": "缺少参数url"})
		return
	}
	m32 := mmh3.Sum32([]byte(ori))
	sub := from10To64(m32)
	x := Wz{Sub: sub, Origin: ori}
	err = db.Create(&x).Error
	if err != nil {
		println("err", err)
	}
	cs := fmt.Sprintf("<a href='%s'>点我</a>",hostName+"/" + sub)
	if c.Query("format") == "button" {
		c.Writer.WriteString(cs)
		return
	}
	c.JSON(200,map[string]interface{}{"dwz":hostName+"/" + sub})
}
func handlerRedirect(c *gin.Context) {
	wz := c.Param("name")
	if wz == "n" {
		handlerNew(c)
		return
	}

	x := Wz{Sub: wz}
	result := db.Where("sub=?", x.Sub).Find(&x)
	if result.RowsAffected < 1 {
		c.JSON(400, "not found")
	} else {
		if !strings.Contains(x.Origin, "http") {
			x.Origin = "http://" + x.Origin
		}
		c.Redirect(301, x.Origin)
	}
}

const dict = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_-"

func from10To64(num uint32) string {
	var str64 []byte
	for {
		var result byte
		var tmp []byte

		number := num % 64    // 100%62 = 38
		result = dict[number] // C
		// 临时变量，为了追加到头部
		tmp = append(tmp, result)
		str64 = append(tmp, str64...)
		num = num / 64
		if num == 0 {
			break
		}
	}
	return string(str64)
}
