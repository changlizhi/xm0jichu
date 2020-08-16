package ml1gongjustests

import(
  "testing"
  "xm0jichu/ml0gongjus"
  "log"
)

func TestDuQuMoXingJson(t *testing.T){
  ret := ml0gongjus.DuQuMoXingJson()
  for _,r:=range ret{
    log.Println("Wj10MoxingJson_test.go---TestDuQuMoXingJson,r",r)
  }
}
