package ml5kus

import (
	"xm0jichu/ml2changliangs"
	"xm0jichu/ml3moxings"
	"xm0jichu/ml0gongjus"
)
func ChuangJianBm7RuCans() ml3moxings.CanShu{
//入参字段表：主键，编码，名称，方法名（不用方法主键是为了清晰），类型，是否必须，父编码，描述。
  canShu := ml3moxings.CanShu{}
  canShu.ShuJu = []map[string]interface{}{}
  
  shuJu0 := map[string]interface{}{
  }
  shuJu0[ml2changliangs.ShuJuKu]=ml2changliangs.XM0JICHU
  shuJu0[ml2changliangs.BiaoMing]=ml2changliangs.Bm7RuCans
  shuJu0[ml2changliangs.SuoYin]=ml2changliangs.BianMa
  shuJu0[ml2changliangs.ZhuJian]=ml2changliangs.ZhuJian
  
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

  shuJu0[ml2changliangs.ZiDuans]=ziDuans
  
  canShu.ShuJu = append(canShu.ShuJu, shuJu0)
  ret := ChuangJianBiao(canShu)
  return ret
}
