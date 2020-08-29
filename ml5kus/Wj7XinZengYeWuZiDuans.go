package ml5kus

import (
	"xm0jichu/ml2changliangs"
  "log"
)
//这个也是管理员操作的，当有一个表时必须选中表才能进行字段的添加
//所以新增字段这个还是必须的，因为要从数据库里查出来进行数据库的新增和修改等
func XinZengYeWuZiDuans() {
  shuJu0 := ml2changliangs.YeWuBiaoJieGous(ml2changliangs.FhKongZiFu)
  for k,v := range shuJu0{
    log.Println()
    log.Println("k,v---",k,v)
    XinZeng(v.(map[string]interface{}))
    map2:=ml2changliangs.YeWuBiaoZiDuans(k)//为了避免在结构中多次使用CaoZuoKu和CaoZuoBiao，所以在这里重新分配这个值
    
    for _,vziduan := range map2[ml2changliangs.ZiDuans].([]map[string]interface{}){
      map2[ml2changliangs.CaoZuoZhis]=vziduan
      XinZeng(map2)
    }
  }

}
