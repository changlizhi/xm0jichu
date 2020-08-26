package ml2changliangs

// 所有进入字段表的数据（也就是业务）都要遵循这个结构
// 从字段角度来说，合在一起返回是没有意义的，如果一定要合在一起那就循环所有的表名表再来拿这个集合
// ZuZhuangVARCHAR(BianMa,"50"),
// ZuZhuangVARCHAR(MingCheng,"50"),
// ZuZhuangVARCHAR(ZhiBiao,"50"),
// ZuZhuangVARCHAR(ZhengZe,"50"),
// ZuZhuangINT(ShiFouZhiDing,"1"),
// ZuZhuangINT(ShiFouYouXingWei,"1"),
// ZuZhuangVARCHAR(ChangDu,"5"),
// ZuZhuangVARCHAR(SuoShuBiao,"5"),


func YeWuBiaoZiDuan(biaoMing string) map[string]interface{} {
	ret := map[string]interface{}{
    Ywb1YongHus:Ywb1YongHusZiDuans(),
    Ywb2JueSes:Ywb2JueSesZiDuans(),
    Ywb3ZiYuans:Ywb3ZiYuansZiDuans(),
    Ywb4YongHuJueSes:Ywb4YongHuJueSesZiDuans(),
    Ywb5JueSeZiYuans:Ywb5JueSeZiYuansZiDuans(),
  }
	return ret[biaoMing].(map[string]interface{})
}
