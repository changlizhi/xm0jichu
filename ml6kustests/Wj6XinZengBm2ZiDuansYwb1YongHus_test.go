package ml6kustests

import (
	"testing"
	"xm0jichu/ml0gongjus"
	"xm0jichu/ml2changliangs"
	"xm0jichu/ml3moxings"
	"xm0jichu/ml5kus"
)

func TestXinZengBm2ZiDuans(t *testing.T) {
	//做这个测试之前先创建表
	biaoMingBiaocanShu := ml3moxings.CanShu{
		ShuJu: []map[string]interface{}{
			ml2changliangs.JiChuBiao[ml2changliangs.Bm1BiaoMings].(map[string]interface{}),
		},
	}
	ml5kus.ChuangJianBiao(biaoMingBiaocanShu)
	//创建后先把预设的数据录入，由于基础表的数据表是基本定在常量里的，所以直接就从常量里拿到并生成，基础表基本上不会有什么变化才对
	ml5kus.XinZengBm1BiaoMings()

	//数据库会初始化所以不需要自己操作，但是表不会，所以还要创建数据表之后再进行新增，新增完之后删除数据库
	canShu := ml3moxings.CanShu{
		ShuJu: []map[string]interface{}{
			ml2changliangs.JiChuBiao[ml2changliangs.Bm2ZiDuans].(map[string]interface{}),
		},
	}
	ml5kus.ChuangJianBiao(canShu)
	ml5kus.XinZengBm2ZiDuansYwb1YongHus()
	ml5kus.ShanChuJiChuKu() //测试完成后删除数据库，基础测试在所有的测试方法中都不需要前置测试存在。
}
