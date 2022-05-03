// 使用博客系统提供的接口
const widget =xBlog.widget
const tools =xBlog.tools

// 定义字段
const keyNotice = "tool_notice"

// 添加卡片
widget.addSide(false,"公告栏","index.html",function () {
    return {
        notice: tools.getSetting(keyNotice),
    }
},true)
