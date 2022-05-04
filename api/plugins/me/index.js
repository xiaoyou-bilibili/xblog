// 使用博客系统提供的接口
const router =xBlog.router
const database =xBlog.database
const net =xBlog.net
const tools =xBlog.tools
const widget = xBlog.widget

// todo 记得加上图片字段

// 一些字段
const keyNickname = "plugin_me_nickname"
const keyAvatar = "plugin_me_avatar"
const keyIntroduce = "plugin_me_introduce"
const keyProfile = "plugin_me_profile"
const keySkillPoint = "plugin_me_skill_point"
const keySkillDetail = "plugin_me_skill_detail"


// 注册我的导航接口
widget.addPage({
    background: "",
    file:"index.html",
    headMeta: {
        title: "个人介绍",
    },
    css: ["element"],
    script: ["vue","element","jquery"],
    url: "",
    full: false,
    side: false
},function (){
    return {
        nickname: tools.getSetting(keyNickname),
        avatar: tools.getSetting(keyAvatar),
        introduce: tools.getSetting(keyIntroduce),
        profile: tools.getSetting(keyProfile),
        skillPoint: tools.getSetting(keySkillPoint),
        skillDetail: tools.getSetting(keySkillDetail),
        server: '/plugins/static/me'
    }
})

// 我的介绍
widget.addSetting("个人介绍",1,[
    {title:"昵称",type: "input",key: keyNickname,default:"小游"},
    {title:"头像",type: "upload",key: keyAvatar,default:"https://cdn.jsdelivr.net/gh/xiaoyou66/imgbed@1.0/xblog/tTSY.jpg"},
    {title:"个人介绍",type: "text",key: keyIntroduce,default: "学渣一个，沉迷于二次元无法自拔，偶尔也玩点游戏。原本喜欢搞硬件，现在大部分时间花在搞软件上，喜欢捣鼓各种稀奇的小玩意。没什么远大的理想，只希望可以做自己喜欢做的事情。"},
    {title:"个人资料（必须按照默认的json格式来）",type: "text",key: keyProfile,default: '[{"item":"常用昵称","content":"小游"},{"item":"爱好","content":"二次元"},{"item":"个人博客","content":"<a target=\\"_blank\\" href=\\"https://xiaoyou66.com\\">xiaoyou66.com</a>"},{"item":"GitHub","content":" <a target=\\"_blank\\" href=\\"https://github.com/xiaoyou66\\">github.com/xiaoyou66</a>"},{"item":"B站","content":"<a target=\\"_blank\\" href=\\"https://space.bilibili.com/343147393\\">UID:343147393</a>"}]'},
    {title:"技能点（必须按照默认的json格式来）",type: "text",key: keySkillPoint,default: '[{"name":"HTML/CSS","process":70,"color":"#8CC7B5"},{"name":"JavaScript","process":65,"color":"#19CAAD"},{"name":"PHP","process":60,"color":"#A0EEE1"},{"name":"JAVA","process":65,"color":"#BEE7E9"},{"name":"Go","process":60,"color":"#BEEDC7"},{"name":"Python","process":60,"color":"#D6D5B7"},{"name":"linux","process":65,"color":"#D1BA74"},{"name":"C#","process":5,"color":"#E6CEAC"},{"name":"node.js","process":30,"color":"#ECAD9E"}]'},
    {title:"技能点详细介绍（必须按照默认的json格式来）",type: "text",key: keySkillDetail,default: '[{"text":"前端方面:简单学习过html、css、JavaScript的语法，学习过jQuery，使用过bootstrap、element-ui、layui等前端框架","color":"#8CC7B5"},{"text":"前端框架:学习过vue、nuxt等主流前端框架，同时还学习过uniapp跨端框架","color":"#19CAAD"},{"text":"后端框架:学习过Echo(go语言开发的轻量后端框架)、thinkPHP(php开发的高性能框架)、简单了解过ssm框架","color":"#A0EEE1"},{"text":"软件开发:学习过android开发，开发过多个APP，同时也会微信小程序开发，擅长使用colorUI、vant等常用微信小程序前端框架","color":"#BEE7E9"},{"text":"爬虫方面:学习过Python的scrapy框架和go的colly框架、会使用xpath、jQuery、正则来查找内容","color":"#BEEDC7"},{"text":"系统方面:使用过centos、ubuntu、kali等linux系统，可以简单的操作linux","color":"#D6D5B7"},{"text":"数据库:学习过MySQL","color":"#D1BA74"},{"text":"设计软件:可以简单的使用PS、Pr、AE等图像和视频处理软件，简单学习过SolidWorks建模以及Altium Designer画电路图","color":"#E6CEAC"},{"text":"硬件方面:使用过51、arduino、树莓派、stm32、stc15等常见硬件环境开发","color":"#ECAD9E"}]'}
])
