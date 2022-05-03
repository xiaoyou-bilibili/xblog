// 使用博客系统提供的接口
const router =xBlog.router
const tools =xBlog.tools
const net =xBlog.net
const widget =xBlog.widget

// 一些设置的值
// B站uid
const keyUID = 'bili_uid'
// B站cookie
const keyCookie = 'bili_cookie'

// 获取动画数据
function getAnimation(id,okHandle,errHandle){
    // 获取B站UID
    let uid = tools.getSetting(keyUID)
    let cookie = tools.getSetting(keyCookie)
    // 设置头部信息
    let head = {
        origin : "https://api.bilibili.com",
        Referer : "https://api.bilibili.com",
        Cookie : cookie,
    }
    // 发送请求
    net.get('https://api.bilibili.com/x/space/bangumi/follow/list?type=1&pn='+id+'&vmid='+uid,head,function (err, res) {
        let response = {
            total: 0, // 总页数
            num: 0, // 番剧数
            current: 0, // 当前第几页
            contents: [], // 番剧内容
        }
        if (err==null){
            try {
                // 解析json数据
                let data = JSON.parse(res).data
                // 设置一些基本数据
                response.num = data.total
                response.total = Math.ceil(response.num/data.ps)
                response.current = parseInt(id)
                // 开始遍历获取追番数据
                data.list.forEach(function(item){
                    // 设置追番信息
                    let animation = {}
                    animation.title = item.title
                    animation.cover = item.cover
                    animation.dec = item.evaluate
                    animation.current = item.progress
                    animation.url = item.url
                    // 如果有进度，我们就计算一下进度
                    if (item.new_ep!==undefined){
                        animation.total = item.new_ep.index_show
                        // 进度我们使用正则来进行匹配获取当前第几集，然后除于总集数
                        let total = item.new_ep.title
                        // 判断total是否为数组，如果不是数字就设置值
                        if (isNaN(total)) { total=1 }
                        // 计算当前进度
                        animation.percent = Math.floor(tools.findMatch('([0-9]+)',animation.current) / total *100)
                    } else {
                        // 为了避免解析出错，这里我们也进行一下赋值
                        animation.total = 0
                        animation.percent = 0
                    }
                    // 把数据加到数组中
                    response.contents.push(animation)
                })
                // 正常返回数据
                okHandle(response)
            } catch (e){
                // 触发错误回调
                errHandle("解析数据异常")
            }
        } else {
            errHandle("获取数据失败")
        }
    })
}

// 注册路由，当获取追番信息时调用此接口
router.registerRouter("GET","",function(context){
    // 用户Bid
    let id = context.Query("page")
    if (id==="") id=1
    getAnimation(id, function (data){
        router.response.ResponseOk(context,data)
    },function (msg){
        router.response.ResponseServerError(context,msg)
    })
})

// 注册追番页面
widget.addPage({
    background:"",
    file:"index.html",
    headMeta: {
        title: "我的追番",
    },
    css: ["element"],
    script: ["vue","jquery","element","xiaoyou"],
    url: "",
    full: false,
    side: false
},function (){
    let a = {}
    return a
})
