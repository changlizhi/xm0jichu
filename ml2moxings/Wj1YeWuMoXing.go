package ml2moxings

//这两个对象将贯穿全部方法的参数和返回

type CanShu struct{
//XuHao的意思就是调用的业务序号，比如注册或者登录，然后这个交易码可以直接在后台用map指定到业务层的方法map中
  Jwt string
  XuHao string
  ShuJu []map[string]interface{}//goframe无法解析map型数据只能设置为interface，但实际使用时里面仍然是map结构
}

type FanHui struct{
  ZhuangTai string
  ShuJu []map[string][]map[string]string
}
