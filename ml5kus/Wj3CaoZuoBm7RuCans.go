package ml5kus

import (
	"xm0jichu/ml2changliangs"
	"xm0jichu/ml3moxings"
	"xm0jichu/ml0gongjus"
)
func ChuangJianBm7RuCansSuoYin()ml3moxings.CanShu{
  canShu := ml3moxings.CanShu{}
  canShu.ShuJu = []map[string]interface{}{}
  
  biaoMing := map[string]interface{}{
  	ml2changliangs.Ceng1: ml2changliangs.Bm7RuCans,
  }
  canShu.ShuJu = append(canShu.ShuJu, biaoMing)
  
  suoYin := map[string]interface{}{
  	ml2changliangs.Ceng1: ml2changliangs.BianMa,
  }
  canShu.ShuJu = append(canShu.ShuJu, suoYin)
	ret := SheZhiWeiYiSuoYin(canShu)
  return ret
}

func ChuangJianBm7RuCans() ml3moxings.CanShu{
//入参字段表：主键，编码，名称，方法名（不用方法主键是为了清晰），类型，是否必须，父编码，描述。
  
	canShu := ml3moxings.CanShu{}
	canShu.ShuJu = []map[string]interface{}{}

	biaoMing := map[string]interface{}{
		ml2changliangs.Ceng1: ml2changliangs.Bm7RuCans,
	}
	canShu.ShuJu = append(canShu.ShuJu, biaoMing)

	zhuJian := map[string]interface{}{
		ml2changliangs.Ceng1: ml2changliangs.ZhuJian,
	}
	canShu.ShuJu = append(canShu.ShuJu, zhuJian)

	ziDuans := []map[string]interface{}{}

	ziDuan1 := ml0gongjus.ZuZhuangBIGINT(ml2changliangs.ZhuJian,"20")
	ziDuans = append(ziDuans, ziDuan1)

  ziDuan2 := ml0gongjus.ZuZhuangVARCHAR(ml2changliangs.BianMa,"50")
  ziDuans = append(ziDuans, ziDuan2)

  ziDuan3 := ml0gongjus.ZuZhuangVARCHAR(ml2changliangs.MingCheng,"50")
  ziDuans = append(ziDuans, ziDuan3)

  ziDuan4 := ml0gongjus.ZuZhuangVARCHAR(ml2changliangs.FangFaMing,"50")
  ziDuans = append(ziDuans, ziDuan4)

  ziDuan5 := ml0gongjus.ZuZhuangVARCHAR(ml2changliangs.LeiXing,"50")
  ziDuans = append(ziDuans, ziDuan5)

  ziDuan6 := ml0gongjus.ZuZhuangINT(ml2changliangs.BiXu,"1")
  ziDuans = append(ziDuans, ziDuan6)

  ziDuan7 := ml0gongjus.ZuZhuangVARCHAR(ml2changliangs.FuBianMa,"50")
  ziDuans = append(ziDuans, ziDuan7)

  ziDuan8 := ml0gongjus.ZuZhuangVARCHAR(ml2changliangs.MiaoShu,"50")
  ziDuans = append(ziDuans, ziDuan8)

	ziDuansKeyMap := map[string]interface{}{
		ml2changliangs.Ceng1: ziDuans,
	}

	canShu.ShuJu = append(canShu.ShuJu, ziDuansKeyMap)

	ret := ChuangJianBiao(canShu)
  return ret
}
