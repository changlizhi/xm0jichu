package ml5kus

import (
	"xm0jichu/ml2changliangs"
	"xm0jichu/ml3moxings"
	"xm0jichu/ml0gongjus"
)

func ChuangJianBm2ZiDuansSuoYin()ml3moxings.CanShu{
  canShu := ml3moxings.CanShu{}
  canShu.ShuJu = []map[string]interface{}{}
  
  biaoMing := map[string]interface{}{
  	ml2changliangs.Ceng1: ml2changliangs.Bm2ZiDuans,
  }
  canShu.ShuJu = append(canShu.ShuJu, biaoMing)
  
  suoYin := map[string]interface{}{
  	ml2changliangs.Ceng1: ml2changliangs.BianMa,
  }
  canShu.ShuJu = append(canShu.ShuJu, suoYin)
	ret := SheZhiWeiYiSuoYin(canShu)
  return ret
}

func ChuangJianBm2ZiDuans() ml3moxings.CanShu{
  // 字段表：主键，名称，编码，字段值表，正则，是否指定，是否有行为，长度
	
	canShu := ml3moxings.CanShu{}
	canShu.ShuJu = []map[string]interface{}{}

	biaoMing := map[string]interface{}{
		ml2changliangs.Ceng1: ml2changliangs.Bm2ZiDuans,
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

	ziDuan4 := ml0gongjus.ZuZhuangVARCHAR(ml2changliangs.ZhiBiao,"50")
	ziDuans = append(ziDuans, ziDuan4)

	ziDuan5 := ml0gongjus.ZuZhuangVARCHAR(ml2changliangs.ZhengZe,"50")
	ziDuans = append(ziDuans, ziDuan5)

	ziDuan6 := ml0gongjus.ZuZhuangINT(ml2changliangs.ShiFouZhiDing,"1")
	ziDuans = append(ziDuans, ziDuan6)

	ziDuan7 := ml0gongjus.ZuZhuangINT(ml2changliangs.ShiFouYouXingWei,"1")
	ziDuans = append(ziDuans, ziDuan7)

	ziDuan8 := ml0gongjus.ZuZhuangVARCHAR(ml2changliangs.ChangDu,"5")
	ziDuans = append(ziDuans, ziDuan8)

	ziDuansKeyMap := map[string]interface{}{
		ml2changliangs.Ceng1: ziDuans,
	}

	canShu.ShuJu = append(canShu.ShuJu, ziDuansKeyMap)

	ret := ChuangJianBiao(canShu)
  return ret
}

