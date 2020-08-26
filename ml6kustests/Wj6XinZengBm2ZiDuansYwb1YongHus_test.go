package ml6kustests

import (
	"testing"
	"xm0jichu/ml2changliangs"
	"xm0jichu/ml5kus"
)

//思考这个模块的创建方式，一方面是基础表已经完全在changliangs中表现了出来，但是业务ziduans没有展现，至少用户1这个没有，也就是说这些数据都在测试环境中，具体应该从哪里来呢？
//而像用户字段说明这种数据很明显是根据业务需求来的，我发现把数据存放在代码里的风险，比如yonghu库需要新加一个表的话得重新修改新增基础表的这个文件，这是极其不合适的，而放到数据库中又是很不希望出现的状态
//不过从现在的状态来说yonghu库基本上不会有什么变化，但为了能够只在一处修改就能完成所有表的创建，那么应该如何操作呢？

func TestXinZengBm2ZiDuans(t *testing.T) {
	//做这个测试之前先创建表
	shuJu0 := ml2changliangs.HuoQuJiChuBiaoJieGou(ml2changliangs.Bm1BiaoMings)
  ml5kus.ChuangJianBiao(shuJu0)
	//创建后先把预设的数据录入，由于基础表的数据表是基本定在常量里的，所以直接就从常量里拿到并生成，基础表基本上不会有什么变化才对
	ml5kus.XinZengBm1BiaoMings()

	//数据库会初始化所以不需要自己操作，但是表不会，所以还要创建数据表之后再进行新增，新增完之后删除数据库
	shuJu1 := ml2changliangs.HuoQuJiChuBiaoJieGou(ml2changliangs.Bm2ZiDuans)
  ml5kus.ChuangJianBiao(shuJu1)
	ml5kus.XinZengBm2ZiDuansYwb1YongHus()
	ml5kus.ShanChuJiChuKu() //测试完成后删除数据库，基础测试在所有的测试方法中都不需要前置测试存在。
}
