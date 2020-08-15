package ml3kus

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"sync"
)

func init() {
	log.Println("初始化连接池开始")
	danShiLi := huoQuShiLi()
	if danShiLi.lianJieChi == nil {
		log.Println("初始化连接池失败...")
		os.Exit(1)
	} else {
		log.Println("初始化连接池成功...", danShiLi.lianJieChi)
	}
}

func HuoQuLianJieChi() *sql.DB {
	ljc := huoQuShiLi()
	return ljc.lianJieChi
}

type GoHouTaiFuWuChi struct {
	lianJieChi *sql.DB
}

var danShiLi *GoHouTaiFuWuChi
var suoShiLi sync.Once
var err error

func huoQuShiLi() *GoHouTaiFuWuChi {
	suoShiLi.Do(func() {
		danShiLi = &GoHouTaiFuWuChi{}
		danShiLi.lianJieChi, err = sql.Open("mysql", "root:rootclz@tcp(127.0.0.1:3306)/hfx")
		if err != nil {
			log.Fatal(err)
		}
		danShiLi.lianJieChi.SetMaxOpenConns(10)
		danShiLi.lianJieChi.SetMaxIdleConns(5)
	})
	return danShiLi
}
