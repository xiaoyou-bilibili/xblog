// 为了不让index.js报错同时进行提示，我们使用一个js文件来模拟
// XBlog对象
const xBlog = {
    // 和路由有关的配置
    router: {
        // 注册路由（请求方式，对应路径，回调函数）
        registerRouter(method,path,RouterContext){},
        // 注册管理路由
        registerAdminRouter(method,path,RouterContext){},
        // 获取post请求的json数据
        getPostJson(context){},
        // 返回结果
        response: {
            // 200返回数据成功
            ResponseOk(context,data){},
            // 201创建数据成功
            ResponseCreated(context,data){},
            // 204删除数据成功(删除数据一般不需要返回数据)
            ResponseNoContent(context){},
            // 400错误，一般是用户输入的参数有问题
            ResponseBadRequest(context,msg = ''){},
            // 401错误，用户没有权限
            ResponseUnauthorized(context,msg = ''){},
            // 403错误，用户禁止访问
            ResponseForbidden(context,msg = ''){},
            // 404错误，没有这个资源
            ResponseNotFound(context,msg = ''){},
            // 422错误，用户在创建对象的时候发生错误
            ResponseUnProcessEntity(context,msg = ''){},
            // 500错误，服务器错误
            ResponseServerError(context,msg = ''){},
            // 返回html对象
            ResponseHtml(context,html){},
        }
    },
    // 数据库操作接口
    database: {
        // 初始化数据库链接（集合名），数据库操作对象
        newDb: (collection)=>DatabaseStruct,
        // 正则搜索
        regex(str){},
        // 管理员删除object数据
        adminDeleteObject(context,collection,field){}
    },
    // 和网络请求相关的接口，比如发送get或者post请求
    net: {
        // 发送get请求（url，头部信息，回调函数）
        get(url,head,callback){},
        // 发送post请求 (url，头部信息，post参数，回调函数)
        post(url,head,pram,callback){}
    },
    // 网络爬虫相关的接口，这里主要是封装一个colly对象
    spider: {
        // 初始化colly对象
        init:(head)=>SpiderStruct
    },
    // 定时任务相关的接口，可以执行定时任务
    cron: {
        // 启动一个定时任务 定时任务参考 https://cron.qqe2.com/
        start(spec,callback){}
    },
    // 文件操作的相关接口，用于操作文件
    file: {
        read(filename,callback){}
    },
    // 当执行某一个事件的时候，我们进行挂载
    action: {

    },
    // 对返回的数据进行过滤操作
    filter: {

    },
    // 用于发送邮件
    mail: {
        // 发送邮件 （发送的对象（字符串数组），主题，内容）
        sendMail(mailTo,subject,body){}
    },
    // tool里面提供了一些常用的工具包
    tools: {
        // 打印日志
        log(data){},
        // 获取网站设置
        getSetting(key){},
        // 对网站进行设置
        setSetting(key,value){},
        // 获取管理员插件设置
        getAdminPluginSetting(option){},
        // 存储key值
        setKey(key,value){},
        // 获取key值
        getKey(key){},
        // 寻找匹配的正则表达式
        findMatch(reg,content){},
        // 替换字段避免XSS攻击
        replaceXSS(content){},
        // 验证字段（这里可以验证字段是否存在，并且如果为字符串类型的话，判断是否不为空）
        verifyField(field){},
        // 验证电子邮箱是否正确
        verifyEmail(mail){},
        // 替换评论表情
        changeCommentSmile(content){},
        // 获取随机头像
        getRandomAvatar(){},
        // 时间转字符串（当前时间，是否显示小时）
        time2String(time,showHour){},
        // 字符串转时间
        string2time(time,showHour){},
        // 字符串转int
        str2int(n){},
        // 字符转bool
        str2bool(n){},
        // 字符串转objectId数组
        string2objetIdArray(data,sep){},
        // 字符串转objectId
        str2objectId:(id)=>objectID,
        // 字符串替换
        strReplace(s,old,newS,n){},
        // 获取B站个人信息
        getBiliPersonInfo(uid,cookie){}
    },
    // 用于静态资源映射
    static: {
        // 注册一个静态路径
        staticPath(name,path){},
        // 注册一个静态文件
        staticFile(name,path){}
    },
    // 全局变量
    global: {
        // 插件名字
        PluginName: ""
    },
    // 页面部件有关的函数
    widget: {
        // 添加侧边栏卡片
        addSide(title,file,content,debug)  {},
        // 添加页面
        addPage(option,call){},
        // 添加设置界面
        addSetting(title,type,setting){}
    }
}
// 路由的context对象
xBlog.router.registerRouter.prototype = {
    // 获取query对象
    Query(key){},
    // 获取URl的头部信息
    Param(key){},
    // 获取form格式
    PostForm(key){}
}
// 数据库newDB的对象
const DatabaseStruct = {
    // 获取单条记录
    FindOne(option,callback){},
    // 获取多条记录
    FindMany(option,callback){},
    // 获取记录条数
    GetCount(filter){},
    // 分页查询(查询条件，当前第几页，一页多少条，回调函数)
    Paginate(option,now,limit,callback){},
    // 插入一条数据
    InsertOne(data,callback){},
    // 插入数据，ID自增
    InsertOneIncrease(data,key,callback){},
    // 更新一条数据
    UpdateOne(option,callback){},
    // 更新多条数据
    UpdateMany(option,callback){},
    // 删除一条数据
    DeleteOne(filter,callback){},
    // 删除多条数据
    DeleteMany(filter,callback){}
}
// 爬虫初始对象
const SpiderStruct= {
    // 解析html数据
    OnHTML(goquerySelector,htmlCallback=htmlCallback){},
    // 返回网站
    Visit(url){}
}
// 爬虫的回调函数
const htmlCallback = {
    // 循环遍历
    ForEach(goquerySelector,foreachCallback){},
    // 获取子元素
    ChildText(goquerySelector){},
    // 获取子元素属性
    ChildAttr(goquerySelector,attrName){},
    Request: {
        Visit(url) {}
    }
}
// objectID
const objectID = {
    // objectId是否为0
    IsZero(){},
    // 获取string值
    String(){}
}