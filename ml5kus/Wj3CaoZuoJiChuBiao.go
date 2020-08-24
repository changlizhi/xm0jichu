package ml5kus

import (
	"xm0jichu/ml2changliangs"
  "strings"
)

func ChuangJianJiChuBiao(){
  ChuangJianBm1BiaoMings()
  ChuangJianBm2ZiDuans()
  ChuangJianBm3YeWus()
  ChuangJianBm4FangFas()
  ChuangJianBm5ZiDuanXingWeiLius()
  ChuangJianBm6ZiDuanZhiDings()
  ChuangJianBm7RuCans()
  ChuangJianBm8ChuCans()
}

func ShanChuSuoYouKu(){
  for _,v :=range strings.Split(ml2changliangs.Kus,ml2changliangs.FhDouHao){
    ShanChuKu(v)
  }
}
