package ml5kus

import (
	"xm0jichu/ml2changliangs"
	"xm0jichu/ml3moxings"
	"xm0jichu/ml0gongjus"
)

func ChuangJianBm1BiaoMings() ml3moxings.CanShu{
	// 表名表：主键，名称，编码，主键表
	canShu := ml3moxings.CanShu{}
	canShu.ShuJu = []map[string]interface{}{}

  shuJu0 := map[string]interface{}{
  }
  shuJu0[ml2changliangs.ShuJuKu]=ml2changliangs.XM0JICHU
  shuJu0[ml2changliangs.BiaoMing]=ml2changliangs.Bm1BiaoMings
  shuJu0[ml2changliangs.SuoYin]=ml2changliangs.BianMa
  shuJu0[ml2changliangs.ZhuJian]=ml2changliangs.ZhuJian

	ziDuans := []map[string]interface{}{}

	ziDuan1 := ml0gongjus.ZuZhuangBIGINT(ml2changliangs.ZhuJian,"20")
  ziDuans = append(ziDuans, ziDuan1)

	ziDuan2 := ml0gongjus.ZuZhuangVARCHAR(ml2changliangs.BianMa,"50")
	ziDuans = append(ziDuans, ziDuan2)

	ziDuan3 := ml0gongjus.ZuZhuangVARCHAR(ml2changliangs.MingCheng,"50")
	ziDuans = append(ziDuans, ziDuan3)

	ziDuan4 := ml0gongjus.ZuZhuangVARCHAR(ml2changliangs.ZhuJianBiao,"50")
	ziDuans = append(ziDuans, ziDuan4)

  ziDuan5 := ml0gongjus.ZuZhuangVARCHAR(ml2changliangs.ShuJuKu,"50")
	ziDuans = append(ziDuans, ziDuan5)
	
  shuJu0[ml2changliangs.ZiDuans]=ziDuans

  canShu.ShuJu = append(canShu.ShuJu, shuJu0)
  ret := ChuangJianBiao(canShu)
  return ret
}
