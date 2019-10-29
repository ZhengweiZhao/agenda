# Sevice Computing：CLI 命令行实用程序开发实战Agenda 及 Linux下安装配置Cobra教程
# 1、概述
命令行实用程序并不是都象 cat、more、grep 是简单命令。go 项目管理程序，类似 java 项目管理 maven、Nodejs 项目管理程序 npm、git 命令行客户端、 docker 与 kubernetes 容器管理工具等等都是采用了较复杂的命令行。即一个实用程序同时支持多个子命令，每个子命令有各自独立的参数，命令之间可能存在共享的代码或逻辑，同时随着产品的发展，这些命令可能发生功能变化、添加新命令等。因此，符合 OCP 原则 的设计是至关重要的编程需求。

# 2、安装配置Cobra
## cobra的安装
首先使用命令 `go get -v github.com/spf13/cobra/cobra`下载cobra
但是这里会出提示如下错误：
```javascript
Fetching https://golang.org/x/sys/unix?go-get=1
https fetch failed: Get https://golang.org/x/sys/unix?go-get=1: dial tcp 216.239.37.1:443: i/o timeout
```
这是熟悉的错误，解决办法是需要安装golang的项目依赖test和sys。
首先cd到`$GOPATH/src/golang.org/x `文件夹下，用 `git clone`下载 sys 和 text 项目，命令如下：
```javascript
cd $GOPATH/src/golang.org/x
git clone https://github.com/golang/sys
git clone https://github.com/golang/text
```
可以看到该目录下出现了sys和text两个文件夹如下：
![在这里插入图片描述](https://img-blog.csdnimg.cn/2019102913385595.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3dlaXhpbl80MDM3NzY5MQ==,size_16,color_FFFFFF,t_70)
然后再重新执行命令 `go get -v github.com/spf13/cobra/cobra`下载cobra：
![在这里插入图片描述](https://img-blog.csdnimg.cn/20191029134047891.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3dlaXhpbl80MDM3NzY5MQ==,size_16,color_FFFFFF,t_70)
![在这里插入图片描述](https://img-blog.csdnimg.cn/20191029134143180.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3dlaXhpbl80MDM3NzY5MQ==,size_16,color_FFFFFF,t_70)
然后执行命令`go install github.com/spf13/cobra/cobra`, 安装后在 $GOBIN 下出现了 cobra 可执行程序。
![在这里插入图片描述](https://img-blog.csdnimg.cn/2019102913421380.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3dlaXhpbl80MDM3NzY5MQ==,size_16,color_FFFFFF,t_70)
输入命令`cobra help`也可以看到有帮助文档的输出，说明安装成功。
![在这里插入图片描述](https://img-blog.csdnimg.cn/20191029135004959.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3dlaXhpbl80MDM3NzY5MQ==,size_16,color_FFFFFF,t_70)
## cobra的初始化
在新建的项目文件夹下使用`cobra init --pkg-name agenda`命令可以初始化一个新的项目
![在这里插入图片描述](https://img-blog.csdnimg.cn/20191029134506814.png)
其中agenda是你自定义的项目名称，可以自行修改的，成功后初始化的项目结构如下：
![在这里插入图片描述](https://img-blog.csdnimg.cn/20191029134413796.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3dlaXhpbl80MDM3NzY5MQ==,size_16,color_FFFFFF,t_70)
## cobra添加命令
在项目文件夹下使用`Cobra add`命令可以为你的程序添加新的命令，在本次实验中，我添加的两条命令分别是register和login那么这里添加命令的语句就是：
```javascript
Cobra add register
Cobra add login
```
![在这里插入图片描述](https://img-blog.csdnimg.cn/20191029135143419.png)
![在这里插入图片描述](https://img-blog.csdnimg.cn/20191029135155578.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3dlaXhpbl80MDM3NzY5MQ==,size_16,color_FFFFFF,t_70)![在这里插入图片描述](https://img-blog.csdnimg.cn/20191029135214973.png)
![在这里插入图片描述](https://img-blog.csdnimg.cn/20191029135226903.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3dlaXhpbl80MDM3NzY5MQ==,size_16,color_FFFFFF,t_70)
可以看到在cmd文件夹下除了最开始的root.go还出现了我们新建的两条命令的.go文件，然后在这些文件当中进行编写开发就可以了。
# 3、开发实践
## 需求描述
- 功能需求： 设计一组命令完成 agenda 的管理，要求包括2条命令：
	- agenda register -uUserName –password pass –email=a@xxx.com -f phone 
	>用户注册
	>- 注册新用户时，用户需设置一个唯一的用户名和一个密码。另外，还需登记邮箱及电话信息。
	>- 如果注册时提供的用户名已由其他用户使用，应反馈一个适当的出错信息；成功注册后，亦应反馈一个成功注册的信息。
	- agenda login -uUserName –password pass
	>用户登录
	>- 用户使用用户名和密码登录 Agenda 系统。
	>- 用户名和密码同时正确则登录成功并反馈一个成功登录的信息。否则，登录失败并反馈一个失败登录的信息。
	
- 持久化要求：
	-	使用 json 存储 User 和 Meeting 实体
	- 当前用户信息存储在 curUser.txt 中
- 项目目录
	- cmd ：存放命令实现代码
	- entity ：存放 User 和 Meeting 对象读写与处理逻辑
	- 其他目录 ： 自由添加
- 日志服务
	- 使用 log 包记录命令执行情况
## 代码分析
### register.go
通过init函数获取用户输入的命令行的不同参数，包括用户名，密码，邮箱，电话号码四种string类型的变量。
```javascript
rootCmd.AddCommand(registerCmd)
	registerCmd.Flags().StringP("user", "u", "Anonymous", "help message for username")
	registerCmd.Flags().StringP("password", "p", "123", "help message for password")
	registerCmd.Flags().StringP("email", "e", "123@xxx.com", "help message for email")
	registerCmd.Flags().StringP("phone", "f", "13611263068", "help message for phone")
```
读取用户信息后，由于我们读取用户信息使用到的是 json 的框架。
>JSON (JavaScript对象表示法)是一种简单的数据交换格式。在语法上，它类似于JavaScript的对象和列表。以文字为基础，具有自我描述性且易于让人阅读。尽管JSON是JavaScript的一个子集，但JSON是独立于语言的文本格式，并且采用了类似于C语言家族的一些习惯。JSON与XML最大的不同在于XML是一个完整的标记语言，而JSON不是。JSON由于比XML更小、更快，更易解析，以及浏览器的內建快速解析支持，使得其更适用于网络数据传输领域。

Go的JSON包中有如下函数:
```javascript
// json.go
package main
import (
    "encoding/json"
    "fmt"
)
type Server struct {
    ServerName string
    ServerIP   string
}
type Serverslice struct {
    Servers []Server
}
func main() {
    var s Serverslice
    str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},
            {"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
 
    json.Unmarshal([]byte(str), &s)
    fmt.Println(s)
    fmt.Println(s.Servers[0].ServerIP)
 
}
```
利用go中的json包我们可以实现将结构体序列化，我们用函数marshal来encode JSON data ，将用户输入转换成json支持的格式以完成对json文件的读写操作，具体关于go中json的用法可以参考这篇 [博客](https://blog.csdn.net/tennysonsky/article/details/79084919)。

Run 匿名回调函数中我们需要遍历所有用户的名字，检查是否与当前的用户信息冲突。若不冲突则可以注册，返回注册信息及注册成功，否则则返回返回错误信息，然后将命令信息记录在log文件中。以下为run匿名回调函数：
```javascript
	Run: func(cmd *cobra.Command, args []string) {
		//实例化
		str := userinfo{
			Name: "",
			Password: "",
			Email: "",
			Phone: "",
		}
		
		username, _:=cmd.Flags().GetString("user")
		password, _:=cmd.Flags().GetString("password")
		email, _:=cmd.Flags().GetString("email")
		phone, _:=cmd.Flags().GetString("phone")
		fmt.Println("register name : "+username)
		fmt.Println("password : "+password)
		fmt.Println("email : "+email)
		fmt.Println("phone : "+phone)
		fmt.Println("Register success!")
		str.Name=username
		str.Password=password
		str.Email=email
		str.Phone=phone
		filename:="./entity/log.log"
		logfile, errr:=os.OpenFile(filename,os.O_RDWR|os.O_APPEND,7)
		if(errr!=nil){
			fmt.Println("openfile fail")
		}
		defer logfile.Close()
		debuglog:=log.New(logfile,"",log.LstdFlags)
	
		//判断是否已有已注册同名用户来判断注册是否成功
		if checkuser(username)==true {
			fmt.Println("username exist, create account fail")
			debuglog.Println("register: username "+username+" exist, create account fail")
		} else {
			debuglog.Println("register: username: "+username+" password: "+password+" email: "+email+" phone:"+phone);
			savecuruser(str)
			input := readinfo()
			input.Id=append(input.Id,str)
			data, _:= json.Marshal(input)
			saveinfo(data)
		}
	},
}
```
在register.go文件中我们还实现了一些功能函数以完成上述逻辑框架，包括保存信息，读取信息等等，用来读取整个data.json文件当中的内容以及将用户新的输入信息保存到data.json当中。

读取之后，我们利用checkuser来对全部的信息进行一个遍历比对，来查询是否用户新注册的用户名已经存在。
```javascript
func checkuser(username string) bool{
	input := readinfo()
	l :=len(input.Id)
	for i:=0;i<l;i++ {
		if(input.Id[i].Name==username){
			return true
		}
	}
	return false
}
```
完整代码见GitHub。
### login.go
login的代码要简单很多，以下为run回调函数，这里的核心其实就是用户输入的用户名密码是否存在且是否正确，调用了一个checkpasswd函数返回check值。
```javascript
Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("login called")
		username, _:=cmd.Flags().GetString("username")
		password, _:=cmd.Flags().GetString("password")
		check := checkpasswd(username, password)
		filename:="./entity/log.log"
		logfile, errr:=os.OpenFile(filename,os.O_RDWR|os.O_APPEND,7)
		if(errr!=nil){
			fmt.Println("openfile fail")
		}
		defer logfile.Close()
		debuglog:=log.New(logfile,"",log.LstdFlags)
		if check==true {
			fmt.Println("login success")
			debuglog.Println("login: username: "+username+" login success")
		} else {
			fmt.Println("login fail")
			debuglog.Println("login: username: "+username+" login fail")
		}
	},
```
主要就是进行两个判断：一个是判断用户是否存在，第二个是判断密码是否匹配，以此来进行错误反馈。
```javascript
func checkpasswd(username string, password string) bool{
	input := readinfo()
	l :=len(input.Id)
	for i:=0;i<l;i++ {
		if(input.Id[i].Name==username && input.Id[i].Password==password){
			return true
		}
	}
	return false
}
```
### root.go
root文件参考了这篇博客的内容，详情请移步[参考文档：Golang: Cobra命令行参数库的使用](https://www.cnblogs.com/welhzh/p/8962489.html)
代码见GitHub。
##  运行结果测试
测试环境：centos
![在这里插入图片描述](https://img-blog.csdnimg.cn/20191029141949255.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3dlaXhpbl80MDM3NzY5MQ==,size_16,color_FFFFFF,t_70)
- register

成功注册：
![在这里插入图片描述](https://img-blog.csdnimg.cn/20191029141610734.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3dlaXhpbl80MDM3NzY5MQ==,size_16,color_FFFFFF,t_70)
![在这里插入图片描述](https://img-blog.csdnimg.cn/20191029141659461.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3dlaXhpbl80MDM3NzY5MQ==,size_16,color_FFFFFF,t_70)
不成功注册，用户名重复：
![在这里插入图片描述](https://img-blog.csdnimg.cn/20191029141856465.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3dlaXhpbl80MDM3NzY5MQ==,size_16,color_FFFFFF,t_70)
- login
![在这里插入图片描述](https://img-blog.csdnimg.cn/20191029141622771.png)
![在这里插入图片描述](https://img-blog.csdnimg.cn/20191029141713371.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3dlaXhpbl80MDM3NzY5MQ==,size_16,color_FFFFFF,t_70)
- log
![在这里插入图片描述](https://img-blog.csdnimg.cn/20191029144148778.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3dlaXhpbl80MDM3NzY5MQ==,size_16,color_FFFFFF,t_70)
- curUser
![在这里插入图片描述](https://img-blog.csdnimg.cn/20191029143948291.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3dlaXhpbl80MDM3NzY5MQ==,size_16,color_FFFFFF,t_70)
- data.json
![在这里插入图片描述](https://img-blog.csdnimg.cn/20191029144044465.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3dlaXhpbl80MDM3NzY5MQ==,size_16,color_FFFFFF,t_70)

至此，实验完成。


