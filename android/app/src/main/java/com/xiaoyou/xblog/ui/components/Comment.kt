package com.xiaoyou.xblog.ui.components

import android.content.Context
import android.util.AttributeSet
import android.view.Gravity
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.LinearLayout
import android.widget.TextView
import androidx.core.content.ContextCompat
import com.facebook.drawee.generic.RoundingParams
import com.facebook.drawee.view.SimpleDraweeView
import com.scwang.smart.refresh.layout.util.SmartUtil.dp2px
import com.xiaoyou.library.common.components.NoScrollWebView
import com.xiaoyou.library.net.entity.response.CommentDetail
import com.xiaoyou.xblog.R

/**
 * @description 自定义评论显示组件
 * @author 小游
 * @data 2021/02/18
 */
class Comment(context: Context, attrs: AttributeSet) : LinearLayout(context, attrs) {

    // 只能使用这种方式来加载布局文件，不能使用binding
    private val view = LayoutInflater.from(context).inflate(R.layout.commponent_comment, this)
    // 绑定布局
    private val root:LinearLayout = view.findViewById(R.id.comment)

    // kotlin方式的回调函数，我们只传递了当个函数
    // 参考 https://www.jianshu.com/p/39d45cd853b5
    private lateinit var replay: (View) -> Unit

    /**
     * 设置reply
     * @param replay Replay
     */
    fun setReplay(replay: (View) -> Unit) {
        this.replay = replay
    }

    /**
     * 根据viewModel跟新comment
     */
    fun addComment(commentDetailList: List<CommentDetail>){
        for (i in commentDetailList.indices) {
                if (commentDetailList[i].parent == 0) {
                    //第一层的parent节点
                    val pComment = LayoutParams(ViewGroup.LayoutParams.MATCH_PARENT, ViewGroup.LayoutParams.WRAP_CONTENT)
                    //每个comment的间距
                    pComment.topMargin = dp2px(5f)
                    pComment.leftMargin = dp2px(5f)
                    pComment.rightMargin = dp2px(5f)
                    //新建布局
                    val comment = LinearLayout(context)
                    //评论背景
                    comment.background = ContextCompat.getDrawable(context, R.drawable.comment_round)
                    //背景透明
//                    comment.background.alpha = if (App.getImg(context) != null) App.getAlpha(context) else 255//透明度
                    comment.orientation = VERTICAL
                    //评论加上评论内容
                    comment.addView(addComment(context, commentDetailList[i], true))
                    //开始获取第二层的内容
                    for (i2 in commentDetailList.indices) {
                        if (commentDetailList[i2].parent == commentDetailList[i].id) {
                            //第一层的parent节点
                            //新建布局
                            val comment2 = LinearLayout(context)
                            //评论背景
                            comment2.background = ContextCompat.getDrawable(context, R.drawable.comment_child_round)
                            //背景透明
//                            comment2.background.alpha = if (App.getImg(context) != null) App.getAlpha(context) else 255//透明度
                            comment2.orientation = VERTICAL
                            //评论加上评论内容
                            comment2.addView(addComment(context, commentDetailList[i2], true))
                            //开始获取第三层的内容
                            for (i3 in commentDetailList.indices) {
                                if (commentDetailList[i3].parent == commentDetailList[i2].id) {
                                    //第一层的parent节点
                                    //新建布局
                                    val comment3 = LinearLayout(context)
                                    //评论背景
                                    comment3.background = ContextCompat.getDrawable(context, R.drawable.comment_child_round)
                                    //背景透明
//                                    comment3.background.alpha = if (App.getImg(context) != null) App.getAlpha(context) else 255//透明度
                                    comment3.orientation = VERTICAL
                                    //评论加上评论内容
                                    comment3.addView(addComment(context, commentDetailList[i3], true))
                                    //开始获取第四层的内容
                                    for (i4 in commentDetailList.indices) {
                                        if (commentDetailList[i4].parent == commentDetailList[i3].id) {
                                            //第一层的parent节点
                                            //新建布局
                                            val comment4 = LinearLayout(context)
                                            //评论背景
                                            comment4.background = ContextCompat.getDrawable(context, R.drawable.comment_child_round)
                                            //背景透明
//                                            comment4.background.alpha = if (App.getImg(context) != null) App.getAlpha(context) else 255//透明度
                                            comment4.orientation = VERTICAL
                                            //评论加上评论内容
                                            comment4.addView(addComment(context, commentDetailList[i4], true))
                                            //开始获取第五层的内容
                                            for (i5 in commentDetailList.indices) {
                                                if (commentDetailList[i5].parent == commentDetailList[i4].id) {
                                                    //第一层的parent节点
                                                    //新建布局
                                                    val comment5 = LinearLayout(context)
                                                    //评论背景
                                                    comment5.background = ContextCompat.getDrawable(context, R.drawable.comment_child_round)
                                                    //背景透明
//                                                    comment5.background.alpha = if (App.getImg(context) != null) App.getAlpha(context) else 255//透明度
                                                    comment5.orientation = VERTICAL
                                                    //评论加上评论内容
                                                    comment5.addView(addComment(context, commentDetailList[i5], false))
                                                    //加上第五层的内容到第四层
                                                    comment4.addView(comment5, pComment)
                                                }
                                            }
                                            //加上第四层的内容到第三层
                                            comment3.addView(comment4, pComment)
                                        }
                                    }
                                    //加上第三层的内容到第二层
                                    comment2.addView(comment3, pComment)
                                }
                            }
                            //加上第二层的内容到第一层
                            comment.addView(comment2, pComment)
                        }
                    }
                    //加到父布局里面去
                    root.addView(comment, pComment)
                }
            }
    }

    /**
     *
     * @param context Context对象
     * @param data 评论数据
     * @param showReplay 是否显示回复
     * @return View
     */
    private fun addComment(context: Context, data: CommentDetail, showReplay: Boolean): View {
        val lp = LayoutParams(ViewGroup.LayoutParams.WRAP_CONTENT, ViewGroup.LayoutParams.WRAP_CONTENT)
        val lp2 = LayoutParams(ViewGroup.LayoutParams.MATCH_PARENT, ViewGroup.LayoutParams.WRAP_CONTENT)
        val m: Int = dp2px(5f)
        lp2.setMargins(m, m, m, m)
        val comment = LinearLayout(context)
        comment.orientation = VERTICAL
        //头像昵称回复信息
        val commentInfo = LinearLayout(context)
        //添加头像
        val avatar = SimpleDraweeView(context)
        //圆形头像效果
        val roundingParams = RoundingParams.asCircle()
        avatar.hierarchy.roundingParams = roundingParams
        val url = data.avatar
        if (url.contains("http")) {
            avatar.setImageURI(url)
        } else {
            avatar.setImageURI("https:${url}")
        }
        val lAvatar = LayoutParams(dp2px(45f), dp2px(45f))
        commentInfo.addView(avatar, lAvatar)
        //头像旁边的昵称布局
        val commentRight = LinearLayout(context)
        commentRight.orientation = VERTICAL
        //内容稍微居中一点
        commentRight.setPadding(dp2px(5f), dp2px(2f), 0, 0)
        //添加昵称
        val nickname = TextView(context)
        nickname.text = data.nickname
        nickname.textSize = 18f
        commentRight.addView(nickname, lp)
        //添加时间
        val time = TextView(context)
        time.text = data.date
        time.textSize = 15f
        time.setTextColor(context.getColor(R.color.comment_time))
        commentRight.addView(time, lp)
        commentInfo.addView(commentRight, lp)
        if (showReplay) {
            val lReplay = LinearLayout(context)
            //添加回复按钮
            val replay = TextView(context)
            replay.text = "回复"
            replay.background = ContextCompat.getDrawable(context, R.drawable.comment_replay_round)
            replay.setTextColor(context.getColor(R.color.white))
            replay.setPadding(5, 5, 5, 5)
            replay.tag = data.id.toString() + "&&" + data.nickname
            replay.setOnClickListener {
                this.replay(it)
            }
            val pReplay = LayoutParams(ViewGroup.LayoutParams.WRAP_CONTENT, ViewGroup.LayoutParams.WRAP_CONTENT)
            pReplay.topMargin = dp2px(5f)
            pReplay.rightMargin = dp2px(5f)
            lReplay.gravity = Gravity.END
            lReplay.addView(replay, pReplay)
            commentInfo.addView(lReplay, lp2)
        }
        comment.addView(commentInfo, lp2)
        //添加内容
        val content = NoScrollWebView(context)
        content.setLayerType(LAYER_TYPE_SOFTWARE, null)
        content.setBackgroundColor(context.getColor(R.color.transplant))
        val commentContent = "<style>img{max-width:45px;}</style>" + data.content
        content.loadDataWithBaseURL(null, commentContent, "text/html", "UTF-8", null)
        comment.addView(content, lp2)
        return comment
    }


    /**
     *  清除评论数据
     */
    fun clearComments() {
        //清除所有的评论
        root.removeAllViews()
    }
}
