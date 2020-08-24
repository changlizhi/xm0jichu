package ml6kustests

import(
"testing"
"xm0jichu/ml3moxings"
"xm0jichu/ml2changliangs"
"xm0jichu/ml0gongjus"
"xm0jichu/ml5kus"
)
//测试新增基础表数据时比较特殊，需要明确的知道
//1.新增的数据库在哪，这个必须在系统启动时配置好，新建数据库比较麻烦，因为能拿到的链接目前只有一个，后续可以通过map的方式把链接放进去然后根据数据库名来获取新的链接，但这样真的太动态化了，很多错误需要识别。其实也不是不行，每个测试先加载数据库实例其实也是合适的。
//2.新增测试时需要知道表名表指出的主键，这个其实也没有绝对的必要性，因为加了索引所以可以根据唯一索引来控制查询出来之后再新增
//3.新增字段时需要新增相应的业务表，这个很重要。根据此时创建的主键表和值表可以控制住后续的业务数据在插入时的真实显示
//4.新增表名时需要增加相应的主键表
//这一层控制将是绝对的关键位置，所有后续业务都依赖这一步的创建绝对成功和没有纰漏。

//做这个的时候需要操作中开启事务

//流程：首先拿到数据库连接池
//1.新增表名数据
//2.根据表名编码获取主键新增字段数据
//3.查询表名表新增表名主键表
//4.查询字段表新增字段值表
//5.会存在表没有建好但开始插入数据的情况吗？如果新建了一个业务，必然会新建相应的表。如果是基础库没有添加表呢？一般不应该出现，如果出现迁移情况也应该把这些测试方法都执行一次之后让表存在。

func zuZhuangBiaoMingShuJu(shuJuKu,bianMa,mingCheng,ZhuJianBiao string)ml3moxings.CanShu{
  canShu := ml3moxings.CanShu{}
  canShusArr := []map[string]interface{}{}

  canShus := map[string]interface{}{}
  canShus[ml2changliangs.ShuJuKu]=shuJuKu
  canShus[ml2changliangs.ZhuJian]=ml0gongjus.HuoQuId()
  canShus[ml2changliangs.BianMa]=bianMa
  canShus[ml2changliangs.MingCheng]=mingCheng
  canShus[ml2changliangs.ZhuJianBiao]=ZhuJianBiao
  
  canShusArr=append(canShusArr,canShus)
  canShu.ShuJu = canShusArr
  return canShu
}

func TestXinZengJiChuBiao(t *testing.T){

  yongHuBiao := zuZhuangBiaoMingShuJu(ml2changliangs.XM0JICHU,ml2changliangs.Ywb1YongHus,"用户",ml2changliangs.ZjBiao+ml2changliangs.Zf1+ml2changliangs.XiaoXies)
  ml5kus.XinZengXm0JiChuShuJu(yongHuBiao)
  
  jueSeBiao := zuZhuangBiaoMingShuJu(ml2changliangs.XM0JICHU,ml2changliangs.Ywb2JueSes,"角色",ml2changliangs.ZjBiao+ml2changliangs.Zf2+ml2changliangs.XiaoXies)
  ml5kus.XinZengXm0JiChuShuJu(jueSeBiao)
  
  ziYuanBiao := zuZhuangBiaoMingShuJu(ml2changliangs.XM0JICHU,ml2changliangs.Ywb3ZiYuans,"资源",ml2changliangs.ZjBiao+ml2changliangs.Zf3+ml2changliangs.XiaoXies)
  ml5kus.XinZengXm0JiChuShuJu(ziYuanBiao)
  
  yongHuJueSeBiao := zuZhuangBiaoMingShuJu(ml2changliangs.XM0JICHU,ml2changliangs.Ywb4YongHuJueSes,"用户角色",ml2changliangs.ZjBiao+ml2changliangs.Zf4+ml2changliangs.XiaoXies)
  ml5kus.XinZengXm0JiChuShuJu(yongHuJueSeBiao)
  
  jueSeZiYuanBiao := zuZhuangBiaoMingShuJu(ml2changliangs.XM0JICHU,ml2changliangs.Ywb5JueSeZiYuans,"角色资源",ml2changliangs.ZjBiao+ml2changliangs.Zf5+ml2changliangs.XiaoXies)
  ml5kus.XinZengXm0JiChuShuJu(jueSeZiYuanBiao)
}
