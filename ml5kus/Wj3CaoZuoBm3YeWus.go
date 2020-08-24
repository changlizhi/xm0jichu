package ml5kus

import (
	"xm0jichu/ml2changliangs"
	"xm0jichu/ml3moxings"
	"xm0jichu/ml0gongjus"
)

func ChuangJianBm3YeWus() ml3moxings.CanShu{
  // 业务表：主键，业务编码，业务名称，描述。//这是针对一个业务进行业务流向梳理,业务是必须有方法流的，但方法流必须也是一对多的，所以这里不需要去指定方法流字段
  canShu := ml3moxings.CanShu{}
  canShu.ShuJu = []map[string]interface{}{}
  shuJu0 := map[string]interface{}{
    ml2changliangs.ShuJuKu:ml2changliangs.XM0JICHU,
    ml2changliangs.BiaoMing:ml2changliangs.Bm3YeWus,
    ml2changliangs.SuoYin:ml2changliangs.BianMa,
    ml2changliangs.ZhuJian:ml2changliangs.ZhuJian,
    ml2changliangs.ZiDuans:[]map[string]interface{}{
      ml0gongjus.ZuZhuangBIGINT(ml2changliangs.ZhuJian,"20"),
      ml0gongjus.ZuZhuangVARCHAR(ml2changliangs.MingCheng,"50"),
      ml0gongjus.ZuZhuangVARCHAR(ml2changliangs.BianMa,"50"),
      ml0gongjus.ZuZhuangVARCHAR(ml2changliangs.MiaoShu,"150"),
    },
  }
  canShu.ShuJu = append(canShu.ShuJu, shuJu0)
  ret := ChuangJianBiao(canShu)
  return ret
}

