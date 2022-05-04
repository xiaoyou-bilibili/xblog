// 使用博客系统提供的接口
const widget =xBlog.widget
const file =xBlog.static
const database =xBlog.database
const cron =xBlog.cron
const tools =xBlog.tools
const router =xBlog.router

// 一些字段
// B站uid
const keyUID = 'bili_uid'
// B站cookie
const keyCookie = 'bili_cookie'
// B站个人信息key值
const keyBiliPerson = 'side_bili_card_person_info'
// B站认证信息
const keyBiliAuth = 'bili_show'
const keySyncNow = 'plugins_bili_sync_now'

// 注册一个定时任务，每20分钟更新一下B站信息
cron.start('0 0/20 * * * ? *',function (){
    updateBliInfo()
})

// 添加管理员路由
router.registerAdminRouter("GET","",function (context){
    if (updateBliInfo()){
        router.response.ResponseCreated(context,{})
    } else {
        router.response.ResponseServerError(context,"更新失败")
    }
})

// 添加卡片
widget.addSide("","index.html",function () {
    // 获取个人信息
    let info = tools.getKey(keyBiliPerson)
    if (!info){
        tools.setKey(keyBiliPerson,tools.getBiliPersonInfo(tools.getSetting(keyUID),tools.getSetting(keyCookie)))
        info = tools.getKey(keyBiliPerson)
    }
    return {
        uid: info.uid,
        nickname: info.nickname,
        sign: info.sign,
        avatar: info.avatar,
        hang: info.hang,
        level: info.level,
        sex: info.sex,
        isVip: info.isvip,
        auth: tools.getSetting(keyBiliAuth),
        fans: info.fans,
        watch: info.watch,
        view: info.view,
        server: "/plugins/static/bili-card"
    }
},true)

// 添加B站卡片设置
widget.addSetting("B站个人信息卡片",1,[
    {title:"B站uid",type: "input",key: keyUID},
    {title:"个人认证",type: "input",key: keyBiliAuth},
    {title:"B站cookie",type: "text",key: keyCookie},
    {title:"立即更新数据",type: "row",key: keySyncNow, default: "admin/plugins/sideBiliCard"}
])

function updateBliInfo(){
    let uid = tools.getSetting(keyUID)
    // uid不为空才更新
    if (uid){
        // 定时把B站个人信息更新到数据库中
        tools.setKey(keyBiliPerson,tools.getBiliPersonInfo(tools.getSetting(keyUID),tools.getSetting(keyCookie)))
        return true
    } else {
        return false
    }
}