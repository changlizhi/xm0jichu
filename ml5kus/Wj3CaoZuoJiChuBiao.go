package ml5kus

import (
	"xm0jichu/ml3moxings"
	"xm0jichu/ml2changliangs"
)

func ChuangJianJiChuBiao(){
  canShu := ml3moxings.CanShu{}
  
  for _,v := range ml2changliangs.JiChuBiao{
    canShu.ShuJu = []map[string]interface{}{
      v.(map[string]interface{}),
    }
    ChuangJianBiao(canShu)
  }
}
