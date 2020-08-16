package ml6kustests

import (
	"log"
	"testing"
	"xm0jichu/ml2changliangs"
	"xm0jichu/ml3moxings"
	"xm0jichu/ml5kus"
)

// 字段表：主键，名称，编码，是否指定，字段值表，是否有行为，校验
// 字段指定：主键，字段表主键，可能值
// 字段行为：主键，字段表主键，编码，名称，方法流，描述。//这是针对一个字段进行业务流向进行梳理
// 业务表：主键，业务编码，业务名称，方法流，描述。//这是针对一个业务进行业务流向进行梳理
// 方法表：主键，方法名，必须参数，可选参数，返回数据，描述。
// 这几个基础表将会是整个系统的产生地。

func zuZhuang20BIGINT(mingCheng string)map[string]interface{}{
  ret := map[string]interface{}{}
	ret[ml2changliangs.BiaoMing] = mingCheng
	ret[ml2changliangs.LeiXing] = ml2changliangs.BIGINT
	ret[ml2changliangs.ChangDu] = "20"
	ret[ml2changliangs.MoRenZhi] = ml2changliangs.Zf0
  return ret
}
func zuZhuang50VARCHAR(mingCheng string)map[string]interface{}{
  ret := map[string]interface{}{}
  ret[ml2changliangs.BiaoMing] = mingCheng
  ret[ml2changliangs.LeiXing] = ml2changliangs.VARCHAR
  ret[ml2changliangs.ChangDu] = "50"
  ret[ml2changliangs.MoRenZhi] = ml2changliangs.FhDanYinHao + ml2changliangs.HFXXIAOXIE + ml2changliangs.FhDanYinHao
  return ret
}

func TestChuangJianBiaoMingBiao(t *testing.T) {
	// 表名表：主键，名称，编码，主键表

	canShu := ml3moxings.CanShu{}
	canShu.ShuJu = []map[string]interface{}{}

	biaoMing := map[string]interface{}{
		ml2changliangs.Ceng1: ml2changliangs.BiaoMing,
	}
	canShu.ShuJu = append(canShu.ShuJu, biaoMing)

	zhuJian := map[string]interface{}{
		ml2changliangs.Ceng1: ml2changliangs.ZhuJian,
	}
	canShu.ShuJu = append(canShu.ShuJu, zhuJian)

	ziDuans := []map[string]interface{}{}

	ziDuan1 := zuZhuang20BIGINT(ml2changliangs.ZhuJian)
	ziDuans = append(ziDuans, ziDuan1)

	ziDuan2 := zuZhuang50VARCHAR(ml2changliangs.MingCheng)
	ziDuans = append(ziDuans, ziDuan2)

	ziDuan3 := zuZhuang50VARCHAR(ml2changliangs.BianMa)
	ziDuans = append(ziDuans, ziDuan3)

	ziDuan4 := zuZhuang50VARCHAR(ml2changliangs.ZhuJianBiao)
	ziDuans = append(ziDuans, ziDuan4)

	ziDuansKeyMap := map[string]interface{}{
		ml2changliangs.Ceng1: ziDuans,
	}

	canShu.ShuJu = append(canShu.ShuJu, ziDuansKeyMap)

	ret := ml5kus.ChuangJianBiao(canShu)
	log.Println("TestChuangJianBiao,ret---", ret)
}
