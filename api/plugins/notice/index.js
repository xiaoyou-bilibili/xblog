// 使用博客系统提供的接口
const widget =xBlog.widget
const tools =xBlog.tools

// 定义字段
const keyNotice = "tool_notice"

// 添加卡片
widget.addSide("公告栏","index.html",function () {
    return {
        notice: tools.getSetting(keyNotice),
    }
},true)

// 添加设置信息
widget.addSetting("公告栏设置",1,[
    {title:"公告栏内容(支持html标签)",type: "text",key: keyNotice}
])
