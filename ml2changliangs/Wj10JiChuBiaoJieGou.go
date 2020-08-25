package ml2changliangs

import (
	"sync"
)

func init() {
	chuShiHuaJiChuBiaoJieGou()
}

var (
	SuoShiLi  sync.Once
	JiChuBiao = map[string]interface{}{}
)

func tianJiaBm9YeWuXingWeiLius() {
	// 字段行为流表：主键，字段主键，顺序，总个数，名称，行为编码（虽然感觉多余，但是为了统一记录加上），方法名（方法名在方法表里是唯一的，所以这里其实是可以作为一个索引的，但是方法流里不同的方法流是可以有多个甚至相同的方法组的，所以不能作为索引），描述。//方法流跟字段是一对多的，查出来之后根据顺序排序即可，即使两个字段的业务方法流是一样的也必须拆成两个值
	shuJu0 := map[string]interface{}{
		CaoZuoKu:   XM0JICHU,
		CaoZuoBiao: Bm9YeWuXingWeiLius,
		SuoYin:     XingWeiBianMa,
		ZhuJian:    ZhuJian,
		ZiDuans: []map[string]interface{}{
			ZuZhuangBIGINT(ZhuJian, "20"),
			ZuZhuangBIGINT(YeWuZhuJian, "20"),
			ZuZhuangINT(FangFaGeShu, "2"),
			ZuZhuangINT(ShunXu, "5"),
			ZuZhuangVARCHAR(MingCheng, "50"),
			ZuZhuangVARCHAR(FangFaMing, "50"),
			ZuZhuangVARCHAR(MiaoShu, "50"),
			ZuZhuangVARCHAR(XingWeiBianMa, "50"),
		},
	}
	JiChuBiao[Bm9YeWuXingWeiLius] = shuJu0
}

func tianJiaBm8ChuCans() {
	//方法出参字段表：主键，编码，名称，方法名（不用方法主键是为了清晰），类型，是否必须，父编码，描述。
	shuJu0 := map[string]interface{}{
		CaoZuoKu:   XM0JICHU,
		CaoZuoBiao: Bm8ChuCans,
		SuoYin:     BianMa,
		ZhuJian:    ZhuJian,
		ZiDuans: []map[string]interface{}{
			ZuZhuangBIGINT(ZhuJian, "20"),
			ZuZhuangVARCHAR(BianMa, "50"),
			ZuZhuangVARCHAR(MingCheng, "50"),
			ZuZhuangVARCHAR(FangFaMing, "50"),
			ZuZhuangVARCHAR(LeiXing, "50"),
			ZuZhuangINT(BiXu, "1"),
			ZuZhuangVARCHAR(FuBianMa, "50"),
			ZuZhuangVARCHAR(MiaoShu, "50"),
		},
	}
	JiChuBiao[Bm8ChuCans] = shuJu0
}

func tianJiaBm7RuCans() {
	//入参字段表：主键，编码，名称，方法名（不用方法主键是为了清晰），类型，是否必须，父编码，描述。
	shuJu0 := map[string]interface{}{
		CaoZuoKu:   XM0JICHU,
		CaoZuoBiao: Bm7RuCans,
		SuoYin:     BianMa,
		ZhuJian:    ZhuJian,
		ZiDuans: []map[string]interface{}{
			ZuZhuangBIGINT(ZhuJian, "20"),
			ZuZhuangVARCHAR(BianMa, "50"),
			ZuZhuangVARCHAR(MingCheng, "50"),
			ZuZhuangVARCHAR(FangFaMing, "50"),
			ZuZhuangVARCHAR(LeiXing, "50"),
			ZuZhuangINT(BiXu, "1"),
			ZuZhuangVARCHAR(FuBianMa, "50"),
			ZuZhuangVARCHAR(MiaoShu, "50"),
		},
	}
	JiChuBiao[Bm7RuCans] = shuJu0
}

func tianJiaBm6ZiDuanZhiDings() {
	// 字段指定表：主键，字段主键，指定编码，可能值，可能值描述。//字段指定跟字段是一对多的，一个字段有多个指定值
	shuJu0 := map[string]interface{}{
		CaoZuoKu:   XM0JICHU,
		CaoZuoBiao: Bm6ZiDuanZhiDings,
		SuoYin:     ZhiDingBianMa,
		ZhuJian:    ZhuJian,
		ZiDuans: []map[string]interface{}{
			ZuZhuangBIGINT(ZhuJian, "20"),
			ZuZhuangBIGINT(ZiDuanZhuJian, "20"),
			ZuZhuangVARCHAR(KeNengZhi, "50"),
			ZuZhuangVARCHAR(KeNengZhiMiaoShu, "50"),
			ZuZhuangVARCHAR(ZhiDingBianMa, "50"),
		},
	}
	JiChuBiao[Bm6ZiDuanZhiDings] = shuJu0
}

func tianJiaBm5ZiDuanXingWeiLius() {
	// 字段行为流表：主键，字段主键，顺序，总个数，名称，行为编码（虽然感觉多余，但是为了统一记录加上），方法名（方法名在方法表里是唯一的，所以这里其实是可以作为一个索引的，但是方法流里不同的方法流是可以有多个甚至相同的方法组的，所以不能作为索引），描述。//方法流跟字段是一对多的，查出来之后根据顺序排序即可，即使两个字段的业务方法流是一样的也必须拆成两个值
	shuJu0 := map[string]interface{}{
		CaoZuoKu:   XM0JICHU,
		CaoZuoBiao: Bm5ZiDuanXingWeiLius,
		SuoYin:     XingWeiBianMa,
		ZhuJian:    ZhuJian,
		ZiDuans: []map[string]interface{}{
			ZuZhuangBIGINT(ZhuJian, "20"),
			ZuZhuangBIGINT(ZiDuanZhuJian, "20"),
			ZuZhuangINT(FangFaGeShu, "2"),
			ZuZhuangINT(ShunXu, "5"),
			ZuZhuangVARCHAR(MingCheng, "50"),
			ZuZhuangVARCHAR(FangFaMing, "50"),
			ZuZhuangVARCHAR(MiaoShu, "50"),
			ZuZhuangVARCHAR(XingWeiBianMa, "50"),
		},
	}
	JiChuBiao[Bm5ZiDuanXingWeiLius] = shuJu0
}

func tianJiaBm4FangFas() {
	// 方法表：主键，方法名，名称（带上名称是因为中文更容易理解），描述。//所有方法流里有的方法都能在这里找到。
	shuJu0 := map[string]interface{}{
		CaoZuoKu:   XM0JICHU,
		CaoZuoBiao: Bm4FangFas,
		SuoYin:     FangFaMing,
		ZhuJian:    ZhuJian,
		ZiDuans: []map[string]interface{}{
			ZuZhuangBIGINT(ZhuJian, "20"),
			ZuZhuangVARCHAR(MingCheng, "50"),
			ZuZhuangVARCHAR(FangFaMing, "50"),
			ZuZhuangVARCHAR(MiaoShu, "50"),
		},
	}
	JiChuBiao[Bm4FangFas] = shuJu0
}

func tianJiaBm3YeWus() {
	// 业务表：主键，业务编码，业务名称，描述。//这是针对一个业务进行业务流向梳理,业务是必须有方法流的，但方法流必须也是一对多的，所以这里不需要去指定方法流字段
	shuJu0 := map[string]interface{}{
		CaoZuoKu:   XM0JICHU,
		CaoZuoBiao: Bm3YeWus,
		SuoYin:     BianMa,
		ZhuJian:    ZhuJian,
		ZiDuans: []map[string]interface{}{
			ZuZhuangBIGINT(ZhuJian, "20"),
			ZuZhuangVARCHAR(MingCheng, "50"),
			ZuZhuangVARCHAR(BianMa, "50"),
			ZuZhuangVARCHAR(MiaoShu, "150"),
		},
	}
	JiChuBiao[Bm3YeWus] = shuJu0
}

func tianJiaBm2ZiDuans() {
	// 字段表：主键，名称，编码，字段值表，正则，是否指定，是否有行为，长度
	shuJu0 := map[string]interface{}{
		CaoZuoKu:   XM0JICHU,
		CaoZuoBiao: Bm2ZiDuans,
		SuoYin:     BianMa,
		ZhuJian:    ZhuJian,
		ZiDuans: []map[string]interface{}{
			ZuZhuangBIGINT(ZhuJian, "20"),
			ZuZhuangVARCHAR(MingCheng, "50"),
			ZuZhuangVARCHAR(BianMa, "50"),
			ZuZhuangVARCHAR(ZhiBiao, "50"),
			ZuZhuangVARCHAR(ZhengZe, "50"),
			ZuZhuangINT(ShiFouZhiDing, "1"),
			ZuZhuangINT(ShiFouYouXingWei, "1"),
			ZuZhuangVARCHAR(ChangDu, "5"),
		},
	}
	JiChuBiao[Bm2ZiDuans] = shuJu0
}

func tianJiaBm1BiaoMings() {
	shuJu0 := map[string]interface{}{
		CaoZuoKu:   XM0JICHU,
		CaoZuoBiao: Bm1BiaoMings,
		SuoYin:     BianMa,
		ZhuJian:    ZhuJian,
		ZiDuans: []map[string]interface{}{
			ZuZhuangBIGINT(ZhuJian, "20"),
			ZuZhuangVARCHAR(BianMa, "50"),
			ZuZhuangVARCHAR(MingCheng, "50"),
			ZuZhuangVARCHAR(ZhuJianBiao, "50"),
			ZuZhuangVARCHAR(ShuJuKu, "50"),
		},
	}
	JiChuBiao[Bm1BiaoMings] = shuJu0
}

func chuShiHuaJiChuBiaoJieGou() {
	SuoShiLi.Do(func() { //这里需要把已存在的都纳入进来，所以需要新建一个配置文件，这个配置文件用go写成
		tianJiaBm1BiaoMings()
		tianJiaBm2ZiDuans()
		tianJiaBm3YeWus()
		tianJiaBm4FangFas()
		tianJiaBm5ZiDuanXingWeiLius()
		tianJiaBm6ZiDuanZhiDings()
		tianJiaBm7RuCans()
		tianJiaBm8ChuCans()
		tianJiaBm9YeWuXingWeiLius()
	})
}
func ZuZhuangBIGINT(mingCheng, kuandu string) map[string]interface{} {
	ret := map[string]interface{}{
		ZiDuanMing: mingCheng,
		LeiXing:    BIGINT,
		ChangDu:    kuandu,
		MoRenZhi:   Zf0,
	}
	return ret
}

func ZuZhuangINT(mingCheng, kuandu string) map[string]interface{} {
	ret := map[string]interface{}{
		ZiDuanMing: mingCheng,
		LeiXing:    INT,
		ChangDu:    kuandu,
		MoRenZhi:   Zf0,
	}
	return ret
}

func ZuZhuangVARCHAR(mingCheng, ziFuShu string) map[string]interface{} {
	ret := map[string]interface{}{
		ZiDuanMing: mingCheng,
		LeiXing:    VARCHAR,
		ChangDu:    ziFuShu,
		MoRenZhi:   FhDanYinHao + HFXXIAOXIE + FhDanYinHao,
	}
	return ret
}
