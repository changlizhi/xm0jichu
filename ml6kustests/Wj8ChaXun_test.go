package ml6kustests

import (
	"testing"
	"xm0jichu/ml2changliangs"
	"xm0jichu/ml3moxings"
	"xm0jichu/ml5kus"
)

func TestChaXun(t *testing.T) {
	biao := ml2changliangs.JiChuBiao[ml2changliangs.Bm1BiaoMings].(map[string]interface{})

	bm1BiaoMingsZiDuans := biao[ml2changliangs.ZiDuans].([]map[string]interface{})

	canShuXinZeng := ml3moxings.CanShu{
		ShuJu: []map[string]interface{}{
			biao,
		},
	}
	ml5kus.ChuangJianBiao(canShuXinZeng)
	//初始化一些数据进去
	ml5kus.XinZengBm1BiaoMings()

	chaXunZiDuans := []interface{}{}

	for _, v := range bm1BiaoMingsZiDuans {
		chaXunZiDuans = append(chaXunZiDuans, v[ml2changliangs.ZiDuanMing])
	}

	canShuChaXun := ml3moxings.CanShu{
		ShuJu: []map[string]interface{}{
			map[string]interface{}{
				ml2changliangs.CaoZuoKu:   ml2changliangs.XM0JICHU,
				ml2changliangs.CaoZuoBiao: ml2changliangs.Bm1BiaoMings,
				ml2changliangs.ZiDuans:    chaXunZiDuans,
				ml2changliangs.TiaoJianHeZhis: map[string]interface{}{
					ml2changliangs.BianMa: ml2changliangs.Ywb1YongHus, //查询这个条件的数据
				},
			},
		},
	}
	ml5kus.ChaXun(canShuChaXun)
}
