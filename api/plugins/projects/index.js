// 使用博客系统提供的接口
const router =xBlog.router
const database =xBlog.database

// 一些字段
const dbProject = "project_card"

// 获取所有的项目
router.registerRouter("GET","",function(context){
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
            router.response.ResponseOk(context,response)
        } else {
            router.response.ResponseServerError(context,"获取项目数据失败")
        }
    })
})

