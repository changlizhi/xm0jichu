package ml5kus

import (
	"log"
	"xm0jichu/ml2changliangs"
	"xm0jichu/ml3moxings"
	"xm0jichu/ml0gongjus."
)

// 字段指定：主键，字段表主键，可能值
// 字段指定中间表：主键，字段主键，指定主键
// 字段行为：主键，字段表主键，编码，名称，方法流，描述。//这是针对一个字段进行业务流向进行梳理
// 字段行为中间表：主键，字段主键，行为主键
// 业务表：主键，业务编码，业务名称，方法流，描述。//这是针对一个业务进行业务流向进行梳理
// 字段方法流表：主键，字段主键，方法名，顺序，总个数，描述。//方法流跟字段是一一对应的，即使两个字段的业务方法流是一样的也必须拆成两个值
// 业务方法流表：主键，业务主键，方法名，顺序，总个数，描述。//方法流跟字段是一一对应的，即使两个业务的业务方法流是一样的也必须拆成两个值
// 方法表：主键，方法名，必须参数，可选参数，返回数据，描述。//所有方法流里有的方法都能在这里找到。

// 主键表：主键
// 字段值表：主键（同主键表的主键），值

func ChuangJianZiDuanBiao() {
  // 字段表：主键，名称，编码，是否指定，字段值表，是否有行为，校验
	
	canShu := ml3moxings.CanShu{}
	canShu.ShuJu = []map[string]interface{}{}

	biaoMing := map[string]interface{}{
		ml2changliangs.Ceng1: ml2changliangs.ZiDuan,
	}
	canShu.ShuJu = append(canShu.ShuJu, biaoMing)

	zhuJian := map[string]interface{}{
		ml2changliangs.Ceng1: ml2changliangs.ZhuJian,
	}
	canShu.ShuJu = append(canShu.ShuJu, zhuJian)

	ziDuans := []map[string]interface{}{}

	ziDuan1 := ml0gongjus.ZuZhuang20BIGINT(ml2changliangs.ZhuJian)
	ziDuans = append(ziDuans, ziDuan1)

	ziDuan2 := ml0gongjus.ZuZhuang50VARCHAR(ml2changliangs.MingCheng)
	ziDuans = append(ziDuans, ziDuan2)

	ziDuan3 := ml0gongjus.ZuZhuang50VARCHAR(ml2changliangs.BianMa)
	ziDuans = append(ziDuans, ziDuan3)

	ziDuan4 := ml0gongjus.ZuZhuang50VARCHAR(ml2changliangs.ZhuJianBiao)
	ziDuans = append(ziDuans, ziDuan4)

	ziDuansKeyMap := map[string]interface{}{
		ml2changliangs.Ceng1: ziDuans,
	}

	canShu.ShuJu = append(canShu.ShuJu, ziDuansKeyMap)

	ret := ChuangJianBiao(canShu)
	log.Println("TestChuangJianBiao,ret---", ret)
}


func ChuangJianBiaoMingBiao() {
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

	ziDuan1 := ml0gongjus.ZuZhuang20BIGINT(ml2changliangs.ZhuJian)
	ziDuans = append(ziDuans, ziDuan1)

	ziDuan2 := ml0gongjus.ZuZhuang50VARCHAR(ml2changliangs.MingCheng)
	ziDuans = append(ziDuans, ziDuan2)

	ziDuan3 := ml0gongjus.ZuZhuang50VARCHAR(ml2changliangs.BianMa)
	ziDuans = append(ziDuans, ziDuan3)

	ziDuan4 := ml0gongjus.ZuZhuang50VARCHAR(ml2changliangs.ZhuJianBiao)
	ziDuans = append(ziDuans, ziDuan4)

	ziDuansKeyMap := map[string]interface{}{
		ml2changliangs.Ceng1: ziDuans,
	}

	canShu.ShuJu = append(canShu.ShuJu, ziDuansKeyMap)

	ret := ChuangJianBiao(canShu)
	log.Println("TestChuangJianBiao,ret---", ret)
}
