package ml5kus

import (
	"xm0jichu/ml2changliangs"
	"xm0jichu/ml3moxings"
	"xm0jichu/ml0gongjus"
)

func ChuangJianBm5ZiDuanXingWeiLius() ml3moxings.CanShu{
  // 字段行为流表：主键，字段主键，顺序，总个数，名称，行为编码（虽然感觉多余，但是为了统一记录加上），方法名（方法名在方法表里是唯一的，所以这里其实是可以作为一个索引的，但是方法流里不同的方法流是可以有多个甚至相同的方法组的，所以不能作为索引），描述。//方法流跟字段是一对多的，查出来之后根据顺序排序即可，即使两个字段的业务方法流是一样的也必须拆成两个值
  canShu := ml3moxings.CanShu{}
  canShu.ShuJu = []map[string]interface{}{}
  shuJu0 := map[string]interface{}{
    ml2changliangs.ShuJuKu:ml2changliangs.XM0JICHU,
    ml2changliangs.BiaoMing:ml2changliangs.Bm5ZiDuanXingWeiLius,
    ml2changliangs.SuoYin:ml2changliangs.XingWeiBianMa,
    ml2changliangs.ZhuJian:ml2changliangs.ZhuJian,
    ml2changliangs.ZiDuans:[]map[string]interface{}{
      ml0gongjus.ZuZhuangBIGINT(ml2changliangs.ZhuJian,"20"),
      ml0gongjus.ZuZhuangBIGINT(ml2changliangs.ZiDuanZhuJian,"20"),
      ml0gongjus.ZuZhuangINT(ml2changliangs.FangFaGeShu,"2"),
      ml0gongjus.ZuZhuangINT(ml2changliangs.ShunXu,"5"),
      ml0gongjus.ZuZhuangVARCHAR(ml2changliangs.MingCheng,"50"),
      ml0gongjus.ZuZhuangVARCHAR(ml2changliangs.FangFaMing,"50"),
      ml0gongjus.ZuZhuangVARCHAR(ml2changliangs.MiaoShu,"50"),
      ml0gongjus.ZuZhuangVARCHAR(ml2changliangs.XingWeiBianMa,"50"),
    },
  }
  canShu.ShuJu = append(canShu.ShuJu, shuJu0)
  ret := ChuangJianBiao(canShu)
  return ret
}
