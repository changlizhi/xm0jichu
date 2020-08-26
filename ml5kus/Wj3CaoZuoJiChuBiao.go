package ml5kus

import (
	"xm0jichu/ml2changliangs"
)

func ChuangJianJiChuBiao() {
	for _, v := range ml2changliangs.HuoQuJiChuBiaoJieGou(ml2changliangs.FhKongZiFu) {
		ChuangJianBiao(v.(map[string]interface{}))
	}
}
