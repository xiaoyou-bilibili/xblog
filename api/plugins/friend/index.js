// 使用博客系统提供的接口
const router =xBlog.router
const database =xBlog.database
const tools = xBlog.tools
const mail = xBlog.mail
const widget = xBlog.widget
// 一些字段
const dbFriend = "friend"
// 管理员邮箱
const keyAdminEmail = "site_email"
// api地址
const keyServer = 'site_api_server'
// 友链名字
const keyFriendName = "friend_name"
const keyFriendDec = "friend_dec"
const keyFriendLink = "friend_link"
const keyFriendAvatar = "friend_avatar"
// 背景
const keyBackground = "friend_img"

// 获取所有友链
router.registerRouter("GET","",function(context){
    let db = database.newDb(dbFriend)
    db.FindMany({
        "sort": { "_id":-1 },
        "filter": { "status":1 }
    },function (err,data){
        if (err == null){
            let response = []
            // 遍历查询到的数据
            data.forEach(function (item){
                // 只放入我们需要的值
                response.push({
                    name : item.name,
                    url : item.url,
                    avatar : item.avatar,
                    dec : item.description
                })
            })
            // 返回友链数据
            router.response.ResponseOk(context,response)
        } else {
            router.response.ResponseServerError(context,"获取友链数据失败")
        }
    })
})
// 用户申请友链
router.registerRouter("POST","",function (context){
    // 先获取请post请求的数据
    let friend = router.getPostJson(context)
    // 验证关键字段是否为空
    if (tools.verifyField(friend.name) && tools.verifyField(friend.email) && tools.verifyEmail(friend.email)){
        // 设置需要插入的数据
        let data = {
            url: friend.site,
            name: friend.name,
            avatar: friend.avatar,
            description: friend.dec,
            email: friend.email,
            status: 0,
            application_time: new Date(),
        }
        // 插入数据
        let db = database.newDb(dbFriend)
        db.InsertOne(data,function (err,res){
            // 判断是否插入成功
            if (err==null){
                // 为了兼容旧版接口，我们还需要返回值
                data.id = res.InsertedID
                // 顺便发送一份邮件给管理员
                let body = "名字:" + data.name + "<br>站点地址:" + data.url +"<br>用户邮箱:" + data.email + "<br>描述:" + data.description
                mail.sendMail([tools.getSetting(keyAdminEmail)],"友链申请",body)
                router.response.ResponseCreated(context,data)
            } else {
                router.response.ResponseServerError(context)
            }
        })
    } else {
        router.response.ResponseBadRequest(context,"请检查名字、网址。邮箱是否填写并正确！")
    }
})
// 友人帐界面
widget.addPage({
    background: tools.getSetting(keyBackground),
    file:"index.html",
    headMeta: {
        title: "友人帐",
    },
    css: ["element"],
    script: ["vue","element","jquery"],
    url: "",
    full: false,
    side: false
},function (){
    // 友链设置
    return {
        name: tools.getSetting(keyFriendName),
        dec: tools.getSetting(keyFriendDec),
        link: tools.getSetting(keyFriendLink),
        avatar: tools.getSetting(keyFriendAvatar),
    }
})

// 添加友人帐设置界面
widget.addSetting("友链管理",2,"friend")

widget.addSetting("友链设置",1,[
    {title:"友链名字",type: "input",key: keyFriendName},
    {title:"友链介绍",type: "input",key: keyFriendDec},
    {title:"友链地址",type: "input",key: keyFriendLink},
    {title:"友链头像",type: "upload",key: keyFriendAvatar}
])

// 友链管理功能
// 获取所有友链
router.registerAdminRouter("GET","",function (context){
    // 获取基本参数
    let id = tools.str2int(context.Query('page'))
    let size = tools.str2int(context.Query('page_size'))
    let search = context.Query('search_type')
    let key = context.Query('search_key')
    if (id===0) { id=1 }
    if (size===0) { size=10 }
    // 设置关键词过滤
    let filter = {}
    if (search!=="" && key!==""){
        // 这里如果是友链，需要额外判断
        if (search === "status") {
            filter[search] = tools.str2int(key)
        } else {
            filter[search] = database.regex(key)
        }
    }
    // 开始搜索
    let db = database.newDb(dbFriend)
    db.Paginate({filter,sort:{_id:-1}},id,size,function (err,page,total,data){
        if (err==null){
            let response = {
                total_num: total,
                total: page,
                current: id,
                contents: []
            }
            data.forEach(function (item){
                response.contents.push({
                    id: item._id,
                    name: item.name,
                    url: item.url,
                    status: item.status,
                    email: item.email,
                    description: item.description,
                    avatar: item.avatar
                })
            })
            router.response.ResponseOk(context,response)
        } else {
            router.response.ResponseServerError(context)
        }
    })
})
// 添加友链
router.registerAdminRouter("POST","",function (context) {
    // 先获取请post请求的数据
    let friend = router.getPostJson(context)
    // 验证关键字段是否为空
    if (tools.verifyField(friend.name) && tools.verifyField(friend.url)) {
        let data = {
            name: friend.name,
            url: friend.url,
            avatar: friend.avatar,
            description: friend.description,
            email: friend.email,
            application_time: new Date(),
            status: tools.str2int(friend.name),
        }
        // 插入数据
        database.newDb(dbFriend).InsertOne(data,function (err,res){
            // 判断是否插入成功
            if (err==null){
                // 为了兼容旧版接口，我们还需要返回值
                data.id = res.InsertedID
                router.response.ResponseCreated(context,data)
            } else {
                router.response.ResponseServerError(context)
            }
        })
    } else {
        router.response.ResponseBadRequest(context,"请检查名字、网址是否填写并正确！")
    }
})
// 更新友链
router.registerAdminRouter("PUT","/:id",function (context){
    // 获取id
    let id = context.Param("id")
    if (id!==""){
        let param = router.getPostJson(context)
        let set = {}
        // 判断不同的参数，设置不同的set
        if (tools.verifyField(param.name)) set.name = param.name
        if (tools.verifyField(param.url)) set.url = param.url
        if (tools.verifyField(param.status)) set.status = tools.str2int(param.status)
        if (tools.verifyField(param.avatar)) set.avatar = param.avatar
        if (tools.verifyField(param.email)) set.email = param.email
        if (tools.verifyField(param.description)) set.description = param.description
        // 获取id
        let filter = {"_id":{"$in":tools.string2objetIdArray(id,",")}}
        // 更新数据
        database.newDb(dbFriend).UpdateMany({update: {"$set":set},filter},function (err,res){
            if (err==null){
                router.response.ResponseCreated(context,param)
            } else {
                router.response.ResponseServerError(context)
            }
        })
    } else {
        router.response.ResponseBadRequest(context,"id格式错误")
    }
})
// 删除友链
router.registerAdminRouter("DELETE","/:id",function (context){
    database.adminDeleteObject(context,dbFriend,"_id")
})
// 发送友链通知
router.registerAdminRouter("PUT","/:id/notification",function (context){
    // 获取id
    let id = tools.str2objectId(context.Param("id"))
    if (id.IsZero()){
        router.response.ResponseBadRequest(context,"id格式错误")
    } else {
        // 查找用户的邮箱信息
        database.newDb(dbFriend).FindOne({ filter: {_id:id}  },function (err,data){
            if (err==null){
                let body = "你好，你的友链已成功通过审核，<a href='" + tools.getSetting(keyServer) + "/more/friend'>点击这里</a>查看所有友链。<br>本邮件由系统直接发出，请勿直接回复"
                mail.sendMail([data.email],"友链申请成功通知",body)
                router.response.ResponseCreated(context,{id:id.String()})
            } else {
                router.response.ResponseNotFound(context,"没有找到用户信息")
            }
        })
    }
})
