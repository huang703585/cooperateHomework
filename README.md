# cooperateHomework
综合设计

目前是做到能运行的程度。由于做数据库那边还没有发来接口，所以每次调用功能都只会返回同样的结果。目前稳定性很差，所以近期会不断修改（但接口不变）。

server.go 是服务器端
server.exe 是服务器端的可执行文件
csInterface.java 实现接口类
csInterface.class 是编译后的接口类
client.java 是我做测试用的客户端
由于在写接口类时我没有写包组织相关的内容，所以请把csInterface.class和客户端程序放在同一文件夹里。

接口提供的方法：
1.boolean SignIn(String username, String password)
登录功能，目前总返回true
2.boolean SignUp(String username, String password)
注册功能，目前总返回true
请保证在成功登录或注册后才使用以下四个功能：
3.String[] getAnnouncement(String lasttime)
目前不可用，参数lasttime为收到的上一个公告的时间
4.String getMaintain()
报修功能，目前总返回“server Maintain”
5.float getAccount()
费用查询功能，目前总返回1.7
6.String getContact()
联系功能，目前总返回“server Contact”
