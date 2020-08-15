package ml3kus

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"sync"
)

func init() {
	huoQuShiLi()
	if danShiLi.lianJieChi == nil {
		os.Exit(1)
	}
}

type GoHouTaiFuWuChi struct {
	lianJieChi *sql.DB
}
var(
  err error
  suoShiLi sync.Once
  danShiLi *GoHouTaiFuWuChi
)

func huoQuShiLi() *GoHouTaiFuWuChi {
	suoShiLi.Do(func() {
		danShiLi = &GoHouTaiFuWuChi{}
		danShiLi.lianJieChi, err = sql.Open("mysql", "root:rootclz@tcp(127.0.0.1:3306)/hfx")
		if err != nil {
			log.Println("建立链接失败",err)
      os.Exit(1)
		}
		danShiLi.lianJieChi.SetMaxOpenConns(50)
		danShiLi.lianJieChi.SetMaxIdleConns(5)
	})
	return danShiLi
}
func GaoBingFaHuoQu() *sql.DB {
  return danShiLi.lianJieChi
}
