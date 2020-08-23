package ml5kus

import (
	"xm0jichu/ml2changliangs"
	"xm0jichu/ml3moxings"
	"xm0jichu/ml0gongjus"
)
func ChuangJianBm4FangFas() ml3moxings.CanShu{
// 方法表：主键，方法名，名称（带上名称是因为中文更容易理解），描述。//所有方法流里有的方法都能在这里找到。
  canShu := ml3moxings.CanShu{}
  canShu.ShuJu = []map[string]interface{}{}
  
  shuJu0 := map[string]interface{}{
  }
  shuJu0[ml2changliangs.ShuJuKu]=ml2changliangs.XM0JICHU
  shuJu0[ml2changliangs.BiaoMing]=ml2changliangs.Bm4FangFas
  shuJu0[ml2changliangs.SuoYin]=ml2changliangs.FangFaMing
  shuJu0[ml2changliangs.ZhuJian]=ml2changliangs.ZhuJian
  
	ziDuans := []map[string]interface{}{}

	ziDuan1 := ml0gongjus.ZuZhuangBIGINT(ml2changliangs.ZhuJian,"20")
	ziDuans = append(ziDuans, ziDuan1)

  ziDuan2 := ml0gongjus.ZuZhuangVARCHAR(ml2changliangs.MingCheng,"50")
  ziDuans = append(ziDuans, ziDuan2)

  ziDuan3 := ml0gongjus.ZuZhuangVARCHAR(ml2changliangs.FangFaMing,"50")
  ziDuans = append(ziDuans, ziDuan3)

  ziDuan4 := ml0gongjus.ZuZhuangVARCHAR(ml2changliangs.MiaoShu,"50")
  ziDuans = append(ziDuans, ziDuan4)

  shuJu0[ml2changliangs.ZiDuans]=ziDuans
  
  canShu.ShuJu = append(canShu.ShuJu, shuJu0)
  ret := ChuangJianBiao(canShu)
  return ret
}

