package ml5kus

import (
	"database/sql"
	"log"
	"strings"
	"xm0jichu/ml2changliangs"
)

func XinZeng(canShu map[string]interface{}) map[string]interface{} {
	//insert into caoZuoKu.caoZuoBiao (columns) values(values)
	caoZuoKu := canShu[ml2changliangs.CaoZuoKu].(string)
	caoZuoBiao := canShu[ml2changliangs.CaoZuoBiao].(string)
	ziDuans := canShu[ml2changliangs.ZiDuans].(map[string]interface{}) //把字段拿出来

	keys := []string{}
	wenHaos := []string{}
	values := []interface{}{}
	for k, v := range ziDuans {
		keys = append(keys, k)
		values = append(values, v)
		wenHaos = append(wenHaos, ml2changliangs.FhWenHao)
	}

	builder := strings.Builder{}
	builder.WriteString(" INSERT INTO ")

	builder.WriteString(caoZuoKu)
	builder.WriteString(ml2changliangs.FhDianHao)
	builder.WriteString(caoZuoBiao)

	builder.WriteString(" ( ")
	builder.WriteString(strings.Join(keys, ","))
	builder.WriteString(" )values( ")
	builder.WriteString(strings.Join(wenHaos, ","))
	builder.WriteString(" ) ")

	sqlStr := builder.String()
	db := HuoQuLianJieChi(caoZuoKu)[ml2changliangs.Ceng1].(*sql.DB)
	result, err := db.Exec(sqlStr, values...)
	log.Println("XinZeng:sqlStr,values,result,err---", sqlStr, values, result, err)
	return canShu
}
