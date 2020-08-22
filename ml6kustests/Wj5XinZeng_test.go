package ml6kustests

import(
"testing"
"xm0jichu/ml3moxings"
"xm0jichu/ml2changliangs"
"xm0jichu/ml5kus"
"xm0jichu/ml0gongjus"
)


func TestXinZeng(t *testing.T){
  yongHuKu := map[string]interface{}{}
  yongHuKu[ml2changliangs.ZhuJian]=ml0gongjus.HuoQuId()
  yongHuKu[ml2changliangs.MingCheng]="用户"
  yongHuKu[ml2changliangs.BianMa]=ml2changliangs.Ywb1YongHus
  yongHuKu[ml2changliangs.ZhuJianBiao]=ml2changliangs.ZjBiao+ml2changliangs.Zf1+ml2changliangs.XiaoXies
  yongHuKu[ml2changliangs.ShuJuKu]=ml2changliangs.XM1YONGHU
  
  canShu := ml3moxings.CanShu{}
  canShu.ShuJu = []map[string]interface{}{}
  
  shuJuKuMing := map[string]interface{}{
  	ml2changliangs.Ceng1: ml2changliangs.XM0JICHU,
  }
  canShu.ShuJu = append(canShu.ShuJu, shuJuKuMing)
  
  biaoMing := map[string]interface{}{
  	ml2changliangs.Ceng1: ml2changliangs.Bm1BiaoMings,
  }
  canShu.ShuJu = append(canShu.ShuJu, biaoMing)
  
  ziDuans := []map[string]interface{}{}
  
  ziDuans=append(ziDuans,yongHuKu)
  ziDuansKeyMap := map[string]interface{}{
  	ml2changliangs.Ceng1: ziDuans,
  }
  
  canShu.ShuJu = append(canShu.ShuJu, ziDuansKeyMap)
  //这里可以进行一波并发插入看下能耗时多久
  ml5kus.XinZeng(canShu)

}
