// 使用博客系统提供的接口
const router =xBlog.router
const database =xBlog.database
const widget =xBlog.widget
// const file =xBlog.static
const tools =xBlog.tools

// 一些字段
const dbSponsor = "donate"
const keyAlipay = "alipay"
const keyWechat = "wechat"
const keyBackground = "donate_img"

// file.staticPath("images","images")

// 获取所有赞助
router.registerRouter("GET","",function(context){
    let db = database.newDb(dbSponsor)
    db.FindMany({
        "sort": { "_id":-1 }
    },function (err,data){
        if (err == null){
            let response = []
            // 遍历查询到的数据
            data.forEach(function (item){
                // 只放入我们需要的值
                response.push({
                    nickname : item.nickname,
                    donate : item.amount,
                    comment : item.comment
                })
            })
            // 返回赞助数据
            router.response.ResponseOk(context,response)
        } else {
            router.response.ResponseServerError(context,"获取赞助数据失败")
        }
    })
})
// 获取赞助数据
widget.addPage({
    background: tools.getSetting(keyBackground),
    file:"index.html",
    headMeta: {
        title: "赞助博主",
    },
    css: ["element"],
    script: ["vue","element","jquery"],
    url: "",
    full: false,
    side: false
},function (){
    // 获取追番数据
    return {
        server: "/plugins/static/sponsors",
        alipay: tools.getSetting(keyAlipay),
        wechat: tools.getSetting(keyWechat)
    }
})

// 添加设置界面
widget.addSetting("赞助设置",2,"donate")

// 管理界面
// 管理员管理相关
// 获取所有赞助
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
        filter[search] = database.regex(key)
    }
    // 开始搜索
    let db = database.newDb(dbSponsor)
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
                    nickname: item.nickname,
                    amount: item.amount,
                    comment: item.comment,
                    donate_time: tools.time2String(item.donate_time,false)
                })
            })
            router.response.ResponseOk(context,response)
        } else {
            router.response.ResponseServerError(context)
        }
    })
})
// 添加赞助
router.registerAdminRouter("POST","",function (context) {
    // 先获取请post请求的数据
    let param = router.getPostJson(context)
    // 验证关键字段是否为空
    if (tools.verifyField(param.nickname) && tools.verifyField(param.amount)) {
        param.donate_time = tools.time2String(param.donate_time)
        // 插入数据
        database.newDb(dbSponsor).InsertOne(param,function (err,res){
            // 判断是否插入成功
            if (err==null){
                // 为了兼容旧版接口，我们还需要返回值
                param.id = res.InsertedID
                router.response.ResponseCreated(context,param)
            } else {
                router.response.ResponseServerError(context)
            }
        })
    } else {
        router.response.ResponseBadRequest(context,"请检查名字是否填写并正确！")
    }
})
// 更新项目
router.registerAdminRouter("PUT","/:id",function (context){
    // 获取id
    let id = context.Param("id")
    tools.log(id)
    if (id!==""){
        let param = router.getPostJson(context)
        param.donate_time = tools.string2time(param.donate_time,false)
        // 获取id
        let filter = {"_id":{"$in":tools.string2objetIdArray(id,",")}}
        // 更新数据
        database.newDb(dbSponsor).UpdateMany({update: {"$set":param},filter},function (err,res){
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
// 删除项目
router.registerAdminRouter("DELETE","/:id",function (context){
    database.adminDeleteObject(context,dbSponsor,"_id")
})

