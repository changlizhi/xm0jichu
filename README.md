# 用户系统

此系统提供一个基本的用户增删改查功能，同时提供其他模块是否允许操作的校验服务，提供校验服务时根据请求的资源分析出这些资源所需要的角色，再根据这些角色反查出是否允许使用该资源，如果不允许使用则校验失败，如果校验成功，则继续用用户往下查询这条数据是否属于该用户，如果属于该用户所拥有的权限范围则可以被该用户进行增删改查？
我感觉到有一点问题，如果基于用户角色访问权限的数据直接可以获取到的话就不用这么麻烦了。也就是说其实基于角色的资源控制只是为了展示哪些功能是可以被用户使用的一种简化方式，按照现有的数据访问格式，数据必然有一个字段是用户编码或者主键，那么这个数据已经是行级限制了，根本不需要进行角色查询控制。比如某个人创建了一个活动，这个活动必然有个所属人字段，那么根据这个字段是可以完全确定到底能不能被用户增删改的也就没必要再校验一次角色问题。这其实是公共接口问题。

上传文件时如果没有当前用户编号是无法确定上传上来的文件到底给谁的。

































