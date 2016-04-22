## beego admin

基于beego，jquery easyui ,
16年04月22日更新
#####
数据库目前仅支持MySQL,PostgreSQL,sqlite3,后续会添加更多的数据库支持。

数据库的配置信息需要填写，程序会根据配置自动建库
MySQL数据库链接信息
```
db_host = localhost
db_port = 3306
db_user = root
db_pass = root
db_name = admin
db_type = mysql
```
postgresql数据库链接信息
```
db_host = localhost
db_port = 5432
db_user = postgres
db_pass = postgres
db_name = admin
db_type = postgres
db_sslmode=disable
```
sqlite3数据库链接信息
```
###db_path 是指数据库保存的路径，默认是在项目的根目录
db_path = ./
db_name = admin
db_type = sqlite3
```
把以上信息配置成你自己数据库的信息。

还有一部分权限系统需要配置的信息
```
sessionon = true
rbac_role_table = role
rbac_node_table = node
rbac_group_table = group
rbac_user_table = user
#admin用户名 此用户登录不用认证
rbac_admin_user = admin

#默认不需要认证模块
not_auth_package = public,static
#默认认证类型 0 不认证 1 登录认证 2 实时认证
user_auth_type = 1
#默认登录网关
rbac_auth_gateway = /public/login
#默认模版
template_type=easyui
```
以上配置信息都需要加入到app.conf文件中。

###复制静态文件

最后还需要把js，css，image，tpl这些文件复制过来。
```bash
$ cd $GOPATH/src/hello
$ cp -R ../github.com/beego/admin/static ./
$ cp -R ../github.com/beego/admin/views ./

```
###编译项目

全部做好了以后。就可以编译了,进入hello目录
```
$ go build
```
首次启动需要创建数据库、初始化数据库表。
```bash
$ ./admin -syncdb
```
好了，现在可以通过浏览器地址访问了[`http://localhost:8080/`](http://localhost:8080/)

默认得用户名密码都是admin

# beegoadmin
