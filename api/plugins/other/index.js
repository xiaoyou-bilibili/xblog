// 使用博客系统提供的接口
const widget =xBlog.widget
const tools =xBlog.tools

// 定义字段
const keyNotice = "tool_notice"

// 添加卡片
widget.addSide(false,"其他客户端","index.html",function () {
    const functions = [
        { title: "日记",icon: "far fa-calendar-alt",link: "/more/diary",color: "#5FB878" },
        { title: "赞助博主",icon: "fas fa-donate",link: "/more/diary",color: "#FFB800" },
        { title: "友人帐",icon: "fas fa-child",link: "/more/diary",color: "#01AAED" },
        { title: "我的追番",icon: "fas fa-ghost",link: "/more/diary",color: "#f25d8e" },
    ]
    return {
        functions
    }
},true)
