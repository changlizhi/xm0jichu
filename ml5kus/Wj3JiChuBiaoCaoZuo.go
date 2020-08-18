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

	biaoMing := map[string]interface{}{
		ml2changliangs.Ceng1: ml2changliangs.Bm7RuCans,
	}
	canShu.ShuJu = append(canShu.ShuJu, biaoMing)

	zhuJian := map[string]interface{}{
		ml2changliangs.Ceng1: ml2changliangs.ZhuJian,
	}
	canShu.ShuJu = append(canShu.ShuJu, zhuJian)

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

	ziDuansKeyMap := map[string]interface{}{
		ml2changliangs.Ceng1: ziDuans,
	}

	canShu.ShuJu = append(canShu.ShuJu, ziDuansKeyMap)

	ret := ChuangJianBiao(canShu)
  return ret
}
func ChuangJianBm8ChuCans() ml3moxings.CanShu{
//方法出参字段表：主键，编码，名称，方法名（不用方法主键是为了清晰），类型，是否必须，父编码，描述。
  
	canShu := ml3moxings.CanShu{}
	canShu.ShuJu = []map[string]interface{}{}

	biaoMing := map[string]interface{}{
		ml2changliangs.Ceng1: ml2changliangs.Bm8ChuCans,
	}
	canShu.ShuJu = append(canShu.ShuJu, biaoMing)

	zhuJian := map[string]interface{}{
		ml2changliangs.Ceng1: ml2changliangs.ZhuJian,
	}
	canShu.ShuJu = append(canShu.ShuJu, zhuJian)

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

	ziDuansKeyMap := map[string]interface{}{
		ml2changliangs.Ceng1: ziDuans,
	}

	canShu.ShuJu = append(canShu.ShuJu, ziDuansKeyMap)

	ret := ChuangJianBiao(canShu)
  return ret
}

func ChuangJianBm4FangFas() ml3moxings.CanShu{
// 方法表：主键，方法名，名称（带上名称是因为中文更容易理解），描述。//所有方法流里有的方法都能在这里找到。
  
	canShu := ml3moxings.CanShu{}
	canShu.ShuJu = []map[string]interface{}{}

	biaoMing := map[string]interface{}{
		ml2changliangs.Ceng1: ml2changliangs.Bm4FangFas,
	}
	canShu.ShuJu = append(canShu.ShuJu, biaoMing)

	zhuJian := map[string]interface{}{
		ml2changliangs.Ceng1: ml2changliangs.ZhuJian,
	}
	canShu.ShuJu = append(canShu.ShuJu, zhuJian)

	ziDuans := []map[string]interface{}{}

	ziDuan1 := ml0gongjus.ZuZhuangBIGINT(ml2changliangs.ZhuJian,"20")
	ziDuans = append(ziDuans, ziDuan1)

  ziDuan2 := ml0gongjus.ZuZhuangVARCHAR(ml2changliangs.MingCheng,"50")
  ziDuans = append(ziDuans, ziDuan2)

  ziDuan3 := ml0gongjus.ZuZhuangVARCHAR(ml2changliangs.FangFaMing,"50")
  ziDuans = append(ziDuans, ziDuan3)

  ziDuan4 := ml0gongjus.ZuZhuangVARCHAR(ml2changliangs.MiaoShu,"50")
  ziDuans = append(ziDuans, ziDuan4)

	ziDuansKeyMap := map[string]interface{}{
		ml2changliangs.Ceng1: ziDuans,
	}

	canShu.ShuJu = append(canShu.ShuJu, ziDuansKeyMap)

	ret := ChuangJianBiao(canShu)
  return ret
}


func ChuangJianBm3YeWus() ml3moxings.CanShu{
  // 业务表：主键，业务编码，业务名称，描述。//这是针对一个业务进行业务流向梳理,业务是必须有方法流的，但方法流必须也是一对多的，所以这里不需要去指定方法流字段

	canShu := ml3moxings.CanShu{}
	canShu.ShuJu = []map[string]interface{}{}

	biaoMing := map[string]interface{}{
		ml2changliangs.Ceng1: ml2changliangs.Bm3YeWus,
	}
	canShu.ShuJu = append(canShu.ShuJu, biaoMing)

	zhuJian := map[string]interface{}{
		ml2changliangs.Ceng1: ml2changliangs.ZhuJian,
	}
	canShu.ShuJu = append(canShu.ShuJu, zhuJian)

	ziDuans := []map[string]interface{}{}

	ziDuan1 := ml0gongjus.ZuZhuangBIGINT(ml2changliangs.ZhuJian,"20")
	ziDuans = append(ziDuans, ziDuan1)

	ziDuan2 := ml0gongjus.ZuZhuangVARCHAR(ml2changliangs.MingCheng,"50")
	ziDuans = append(ziDuans, ziDuan2)

	ziDuan3 := ml0gongjus.ZuZhuangVARCHAR(ml2changliangs.BianMa,"50")
	ziDuans = append(ziDuans, ziDuan3)

	ziDuan4 := ml0gongjus.ZuZhuangVARCHAR(ml2changliangs.MiaoShu,"150")
	ziDuans = append(ziDuans, ziDuan4)

	ziDuansKeyMap := map[string]interface{}{
		ml2changliangs.Ceng1: ziDuans,
	}

	canShu.ShuJu = append(canShu.ShuJu, ziDuansKeyMap)

	ret := ChuangJianBiao(canShu)
  return ret
}

func ChuangJianBm5ZiDuanXingWeiLius() ml3moxings.CanShu{
  // 字段行为流表：主键，字段主键，顺序，总个数，名称，方法名，描述。//方法流跟字段是一对多的，查出来之后根据顺序排序即可，即使两个字段的业务方法流是一样的也必须拆成两个值
  
	canShu := ml3moxings.CanShu{}
	canShu.ShuJu = []map[string]interface{}{}

	biaoMing := map[string]interface{}{
		ml2changliangs.Ceng1: ml2changliangs.Bm5ZiDuanXingWeiLius,
	}
	canShu.ShuJu = append(canShu.ShuJu, biaoMing)

	zhuJian := map[string]interface{}{
		ml2changliangs.Ceng1: ml2changliangs.ZhuJian,
	}
	canShu.ShuJu = append(canShu.ShuJu, zhuJian)

	ziDuans := []map[string]interface{}{}

	ziDuan1 := ml0gongjus.ZuZhuangBIGINT(ml2changliangs.ZhuJian,"20")
	ziDuans = append(ziDuans, ziDuan1)

	ziDuan2 := ml0gongjus.ZuZhuangBIGINT(ml2changliangs.ZiDuanZhuJian,"20")
	ziDuans = append(ziDuans, ziDuan2)

  ziDuan3 := ml0gongjus.ZuZhuangINT(ml2changliangs.FangFaGeShu,"2")
  ziDuans = append(ziDuans, ziDuan3)

  ziDuan4 := ml0gongjus.ZuZhuangINT(ml2changliangs.ShunXu,"5")
  ziDuans = append(ziDuans, ziDuan4)

  ziDuan5 := ml0gongjus.ZuZhuangVARCHAR(ml2changliangs.MingCheng,"50")
  ziDuans = append(ziDuans, ziDuan5)

  ziDuan6 := ml0gongjus.ZuZhuangVARCHAR(ml2changliangs.FangFaMing,"50")
  ziDuans = append(ziDuans, ziDuan6)

  ziDuan7 := ml0gongjus.ZuZhuangVARCHAR(ml2changliangs.MiaoShu,"50")
  ziDuans = append(ziDuans, ziDuan7)

	ziDuansKeyMap := map[string]interface{}{
		ml2changliangs.Ceng1: ziDuans,
	}

	canShu.ShuJu = append(canShu.ShuJu, ziDuansKeyMap)

	ret := ChuangJianBiao(canShu)
  return ret
}


func ChuangJianBm6ZiDuanZhiDings() ml3moxings.CanShu{
  // 字段指定表：主键，字段主键，可能值，可能值描述。//字段指定跟字段是一对多的，一个字段有多个指定值
  
	canShu := ml3moxings.CanShu{}
	canShu.ShuJu = []map[string]interface{}{}

	biaoMing := map[string]interface{}{
		ml2changliangs.Ceng1: ml2changliangs.Bm6ZiDuanZhiDings,
	}
	canShu.ShuJu = append(canShu.ShuJu, biaoMing)

	zhuJian := map[string]interface{}{
		ml2changliangs.Ceng1: ml2changliangs.ZhuJian,
	}
	canShu.ShuJu = append(canShu.ShuJu, zhuJian)

	ziDuans := []map[string]interface{}{}

	ziDuan1 := ml0gongjus.ZuZhuangBIGINT(ml2changliangs.ZhuJian,"20")
	ziDuans = append(ziDuans, ziDuan1)

	ziDuan2 := ml0gongjus.ZuZhuangBIGINT(ml2changliangs.ZiDuanZhuJian,"20")
	ziDuans = append(ziDuans, ziDuan2)

  ziDuan3 := ml0gongjus.ZuZhuangVARCHAR(ml2changliangs.KeNengZhi,"50")
  ziDuans = append(ziDuans, ziDuan3)

  ziDuan4 := ml0gongjus.ZuZhuangVARCHAR(ml2changliangs.KeNengZhiMiaoShu,"50")
  ziDuans = append(ziDuans, ziDuan4)

	ziDuansKeyMap := map[string]interface{}{
		ml2changliangs.Ceng1: ziDuans,
	}

	canShu.ShuJu = append(canShu.ShuJu, ziDuansKeyMap)

	ret := ChuangJianBiao(canShu)
  return ret
}

func ChuangJianBm2ZiDuans() ml3moxings.CanShu{
  // 字段表：主键，名称，编码，字段值表，正则，是否指定，是否有行为
	
	canShu := ml3moxings.CanShu{}
	canShu.ShuJu = []map[string]interface{}{}

	biaoMing := map[string]interface{}{
		ml2changliangs.Ceng1: ml2changliangs.Bm2ZiDuans,
	}
	canShu.ShuJu = append(canShu.ShuJu, biaoMing)

	zhuJian := map[string]interface{}{
		ml2changliangs.Ceng1: ml2changliangs.ZhuJian,
	}
	canShu.ShuJu = append(canShu.ShuJu, zhuJian)

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

	ziDuansKeyMap := map[string]interface{}{
		ml2changliangs.Ceng1: ziDuans,
	}

	canShu.ShuJu = append(canShu.ShuJu, ziDuansKeyMap)

	ret := ChuangJianBiao(canShu)
  return ret
}


func ChuangJianBm1BiaoMings() ml3moxings.CanShu{
	// 表名表：主键，名称，编码，主键表

	canShu := ml3moxings.CanShu{}
	canShu.ShuJu = []map[string]interface{}{}

	biaoMing := map[string]interface{}{
		ml2changliangs.Ceng1: ml2changliangs.Bm1BiaoMings,
	}
	canShu.ShuJu = append(canShu.ShuJu, biaoMing)

	zhuJian := map[string]interface{}{
		ml2changliangs.Ceng1: ml2changliangs.ZhuJian,
	}
	canShu.ShuJu = append(canShu.ShuJu, zhuJian)

	ziDuans := []map[string]interface{}{}

	ziDuan1 := ml0gongjus.ZuZhuangBIGINT(ml2changliangs.ZhuJian,"20")
	ziDuans = append(ziDuans, ziDuan1)

	ziDuan2 := ml0gongjus.ZuZhuangVARCHAR(ml2changliangs.MingCheng,"50")
	ziDuans = append(ziDuans, ziDuan2)

	ziDuan3 := ml0gongjus.ZuZhuangVARCHAR(ml2changliangs.BianMa,"50")
	ziDuans = append(ziDuans, ziDuan3)

	ziDuan4 := ml0gongjus.ZuZhuangVARCHAR(ml2changliangs.ZhuJianBiao,"50")
	ziDuans = append(ziDuans, ziDuan4)

	ziDuansKeyMap := map[string]interface{}{
		ml2changliangs.Ceng1: ziDuans,
	}

	canShu.ShuJu = append(canShu.ShuJu, ziDuansKeyMap)

	ret := ChuangJianBiao(canShu)
  return ret
}
