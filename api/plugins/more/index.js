// 使用博客系统提供的接口
const widget =xBlog.widget
const tools =xBlog.tools

const keyMore = "plugin_more_functions"

// 添加卡片
widget.addSide("更多功能","index.html",function () {
    // 因为系统API限制，需要自己进行手动设置
    // 不需要的自己注释掉就行了
    const functions = JSON.parse(tools.getSetting(keyMore))
    return {
        functions
    }
},true)


// 我的介绍
widget.addSetting("更多设置",1,[
    {title:"设置内容（必须按照默认的json格式来）",type: "text",key: keyMore,default: '[{"title":"日记","icon":"far fa-calendar-alt","link":"/diary","color":"#5FB878"},{"title":"赞助博主","icon":"fas fa-donate","link":"/more/sponsors","color":"#FFB800"},{"title":"友人帐","icon":"fas fa-child","link":"/more/friends","color":"#01AAED"},{"title":"我的追番","icon":"fas fa-ghost","link":"/more/animations","color":"#f25d8e"}]'}
])