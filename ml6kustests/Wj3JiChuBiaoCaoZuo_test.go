package ml5kutests

import(
"testing"
"xm0jichu/ml5kus"
)

func TestChuangJianJiChuBiaos(t *testing.T){
  ml5kus.ChuangJianJiChuBiao()
  ml5kus.ChuangJianJiChuBiao()//测试第二次使用时会不会再次创建数据库
  ml5kus.ShanChuJiChuKu()//测试完成后删除数据库，基础测试在所有的测试方法中都不需要前置测试存在。
}
