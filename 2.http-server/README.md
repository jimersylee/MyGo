#Go如何使得Web工作
使用net/http包就可以搭建web服务器

##web工作方式的几个概念
Request:用户请求的信息,包括param,cookie,url
Response:服务器需要反馈给客户端的信息
Conn:用户的每次请求连接
Handle:处理请求和生成返回信息的处理逻辑



##Go的http包详解
Go的http有连个核心功能,Conn,ServeMux

###Conn的goroutine
与一般http服务器不同,如Nginx使用单工作进程,NodeJs使用单线程,Go使用协程(goroutine)来处理Conn的读写事件,相互独立,不会阻塞
Go在等待客户端的请求里面是这样写的
```
c,err:=srv.newConn(rw)
if err!=nil{
    continue
}
go c.serve()
```