package ml5kus

import (
	"xm0jichu/ml2changliangs"
	"xm0jichu/ml3moxings"
	"xm0jichu/ml0gongjus"
)
func ChuangJianBm5ZiDuanXingWeiLiusSuoYin()ml3moxings.CanShu{
  canShu := ml3moxings.CanShu{}
  canShu.ShuJu = []map[string]interface{}{}
  
  biaoMing := map[string]interface{}{
  	ml2changliangs.Ceng1: ml2changliangs.Bm5ZiDuanXingWeiLius,
  }
  canShu.ShuJu = append(canShu.ShuJu, biaoMing)
  
  suoYin := map[string]interface{}{
  	ml2changliangs.Ceng1: ml2changliangs.XingWeiBianMa,
  }
  canShu.ShuJu = append(canShu.ShuJu, suoYin)
	ret := SheZhiWeiYiSuoYin(canShu)
  return ret
}

func ChuangJianBm5ZiDuanXingWeiLius() ml3moxings.CanShu{
  // 字段行为流表：主键，字段主键，顺序，总个数，名称，方法名，描述。//方法流跟字段是一对多的，查出来之后根据顺序排序即可，即使两个字段的业务方法流是一样的也必须拆成两个值
  
	canShu := ml3moxings.CanShu{}
	canShu.ShuJu = []map[string]interface{}{}

	biaoMing := map[string]interface{}{
		ml2changliangs.Ceng1: ml2changliangs.Bm5ZiDuanXingWeiLius,
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

  ziDuan3 := ml0gongjus.ZuZhuangINT(ml2changliangs.FangFaGeShu,"2")
  ziDuans = append(ziDuans, ziDuan3)

  ziDuan4 := ml0gongjus.ZuZhuangINT(ml2changliangs.ShunXu,"5")
  ziDuans = append(ziDuans, ziDuan4)

  ziDuan5 := ml0gongjus.ZuZhuangVARCHAR(ml2changliangs.MingCheng,"50")
  ziDuans = append(ziDuans, ziDuan5)

  ziDuan6 := ml0gongjus.ZuZhuangVARCHAR(ml2changliangs.FangFaMing,"50")
  ziDuans = append(ziDuans, ziDuan6)

  ziDuan7 := ml0gongjus.ZuZhuangVARCHAR(ml2changliangs.MiaoShu,"50")
  ziDuans = append(ziDuans, ziDuan7)

  ziDuan8 := ml0gongjus.ZuZhuangVARCHAR(ml2changliangs.XingWeiBianMa,"50")
  ziDuans = append(ziDuans, ziDuan8)

	ziDuansKeyMap := map[string]interface{}{
		ml2changliangs.Ceng1: ziDuans,
	}

	canShu.ShuJu = append(canShu.ShuJu, ziDuansKeyMap)

	ret := ChuangJianBiao(canShu)
  return ret
}
