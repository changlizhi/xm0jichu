package ml5kus

import(
  "log"
  "xm0jichu/ml2changliangs"
  "xm0jichu/ml3moxings"
  "strings"
  "database/sql"
)
func XinZeng(canShu ml3moxings.CanShu)ml3moxings.CanShu{
  //insert into table (columns) values(values)
  biaoMing := canShu.ShuJu[ml2changliangs.Sz0][ml2changliangs.Ceng1].(string)                  //把数据第一个拿出来当作表名
  ziDuans := canShu.ShuJu[ml2changliangs.Sz1][ml2changliangs.Ceng1].([]map[string]interface{}) //把字段拿出来
  //简单插入数据没必要做事务和ps了
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
  builder.WriteString(biaoMing)
  builder.WriteString(" ( ")
  builder.WriteString(strings.Join(keys,","))
  builder.WriteString(" )values( ")
  builder.WriteString(strings.Join(wenHaos,","))
  builder.WriteString(" ) ")
  
  sqlStr := builder.String()
  dbCanShuRet := HuoQuJiChuLianJieChi()
  db := ml3moxings.HuoQuCeng1YiGe(dbCanShuRet).(*sql.DB)
  result, err := db.Exec(sqlStr,values...)
  
  log.Println("sqlStr,result,err---", sqlStr,result, err,ziDuans)
  return canShu
}
