// 使用博客系统提供的接口
const router =xBlog.router
const database =xBlog.database
const widget =xBlog.widget
const file =xBlog.static

// 一些字段
const dbSponsor = "donate"

file.staticPath("images","images")

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
    background:"",
    file:"index.html",
    headMeta: {
        title: "赞助博主",
    },
    css: ["element"],
    script: ["vue","element","jquery","xiaoyou"],
    url: "",
    full: false,
    side: false
},function (){
    // 获取追番数据
    return {
        server: "http://127.0.0.1:2333/static/sponsors"
    }
})

