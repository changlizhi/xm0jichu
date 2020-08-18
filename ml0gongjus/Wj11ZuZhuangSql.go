package ml0gongjus

import(
"xm0jichu/ml2changliangs"
)

func ZuZhuangBIGINT(mingCheng,kuandu string)map[string]interface{}{
  ret := map[string]interface{}{}
	ret[ml2changliangs.ZiDuanMing] = mingCheng
	ret[ml2changliangs.LeiXing] = ml2changliangs.BIGINT
	ret[ml2changliangs.ChangDu] = kuandu
	ret[ml2changliangs.MoRenZhi] = ml2changliangs.Zf0
  return ret
}

func ZuZhuangINT(mingCheng,kuandu string)map[string]interface{}{
  ret := map[string]interface{}{}
	ret[ml2changliangs.ZiDuanMing] = mingCheng
	ret[ml2changliangs.LeiXing] = ml2changliangs.INT
	ret[ml2changliangs.ChangDu] = kuandu
	ret[ml2changliangs.MoRenZhi] = ml2changliangs.Zf0
  return ret
}

func ZuZhuangVARCHAR(mingCheng,ziFuShu string)map[string]interface{}{
  ret := map[string]interface{}{}
  ret[ml2changliangs.ZiDuanMing] = mingCheng
  ret[ml2changliangs.LeiXing] = ml2changliangs.VARCHAR
  ret[ml2changliangs.ChangDu] = ziFuShu
  ret[ml2changliangs.MoRenZhi] = ml2changliangs.FhDanYinHao + ml2changliangs.HFXXIAOXIE + ml2changliangs.FhDanYinHao
  return ret
}
