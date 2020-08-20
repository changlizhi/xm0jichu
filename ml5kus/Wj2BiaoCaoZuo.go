package ml5kus

import (
	"database/sql"
	"log"
	"strings"
	"xm0jichu/ml2changliangs"
	"xm0jichu/ml3moxings"
)
//后续需要实现一下，在创建表时不做那么多限制

func SheZhiWeiYiSuoYin(canShu ml3moxings.CanShu)ml3moxings.CanShu{
  //ALTER TABLE `bm1biaomings` ADD UNIQUE INDEX `BianMa` (`BianMa`);
  biaoMing := canShu.ShuJu[ml2changliangs.Sz0][ml2changliangs.Ceng1].(string)                  //把数据第一个拿出来当作表名
  suoYin := canShu.ShuJu[ml2changliangs.Sz1][ml2changliangs.Ceng1].(string)                  //索引名
  builder := strings.Builder{}
  
  builder.WriteString(" ALTER TABLE ")
  builder.WriteString(biaoMing)
  builder.WriteString(" ADD UNIQUE INDEX ")
  builder.WriteString(suoYin)
  builder.WriteString(" ( ")
  builder.WriteString(suoYin)
  builder.WriteString(" ) ")
  
  sqlStr:=builder.String()
	dbCanShuRet := HuoQuLianJieChi(ml2changliangs.XM0JICHU)

	db := ml3moxings.HuoQuCeng1YiGe(dbCanShuRet).(*sql.DB)
	result, err := db.Exec(sqlStr)

	log.Println("SheZhiWeiYiSuoYin:sqlStr,result,err---", sqlStr, result, err)
  return canShu
}

func ChuangJianBiao(canShu ml3moxings.CanShu) ml3moxings.CanShu {
	//CREATE TABLE `BiaoMing` (
	// `ZhuJian` BIGINT(20) NOT NULL,
	// `MingCheng` VARCHAR(50) NOT NULL DEFAULT 'hfx',
	// `BianMa` VARCHAR(50) NOT NULL DEFAULT 'hfx',
	// `ZhuJianZhiBiao` VARCHAR(50) NOT NULL DEFAULT 'hfx',
	// PRIMARY KEY (`ZhuJian`)
	//  )
	//  ENGINE=InnoDB
  // 这里要求传入的ShuJu第一个必须是表名字符串，
  // 第二个必须是主键名指定
  // 第三个必须是字段列表
	// 后续强化这个校验
  biaoMing := canShu.ShuJu[ml2changliangs.Sz0][ml2changliangs.Ceng1].(string)                  //把数据第一个拿出来当作表名
	zhuJian := canShu.ShuJu[ml2changliangs.Sz1][ml2changliangs.Ceng1].(string)                  //主键名，也可能是组合主键
	ziDuans := canShu.ShuJu[ml2changliangs.Sz2][ml2changliangs.Ceng1].([]map[string]interface{}) //把字段拿出来

	builder := strings.Builder{}

	builder.WriteString("CREATE TABLE " + biaoMing + " (")
	for _, v := range ziDuans {
		//`MingCheng` VARCHAR(50) NOT NULL DEFAULT 'hfx',
		builder.WriteString(v[ml2changliangs.ZiDuanMing].(string))
		builder.WriteString(" ")
		builder.WriteString(v[ml2changliangs.LeiXing].(string))
		builder.WriteString("(")
		builder.WriteString(v[ml2changliangs.ChangDu].(string))
		builder.WriteString(") NOT NULL DEFAULT ")
		builder.WriteString(v[ml2changliangs.MoRenZhi].(string))
		builder.WriteString(",")
	}

	builder.WriteString("PRIMARY KEY ("+zhuJian+"))COLLATE='utf8mb4_general_ci' ENGINE=InnoDB")
	sqlStr := builder.String()

	dbCanShuRet := HuoQuLianJieChi(ml2changliangs.XM0JICHU)

	db := ml3moxings.HuoQuCeng1YiGe(dbCanShuRet).(*sql.DB)
	result, err := db.Exec(sqlStr)

	log.Println("ChuangJianBiao:sqlStr,result,err---", sqlStr, result, err)
	return canShu
}
