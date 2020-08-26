package ml6kustests

// import (
// 	"testing"
// 	"xm0jichu/ml0gongjus"
// 	"xm0jichu/ml2changliangs"
// 	"xm0jichu/ml5kus"
// )

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

// func TestXinZengJueSeZiYuanJieGou(t *testing.T){
//   jueSeZiYuanKu := map[string]interface{}{}
//   jiChuKu=append(jiChuKu,jueSeZiYuanKu)
//   jueSeZiYuanKu[ml2changliangs.MingCheng]="角色资源"
//   jueSeZiYuanKu[ml2changliangs.BianMa]="JueSeZiYuan"
//   jueSeZiYuanKu[ml2changliangs.ZhuJianBiao]=ml2changliangs.ZjBiao+"5"

//   jszyzd := []map[string]string{}

//   jueSeZiYuanZiDuan1 := map[string]string{}
//   jueSeZiYuanZiDuan1[ml2changliangs.BianMa]="JueSeZiYuanZhuJian"
//   jueSeZiYuanZiDuan1[ml2changliangs.MingCheng]="角色资源主键"
//   jueSeZiYuanZiDuan1[ml2changliangs.ZhiBiao]=ml2changliangs.ZdBiao+"21"
//   jszyzd=append(jszyzd,jueSeZiYuanZiDuan1)

//   jueSeZiYuanZiDuan2 := map[string]string{}
//   jueSeZiYuanZiDuan2[ml2changliangs.BianMa]="JueSeZhuJian"
//   jueSeZiYuanZiDuan2[ml2changliangs.MingCheng]="角色主键"
//   jueSeZiYuanZiDuan2[ml2changliangs.ZhiBiao]=ml2changliangs.ZdBiao+"22"
//   jszyzd=append(jszyzd,jueSeZiYuanZiDuan1)

//   jueSeZiYuanZiDuan3 := map[string]string{}
//   jueSeZiYuanZiDuan3[ml2changliangs.BianMa]="ZiYuanZhuJian"
//   jueSeZiYuanZiDuan3[ml2changliangs.MingCheng]="资源主键"
//   jueSeZiYuanZiDuan3[ml2changliangs.ZhiBiao]=ml2changliangs.ZdBiao+"23"
//   jszyzd=append(jszyzd,jueSeZiYuanZiDuan3)

//   jueSeZiYuanKu[ml2changliangs.ZiDuans]=jszyzd

// }

// func TestXinZengYongHuJueSeJieGou(t *testing.T){
//   yongHuJueSeKu := map[string]interface{}{}
//   jiChuKu=append(jiChuKu,yongHuJueSeKu)
//   yongHuJueSeKu[ml2changliangs.MingCheng]="用户角色"
//   yongHuJueSeKu[ml2changliangs.BianMa]="YongHuJueSe"
//   yongHuJueSeKu[ml2changliangs.ZhuJianBiao]=ml2changliangs.ZjBiao+"4"

//   yhjszd := []map[string]string{}

//   yongHuJueSeZiDuan1 := map[string]string{}
//   yongHuJueSeZiDuan1[ml2changliangs.BianMa]="YongHuJueSeZhuJian"
//   yongHuJueSeZiDuan1[ml2changliangs.MingCheng]="用户角色主键"
//   yongHuJueSeZiDuan1[ml2changliangs.ZhiBiao]=ml2changliangs.ZdBiao+"21"
//   yhjszd=append(yhjszd,yongHuJueSeZiDuan1)

//   yongHuJueSeZiDuan2 := map[string]string{}
//   yongHuJueSeZiDuan2[ml2changliangs.BianMa]="YongHuZhuJian"
//   yongHuJueSeZiDuan2[ml2changliangs.MingCheng]="用户主键"
//   yongHuJueSeZiDuan2[ml2changliangs.ZhiBiao]=ml2changliangs.ZdBiao+"22"
//   yhjszd=append(yhjszd,yongHuJueSeZiDuan2)

//   yongHuJueSeZiDuan3 := map[string]string{}
//   yongHuJueSeZiDuan3[ml2changliangs.BianMa]="JueSeZhuJian"
//   yongHuJueSeZiDuan3[ml2changliangs.MingCheng]="角色主键"
//   yongHuJueSeZiDuan3[ml2changliangs.ZhiBiao]=ml2changliangs.ZdBiao+"23"
//   yhjszd=append(yhjszd,yongHuJueSeZiDuan3)

//   yongHuJueSeKu[ml2changliangs.ZiDuans]=yhjszd

// }

// func TestXinZengZiYuanJieGou(t *testing.T){
//   ziYuanKu := map[string]interface{}{}
//   jiChuKu=append(jiChuKu,ziYuanKu)
//   ziYuanKu[ml2changliangs.MingCheng]="资源"
//   ziYuanKu[ml2changliangs.BianMa]="ZiYuan"
//   ziYuanKu[ml2changliangs.ZhuJianBiao]=ml2changliangs.ZjBiao+"3"

//   zyzd := []map[string]string{}

//   ziYuanZiDuan1 := map[string]string{}
//   ziYuanZiDuan1[ml2changliangs.BianMa]="ZiYuanZhuJian"
//   ziYuanZiDuan1[ml2changliangs.MingCheng]="资源主键"
//   ziYuanZiDuan1[ml2changliangs.ZhiBiao]=ml2changliangs.ZdBiao+"15"
//   zyzd=append(zyzd,ziYuanZiDuan1)

//   ziYuanZiDuan2 := map[string]string{}
//   ziYuanZiDuan2[ml2changliangs.BianMa]="ZiYuanLeiXing"
//   ziYuanZiDuan2[ml2changliangs.MingCheng]="资源类型"
//   ziYuanZiDuan2[ml2changliangs.ZhiBiao]=ml2changliangs.ZdBiao+"16"
//   zyzd=append(zyzd,ziYuanZiDuan2)

//   ziYuanZiDuan3 := map[string]string{}
//   ziYuanZiDuan3[ml2changliangs.BianMa]="ZiYuanBianMa"
//   ziYuanZiDuan3[ml2changliangs.MingCheng]="资源编码"
//   ziYuanZiDuan3[ml2changliangs.ZhiBiao]=ml2changliangs.ZdBiao+"17"
//   zyzd=append(zyzd,ziYuanZiDuan3)

//   ziYuanZiDuan4 := map[string]string{}
//   ziYuanZiDuan4[ml2changliangs.BianMa]="ZiYuanMingCheng"
//   ziYuanZiDuan4[ml2changliangs.MingCheng]="资源名称"
//   ziYuanZiDuan4[ml2changliangs.ZhiBiao]=ml2changliangs.ZdBiao+"18"
//   zyzd=append(zyzd,ziYuanZiDuan4)

//   ziYuanZiDuan5 := map[string]string{}
//   ziYuanZiDuan5[ml2changliangs.BianMa]="ZiYuanLianJie"
//   ziYuanZiDuan5[ml2changliangs.MingCheng]="资源链接"
//   ziYuanZiDuan5[ml2changliangs.ZhiBiao]=ml2changliangs.ZdBiao+"19"
//   zyzd=append(zyzd,ziYuanZiDuan5)

//   ziYuanZiDuan6 := map[string]string{}
//   ziYuanZiDuan6[ml2changliangs.BianMa]="ZiYuanMiaoShu"
//   ziYuanZiDuan6[ml2changliangs.MingCheng]="资源描述"
//   ziYuanZiDuan6[ml2changliangs.ZhiBiao]=ml2changliangs.ZdBiao+"20"
//   zyzd=append(zyzd,ziYuanZiDuan6)

//   ziYuanKu[ml2changliangs.ZiDuans]=zyzd

// }

// func TestXinZengJueSeJieGou(t *testing.T){
//   jueSeKu := map[string]interface{}{}
//   jiChuKu=append(jiChuKu,jueSeKu)
//   jueSeKu[ml2changliangs.MingCheng]="角色"
//   jueSeKu[ml2changliangs.BianMa]="JueSe"
//   jueSeKu[ml2changliangs.ZhuJianBiao]=ml2changliangs.ZjBiao+"2"

//   jszd:=[]map[string]string{}

//   jueSeZiDuan1 := map[string]string{}
//   jueSeZiDuan1[ml2changliangs.BianMa]="JueSeZhuJian"
//   jueSeZiDuan1[ml2changliangs.MingCheng]="角色主键"
//   jueSeZiDuan1[ml2changliangs.ZhiBiao]=ml2changliangs.ZdBiao+"10"
//   jszd=append(jszd,jueSeZiDuan1)

//   jueSeZiDuan2 := map[string]string{}
//   jueSeZiDuan2[ml2changliangs.BianMa]="JueSeBianMa"
//   jueSeZiDuan2[ml2changliangs.MingCheng]="角色编码"
//   jueSeZiDuan2[ml2changliangs.ZhiBiao]=ml2changliangs.ZdBiao+"11"
//   jszd=append(jszd,jueSeZiDuan2)

//   jueSeZiDuan3 := map[string]string{}
//   jueSeZiDuan3[ml2changliangs.BianMa]="JueSeMingCheng"
//   jueSeZiDuan3[ml2changliangs.MingCheng]="角色名称"
//   jueSeZiDuan3[ml2changliangs.ZhiBiao]=ml2changliangs.ZdBiao+"12"
//   jszd=append(jszd,jueSeZiDuan3)

//   jueSeZiDuan4 := map[string]string{}
//   jueSeZiDuan4[ml2changliangs.BianMa]="JueSeZhuangTai"
//   jueSeZiDuan4[ml2changliangs.MingCheng]="角色状态"
//   jueSeZiDuan4[ml2changliangs.ZhiBiao]=ml2changliangs.ZdBiao+"13"
//   jszd=append(jszd,jueSeZiDuan4)

//   jueSeZiDuan5 := map[string]string{}
//   jueSeZiDuan5[ml2changliangs.BianMa]="JueSeShuoMing"
//   jueSeZiDuan5[ml2changliangs.MingCheng]="角色说明"
//   jueSeZiDuan5[ml2changliangs.ZhiBiao]=ml2changliangs.ZdBiao+"14"
//   jszd=append(jszd,jueSeZiDuan5)

//   jueSeKu[ml2changliangs.ZiDuans]=jszd

// }
