// 使用博客系统提供的接口
const router =xBlog.router
const database =xBlog.database
const widget =xBlog.widget
const tools = xBlog.tools

// 一些字段
const dbProject = "project_card"

// 获取项目信息
function getProject(okHandle,errHandle){
    let db = database.newDb(dbProject)
    db.FindMany({},function (err,data){
        if (err == null){
            let response = {
                top_content: [],
                bottom_content: []
            }
            // 遍历查询到的数据
            data.forEach(function (item){
                // 判断是否需要置顶
                if (item.is_top) {
                    response.top_content.push({
                        url:item.link,
                        image:item.img,
                        title:item.name
                    })
                }
                // 获取所有项目
                response.bottom_content.push({
                    name : item.name,
                    image : item.img,
                    dec : item.description,
                    time : item.make_time,
                    video : item.video_url,
                    blog : item.blog_url,
                    code : item.code_url,
                })
            })
            // 返回项目数据
            okHandle(response)
        } else {
            errHandle()
        }
    })
}

// 获取所有的项目
router.registerRouter("GET","",function(context){
    getProject(function (data){
        router.response.ResponseOk(context,data)
    },function (){
        router.response.ResponseServerError(context,"获取项目数据失败")
    })
})

// 注册页面
widget.addPage({
    background:"",
    file:"index.html",
    headMeta: {
        title: "我的项目",
    },
    css: ["element"],
    script: ["vue","element","jquery"],
    url: "",
    full: false,
    side: false
},function (){
    let data = []
    getProject(function (res){
        data = res
        tools.log(res)
    },function (){})
    return data
})

// 添加设置界面
widget.addSetting("我的项目",2,"project")

// 管理员管理相关
// 获取所有项目
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
        // 判断是否为置顶
        if (search === "is_top") {
            filter[search] = tools.str2bool(key)
        } else {
            filter[search] = database.regex(key)
        }
    }
    // 开始搜索
    let db = database.newDb(dbProject)
    db.Paginate({filter,sort:{_id:-1}},id,size,function (err,page,total,data){
        if (err==null){
            let response = {
                total_num: total,
                total: page,
                current: id,
                contents: data
            }
            router.response.ResponseOk(context,response)
        } else {
            router.response.ResponseServerError(context)
        }
    })
})
// 添加项目
router.registerAdminRouter("POST","",function (context) {
    // 先获取请post请求的数据
    let data = router.getPostJson(context)
    // 验证关键字段是否为空
    if (tools.verifyField(data.name)) {
        // 插入数据
        database.newDb(dbProject).InsertOne(data,function (err,res){
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
        // 获取id
        let filter = {"_id":{"$in":tools.string2objetIdArray(id,",")}}
        delete param._id
        // 更新数据
        database.newDb(dbProject).UpdateMany({update: {"$set":param},filter},function (err,res){
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
    database.adminDeleteObject(context,dbProject,"_id")
})
