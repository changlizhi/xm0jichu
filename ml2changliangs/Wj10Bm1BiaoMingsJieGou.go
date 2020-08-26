package ml2changliangs

func Bm1BiaoMingsJieGou() map[string]interface{} {
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
	return shuJu0
}
