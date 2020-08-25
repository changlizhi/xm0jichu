package ml6kustests

import (
	"testing"
	"xm0jichu/ml2changliangs"
	"xm0jichu/ml3moxings"
	"xm0jichu/ml5kus"
)

// 字段表：主键，名称，编码，是否指定，字段值表，是否有行为，校验
// 字段指定：主键，字段表主键，可能值
// 业务行为：主键，业务表主键，编码，名称，方法流，描述。//这是针对一个业务进行业务流向进行梳理
// 字段行为：主键，字段表主键，编码，名称，方法流，描述。//这是针对一个字段进行业务流向进行梳理
// 业务表：主键，业务编码，业务名称，方法流，描述。//这是针对一个业务进行业务流向进行梳理
// 方法表：主键，方法名，必须参数，可选参数，返回数据，描述。
// 这几个基础表将会是整个系统的产生地。

func TestChuangJianBiaoMingBiao(t *testing.T) {
	canShu := ml3moxings.CanShu{}
	canShu.ShuJu = []map[string]interface{}{}
	shuJu0 := ml2changliangs.JiChuBiao[ml2changliangs.Bm1BiaoMings].(map[string]interface{})

	canShu.ShuJu = append(canShu.ShuJu, shuJu0)
	ml5kus.ChuangJianBiao(canShu)
}

func TestSheZhiWeiYiSuoYin(t *testing.T) {
	canShu := ml3moxings.CanShu{}
	canShu.ShuJu = []map[string]interface{}{}
	shuJu0 := ml2changliangs.JiChuBiao[ml2changliangs.Bm1BiaoMings].(map[string]interface{})

	canShu.ShuJu = append(canShu.ShuJu, shuJu0)
	ml5kus.SheZhiWeiYiSuoYin(canShu)
}
