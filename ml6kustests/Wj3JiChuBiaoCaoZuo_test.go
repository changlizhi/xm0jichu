package ml5kutests

import(
"testing"
"log"
"xm0jichu/ml5kus"
)

func TestChuangJianZiDuanBiao(t *testing.T){
  ret := ml5kus.ChuangJianZiDuanBiao()
  log.Println("TestChuangJianZiDuanBiao,ret---",ret)
}
func TestChuangJianBiaoMingBiao(t *testing.T){
  ret := ml5kus.ChuangJianBiaoMingBiao()
  log.Println("TestChuangJianBiaoMingBiao,ret---",ret)
}

