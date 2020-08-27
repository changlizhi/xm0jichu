package ml2changliangs
func Ywb2JueSesZiDuans()map[string]interface{}{
  ret := map[string]interface{}{
    CaoZuoKu:   XT0JICHU,
    CaoZuoBiao: Bm2ZiDuans,//这个表里定义了多少字段就要提供多少字段值
    ZiDuans:[]map[string]interface{}{
      map[string]interface{}{
        SuoShuBiao:Ywb2JueSes,
        BianMa:Ywb2JueSes+ZhuJian,
        MingCheng:"角色主键",
        ZhiBiao:ZdBiao+Zf1+Zf0+XiaoXies,
      },
      map[string]interface{}{
        SuoShuBiao:Ywb2JueSes,
        BianMa:Ywb2JueSes+BianMa,
        MingCheng:"角色编码",
        ZhiBiao:ZdBiao+Zf1+Zf1+XiaoXies,
      },
      map[string]interface{}{
        SuoShuBiao:Ywb2JueSes,
        BianMa:Ywb2JueSes+MingCheng,
        MingCheng:"角色名称",
        ZhiBiao:ZdBiao+Zf1+Zf2+XiaoXies,
      },
      map[string]interface{}{
        SuoShuBiao:Ywb2JueSes,
        BianMa:Ywb2JueSes+ZhuangTai,
        MingCheng:"角色状态",
        ZhiBiao:ZdBiao+Zf1+Zf3+XiaoXies,
      },
      map[string]interface{}{
        SuoShuBiao:Ywb2JueSes,
        BianMa:Ywb2JueSes+ShuoMing,
        MingCheng:"角色说明",
        ZhiBiao:ZdBiao+Zf1+Zf4+XiaoXies,
      },
    },
  }
  return ret
}
