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

func zuZhuangBiaoMingShuJu(shuJuKu,bianMa,mingCheng,ZhuJianBiao string)[]map[string]interface{}{
  canShusArr := []map[string]interface{}{}

  canShus := map[string]interface{}{}
  canShus[ml2changliangs.ShuJuKu]=shuJuKu
  canShus[ml2changliangs.ZhuJian]=ml0gongjus.HuoQuId()
  canShus[ml2changliangs.BianMa]=bianMa
  canShus[ml2changliangs.MingCheng]=mingCheng
  canShus[ml2changliangs.ZhuJianBiao]=ZhuJianBiao
  
  canShusArr=append(canShusArr,canShus)
  return canShusArr
}

func TestXinZengJiChuBiao(t *testing.T){
  yongHuBiao := zuZhuangBiaoMingShuJu(ml2changliangs.XM0JICHU,ml2changliangs.Ywb1YongHus,"用户",ml2changliangs.ZjBiao+ml2changliangs.Zf1+ml2changliangs.XiaoXies)
  canShu := ml3moxings.CanShu{}
  canShu.ShuJu = yongHuBiao
  ml5kus.XinZengXm0JiChuShuJu(canShu)
}

// func TestXinZengJueSeZiYuanJieGou(t *testing.T){
//   jueSeZiYuanKu := map[string]interface{}{}
//   jiChuKu=append(jiChuKu,jueSeZiYuanKu)
//   jueSeZiYuanKu[ml2changliangs.MingCheng]="角色资源"
//   jueSeZiYuanKu[ml2changliangs.BianMa]="JueSeZiYuan"
//   jueSeZiYuanKu[ml2changliangs.ZhuJianBiao]=ml2changliangs.ZjBiao+"5"
//   jueSeZiYuanKu[ml2changliangs.ShuJuKu]=ml2changliangs.ZjBiao+"5"
  
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

// func TestXinZengYongHuJieGou(t *testing.T){
//   //根据情况需要新增的数据有哪些？先把三个权限相关的表和两个中间表做了。从整个系统资源来看，除了界面资源之外其他的资源跟角色都没有多大关系。后台管理时有一定关系，就是根据自己这个角色管理其他角色的用户数据。
//   //首先新增一个用户表，由于数据库对象都被放到了map里，所以没必要定义结构体来进行数据库字段的映射，实际上结构体是用来进行算法优化的，而不是像之前那样拿来进行数据库Orm的，以前完全的本末倒置
//   //测试的时候新增一个对象，传入的时候是没有id的，所以需要保证留下一个参数，当正式环境时这个参数会传回界面保存下来留作字段新增时使用，
//   yongHuKu := map[string]interface{}{}
//   yongHuKu[ml2changliangs.MingCheng]="用户"
//   yongHuKu[ml2changliangs.BianMa]="YongHu"
//   yongHuKu[ml2changliangs.ZhuJianBiao]=ml2changliangs.ZjBiao+"1"
//   yhzd:=[]map[string]string{}
  
//   yongHuZiDuan1 := map[string]string{}
//   yongHuZiDuan1[ml2changliangs.BianMa]="YongHuZhuJian"
//   yongHuZiDuan1[ml2changliangs.MingCheng]="用户主键"
//   yongHuZiDuan1[ml2changliangs.ZhiBiao]=ml2changliangs.ZdBiao+"1"
//   yhzd=append(yhzd,yongHuZiDuan1)

//   yongHuZiDuan2 := map[string]string{}
//   yongHuZiDuan2[ml2changliangs.BianMa]="YongHuBianMa"
//   yongHuZiDuan2[ml2changliangs.MingCheng]="用户编码"
//   yongHuZiDuan2[ml2changliangs.ZhiBiao]=ml2changliangs.ZdBiao+"2"
//   yhzd=append(yhzd,yongHuZiDuan2)
  
//   yongHuZiDuan3 := map[string]string{}
//   yongHuZiDuan3[ml2changliangs.BianMa]="YongHuMingCheng"
//   yongHuZiDuan3[ml2changliangs.MingCheng]="用户名称"
//   yongHuZiDuan3[ml2changliangs.ZhiBiao]=ml2changliangs.ZdBiao+"3"
//   yhzd=append(yhzd,yongHuZiDuan3)
  
//   yongHuZiDuan4 := map[string]string{}
//   yongHuZiDuan4[ml2changliangs.BianMa]="YongHuTouXiang"
//   yongHuZiDuan4[ml2changliangs.MingCheng]="用户头像"
//   yongHuZiDuan4[ml2changliangs.ZhiBiao]=ml2changliangs.ZdBiao+"4"
//   yhzd=append(yhzd,yongHuZiDuan4)
  
//   yongHuZiDuan5 := map[string]string{}
//   yongHuZiDuan5[ml2changliangs.BianMa]="YongHuMiMa"
//   yongHuZiDuan5[ml2changliangs.MingCheng]="用户密码"
//   yongHuZiDuan5[ml2changliangs.ZhiBiao]=ml2changliangs.ZdBiao+"5"
//   yhzd=append(yhzd,yongHuZiDuan5)
  
//   yongHuZiDuan6 := map[string]string{}
//   yongHuZiDuan6[ml2changliangs.BianMa]="YongHuYouXiang"
//   yongHuZiDuan6[ml2changliangs.MingCheng]="用户邮箱"
//   yongHuZiDuan6[ml2changliangs.ZhiBiao]=ml2changliangs.ZdBiao+"6"
//   yhzd=append(yhzd,yongHuZiDuan6)
  
//   yongHuZiDuan7 := map[string]string{}
//   yongHuZiDuan7[ml2changliangs.BianMa]="YongHuDiZhi"
//   yongHuZiDuan7[ml2changliangs.MingCheng]="用户地址"
//   yongHuZiDuan7[ml2changliangs.ZhiBiao]=ml2changliangs.ZdBiao+"7"
//   yhzd=append(yhzd,yongHuZiDuan7)
  
//   yongHuZiDuan8 := map[string]string{}
//   yongHuZiDuan8[ml2changliangs.BianMa]="YongHuZhuangTai"
//   yongHuZiDuan8[ml2changliangs.MingCheng]="用户状态"
//   yongHuZiDuan8[ml2changliangs.ZhiBiao]=ml2changliangs.ZdBiao+"8"
//   yhzd=append(yhzd,yongHuZiDuan8)
  
//   yongHuZiDuan9 := map[string]string{}
//   yongHuZiDuan9[ml2changliangs.BianMa]="YongHuTiZhong"
//   yongHuZiDuan9[ml2changliangs.MingCheng]="用户体重"
//   yongHuZiDuan9[ml2changliangs.ZhiBiao]=ml2changliangs.ZdBiao+"9"
//   yhzd=append(yhzd,yongHuZiDuan9)

//   yongHuKu[ml2changliangs.ZiDuans]=yhzd
  
//   canShu := ml3moxings.CanShu{}
//   canShu.ShuJu = []map[string]interface{}{}
  
//   biaoMing := map[string]interface{}{
//   	ml2changliangs.Ceng1: ml2changliangs.Ywb1YongHus,
//   }
//   canShu.ShuJu = append(canShu.ShuJu, biaoMing)
  
//   ziDuans := []map[string]string{}
//   zd := map[string]string{}
  
  
  
//   ziDuans=append(ziDuans,zd)
//   ziDuansKeyMap := map[string]interface{}{
//   	ml2changliangs.Ceng1: ziDuans,
//   }
  
//   canShu.ShuJu = append(canShu.ShuJu, ziDuansKeyMap)
  
//   ret := ml5kus.XinZeng(canShu)

// }
