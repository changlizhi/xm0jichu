package ml5kus
import(
  "log"
  "xm0jichu/ml2changliangs"
  "xm0jichu/ml3moxings"
  "strings"
  "sort"
  "database/sql"
)

func ChaXun(canShu ml3moxings.CanShu)ml3moxings.CanShu{
  //select ziduans from caoZuoKu.caoZuoBiao where ziDuan=?
  canShuYiGe := canShu.ShuJu[ml2changliangs.Sz0]
  caoZuoKu := canShuYiGe[ml2changliangs.CaoZuoKu].(string)
  caoZuoBiao := canShuYiGe[ml2changliangs.CaoZuoBiao].(string)
  ziDuans := canShuYiGe[ml2changliangs.ZiDuans].([]interface{}) //把要查询的字段拿出来
  tiaoJianHeZhis := canShuYiGe[ml2changliangs.TiaoJianHeZhis].(map[string]interface{}) //把字段拿出来

  chaXunZiDuans:=[]string{}
  for _, v := range ziDuans {
    chaXunZiDuans=append(chaXunZiDuans,v.(string))
  }
  sort.Strings(chaXunZiDuans)
  wenHaos:=[]string{}
  values:=[]interface{}{}
  for k, v := range tiaoJianHeZhis {
    wenHaos=append(wenHaos,k+ml2changliangs.FhDengHao+ml2changliangs.FhWenHao)
    values=append(values,v)
  }
  
  builder := strings.Builder{}
  builder.WriteString(" SELECT ")
  builder.WriteString(strings.Join(chaXunZiDuans,ml2changliangs.FhDouHao))
  builder.WriteString(" FROM ")
  builder.WriteString(caoZuoKu)
  builder.WriteString(ml2changliangs.FhDianHao)
  builder.WriteString(caoZuoBiao)
  if len(wenHaos) > ml2changliangs.Sz0{
    builder.WriteString(" WHERE ")
    builder.WriteString(strings.Join(wenHaos," AND "))
  }
  
  sqlStr := builder.String()
  dbCanShuRet := HuoQuLianJieChi(caoZuoKu)
  db := ml3moxings.HuoQuCeng1YiGe(dbCanShuRet).(*sql.DB)
  
  rows, err := db.Query(sqlStr,values...)
  
  //在没有报错的情况下可以执行Next和Scan方法把所有查到的字段值拿出来放到CanShu里返回回去
  retShuJu := scanRet(chaXunZiDuans,rows)
  ret := ml3moxings.CanShu{
    ShuJu:retShuJu,
  }
  log.Println("XinZeng:sqlStr,values,err,ret---", sqlStr,values, err,ret)
  return canShu
}

func scanRet(lies []string,rows *sql.Rows)[]map[string]interface{}{
	ret := []map[string]interface{}{}
	for rows.Next() {//每一行
		tempLie := make([]string,len(lies))
    tempLieRef := make([]interface{},len(lies))

    for i := ml2changliangs.Sz0;i < len(lies); i++{
      tempLieRef[i] = &tempLie[i]
    }
    rows.Scan(
      tempLieRef...,
    )
		retOne := map[string]interface{}{}
		for xiaBiao,lie := range lies{
      if tempLieRef[xiaBiao] !=  nil{
        retOne[lie]=*tempLieRef[xiaBiao].(*string)
      }
		}
		ret = append(ret,retOne)
	}
	return ret
}
