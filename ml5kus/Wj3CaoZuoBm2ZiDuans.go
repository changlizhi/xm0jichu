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
    ml2changliangs.ShuJuKu:ml2changliangs.XM0JICHU,
    ml2changliangs.BiaoMing:ml2changliangs.Bm2ZiDuans,
    ml2changliangs.SuoYin:ml2changliangs.BianMa,
    ml2changliangs.ZhuJian:ml2changliangs.ZhuJian,
    ml2changliangs.ZiDuans:[]map[string]interface{}{
      ml0gongjus.ZuZhuangBIGINT(ml2changliangs.ZhuJian,"20"),
      ml0gongjus.ZuZhuangVARCHAR(ml2changliangs.MingCheng,"50"),
      ml0gongjus.ZuZhuangVARCHAR(ml2changliangs.BianMa,"50"),
      ml0gongjus.ZuZhuangVARCHAR(ml2changliangs.ZhiBiao,"50"),
      ml0gongjus.ZuZhuangVARCHAR(ml2changliangs.ZhengZe,"50"),
      ml0gongjus.ZuZhuangINT(ml2changliangs.ShiFouZhiDing,"1"),
      ml0gongjus.ZuZhuangINT(ml2changliangs.ShiFouYouXingWei,"1"),
      ml0gongjus.ZuZhuangVARCHAR(ml2changliangs.ChangDu,"5"),
      ml0gongjus.ZuZhuangVARCHAR(ml2changliangs.ShuJuKu,"50"),
    },
  }
  canShu.ShuJu = append(canShu.ShuJu, shuJu0)
  ret := ChuangJianBiao(canShu)
  return ret
}

