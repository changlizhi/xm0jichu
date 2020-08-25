package ml6kustests

import (
	"database/sql"
	"log"
	"testing"
	"time"
	"xm0jichu/ml2changliangs"
	"xm0jichu/ml3moxings"
	"xm0jichu/ml5kus"
)

func TestHuoQuLianJieChi(t *testing.T) {
	count := make(chan int, 1)
	canShu := ml5kus.HuoQuLianJieChi(ml2changliangs.XM0JICHU)
	log.Println("TestHuoQuLianJieChi,canShu:", canShu)
	db := ml3moxings.HuoQuCeng1YiGe(canShu).(*sql.DB)
	log.Println("TestHuoQuLianJieChi,db:", db)
	for i := 0; i < 10; i++ { //十万个并发去拿相同的连接池都不会报错，下一步是真正建立链接进行查询，可以校验maxConn是否设置成功
		go func(i int) {
			count <- i
			err := db.Ping()
			log.Println("db---", err, i, db)
			<-count
		}(i)
	}
	time.Sleep(time.Duration(5) * time.Hour)
}
