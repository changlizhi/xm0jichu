package ml5kus

import (
	"xm0jichu/ml0gongjus"
	"xm0jichu/ml2changliangs"
)

func XinZengBm1BiaoMings() {
	//数据库会初始化所以不需要自己操作，但是表不会，所以还要创建数据表之后再进行新增，新增完之后删除数据库
	shuJu0 := map[string]interface{}{
		CaoZuoKu:   XM0JICHU,
		CaoZuoBiao: Bm1BiaoMings,
	}
	ziDuans := []map[string]interface{}{
		// ZuZhuangBIGINT(ZhuJian,"20"),
		// ZuZhuangVARCHAR(BianMa,"50"),
		// ZuZhuangVARCHAR(MingCheng,"50"),
		// ZuZhuangVARCHAR(ZhuJianBiao,"50"),
		// ZuZhuangVARCHAR(ShuJuKu,"50"),
		map[string]interface{}{ //Bm1BiaoMings中字段定义了多少字段就要声明多少字段。
		},
		map[string]interface{}{ //Bm1BiaoMings中字段定义了多少字段就要声明多少字段。
		},
		map[string]interface{}{ //Bm1BiaoMings中字段定义了多少字段就要声明多少字段。
		},
	}
	for _, v := range ziDuans {
		shuJu0[ZiDuans] = v
		XinZeng(shuJu0)
	}
}
