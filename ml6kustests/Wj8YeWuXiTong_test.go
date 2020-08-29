package ml6kustests

import(
"testing"
"xm0jichu/ml5kus"
"xm0jichu/ml2changliangs"
)
func TestChuangJianYeWuXiTong(t *testing.T){
  ml5kus.ChuangJianJiChuBiao()//基础表必须存在，不需要单独创建
  ml5kus.XinZengYeWuJieGous()
  shuJu0 := ml2changliangs.YeWuBiaoJieGous(ml2changliangs.FhKongZiFu)
  for k,_ := range shuJu0{
    ml5kus.ChuangJianYeWuXiTong(k)
  }
  ml5kus.ShanChuJiChuKu() //测试完成后删除数据库，基础测试在所有的测试方法中都不需要前置测试存在。
  
}
