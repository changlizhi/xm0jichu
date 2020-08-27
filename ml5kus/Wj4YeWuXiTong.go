package ml5kus

import(
"xm0jichu/ml2changliangs"
"log"
)

//从业务操作来说，用户表，角色表，活动表或者言论表都是要在系统提供出去的时候由管理员在界面上进行添加的，所以前面添加的那些常量其实是毫无意义的，都要删掉。
//但现在依托那几个基本的数据进行真正的业务表生成，然后进行业务数据添加真正进入业务开发阶段。应该增加一个通用字段词表，当新增字段或者表的时候需要在这个词表中进行添加然后指定选择，不允许随便添加也不允许重名。

func chaXunBm1BiaoMings(canShu string)map[string]interface{}{
  shuJu0 := ml2changliangs.HuoQuJiChuBiaoJieGou(ml2changliangs.Bm1BiaoMings)
  bm1BiaoMingsZiDuans := shuJu0[ml2changliangs.ZiDuans].([]map[string]interface{})
	chaXunZiDuans := []interface{}{}
  
  for _, v := range bm1BiaoMingsZiDuans {
  	chaXunZiDuans = append(chaXunZiDuans, v[ml2changliangs.ZiDuanMing])
  }
  
  shuJu1 := map[string]interface{}{
  	ml2changliangs.CaoZuoKu:   ml2changliangs.XT0JICHU,
  	ml2changliangs.CaoZuoBiao: ml2changliangs.Bm1BiaoMings,
  	ml2changliangs.ZiDuans:    chaXunZiDuans,
  	ml2changliangs.TiaoJianHeZhis: map[string]interface{}{
  		ml2changliangs.BianMa: canShu, //查询这个条件的数据
  	},
  }
  return ChaXun(shuJu1)
}

func ChuangJianYeWuXiTong(canShu string){//这个系统需要做到的事情是
  //逻辑就是在biaomings里先查到所有的业务表名，然后再根据业务表查出拥有的字段，然后直接创建数据库和表，这时候要根据数据库来定，因为数据库是绝对业务相关的，所以创建业务表就是根据某一个业务进行创建。
  //这时候数据库里必然已经在biaomings和ziduans里有了一个Ywb1YongHus数据相关的结构体系，那么首先根据这个名称从biaomins里拿出所有的业务表声明，然后直接创建一个Ywb1YongHus数据库并创建结构体系一致的主键表和字段表。
  retBiaos:=chaXunBm1BiaoMings(canShu)
  log.Println("retBiaos---",retBiaos)
  //查到了所有的表直接创建数据库和库中的表
}
