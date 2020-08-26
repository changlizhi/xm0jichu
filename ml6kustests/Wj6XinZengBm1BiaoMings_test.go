package ml6kustests

import (
	"testing"
	"xm0jichu/ml2changliangs"
	"xm0jichu/ml5kus"
)

func TestXinZengBm1BiaoMings(t *testing.T) {
	//数据库会初始化所以不需要自己操作，但是表不会，所以还要创建数据表之后再进行新增，新增完之后删除数据库
	shuju0 := ml2changliangs.HuoQuJiChuBiaoJieGou(ml2changliangs.Bm1BiaoMings)
	ml5kus.ChuangJianBiao(shuJu0)
	ml5kus.XinZengBm1BiaoMings()
	ml5kus.ShanChuJiChuKu() //测试完成后删除数据库，基础测试在所有的测试方法中都不需要前置测试存在。
}
