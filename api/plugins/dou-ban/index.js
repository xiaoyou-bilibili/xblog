// 使用博客系统提供的接口
const router =xBlog.router
const database =xBlog.database
const tools =xBlog.tools
const widget = xBlog.widget
const net = xBlog.net
const spider = xBlog.spider

// 一些字段
const dbDouBan = "dou_ban"

const keyBackground = "dou_ban_img"
const keyLastUpdate = "dou_ban_last_update"
const keyDouBanCookie = "dou_ban_cookie"
const keyDouBanUser = "dou_ban_user"

// 获取豆瓣记录
router.registerRouter("GET","/:type",function(context){
    // 获取页数信息
    let id = parseInt(context.Query("page"))
    // id转换为数字
    if (isNaN(id)){
        id = 1
    }
    // 获取请求的类型
    let type = context.Param("type")
    // 返回的数据
    let response = {
        current: id,
        total: 0,
        contents: []
    }
    // 开始获取数据
    let db = database.newDb(dbDouBan)
    db.Paginate({ "filter" : { "item_type": type } },id,20,function (err,page,total,data){
        if (err==null){
            response.total = page
            data.forEach(function (item){
                response.contents.push({
                    name: item.name,
                    picture: item.image,
                    star: item.score,
                    pub: item.pub_info,
                    comment: item.comment,
                    status: item.status,
                    url: item.url
                })
            })
            router.response.ResponseOk(context,response)
        } else {
            router.response.ResponseServerError(context,"查询数据失败")
        }
    })
})
// 注册豆瓣接口
widget.addPage({
    background: tools.getSetting(keyBackground),
    file:"index.html",
    headMeta: {
        title: "我的豆瓣",
    },
    css: ["element"],
    script: ["vue","element","jquery"],
    url: "",
    full: false,
    side: false
},function (){
    let db = database.newDb(dbDouBan)
    //获取一言
    // setting.Talk = tools.HttpGet("https://v1.hitokoto.cn?c=a&c=b&c=c&c=d&c=e&c=f&c=g&c=h&c=i&c=j&c=k&c=l&encode=text")
    let talk = ""
    net.get("https://v1.hitokoto.cn?c=a&c=b&c=c&c=d&c=e&c=f&c=g&c=h&c=i&c=j&c=k&c=l&encode=text",{},function (err,res){
        if (err==null){ talk = res }
    })
    // 友链设置
    return {
        movie: db.GetCount({"item_type":"movie"}),
        book: db.GetCount({"item_type":"book"}),
        music: db.GetCount({"item_type":"music"}),
        last: tools.getSetting(keyLastUpdate),
        talk
    }
})

// 获取图书
function getBook(colly,user,type){
    // 开始爬取
    colly.OnHTML("#content",function (e){
        // 遍历所有的项目
        e.ForEach("[class='interest-list']>.subject-item", function (i,element) {
            let title = element.ChildText(".info>h2>a").toString()
            title = tools.strReplace(title," ","",-1)
            title = tools.strReplace(title,"\n","",-1)
            // 获取内容
            let content = {
                title,
                pic: element.ChildAttr(".pic>a>img","src"),
                pub: element.ChildText(".info>.pub"),
                comment: element.ChildText(".short-note>.comment").trim(),
                url: element.ChildAttr(".info>h2>a","href"),
                star: tools.findMatch("rating([0-5]).*?", element.ChildAttr(".info>.short-note>div>span:first-child","class"))
            }
            tools.log(content)
        })
        // 获取下一页的链接
        let next = e.ChildAttr(".paginator>.next>a","href")
        if (next!==""){
            next = "http://book.douban.com" + next
            e.Request.Visit(next)
        }
    })
   colly.Visit("http://book.douban.com/people/"+user+"/"+type)
}

// 获取视频
function getMovie(){

}

// 获取音乐
function getMusic(){


}


// 爬虫爬取数据
function Spider(){
    let colly = spider.init({"Host":"www.douban.com","Cookie":tools.getSetting(keyDouBanCookie)})
    // 更新一下时间
    // db.SetSiteOption(db.KeyDouBanLastUpdate, tools.Time2String(time.Now(), true))
    // 清空所有数据
    let db = database.newDb(dbDouBan)
    // db.de
    // 获取在读的图书
    getBook(colly,tools.getSetting(keyDouBanUser),"collect")
}

// 注册定时任务爬数据
router.registerRouter("GET","/spider",function(context){
    Spider()
    router.response.ResponseOk(context,{message:"爬取数据中"})
    // db.Paginate({ "filter" : { "item_type": type } },id,20,function (err,page,total,data){
    //     if (err==null){
    //         response.total = page
    //         data.forEach(function (item){
    //             response.contents.push({
    //                 name: item.name,
    //                 picture: item.image,
    //                 star: item.score,
    //                 pub: item.pub_info,
    //                 comment: item.comment,
    //                 status: item.status,
    //                 url: item.url
    //             })
    //         })
    //         router.response.ResponseOk(context,response)
    //     } else {
    //         router.response.ResponseServerError(context,"查询数据失败")
    //     }
    // })
})
