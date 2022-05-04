// 使用博客系统提供的接口
const widget =xBlog.widget
const database =xBlog.database
const tools =xBlog.tools
const router = xBlog.router
const net = xBlog.net

// 常量
const keyMusicContent = 'music_content'
const  keyMusicCsrf = 'music_163_csrf'
const  keyMusicParams = 'music_163_params'
const  keyMusicEncSecKey = 'music_163_enc'
const  keyMusicSyncNow = 'plugin_music_sync_now'
const  keyMusicCookie = 'plugin_music_163_cookie'
const keyWebServer = 'site_api_server'

// 获取音乐歌词
router.registerRouter("GET","/:id/irc",function(context){
    net.get("http://music.163.com/api/song/media?id="+context.Param("id"),{"Host":"music.163.com"},function (err,res){
        if (err==null){
            let data = JSON.parse(res)
            if (data.lyric!==undefined){
                router.response.ResponseHtml(context,data.lyric)
                return
            }
        }
        router.response.ResponseHtml(context,"")
    })
})

// 添加卡片
widget.addSide("","index.html",function () {
    // 初始化数据库链接
    return {
        content: JSON.stringify(tools.getSetting(keyMusicContent))
    }
},true)

// 添加设置信息
widget.addSetting("音乐盒设置",1,[
    {title:"网易云csrf_token",type: "input",key: keyMusicCsrf,default: ''},
    {title:"网易云params",type: "input",key: keyMusicParams,default: ''},
    {title:"网易云encSecKey",type: "input",key: keyMusicEncSecKey,default: ''},
    {title:"网易云cookie",type: "text",key: keyMusicCookie,default: ''},
    {title:"立即同步",type: "row",key: keyMusicSyncNow,default:"admin/plugins/sideMusic"},
])

// 获取网易云歌单
function getMusicContent(context){
    // 设置头部信息
    let head = {
        origin: 'https://music.163.com',
        referer:' https://music.163.com/',
        Cookie: tools.getSetting(keyMusicCookie)
    }
    // 发送请求
    net.post('https://music.163.com/weapi/v6/playlist/detail?csrf_token='+tools.getSetting(keyMusicCsrf),head,{
        'params': tools.getSetting(keyMusicParams),
        'encSecKey': tools.getSetting(keyMusicEncSecKey)
    },function (err,res){
        if (!err) {
            let server = tools.getSetting(keyWebServer)
            let data = []
            res = JSON.parse(res)
            if (res.code !== 200) {
                router.response.ResponseServerError(context, "获取歌曲失败!")
                return false
            }
            // 解析数据
            res.playlist.tracks.forEach(function (item){
                data.push({
                    name: item.name,
                    artist: item.ar[0].name,
                    url: 'https://music.163.com/song/media/outer/url?id='+item.id,
                    cover: tools.strReplace(item.al.picUrl,'http:','',-1),
                    lrc: server + "/api/v3/plugins/sideMusic/" + item.id + "/irc"
                })
            })
            tools.log(JSON.stringify(data))
            // 保存数据库
            tools.setSetting(keyMusicContent,data)
            router.response.ResponseOk(context,data)
        }
    })
}

//  网易云爬虫
router.registerAdminRouter("GET","",function (context){
    getMusicContent(context)
})