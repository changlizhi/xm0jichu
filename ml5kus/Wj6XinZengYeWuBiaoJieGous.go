package ml5kus

import (
	"xm0jichu/ml2changliangs"
  "log"
)

func XinZengYeWuJieGous() {
	//数据库会初始化所以不需要自己操作，但是表不会，所以还要创建数据表之后再进行新增，新增完之后删除数据库
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
