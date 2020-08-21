package ml5kus

import(
  "log"
  "xm0jichu/ml2changliangs"
  "xm0jichu/ml3moxings"
  "strings"
  "database/sql"
)

func XinZengXm0JiChuShuJu(canShu ml3moxings.CanShu)ml3moxings.CanShu{
  canShuUse := ml3moxings.CanShu{}
  canShuUse.ShuJu = []map[string]interface{}{}
  
  shuJuKuMing := map[string]interface{}{
  	ml2changliangs.Ceng1: ml2changliangs.XM0JICHU,
  }
  canShuUse.ShuJu = append(canShuUse.ShuJu, shuJuKuMing)
  
  biaoMing := map[string]interface{}{
  	ml2changliangs.Ceng1: ml2changliangs.Bm1BiaoMings,
  }
  canShuUse.ShuJu = append(canShuUse.ShuJu, biaoMing)
  
  ziDuansKeyMap := map[string]interface{}{
  	ml2changliangs.Ceng1: canShu.ShuJu,
  }
  
  canShuUse.ShuJu = append(canShuUse.ShuJu, ziDuansKeyMap)
  //这里可以进行一波并发插入看下能耗时多久

  ChuangJianBiaoHeSuoYin()//先创建表和索引
  XinZeng(canShuUse)
  return canShuUse
}

func XinZeng(canShu ml3moxings.CanShu)ml3moxings.CanShu{
  //insert into shuJuKuMing.table (columns) values(values)
  shuJuKuMing := canShu.ShuJu[ml2changliangs.Sz0][ml2changliangs.Ceng1].(string)
  biaoMing := canShu.ShuJu[ml2changliangs.Sz1][ml2changliangs.Ceng1].(string)
  ziDuans := canShu.ShuJu[ml2changliangs.Sz2][ml2changliangs.Ceng1].([]map[string]interface{})
  
  keys:=[]string{}
  wenHaos:=[]string{}
  values:=[]interface{}{}
  for k, v := range ziDuans[ml2changliangs.Sz0] {
    keys = append(keys,k)
    values=append(values,v)
    wenHaos=append(wenHaos,ml2changliangs.FhWenHao)
  }
  
  builder := strings.Builder{}
  builder.WriteString(" INSERT INTO ")
  
  builder.WriteString(shuJuKuMing)
  builder.WriteString(ml2changliangs.FhDianHao)
  builder.WriteString(biaoMing)
  
  builder.WriteString(" ( ")
  builder.WriteString(strings.Join(keys,","))
  builder.WriteString(" )values( ")
  builder.WriteString(strings.Join(wenHaos,","))
  builder.WriteString(" ) ")
  
  sqlStr := builder.String()
  dbCanShuRet := HuoQuLianJieChi(shuJuKuMing)
  db := ml3moxings.HuoQuCeng1YiGe(dbCanShuRet).(*sql.DB)
  result, err := db.Exec(sqlStr,values...)
  
  log.Println("XinZeng:sqlStr,values,result,err---", sqlStr,values,result, err)
  return canShu
}
