package ml5kus

import (
	"xm0jichu/ml2changliangs"
	"xm0jichu/ml3moxings"
	"xm0jichu/ml0gongjus"
)

func ChuangJianBm2ZiDuans() ml3moxings.CanShu{
  // 字段表：主键，名称，编码，字段值表，正则，是否指定，是否有行为，长度
  
  canShu := ml3moxings.CanShu{}
  canShu.ShuJu = []map[string]interface{}{}
  
  shuJu0 := map[string]interface{}{
  }
  shuJu0[ml2changliangs.ShuJuKu]=ml2changliangs.XM0JICHU
  shuJu0[ml2changliangs.BiaoMing]=ml2changliangs.Bm2ZiDuans
  shuJu0[ml2changliangs.ZhuJian]=ml2changliangs.ZhuJian
  shuJu0[ml2changliangs.SuoYin]=ml2changliangs.BianMa

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

	ziDuan9 := ml0gongjus.ZuZhuangVARCHAR(ml2changliangs.ShuJuKu,"50")
	ziDuans = append(ziDuans, ziDuan9)
	
  shuJu0[ml2changliangs.ZiDuans]=ziDuans

  canShu.ShuJu = append(canShu.ShuJu, shuJu0)
  ret := ChuangJianBiao(canShu)
  return ret
}

