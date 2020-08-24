package ml6kustests

import(
"testing"
"xm0jichu/ml3moxings"
"xm0jichu/ml2changliangs"
"xm0jichu/ml5kus"
"xm0jichu/ml0gongjus"
)

func TestXinZengBiao(t *testing.T){
  //数据库会初始化所以不需要自己操作，但是表不会，所以还要创建数据表之后再进行新增，新增完之后删除数据库
  canShuChuangJianBiao := ml3moxings.CanShu{
    ShuJu:[]map[string]interface{}{
      ml2changliangs.JiChuBiao[ml2changliangs.Bm1BiaoMings].(map[string]interface{}),
    },
  }
  ml5kus.ChuangJianBiao(canShuChuangJianBiao)

  canShu := ml3moxings.CanShu{}
  canShu.ShuJu = []map[string]interface{}{}
  
  shuJu0 := map[string]interface{}{
    ml2changliangs.CaoZuoKu:ml2changliangs.XM0JICHU,
    ml2changliangs.CaoZuoBiao:ml2changliangs.Bm1BiaoMings,
    
  }
  ziDuans:= []map[string]interface{}{
    map[string]interface{}{//Bm1BiaoMings中字段定义了多少字段就要声明多少字段。
      ml2changliangs.ShuJuKu:ml2changliangs.XM1YONGHU,
      ml2changliangs.ZhuJian:ml0gongjus.HuoQuId(),
      ml2changliangs.BianMa:ml2changliangs.Ywb1YongHus,
      ml2changliangs.MingCheng:"用户",
      ml2changliangs.ZhuJianBiao:ml2changliangs.ZjBiao+ml2changliangs.Zf1+ml2changliangs.XiaoXies,
    },
    map[string]interface{}{//Bm1BiaoMings中字段定义了多少字段就要声明多少字段。
      ml2changliangs.ShuJuKu:ml2changliangs.XM1YONGHU,
      ml2changliangs.ZhuJian:ml0gongjus.HuoQuId(),
      ml2changliangs.BianMa:ml2changliangs.Ywb2JueSes,
      ml2changliangs.MingCheng:"角色",
      ml2changliangs.ZhuJianBiao:ml2changliangs.ZjBiao+ml2changliangs.Zf2+ml2changliangs.XiaoXies,
    },
    map[string]interface{}{//Bm1BiaoMings中字段定义了多少字段就要声明多少字段。
      ml2changliangs.ShuJuKu:ml2changliangs.XM1YONGHU,
      ml2changliangs.ZhuJian:ml0gongjus.HuoQuId(),
      ml2changliangs.BianMa:ml2changliangs.Ywb3ZiYuans,
      ml2changliangs.MingCheng:"资源",
      ml2changliangs.ZhuJianBiao:ml2changliangs.ZjBiao+ml2changliangs.Zf3+ml2changliangs.XiaoXies,
    },
    map[string]interface{}{//Bm1BiaoMings中字段定义了多少字段就要声明多少字段。
      ml2changliangs.ShuJuKu:ml2changliangs.XM1YONGHU,
      ml2changliangs.ZhuJian:ml0gongjus.HuoQuId(),
      ml2changliangs.BianMa:ml2changliangs.Ywb4YongHuJueSes,
      ml2changliangs.MingCheng:"用户角色",
      ml2changliangs.ZhuJianBiao:ml2changliangs.ZjBiao+ml2changliangs.Zf4+ml2changliangs.XiaoXies,
    },
    map[string]interface{}{//Bm1BiaoMings中字段定义了多少字段就要声明多少字段。
      ml2changliangs.ShuJuKu:ml2changliangs.XM1YONGHU,
      ml2changliangs.ZhuJian:ml0gongjus.HuoQuId(),
      ml2changliangs.BianMa:ml2changliangs.Ywb5JueSeZiYuans,
      ml2changliangs.MingCheng:"角色资源",
      ml2changliangs.ZhuJianBiao:ml2changliangs.ZjBiao+ml2changliangs.Zf5+ml2changliangs.XiaoXies,
    },
  }
  for _,v :=range ziDuans{
    shuJu0[ml2changliangs.ZiDuans]=v
    canShu.ShuJu = append(canShu.ShuJu, shuJu0)
    ml5kus.XinZeng(canShu)
  }
  // ml5kus.ShanChuJiChuKu()//测试完成后删除数据库，基础测试在所有的测试方法中都不需要前置测试存在。
}
