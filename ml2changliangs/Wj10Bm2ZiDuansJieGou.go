package ml2changliangs

func Bm2ZiDuansJieGou() map[string]interface{} {
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
	return shuJu0
}
