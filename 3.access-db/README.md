#访问数据库
go没有内置的驱动支持任何数据库,但是Go定义了database/sql接口


##数据库驱动类,定义一些接口

- driver.Driver
- driver.Conn
- driver.Stmt
- driver.Execer
- driver.Result
- driver.Rows


##使用mysql数据库
通过上面的代码我们可以看出，Go操作Mysql数据库是很方便的。

关键的几个函数我解释一下：

sql.Open()函数用来打开一个注册过的数据库驱动，go-sql-driver中注册了mysql这个数据库驱动，第二个参数是DSN(Data Source Name)，它是go-sql-driver定义的一些数据库链接和配置信息。它支持如下格式：

user@unix(/path/to/socket)/dbname?charset=utf8
user:password@tcp(localhost:5555)/dbname?charset=utf8
user:password@/dbname
user:password@tcp([de:ad:be:ef::ca:fe]:80)/dbname
db.Prepare()函数用来返回准备要执行的sql操作，然后返回准备完毕的执行状态。

db.Query()函数用来直接执行Sql返回Rows结果。

stmt.Exec()函数用来执行stmt准备好的SQL语句

我们可以看到我们传入的参数都是=?对应的数据，这样做的方式可以一定程度上防止SQL注入。