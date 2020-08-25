package ml5kus

import (
	"database/sql"
	"log"
	"strings"
	"xm0jichu/ml2changliangs"
	"xm0jichu/ml3moxings"
)

func XinZeng(canShu ml3moxings.CanShu) ml3moxings.CanShu {
	//insert into caoZuoKu.caoZuoBiao (columns) values(values)
	canShuYiGe := canShu.ShuJu[ml2changliangs.Sz0]
	caoZuoKu := canShuYiGe[ml2changliangs.CaoZuoKu].(string)
	caoZuoBiao := canShuYiGe[ml2changliangs.CaoZuoBiao].(string)
	ziDuans := canShuYiGe[ml2changliangs.ZiDuans].(map[string]interface{}) //把字段拿出来

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

	dbCanShuRet := HuoQuLianJieChi(caoZuoKu)

	db := ml3moxings.HuoQuCeng1YiGe(dbCanShuRet).(*sql.DB)

	result, err := db.Exec(sqlStr, values...)

	log.Println("XinZeng:sqlStr,values,result,err---", sqlStr, values, result, err)
	return canShu
}
