# cooperateHomework
综合设计
7.11   
7.12 bingo!

这回连接上了数据库，如果安装了sql server的话可以跑整个后端了
目前还存在的问题：不要使用中文，目前不支持utf8
数据库默认名为“homework”，请不要有重名数据库，否则会被替换。
每运行一次服务器程序都会重置一遍数据库。


database文件夹存放数据库建立和初始化数据和数据库接口说明
server文件夹是服务器端源代码文件


server.exe 是服务器端的可执行文件
csInterface.java 实现接口类
csInterface.class 是编译后的接口类
client.java 是我做测试用的客户端
由于在写接口类时我没有写包组织相关的内容，所以请把csInterface.class和客户端程序放在同一文件夹里。


接口提供的方法：
1.boolean SignIn(String username, String password)
登录功能，如果数据库不存在username则返回false，其余都会返回true
2.boolean SignUp(String username, String password)
注册功能，如果数据库存在username则返回false,其余都会返回true
请保证在成功登录或注册后才使用以下四个功能：
3.String[] getAnnouncement(string lasttime)
公告功能，lasttime指接收的上个公告的时间，将返回在这个时间后发布的所有公共，每个数组元素代表一个公告。
4.String getMaintain(string username)
报修功能，将返回报修信息，同时给相应用户的费用增加50.1。
5.float getAccount(string username)
费用查询功能。
6.String getContact(string username)
联系功能，每调用一次会增加一个公告。
