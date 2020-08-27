package ml2changliangs

func Bm10BiaoCisJieGou() map[string]interface{} {
	// 字段行为流表：主键，字段主键，顺序，总个数，名称，行为编码（虽然感觉多余，但是为了统一记录加上），方法名（方法名在方法表里是唯一的，所以这里其实是可以作为一个索引的，但是方法流里不同的方法流是可以有多个甚至相同的方法组的，所以不能作为索引），描述。//方法流跟字段是一对多的，查出来之后根据顺序排序即可，即使两个字段的业务方法流是一样的也必须拆成两个值
	shuJu0 := map[string]interface{}{
		CaoZuoKu:   XT0JICHU,
		CaoZuoBiao: Bm10BiaoCis,
		SuoYin:     BianMa,
		ZhuJian:    ZhuJian,
		ZiDuans: []map[string]interface{}{
			ZuZhuangBIGINT(ZhuJian, "20"),
			ZuZhuangVARCHAR(BianMa, "50"),
			ZuZhuangVARCHAR(MingCheng, "50"),
			ZuZhuangVARCHAR(MiaoShu, "350"),
		},
	}
	return shuJu0
}
