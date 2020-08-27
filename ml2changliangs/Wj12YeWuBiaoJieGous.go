package ml2changliangs

// 所有进入表名表的数据（也就是业务）都要遵循这个结构
// ZuZhuangBIGINT(ZhuJian,"20"),
// ZuZhuangVARCHAR(BianMa,"50"),
// ZuZhuangVARCHAR(MingCheng,"50"),
// ZuZhuangVARCHAR(ZhuJianBiao,"50"),
// ZuZhuangVARCHAR(ShuJuKu,"50"),

func YeWuBiaoJieGous(biaoMing string) map[string]interface{} {
	ret := map[string]interface{}{
    Ywb1YongHus:Ywb1YongHusJieGou(),
    Ywb2JueSes:Ywb2JueSesJieGou(),
    Ywb3ZiYuans:Ywb3ZiYuansJieGou(),
    Ywb4YongHuJueSes:Ywb4YongHuJueSesJieGou(),
    Ywb5JueSeZiYuans:Ywb5JueSeZiYuansJieGou(),
  }
	if biaoMing == FhKongZiFu {
		return ret
	}
	return ret[biaoMing].(map[string]interface{})
}
