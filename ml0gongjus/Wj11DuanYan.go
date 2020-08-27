package ml0gongjus

import(
"xm0jichu/ml2changliangs"
)

func DuanYanZiFuChuan(canShu interface{})string{
  if canShu != nil{
    return canShu.(string)
  }
  return ml2changliangs.FhKongZiFu
}
