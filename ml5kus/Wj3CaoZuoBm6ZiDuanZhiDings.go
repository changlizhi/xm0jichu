package ml5kus

import (
	"xm0jichu/ml2changliangs"
	"xm0jichu/ml3moxings"
	"xm0jichu/ml0gongjus"
)
func ChuangJianBm6ZiDuanZhiDingsSuoYin()ml3moxings.CanShu{
  canShu := ml3moxings.CanShu{}
  canShu.ShuJu = []map[string]interface{}{}
  
  biaoMing := map[string]interface{}{
  	ml2changliangs.Ceng1: ml2changliangs.Bm6ZiDuanZhiDings,
  }
  canShu.ShuJu = append(canShu.ShuJu, biaoMing)
  
  suoYin := map[string]interface{}{
  	ml2changliangs.Ceng1: ml2changliangs.ZhiDingBianMa,
  }
  canShu.ShuJu = append(canShu.ShuJu, suoYin)
	ret := SheZhiWeiYiSuoYin(canShu)
  return ret
}

func ChuangJianBm6ZiDuanZhiDings() ml3moxings.CanShu{
  // 字段指定表：主键，字段主键，可能值，可能值描述。//字段指定跟字段是一对多的，一个字段有多个指定值
  
	canShu := ml3moxings.CanShu{}
	canShu.ShuJu = []map[string]interface{}{}

	biaoMing := map[string]interface{}{
		ml2changliangs.Ceng1: ml2changliangs.Bm6ZiDuanZhiDings,
	}
	canShu.ShuJu = append(canShu.ShuJu, biaoMing)

	zhuJian := map[string]interface{}{
		ml2changliangs.Ceng1: ml2changliangs.ZhuJian,
	}
	canShu.ShuJu = append(canShu.ShuJu, zhuJian)

	ziDuans := []map[string]interface{}{}

	ziDuan1 := ml0gongjus.ZuZhuangBIGINT(ml2changliangs.ZhuJian,"20")
	ziDuans = append(ziDuans, ziDuan1)

	ziDuan2 := ml0gongjus.ZuZhuangBIGINT(ml2changliangs.ZiDuanZhuJian,"20")
	ziDuans = append(ziDuans, ziDuan2)

  ziDuan3 := ml0gongjus.ZuZhuangVARCHAR(ml2changliangs.KeNengZhi,"50")
  ziDuans = append(ziDuans, ziDuan3)

  ziDuan4 := ml0gongjus.ZuZhuangVARCHAR(ml2changliangs.KeNengZhiMiaoShu,"50")
  ziDuans = append(ziDuans, ziDuan4)

  ziDuan5 := ml0gongjus.ZuZhuangVARCHAR(ml2changliangs.ZhiDingBianMa,"50")
  ziDuans = append(ziDuans, ziDuan5)

	ziDuansKeyMap := map[string]interface{}{
		ml2changliangs.Ceng1: ziDuans,
	}

	canShu.ShuJu = append(canShu.ShuJu, ziDuansKeyMap)

	ret := ChuangJianBiao(canShu)
  return ret
}
