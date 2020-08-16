package ml4moxingstests

import(
"testing"
"xm0jichu/ml3moxings"
"log"
"xm0jichu/ml2changliangs"
)

func TestZuJianCeng1YiGe(t *testing.T){
  ret := ml3moxings.ZuJianCeng1YiGe("123")
  log.Println("ret---",ret.ShuJu[ml2changliangs.Sz0])
  ret2 := ml3moxings.HuoQuCeng1YiGe(ret)
  log.Println("ret2---",ret2.(string))
  ret3 := ml3moxings.ZuJianCeng1YiGe([]string{"123","1234"})
  log.Println("ret3---",ret3.ShuJu[ml2changliangs.Sz0][ml2changliangs.Ceng1].([]string))
}


