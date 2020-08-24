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
    ml2changliangs.ShuJuKu:ml2changliangs.XM0JICHU,
    ml2changliangs.BiaoMing:ml2changliangs.Bm1BiaoMings,
    ml2changliangs.SuoYin:ml2changliangs.BianMa,
    ml2changliangs.ZhuJian:ml2changliangs.ZhuJian,
    ml2changliangs.ZiDuans:[]map[string]interface{}{
      ml0gongjus.ZuZhuangBIGINT(ml2changliangs.ZhuJian,"20"),
      ml0gongjus.ZuZhuangVARCHAR(ml2changliangs.BianMa,"50"),
      ml0gongjus.ZuZhuangVARCHAR(ml2changliangs.MingCheng,"50"),
      ml0gongjus.ZuZhuangVARCHAR(ml2changliangs.ZhuJianBiao,"50"),
      ml0gongjus.ZuZhuangVARCHAR(ml2changliangs.ShuJuKu,"50"),
    },
  }

  canShu.ShuJu = append(canShu.ShuJu, shuJu0)
  ret := ChuangJianBiao(canShu)
  return ret
}
