package ml3kutests

import (
	"testing"
	"xm0yonghu/ml3kus"
	"log"
	"time"
)

func TestPing(t *testing.T) {
	for i := 0; i < 20; i++ {
		go func(i int){
      db := ml3kus.HuoQuLianJieChi()//检测并发时数据库实例是否是一个
      err := db.Ping()
      log.Println("TestPing---",i,db,err)
    }(i)
	}
	time.Sleep(time.Duration(5) * time.Second)
}
