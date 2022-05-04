// 使用博客系统提供的接口
const widget =xBlog.widget
const tools =xBlog.tools

// 定义字段
const keyProgram = "mini_program_img"
const keyApp = "app_img"
const keyH5 = "h5_img"

// 添加卡片
widget.addSide("其他客户端","index.html",function () {
    // 返回客户端数据
    return {
        program: tools.getSetting(keyProgram),
        app: tools.getSetting(keyApp),
        h5: tools.getSetting(keyH5),
    }
},true)
// 添加设置信息
widget.addSetting("更多客户端设置",1,[
    {title:"小程序二维码",type: "upload",key: keyProgram},
    {title:"安卓二维码",type: "upload",key: keyApp},
    {title:"H5二维码",type: "upload",key: keyH5},
])

