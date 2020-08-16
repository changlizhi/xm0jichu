package ml4kus

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"sync"
  "xm0jichu/xm3moxings"
)

func init() {
	chuShiHuaChi()
	if err:= lianJieChi.Ping(); err != nil {
		log.Println("初始化数据库失败", err)
		os.Exit(1)
	}
}

var (
	err        error
	suoShiLi   sync.Once
	lianJieChi *sql.DB
)

func chuShiHuaChi() *sql.DB {
	suoShiLi.Do(func() {
		lianJieChi, err = sql.Open("mysql", "root:rootclz@tcp(127.0.0.1:3306)/hfxbase")
		if err != nil {
			log.Println("建立链接失败", err)
			os.Exit(1)
		}
		lianJieChi.SetMaxOpenConns(50)
		lianJieChi.SetMaxIdleConns(5)
	})
	return lianJieChi
}
func HuoQuLianJieChi() *sql.DB {
	return lianJieChi
}
