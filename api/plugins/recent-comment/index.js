// 使用博客系统提供的接口
const widget =xBlog.widget
const tools =xBlog.tools
const database =xBlog.database

// 定义字段
const dbComment = "comment"
const keyCommentNum = "comment_num"

// 添加卡片
widget.addSide("最近评论","index.html",function () {
    let db = database.newDb(dbComment)
    let comments = []
    db.FindMany({
        sort: { "_id":-1 },
        filter: { "agree": 1 },
        limit: tools.getSetting(keyCommentNum)
    },function (err, data) {
        data.forEach(function (item) {
            comments.push({
                id: item.id,
                postId: item.post_id,
                userID: item.user_id,
                nickname: item.nickname,
                avatar: item.avatar===""?tools.getRandomAvatar():item.avatar,
                content: tools.changeCommentSmile(item.content),
                date: tools.time2String(new Date(item.comment_time),true),
                parent: item.parent,
                hang: item.hang,
                level: item.level,
                uid: item.uid
            })
        })
    })
    return {
        comments
    }
},true)

// 添加设置信息
widget.addSetting("最近评论设置",1,[
    {title:"最近评论显示个数",type: "number",key: keyCommentNum}
])