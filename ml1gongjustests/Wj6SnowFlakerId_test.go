package ml1gongjustests
import(
  "log"
  "testing"
  "xm0jichu/ml0gongjus"
  "time"
)
func TestHuoquId(t *testing.T){
  ret := ml0gongjus.HuoquId()
  log.Println(ret)
}
func TestHuoQuIdZiFu(t *testing.T){
  for i:=0;i<100000;i++{
    go func(i int){
      ret := ml0gongjus.HuoQuIdZiFu()
      log.Println("TestHuoQuIdZiFu---",i,ret,len(ret))
    }(i)
  }
  time.Sleep(time.Duration(5) * time.Hour)
  
}
