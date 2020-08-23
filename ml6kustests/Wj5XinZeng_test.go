package ml6kustests

import(
"testing"
"xm0jichu/ml3moxings"
"xm0jichu/ml2changliangs"
"xm0jichu/ml5kus"
"xm0jichu/ml0gongjus"
)


func TestXinZeng(t *testing.T){
  canShu := ml3moxings.CanShu{}
  canShu.ShuJu = []map[string]interface{}{}
  
  shuJu0 := map[string]interface{}{
  }
  shuJu0[ml2changliangs.ShuJuKu]=ml2changliangs.XM0JICHU
  shuJu0[ml2changliangs.ZhuJian]=ml2changliangs.ZhuJian
  shuJu0[ml2changliangs.BiaoMing]=ml2changliangs.Bm1BiaoMings
  
  ziDuans := map[string]interface{}{}
  ziDuans[ml2changliangs.ShuJuKu]=ml2changliangs.XM1YONGHU
  ziDuans[ml2changliangs.ZhuJian]=ml0gongjus.HuoQuId()
  ziDuans[ml2changliangs.BianMa]=ml2changliangs.Ywb1YongHus
  ziDuans[ml2changliangs.MingCheng]="用户"
  ziDuans[ml2changliangs.ZhuJianBiao]=ml2changliangs.ZjBiao+ml2changliangs.Zf1+ml2changliangs.XiaoXies
  
  shuJu0[ml2changliangs.ZiDuans]=ziDuans
  
  canShu.ShuJu = append(canShu.ShuJu, shuJu0)
  //这里可以进行一波并发插入看下能耗时多久
  ml5kus.XinZeng(canShu)
}
