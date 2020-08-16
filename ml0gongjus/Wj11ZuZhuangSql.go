package ml0gongjus

import(
"xm0jichu/ml2changliangs"
)

func ZuZhuang20BIGINT(mingCheng string)map[string]interface{}{
  ret := map[string]interface{}{}
	ret[ml2changliangs.BiaoMing] = mingCheng
	ret[ml2changliangs.LeiXing] = ml2changliangs.BIGINT
	ret[ml2changliangs.ChangDu] = "20"
	ret[ml2changliangs.MoRenZhi] = ml2changliangs.Zf0
  return ret
}
func ZuZhuang50VARCHAR(mingCheng string)map[string]interface{}{
  ret := map[string]interface{}{}
  ret[ml2changliangs.BiaoMing] = mingCheng
  ret[ml2changliangs.LeiXing] = ml2changliangs.VARCHAR
  ret[ml2changliangs.ChangDu] = "50"
  ret[ml2changliangs.MoRenZhi] = ml2changliangs.FhDanYinHao + ml2changliangs.HFXXIAOXIE + ml2changliangs.FhDanYinHao
  return ret
}

