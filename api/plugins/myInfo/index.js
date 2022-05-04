// 使用博客系统提供的接口
const widget =xBlog.widget
const database =xBlog.database
const tools =xBlog.tools

// 一些字段
const dbArticle = "article"
const dbComment = "comment"
const keyIntroduce = "my_introduce"
const keyBackground = "my_background"
const keyAvatar = "my_avatar"
const keyGithub = "my_github"
const keyTelegram = "my_telegram"
const keyTwitter = "my_twitter"
const keyZhiHu = "my_zhi_hu"
const keySteam = "my_steam"
const keyBili = "my_bili_bili"

// 添加卡片
widget.addSide("","index.html",function () {
    // 初始化数据库链接
    let db = database.newDb(dbArticle)
    let commentDb = database.newDb(dbComment)
    return {
        background: tools.getSetting(keyBackground),
        avatar: tools.getSetting(keyAvatar),
        post: db.GetCount({"status": "publish", "post_type": "post"}),
        diary: db.GetCount({"status": "publish", "post_type": "diary"}),
        comment: commentDb.GetCount({}),
        introduce: tools.getSetting(keyIntroduce),
        bili: tools.getSetting(keyBili),
        github: tools.getSetting(keyGithub),
        twitter: tools.getSetting(keyTwitter),
        telegram: tools.getSetting(keyTelegram),
        steam: tools.getSetting(keySteam),
        zhiHu: tools.getSetting(keyZhiHu)
    }
},true)

// 添加设置信息
widget.addSetting("个人信息设置",1,[
    {title:"个人介绍",type: "input",key: keyIntroduce},
    {title:"个人头像",type: "upload",key: keyAvatar},
    {title:"卡片背景",type: "upload",key: keyBackground},
    {title:"telegram",type: "input",key: keyTelegram},
    {title:"github",type: "input",key: keyGithub},
    {title:"Twitter",type: "input",key: keyTwitter},
    {title:"知乎",type: "input",key: keyZhiHu},
    {title:"steam",type: "input",key: keySteam},
    {title:"B站",type: "input",key: keyBili}
])

