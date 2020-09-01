package ml5kus

func init() {
	huanCunSuoYouJieGou() //初始化时用test作为初始化链接
}

var (
	suoShiLi sync.Once
	chi      = map[string]interface{}{}
)


//这个方法唯一要注意的就是表和字段增到多大的时候会大到内存无法支持,当需要再次新增字段时只是在用户增加了字段的情况下需要再次查库更新
func huanCunSuoYouJieGou(){
	suoShiLi.Do(func() { //这里需要把已存在的都纳入进来，所以需要新建一个配置文件，这个配置文件用go写成
    gengXinXiTongJieGou()
	})
}

func gengXinXiTongJieGou(){
  
}

func HuoQuXiTongJieGou(canShu string)map[string]interface{}{
  //直接传入系统名，如果系统名在常量文件定义过则继续，其实系统表也可以放到数据库。。。
  
  return chi
}

