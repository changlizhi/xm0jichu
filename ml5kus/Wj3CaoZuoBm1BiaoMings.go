package ml5kus

import (
	"xm0jichu/ml2changliangs"
	"xm0jichu/ml3moxings"
	"xm0jichu/ml0gongjus"
)

func ChuangJianBm1BiaoMingsSuoYin()ml3moxings.CanShu{
  canShu := ml3moxings.CanShu{}
  canShu.ShuJu = []map[string]interface{}{}

  shuJuKuMing := map[string]interface{}{
  	ml2changliangs.Ceng1: ml2changliangs.XM0JICHU,
  }
  canShu.ShuJu = append(canShu.ShuJu, shuJuKuMing)
  
  biaoMing := map[string]interface{}{
  	ml2changliangs.Ceng1: ml2changliangs.Bm1BiaoMings,
  }
  canShu.ShuJu = append(canShu.ShuJu, biaoMing)
  
  suoYin := map[string]interface{}{
  	ml2changliangs.Ceng1: ml2changliangs.BianMa,
  }
  canShu.ShuJu = append(canShu.ShuJu, suoYin)
	ret := SheZhiWeiYiSuoYin(canShu)
  return ret
}

func ChuangJianBm1BiaoMings() ml3moxings.CanShu{
	// 表名表：主键，名称，编码，主键表

	canShu := ml3moxings.CanShu{}
	canShu.ShuJu = []map[string]interface{}{}

	shuJuKuMing := map[string]interface{}{
		ml2changliangs.Ceng1: ml2changliangs.XM0JICHU,
	}
	canShu.ShuJu = append(canShu.ShuJu, shuJuKuMing)
	
	biaoMing := map[string]interface{}{
		ml2changliangs.Ceng1: ml2changliangs.Bm1BiaoMings,
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

	ziDuan4 := ml0gongjus.ZuZhuangVARCHAR(ml2changliangs.ZhuJianBiao,"50")
	ziDuans = append(ziDuans, ziDuan4)

	ziDuan5 := ml0gongjus.ZuZhuangVARCHAR(ml2changliangs.ShuJuKu,"50")
	ziDuans = append(ziDuans, ziDuan5)

	ziDuansKeyMap := map[string]interface{}{
		ml2changliangs.Ceng1: ziDuans,
	}

	canShu.ShuJu = append(canShu.ShuJu, ziDuansKeyMap)

	ret := ChuangJianBiao(canShu)
  return ret
}
