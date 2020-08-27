package ml6kustests

import (
	"testing"
	"xm0jichu/ml5kus"
)
//在这里把表和字段数据都添加了
func TestXinZengBm1BiaoMings(t *testing.T) {
	//数据库会初始化所以不需要自己操作，但是表不会，所以还要创建数据表之后再进行新增，新增完之后删除数据库
	ml5kus.ChuangJianJiChuBiao()//基础表必须存在，不需要单独创建
	ml5kus.XinZengYeWuJieGous()
	// ml5kus.ShanChuJiChuKu() //测试完成后删除数据库，基础测试在所有的测试方法中都不需要前置测试存在。
}
