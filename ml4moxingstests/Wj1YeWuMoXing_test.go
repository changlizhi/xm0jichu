package ml4moxingstests

import(
"testing"
"xm0jichu/ml3moxings"
"log"
)

func TestZuJianCeng1YiGe(t *testing.T){
  ret := ml3moxings.ZuJianCeng1YiGe("123")
  log.Println("ret---",ret.ShuJu[0])
  ret2 := ml3moxings.HuoQuCeng1YiGe(ret)
  log.Println("ret2---",ret2.(string))
}


