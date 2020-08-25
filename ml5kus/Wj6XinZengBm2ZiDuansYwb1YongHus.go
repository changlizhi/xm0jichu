package ml5kus

import(
"xm0jichu/ml3moxings"
"xm0jichu/ml2changliangs"
"xm0jichu/ml0gongjus"
)

func XinZengBm2ZiDuansYwb1YongHus(){
  //查出Bm1BiaoMings中Ywb1YongHus指定的主键
  caoZuoBiao:=ml2changliangs.Bm2ZiDuans
  
  //默认有数据库，在调用此方法之前必须自己保证数据库已存在
  canShu := ml3moxings.CanShu{}
  canShu.ShuJu = []map[string]interface{}{}
  
  shuJu0 := map[string]interface{}{
    ml2changliangs.CaoZuoKu:ml2changliangs.XM0JICHU,
    ml2changliangs.CaoZuoBiao:caoZuoBiao,
  }
  
  ziDuans:= []map[string]interface{}{
    // ZuZhuangBIGINT(ZhuJian,"20"),
    // ZuZhuangVARCHAR(BianMa,"50"),
    // ZuZhuangVARCHAR(MingCheng,"50"),
    // ZuZhuangVARCHAR(ZhiBiao,"50"),
    // ZuZhuangVARCHAR(ZhengZe,"50"),
    // ZuZhuangINT(ShiFouZhiDing,"1"),
    // ZuZhuangINT(ShiFouYouXingWei,"1"),
    // ZuZhuangVARCHAR(ChangDu,"5"),
    
    map[string]interface{}{//Bm2ZiDuans中字段定义了多少字段就要声明多少字段。
      ml2changliangs.ZhuJian:ml0gongjus.HuoQuId(),
    },
  }
  for _,v :=range ziDuans{
    shuJu0[ml2changliangs.ZiDuans]=v
    canShu.ShuJu = append(canShu.ShuJu, shuJu0)
    XinZeng(canShu)
  }
  // ml5kus.ShanChuJiChuKu()//测试完成后删除数据库，基础测试在所有的测试方法中都不需要前置测试存在。
}
// func TestXinZengYongHuJieGou(t *testing.T){
//   //根据情况需要新增的数据有哪些？先把三个权限相关的表和两个中间表做了。从整个系统资源来看，除了界面资源之外其他的资源跟角色都没有多大关系。后台管理时有一定关系，就是根据自己这个角色管理其他角色的用户数据。
//   //首先新增一个用户表，由于数据库对象都被放到了map里，所以没必要定义结构体来进行数据库字段的映射，实际上结构体是用来进行算法优化的，而不是像之前那样拿来进行数据库Orm的，以前完全的本末倒置
//   //测试的时候新增一个对象，传入的时候是没有id的，所以需要保证留下一个参数，当正式环境时这个参数会传回界面保存下来留作字段新增时使用，
//   yongHuKu := map[string]interface{}{}
//   yongHuKu[ml2changliangs.MingCheng]="用户"
//   yongHuKu[ml2changliangs.BianMa]="YongHu"
//   yongHuKu[ml2changliangs.ZhuJianBiao]=ml2changliangs.ZjBiao+"1"
//   yhzd:=[]map[string]string{}
  
//   yongHuZiDuan1 := map[string]string{}
//   yongHuZiDuan1[ml2changliangs.BianMa]="YongHuZhuJian"
//   yongHuZiDuan1[ml2changliangs.MingCheng]="用户主键"
//   yongHuZiDuan1[ml2changliangs.ZhiBiao]=ml2changliangs.ZdBiao+"1"
//   yhzd=append(yhzd,yongHuZiDuan1)

//   yongHuZiDuan2 := map[string]string{}
//   yongHuZiDuan2[ml2changliangs.BianMa]="YongHuBianMa"
//   yongHuZiDuan2[ml2changliangs.MingCheng]="用户编码"
//   yongHuZiDuan2[ml2changliangs.ZhiBiao]=ml2changliangs.ZdBiao+"2"
//   yhzd=append(yhzd,yongHuZiDuan2)
  
//   yongHuZiDuan3 := map[string]string{}
//   yongHuZiDuan3[ml2changliangs.BianMa]="YongHuMingCheng"
//   yongHuZiDuan3[ml2changliangs.MingCheng]="用户名称"
//   yongHuZiDuan3[ml2changliangs.ZhiBiao]=ml2changliangs.ZdBiao+"3"
//   yhzd=append(yhzd,yongHuZiDuan3)
  
//   yongHuZiDuan4 := map[string]string{}
//   yongHuZiDuan4[ml2changliangs.BianMa]="YongHuTouXiang"
//   yongHuZiDuan4[ml2changliangs.MingCheng]="用户头像"
//   yongHuZiDuan4[ml2changliangs.ZhiBiao]=ml2changliangs.ZdBiao+"4"
//   yhzd=append(yhzd,yongHuZiDuan4)
  
//   yongHuZiDuan5 := map[string]string{}
//   yongHuZiDuan5[ml2changliangs.BianMa]="YongHuMiMa"
//   yongHuZiDuan5[ml2changliangs.MingCheng]="用户密码"
//   yongHuZiDuan5[ml2changliangs.ZhiBiao]=ml2changliangs.ZdBiao+"5"
//   yhzd=append(yhzd,yongHuZiDuan5)
  
//   yongHuZiDuan6 := map[string]string{}
//   yongHuZiDuan6[ml2changliangs.BianMa]="YongHuYouXiang"
//   yongHuZiDuan6[ml2changliangs.MingCheng]="用户邮箱"
//   yongHuZiDuan6[ml2changliangs.ZhiBiao]=ml2changliangs.ZdBiao+"6"
//   yhzd=append(yhzd,yongHuZiDuan6)
  
//   yongHuZiDuan7 := map[string]string{}
//   yongHuZiDuan7[ml2changliangs.BianMa]="YongHuDiZhi"
//   yongHuZiDuan7[ml2changliangs.MingCheng]="用户地址"
//   yongHuZiDuan7[ml2changliangs.ZhiBiao]=ml2changliangs.ZdBiao+"7"
//   yhzd=append(yhzd,yongHuZiDuan7)
  
//   yongHuZiDuan8 := map[string]string{}
//   yongHuZiDuan8[ml2changliangs.BianMa]="YongHuZhuangTai"
//   yongHuZiDuan8[ml2changliangs.MingCheng]="用户状态"
//   yongHuZiDuan8[ml2changliangs.ZhiBiao]=ml2changliangs.ZdBiao+"8"
//   yhzd=append(yhzd,yongHuZiDuan8)
  
//   yongHuZiDuan9 := map[string]string{}
//   yongHuZiDuan9[ml2changliangs.BianMa]="YongHuTiZhong"
//   yongHuZiDuan9[ml2changliangs.MingCheng]="用户体重"
//   yongHuZiDuan9[ml2changliangs.ZhiBiao]=ml2changliangs.ZdBiao+"9"
//   yhzd=append(yhzd,yongHuZiDuan9)

//   yongHuKu[ml2changliangs.ZiDuans]=yhzd
  
//   canShu := ml3moxings.CanShu{}
//   canShu.ShuJu = []map[string]interface{}{}
  
//   biaoMing := map[string]interface{}{
//   	ml2changliangs.Ceng1: ml2changliangs.Ywb1YongHus,
//   }
//   canShu.ShuJu = append(canShu.ShuJu, biaoMing)
  
//   ziDuans := []map[string]string{}
//   zd := map[string]string{}
  
  
  
//   ziDuans=append(ziDuans,zd)
//   ziDuansKeyMap := map[string]interface{}{
//   	ml2changliangs.Ceng1: ziDuans,
//   }
  
//   canShu.ShuJu = append(canShu.ShuJu, ziDuansKeyMap)
  
//   ret := ml5kus.XinZeng(canShu)

// }
