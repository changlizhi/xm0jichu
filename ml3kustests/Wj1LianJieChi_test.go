package ml3kutests

import (
	"testing"
	"xm0yonghu/ml3kus"
	"log"
	"time"
  "database/sql"
)
type countChan struct{
  dbConn *sql.DB
  count int64
}

func TestGaoBingFaHuoQu(t *testing.T){
  count := make(chan int,1)
  for i := 0; i < 100000; i++ {//十万个并发去拿相同的连接池都不会报错，下一步是真正建立链接进行查询，可以校验maxConn是否设置成功
		go func(i int){
      db:=ml3kus.GaoBingFaHuoQu()
      count <- i
      err := db.Ping()
      log.Println("db---",err,i,db)
      <- count
    }(i)
	}
	time.Sleep(time.Duration(5) * time.Hour)
}

