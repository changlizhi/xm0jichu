package ml6kustests

import(
"testing"
"xm0jichu/ml5kus"
"xm0jichu/ml2changliangs"
)
func TestChuangJianYeWuXiTong(t *testing.T){
  shuJu0 := ml2changliangs.YeWuBiaoJieGous(ml2changliangs.FhKongZiFu)
  for k,_ := range shuJu0{
    ml5kus.ChuangJianYeWuXiTong(k)
  }
  
}
