package ml5kus

import (
	"xm0jichu/ml2changliangs"
	"xm0jichu/ml3moxings"
	"xm0jichu/ml0gongjus"
)

func ChuangJianBm3YeWusSuoYin()ml3moxings.CanShu{
  canShu := ml3moxings.CanShu{}
  canShu.ShuJu = []map[string]interface{}{}
  
  biaoMing := map[string]interface{}{
  	ml2changliangs.Ceng1: ml2changliangs.Bm3YeWus,
  }
  canShu.ShuJu = append(canShu.ShuJu, biaoMing)
  
  suoYin := map[string]interface{}{
  	ml2changliangs.Ceng1: ml2changliangs.BianMa,
  }
  canShu.ShuJu = append(canShu.ShuJu, suoYin)
	ret := SheZhiWeiYiSuoYin(canShu)
  return ret
}


func ChuangJianBm3YeWus() ml3moxings.CanShu{
  // 业务表：主键，业务编码，业务名称，描述。//这是针对一个业务进行业务流向梳理,业务是必须有方法流的，但方法流必须也是一对多的，所以这里不需要去指定方法流字段

	canShu := ml3moxings.CanShu{}
	canShu.ShuJu = []map[string]interface{}{}

	biaoMing := map[string]interface{}{
		ml2changliangs.Ceng1: ml2changliangs.Bm3YeWus,
	}
	canShu.ShuJu = append(canShu.ShuJu, biaoMing)

	zhuJian := map[string]interface{}{
		ml2changliangs.Ceng1: ml2changliangs.ZhuJian,
	}
	canShu.ShuJu = append(canShu.ShuJu, zhuJian)

	ziDuans := []map[string]interface{}{}

	ziDuan1 := ml0gongjus.ZuZhuangBIGINT(ml2changliangs.ZhuJian,"20")
	ziDuans = append(ziDuans, ziDuan1)

	ziDuan2 := ml0gongjus.ZuZhuangVARCHAR(ml2changliangs.MingCheng,"50")
	ziDuans = append(ziDuans, ziDuan2)

	ziDuan3 := ml0gongjus.ZuZhuangVARCHAR(ml2changliangs.BianMa,"50")
	ziDuans = append(ziDuans, ziDuan3)

	ziDuan4 := ml0gongjus.ZuZhuangVARCHAR(ml2changliangs.MiaoShu,"150")
	ziDuans = append(ziDuans, ziDuan4)

	ziDuansKeyMap := map[string]interface{}{
		ml2changliangs.Ceng1: ziDuans,
	}

	canShu.ShuJu = append(canShu.ShuJu, ziDuansKeyMap)

	ret := ChuangJianBiao(canShu)
  return ret
}

