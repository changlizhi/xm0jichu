package ml5kus

import (
	"xm0jichu/ml2changliangs"
	"xm0jichu/ml3moxings"
	"xm0jichu/ml0gongjus"
)

func ChuangJianBm6ZiDuanZhiDings() ml3moxings.CanShu{
  // 字段指定表：主键，字段主键，指定编码，可能值，可能值描述。//字段指定跟字段是一对多的，一个字段有多个指定值
  canShu := ml3moxings.CanShu{}
  canShu.ShuJu = []map[string]interface{}{}
  
  shuJu0 := map[string]interface{}{
  }
  shuJu0[ml2changliangs.ShuJuKu]=ml2changliangs.XM0JICHU
  shuJu0[ml2changliangs.BiaoMing]=ml2changliangs.Bm6ZiDuanZhiDings
  shuJu0[ml2changliangs.SuoYin]=ml2changliangs.ZhiDingBianMa
  shuJu0[ml2changliangs.ZhuJian]=ml2changliangs.ZhuJian
  
	ziDuans := []map[string]interface{}{}

	ziDuan1 := ml0gongjus.ZuZhuangBIGINT(ml2changliangs.ZhuJian,"20")
	ziDuans = append(ziDuans, ziDuan1)

	ziDuan2 := ml0gongjus.ZuZhuangBIGINT(ml2changliangs.ZiDuanZhuJian,"20")
	ziDuans = append(ziDuans, ziDuan2)

  ziDuan3 := ml0gongjus.ZuZhuangVARCHAR(ml2changliangs.KeNengZhi,"50")
  ziDuans = append(ziDuans, ziDuan3)

  ziDuan4 := ml0gongjus.ZuZhuangVARCHAR(ml2changliangs.KeNengZhiMiaoShu,"50")
  ziDuans = append(ziDuans, ziDuan4)

  ziDuan5 := ml0gongjus.ZuZhuangVARCHAR(ml2changliangs.ZhiDingBianMa,"50")
  ziDuans = append(ziDuans, ziDuan5)

  shuJu0[ml2changliangs.ZiDuans]=ziDuans
  
  canShu.ShuJu = append(canShu.ShuJu, shuJu0)
  ret := ChuangJianBiao(canShu)
  return ret
}
