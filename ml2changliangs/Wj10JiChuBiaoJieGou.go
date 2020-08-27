package ml2changliangs

func HuoQuJiChuBiaoJieGou(biaoMing string) map[string]interface{} {
	ret := map[string]interface{}{
		Bm1BiaoMings:         Bm1BiaoMingsJieGou(),
		Bm2ZiDuans:           Bm2ZiDuansJieGou(),
		Bm3YeWus:             Bm3YeWusJieGou(),
		Bm4FangFas:           Bm4FangFasJieGou(),
		Bm5ZiDuanXingWeiLius: Bm5ZiDuanXingWeiLiusJieGou(),
		Bm6ZiDuanZhiDings:    Bm6ZiDuanZhiDingsJieGou(),
		Bm7RuCans:            Bm7RuCansJieGou(),
		Bm8ChuCans:           Bm8ChuCansJieGou(),
		Bm9YeWuXingWeiLius:   Bm9YeWuXingWeiLiusJieGou(),
    Bm10BiaoCis:          Bm10BiaoCisJieGou(),//所有字段和表编码的拆分选择词表
	}
	if biaoMing == FhKongZiFu {
		return ret
	}
	return ret[biaoMing].(map[string]interface{})
}
