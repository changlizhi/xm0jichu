package ml5kus

import (
	"xm0jichu/ml2changliangs"
	"xm0jichu/ml3moxings"
	"xm0jichu/ml0gongjus"
)
func ChuangJianBm8ChuCans() ml3moxings.CanShu{
//方法出参字段表：主键，编码，名称，方法名（不用方法主键是为了清晰），类型，是否必须，父编码，描述。
  canShu := ml3moxings.CanShu{}
  canShu.ShuJu = []map[string]interface{}{}
  shuJu0 := map[string]interface{}{
    ml2changliangs.ShuJuKu:ml2changliangs.XM0JICHU,
    ml2changliangs.BiaoMing:ml2changliangs.Bm8ChuCans,
    ml2changliangs.SuoYin:ml2changliangs.BianMa,
    ml2changliangs.ZhuJian:ml2changliangs.ZhuJian,
    ml2changliangs.ZiDuans:[]map[string]interface{}{
      ml0gongjus.ZuZhuangBIGINT(ml2changliangs.ZhuJian,"20"),
      ml0gongjus.ZuZhuangVARCHAR(ml2changliangs.BianMa,"50"),
      ml0gongjus.ZuZhuangVARCHAR(ml2changliangs.MingCheng,"50"),
      ml0gongjus.ZuZhuangVARCHAR(ml2changliangs.FangFaMing,"50"),
      ml0gongjus.ZuZhuangVARCHAR(ml2changliangs.LeiXing,"50"),
      ml0gongjus.ZuZhuangINT(ml2changliangs.BiXu,"1"),
      ml0gongjus.ZuZhuangVARCHAR(ml2changliangs.FuBianMa,"50"),
      ml0gongjus.ZuZhuangVARCHAR(ml2changliangs.MiaoShu,"50"),
    },
  }
  canShu.ShuJu = append(canShu.ShuJu, shuJu0)
  ret := ChuangJianBiao(canShu)
  return ret
}
