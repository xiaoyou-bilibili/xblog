// 使用博客系统提供的接口
const router =xBlog.router
const database =xBlog.database
const tools =xBlog.tools
const widget = xBlog.widget
const net = xBlog.net
const spider = xBlog.spider
const cron = xBlog.cron

// 一些字段
const dbDouBan = "dou_ban"

const keyBackground = "dou_ban_img"
const keyLastUpdate = "dou_ban_last_update"
const keyDouBanCookie = "dou_ban_cookie"
const keyDouBanUser = "dou_ban_user"
const keyDouBanSync = "dou_ban_sync"
const keyDouBanSyncNow = "dou_ban_sync_now"

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

// 替换豆瓣状态
function replaceDouBanStatus(content,types){
    if (types === "book") {
        switch (content) {
            case "do":
                return "在读"
            case "wish":
                return "想读"
            case "collect":
                return "已读"
        }
    } else if (types === "movie") {
        switch (content) {
            case "do":
                return "在看"
            case "wish":
                return "想看"
            case "collect":
                return "已看"
        }
    } else if (types === "music") {
        switch (content) {
            case "do":
                return "在听"
            case "wish":
                return "想听"
            case "collect":
                return "已听"
        }
    }
    return ""
}

// 获取图书
function getBook(db,cookie,user,type){
    let colly = spider.init({"Host":"www.douban.com","Cookie":cookie})
    // 开始爬取
    colly.OnHTML("#content",function (e){
        // 遍历所有的项目
        e.ForEach("[class='interest-list']>.subject-item", function (i,element) {
            let name = element.ChildText(".info>h2>a").toString()
            name = tools.strReplace(name," ","",-1)
            name = tools.strReplace(name,"\n","",-1)
            let star = tools.findMatch("rating([0-5]).*?", element.ChildAttr(".info>.short-note>div>span:first-child","class"))
            // 获取内容
            let content = {
                name,
                image: element.ChildAttr(".pic>a>img","src"),
                pub_info: element.ChildText(".info>.pub"),
                comment: element.ChildText(".short-note>.comment").trim(),
                url: element.ChildAttr(".info>h2>a","href"),
                score: tools.str2int(star),
                item_type: "book",
                status: replaceDouBanStatus(type,"book")
            }
            db.InsertOne(content,function (err,res){})
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
function getMovie(db,cookie,user,type){
    let colly = spider.init({"Host":"www.douban.com","Cookie":cookie})
    // 开始爬取
    colly.OnHTML(".article",function (e){
        // 遍历所有的项目
        e.ForEach(".grid-view>.item", function (i,element) {
            let name = element.ChildText(".info>ul>.title>a").toString()
            name = tools.strReplace(name," ","",-1)
            name = tools.strReplace(name,"\n","",-1)
            let star = tools.findMatch("rating([0-5]).*?", element.ChildAttr(".info>ul>li>span:first-child","class"))
            // 获取内容
            let content = {
                name,
                image: element.ChildAttr(".pic>a>img","src"),
                pub_info: element.ChildText(".info>ul>.intro"),
                comment: element.ChildText(".info>ul>li>.comment").trim(),
                url: element.ChildAttr(".info>ul>.title>a","href"),
                score: tools.str2int(star),
                item_type: "movie",
                status: replaceDouBanStatus(type,"movie")
            }
            db.InsertOne(content,function (err,res){})
        })
        // 获取下一页的链接
        let next = e.ChildAttr(".paginator>.next>a","href")
        if (next!==""){
            next = "http://movie.douban.com" + next
            e.Request.Visit(next)
        }
    })
    colly.Visit("http://movie.douban.com/people/"+user+"/"+type)
}

// 获取音乐
function getMusic(db,cookie,user,type){
    let colly = spider.init({"Host":"www.douban.com","Cookie":cookie})
    // 开始爬取
    colly.OnHTML(".article",function (e){
        // 遍历所有的项目
        e.ForEach(".grid-view>.item", function (i,element) {
            let name = element.ChildText(".info>ul>.title>a").toString()
            name = tools.strReplace(name," ","",-1)
            name = tools.strReplace(name,"\n","",-1)
            let star = tools.findMatch("rating([0-5]).*?", element.ChildAttr(".info>ul>li>span:first-child","class"))
            // 获取内容
            let content = {
                name,
                image: element.ChildAttr(".pic>a>img","src"),
                pub_info: element.ChildText(".info>ul>.intro"),
                comment: element.ChildText(".info>ul>li:nth-child(4)").trim(),
                url: element.ChildAttr(".info>ul>.title>a","href"),
                score: tools.str2int(star),
                item_type: "music",
                status: replaceDouBanStatus(type,"music")
            }
            db.InsertOne(content,function (err,res){})
        })
        // 获取下一页的链接
        let next = e.ChildAttr(".paginator>.next>a","href")
        if (next!==""){
            next = "http://music.douban.com" + next
            e.Request.Visit(next)
        }
    })
    colly.Visit("http://music.douban.com/people/"+user+"/"+type)
}

// 爬虫爬取数据
function Spider(){
    // 更新一下时间
    tools.setSetting(keyLastUpdate,tools.time2String(new Date(),true))
   // 清空所有数据
    let db = database.newDb(dbDouBan)
    db.DeleteMany({},function (err,res){})
    let user = tools.getSetting(keyDouBanUser)
    let cookie = tools.getSetting(keyDouBanCookie)
    // 获取豆瓣的所有记录
    getBook(db,cookie,user,"collect")
    getBook(db,cookie,user,"do")
    getBook(db,cookie,user,"wish")
    getMovie(db,cookie,user,"collect")
    getMovie(db,cookie,user,"do")
    getMovie(db,cookie,user,"wish")
    getMusic(db,cookie,user,"collect")
    getMusic(db,cookie,user,"do")
    getMusic(db,cookie,user,"wish")
}

// 设置界面
// 注册追番设置
widget.addSetting("豆瓣设置",1,[
    {title:"豆瓣用户ID",type: "input",key: keyDouBanUser,default:""},
    {title:"豆瓣cookie",type: "text",key: keyDouBanCookie,default:""},
    {title:"每日定时同步",type: "switch",key: keyDouBanSync,default: false},
    {title:"立即同步(耗时操作，请勿重复点击！)",type: "row",key: keyDouBanSyncNow,default: "admin/plugins/dou_ban"}
])

// 豆瓣爬虫
router.registerAdminRouter("GET","",function (context){
    Spider()
    router.response.ResponseOk(context,{})
})

// 注册定时任务
cron.start("0 0 0 1/1 * ?",function () {
    if (tools.getSetting(keyDouBanSync)){
        spider()
    }
})