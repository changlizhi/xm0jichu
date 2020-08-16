package ml3moxings

import(
"xm0jichu/ml2changliangs"
)

type CanShu struct{
  Jwt string
  XuHao string
  ZhuangTai string
  ShuJu []map[string]interface{}
}
func HuoQuCeng1YiGe(canShu CanShu)interface{}{
  return canShu.ShuJu[0]["Ceng1"]
}
func ZuJianCeng1YiGe(oneShuJu interface{})CanShu{
  retShuJu:= []map[string]interface{}{}
  retOneShuJu:=map[string]interface{}{}
  retOneShuJu["Ceng1"]=oneShuJu
  retShuJu=append(retShuJu,retOneShuJu)
  ret := CanShu{
    ZhuangTai:ml2changliangs.HFX000000000,
    ShuJu:retShuJu,
  }
  return ret
}