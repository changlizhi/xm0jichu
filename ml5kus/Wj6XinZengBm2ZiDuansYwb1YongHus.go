package ml5kus

import (
	"xm0jichu/ml0gongjus"
	"xm0jichu/ml2changliangs"
)

func XinZengBm2ZiDuansYwb1YongHus() {
	//查出Bm1BiaoMings中Ywb1YongHus指定的主键
	caoZuoBiao := ml2changliangs.Bm2ZiDuans

	//默认有数据库，在调用此方法之前必须自己保证数据库已存在
	shuJu0 := map[string]interface{}{
		ml2changliangs.CaoZuoKu:   ml2changliangs.XM0JICHU,
		ml2changliangs.CaoZuoBiao: caoZuoBiao,
	}

	ziDuans := []map[string]interface{}{
		// ZuZhuangBIGINT(ZhuJian,"20"),
		// ZuZhuangVARCHAR(BianMa,"50"),
		// ZuZhuangVARCHAR(MingCheng,"50"),
		// ZuZhuangVARCHAR(ZhiBiao,"50"),
		// ZuZhuangVARCHAR(ZhengZe,"50"),
		// ZuZhuangINT(ShiFouZhiDing,"1"),
		// ZuZhuangINT(ShiFouYouXingWei,"1"),
		// ZuZhuangVARCHAR(ChangDu,"5"),

		map[string]interface{}{ //Bm2ZiDuans中字段定义了多少字段就要声明多少字段。
			ml2changliangs.ZhuJian: ml0gongjus.HuoQuId(),
		},
	}
	for _, v := range ziDuans {
		shuJu0[ml2changliangs.ZiDuans] = v
		XinZeng(shuJu0)
	}
	// ml5kus.ShanChuJiChuKu()//测试完成后删除数据库，基础测试在所有的测试方法中都不需要前置测试存在。
}

// func TestXinZengYongHuJieGou(t *testing.T){
//   //根据情况需要新增的数据有哪些？先把三个权限相关的表和两个中间表做了。从整个系统资源来看，除了界面资源之外其他的资源跟角色都没有多大关系。后台管理时有一定关系，就是根据自己这个角色管理其他角色的用户数据。
//   //首先新增一个用户表，由于数据库对象都被放到了map里，所以没必要定义结构体来进行数据库字段的映射，实际上结构体是用来进行算法优化的，而不是像之前那样拿来进行数据库Orm的，以前完全的本末倒置
//   //测试的时候新增一个对象，传入的时候是没有id的，所以需要保证留下一个参数，当正式环境时这个参数会传回界面保存下来留作字段新增时使用，
//   yongHuKu := map[string]interface{}{}

//   yongHuKu[ml2changliangs.ZiDuans]=yhzd

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
