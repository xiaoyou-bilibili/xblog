// 使用博客系统提供的接口
const router =xBlog.router
const database =xBlog.database
const net =xBlog.net
const tools =xBlog.tools
const widget = xBlog.widget

// todo 记得加上图片字段

// 一些字段
const dbNavigation = "navigation"
const keyBackground = "animation_img"

// 获取所有项目
router.registerRouter("GET","/links",function(context){
    let db = database.newDb(dbNavigation)
    db.FindMany({},function (err,data){
        if (err == null){
            // 先使用tmp暂存数据
            let tmp = []
            // 第一次遍历获取父节点
            data.forEach(function (item){
                // 只放入我们需要的值
                if (item.parent===0){
                    let parent = {
                        id : item.navigation_id,
                        name : item.name,
                        color : item.value,
                        child : []
                    }
                    tmp[parent.id] = parent
                }
            })
            // 第二次遍历填充子节点
            data.forEach(function (item){
                // 只放入我们需要的值
                if (item.parent!==0){
                    tmp[item.parent].child.push({
                        id : item.navigation_id,
                        name : item.name,
                        color : item.value
                    })
                }
            })
            // 最后我们把对象转换为数组
            // Object.values 可以把对象里面所有的value转换为数组
            let response = Object.values(tmp)
            // 返回数据
            router.response.ResponseOk(context,response)
        } else {
            router.response.ResponseServerError(context,"获取个人导航失败")
        }
    })
})

// 获取百度搜索关键词
router.registerRouter("GET","/keyword",function(context){
    // 获取搜索的关键词
    let q = context.Query("q")
    // 验证关键词是否为空
    if (tools.verifyField(q)){
        // 发送get请求
        net.get("https://www.baidu.com/sugrec?prod=pc&wd="+q,{},function (err,row){
            if (err==null){
                // 我们解析一下json数据
                let data = JSON.parse(row)
                // 最终结果
                let response = []
                // 这里我们只需要获取关键词数组即可
                data.g.forEach(function (item){
                    response.push(item.q)
                })
                router.response.ResponseOk(context,response)
            } else {
               router.response.ResponseServerError(context,"发送请求失败")
            }
        })
    } else {
        router.response.ResponseBadRequest(context,"关键词不能为空")
    }
})

// 注册我的导航接口
widget.addPage({
    background: tools.getSetting(keyBackground),
    file:"index.html",
    headMeta: {
        title: "个人导航",
    },
    css: ["element"],
    script: ["vue","element","jquery"],
    url: "",
    full: false,
    side: false
},function (){
    // let db = database.newDb(dbDouBan)
    return {
        server: '/plugins/static/navigation'
    }
})

// 添加设置界面
widget.addSetting("网址管理",2,"navigation")

// 管理员接口
// 获取所有导航分类
router.registerAdminRouter("GET","",function (context){
    let db = database.newDb(dbNavigation)
    db.FindMany({filter:{"parent":0},sort:{_id:-1}},function (err,data){
        if (err == null){
            let response = []
            data.forEach(function (item){
                // 只放入我们需要的值
                response.push({
                    id : item.navigation_id,
                    name : item.name,
                    color : item.value
                })
            })
            // 返回数据
            router.response.ResponseOk(context,response)
        } else {
            router.response.ResponseServerError(context,"获取个人导航失败")
        }
    })
})

// 获取导航分类下的网址
router.registerAdminRouter("GET","/:id/links",function (context){
    let db = database.newDb(dbNavigation)
    let id = tools.str2int(context.Param("id"))
    let pageId = tools.str2int(context.Query('page'))
    let size = tools.str2int(context.Query('page_size'))
    if (pageId===0) { pageId=1 }
    if (size===0) { size=10 }
    db.Paginate({filter: {parent:id},sort:{_id:-1}},pageId,size,function (err,page,total,data){
        if (err==null){
            let response = {
                total_num: total,
                total: page,
                current: pageId,
                contents: []
            }
            data.forEach(function (item){
                response.contents.push({
                    id : item.navigation_id,
                    name : item.name,
                    url : item.value,
                    parent : item.parent,
                })
            })
            router.response.ResponseOk(context,response)
        } else {
            router.response.ResponseServerError(context)
        }
    })
})

// 添加导航或者网址
router.registerAdminRouter("POST","",function (context){
    // 先获取请post请求的数据
    let navigation = router.getPostJson(context)
    // 验证关键字段是否为空
    if (tools.verifyField(navigation.name) && tools.verifyField(navigation.value)) {
        let data = {
            name: navigation.name,
            value: navigation.value,
            parent: navigation.parent
        }
        // 插入数据
        database.newDb(dbNavigation).InsertOneIncrease(data,"navigation_id",function (err,res){
            // 判断是否插入成功
            if (err==null){
                router.response.ResponseCreated(context,data)
            } else {
                router.response.ResponseServerError(context)
            }
        })
    } else {
        router.response.ResponseBadRequest(context,"请检查名字、网址是否填写并正确！")
    }
})

// 更新导航内容
router.registerAdminRouter("PUT","/:id",function (context){
    // 获取ID
    let id = tools.str2int(context.Param("id"))
    // 先获取请post请求的数据
    let navigation = router.getPostJson(context)
    // 验证关键字段是否为空
    if (tools.verifyField(navigation.name) && tools.verifyField(navigation.value)) {
        let data = {
            name: navigation.name,
            value: navigation.value,
            parent: navigation.parent
        }
        database.newDb(dbNavigation).UpdateOne({filter:{"navigation_id": id},update: {"$set":data}},function (err,res){
            if (!err){
                router.response.ResponseCreated(context,data)
            } else {
                router.response.ResponseServerError(context)
            }
        })
    }

})

// 删除导航
router.registerAdminRouter("DELETE","/:id",function (context){
    database.adminDelete(context,dbNavigation,"navigation_id")
})