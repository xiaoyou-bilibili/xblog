package com.xiaoyou.xblog.ui.activity

import android.annotation.SuppressLint
import android.content.Intent
import android.os.Bundle
import androidx.lifecycle.ViewModelProvider
import com.bumptech.glide.Glide
import com.lxj.xpopup.XPopup
import com.lxj.xpopup.core.BasePopupView
import com.lxj.xpopup.enums.PopupAnimation
import com.xiaoyou.library.common.base.BaseDbActivity
import com.xiaoyou.library.common.base.WEB
import com.xiaoyou.library.common.ext.dismissLoading
import com.xiaoyou.library.common.ext.isVisible
import com.xiaoyou.library.common.ext.visibleOrGone
import com.xiaoyou.library.common.util.MyToast
import com.xiaoyou.library.common.util.XLog
import com.xiaoyou.library.net.core.DataReceive
import com.xiaoyou.library.net.core.Token
import com.xiaoyou.library.net.entity.param.PostComment
import com.xiaoyou.library.net.entity.param.PostsStatus
import com.xiaoyou.library.net.entity.response.CommentDetail
import com.xiaoyou.library.net.entity.response.PostContentDetail
import com.xiaoyou.library.net.entity.response.UserDetail
import com.xiaoyou.xblog.R
import com.xiaoyou.xblog.databinding.ActivityPostBinding
import com.xiaoyou.library.common.util.ImageUtil
import com.xiaoyou.library.net.entity.param.EncryptContent
import com.xiaoyou.library.net.entity.param.SubmitAdvice
import com.xiaoyou.library.plugins.sponsor.SponsorActivity
import com.xiaoyou.xblog.viewmodel.PostsVM
import com.xiaoyou.xblog.viewmodel.UserVM
import kotlinx.android.synthetic.main.activity_post.view.*

/**
 * @description 文章界面
 * @author 小游
 * @data 2021/02/21
 */
class PostActivity: BaseDbActivity<PostsVM,ActivityPostBinding>() {

    //获取相关参数
    private var img = ""
    // 文章id
    private var id = 0
    // 父节点
    private var parent = 0
    // 用户是否登录
    private var isLogin = false
    // 用户信息
    private lateinit var info : UserDetail
    // 密码弹窗对象
    private var dialog : BasePopupView? = null

    // 用户信息的viewModel
    private val userViewModel by lazy { ViewModelProvider(this).get(UserVM::class.java) }

    override fun initView(savedInstanceState: Bundle?) {
        //获取文章id和背景图
        id = intent.getIntExtra("id", -500)
        img = intent.getStringExtra("img")?:""
        // 首先获取一下用户信息
        if (Token.getToken().userId > 0){
            userViewModel.getUserInfo()
            // 监听用户信息变化
            userViewModel.info.observe(this){
                isLogin = true
                info = it
                ImageUtil.setImageByUrl(it.avatar,mDataBind.commentAvatar)
            }
            // 获取用户文章状态
            mViewModel.getPostStatus(id,object :DataReceive<PostsStatus?>(){
                override fun success(data: PostsStatus?) {
                    if (data is PostsStatus) setStatus(data)
                }
            })
        }
        // 获取文章内容
        mViewModel.getPostContent(id)
        // 获取文章评论
        mViewModel.getComment(id)
        // 初始化UI
        initUI()
    }

    /**
     *  UI界面初始化
     */
    @SuppressLint("SetJavaScriptEnabled", "SetTextI18n", "ResourceAsColor")
    private fun initUI() {
        // 设置透明状态栏
//        ImmersionBar.with(this).titleBar(mDataBind.appBar).init();
        // 返回按钮监听
        mDataBind.toolbar.setOnClickListener{finish()}
        // 设置标题颜色
        mDataBind.collapsingLayout.setCollapsedTitleTextColor(getColor(R.color.white))
        mDataBind.collapsingLayout.setExpandedTitleColor(getColor(R.color.white))
        // 监听文章内容变化
        mViewModel.postContent.observe(this){
            if (it is PostContentDetail){
                // 首先绑定数据
                mDataBind.item = it
                // 加载背景图片
                Glide.with(this).load(if (img!="") img else it.image).into(mDataBind.postBackground)
                setWeb(it.content)
                // 判断文章是否加密
                if (it.encrypt){
                    dialog = XPopup.Builder(this )
                            .isDestroyOnDismiss(true) // 关闭后销毁资源
                            .dismissOnBackPressed(false) // 点击取消不提示
                            .popupAnimation(PopupAnimation.ScrollAlphaFromLeft)
                            .autoDismiss(false) // 不自动消失
                            .dismissOnTouchOutside(false) // 点击外部不关闭
                            .asInputConfirm("密码验证", "文章已加密，请输入访问密码！", "访问密码") {
                        password ->
                        mViewModel.getEncryptContent(id,password,object :DataReceive<EncryptContent?>(){
                            override fun success(data: EncryptContent?) {
                                if (data is EncryptContent){
                                    setWeb(data.content)
                                    dialog?.dismiss()
                                }
                            }
                        })
                    }.show()
                }
            }
        }
        // 监听文章评论变化
        mViewModel.postComment.observe(this){
            if (it is List<CommentDetail>) {
                // 先清空评论
                mDataBind.postComment.clearComments()
                // 再显示评论
                mDataBind.postComment.addComment(it)
            }
        }
        // 监听底部工具栏点击事件
        mDataBind.postBottomAdd.setOnClickListener{
            if (mDataBind.postBottom.isVisible()){
                // 隐藏底部工具栏
                mDataBind.postBottom.visibleOrGone(false)
                mDataBind.postBottomAdd.setTextColor(getColor(R.color.text))
            } else {
                // 显示工具栏
                mDataBind.postBottom.visibleOrGone(true)
                mDataBind.postBottomAdd.setTextColor(getColor(R.color.cyanea_primary))
            }
        }
        // 用户点赞
        mDataBind.postGoodAction.setOnClickListener{updateStatus(true)}
        // 用户收藏
        mDataBind.postCollectAction.setOnClickListener{updateStatus(false)}
        // 赞助被点赞
        mDataBind.postPayAction.setOnClickListener{startActivity(Intent(this,SponsorActivity::class.java))}
        // 分析被点击
        mDataBind.postShareAction.setOnClickListener{MyToast.info("开发中")}
        // 表情点击事件
        mDataBind.commentSmile.setOnClickListener{
            if (mDataBind.faceContent.isVisible()) {
                // 隐藏表情包
                mDataBind.faceContent.visibleOrGone(false)
                mDataBind.commentSmile.setTextColor(getColor(R.color.text))
            } else {
                // 显示表情包
                mDataBind.faceContent.visibleOrGone(true)
                mDataBind.commentSmile.setTextColor(getColor(R.color.cyanea_primary))
            }
        }
        // 表情部分回调
        mDataBind.faceContent.setListener {
            // 获取表情数据
            val face = it.tag.toString()
            // 设置编辑框内容，因为表情数据包含单引号，所以需要去掉
            mDataBind.commentEdit.setText(mDataBind.commentEdit.text.toString() + face.substring(1, face.length - 1))
        }
        // 评论回调事件
        mDataBind.postComment.setReplay{
            // 获取评论id和昵称
            val data = it.tag.toString().split("&&")
            // 设置编辑框提示内容
            mDataBind.commentEdit.hint = "@${data[1]}"
            // 设置parent
            parent = data[0].toInt()
        }
        // 点击发送事件
        mDataBind.commentSend.setOnClickListener{
            if (isLogin){
                if (mDataBind.commentEdit.text.isNotEmpty()){
                    // 构造发送参数
                    val data = PostComment(
                        mDataBind.commentEdit.text.toString(),
                        info.nickname,
                        info.user_id,
                        parent,
                        info.avatar,
                        info.email
                    )
                    // 禁用评论按钮，避免重复评论
                    mDataBind.commentSend.isEnabled = false
                    // 发布评论
                    mViewModel.postComment(id,data,object : DataReceive<PostComment?>(){
                        override fun success(data: PostComment?) {
                            mDataBind.commentSend.isEnabled = true
                            MyToast.success("发表成功！")
                            // 重新获取文章评论
                            mViewModel.getComment(id)
                            // 清空评论框
                            mDataBind.commentEdit.text.clear()
                        }
                        // 错误时的回调
                        override fun error(message: String) {
                            XLog.e("自己的错误回调")
                            mDataBind.commentSend.isEnabled = true
                            MyToast.error(message)
                        }
                    })
                } else {
                    MyToast.waring("评论内容不能为空")
                }
            } else {
                MyToast.waring("请先登录")
            }
        }
    }


    /**
     * 更新文章状态
     * @param flag Boolean 如果为true就是点赞，否则就是收藏
     */
    private fun updateStatus(flag: Boolean){
        // 先判断是否登录
        if (isLogin){
            val good = mDataBind.postBottomGoodText.text.toString() == getText(R.string.post_bottom_good_true)
            val collection = mDataBind.postBottomCollectText.text.toString() == getText(R.string.post_bottom_collect_true)
            // 设置更新的文章状态
            val status = if (flag) PostsStatus(!good, collection)
            else PostsStatus(good, !collection)
            // 更新状态并设置更新结果
            mViewModel.updatePostStatus(id, status,object :DataReceive<PostsStatus?>(){
                override fun success(data: PostsStatus?) {
                    if (data is PostsStatus){
                        setStatus(data)
                    }
                }
            })
        } else {
            MyToast.info("请先登录!")
        }
    }

    /**
     *  显示文章界面内容
     * @param web String 文章内容
     */
    @SuppressLint("SetJavaScriptEnabled")
    private fun setWeb(web:String){
        // 加入CSS和JS
        val content = "<!DOCTYPE html><html><head><meta charset=utf-8><link href='${WEB}/static/css/blog-post.css' rel='stylesheet'></head>" +
                "<body><div class='content'>${web}</div><script src='${WEB}/static/js/prism.js'></script></body></html>"
        // webView设置内容
        mDataBind.postContentWeb.apply {
            // 加载内容
            loadDataWithBaseURL(null,content,"text/html","UTF-8",null)
            // 允许js
            settings.javaScriptEnabled = true
        }
    }


    /**
     *  设置文章的点赞和收藏状态
     * @param status PostsStatus
     */
    private fun setStatus(status: PostsStatus){
        //先设置点赞的
        if (status.good) {
            mDataBind.postBottomGoodText.setText(R.string.post_bottom_good_true)
            mDataBind.postBottomGood.typeface = resources.getFont(R.font.solid)
        } else {
            mDataBind.postBottomGoodText.setText(R.string.post_bottom_good)
            mDataBind.postBottomGood.typeface =  resources.getFont(R.font.regular)
        }
        //设置收藏
        if (status.collection) {
            mDataBind.postBottomCollectText.setText(R.string.post_bottom_collect_true)
            mDataBind.postBottomCollect.typeface =  resources.getFont(R.font.solid)
        } else {
            mDataBind.postBottomCollectText.setText(R.string.post_bottom_collect)
            mDataBind.postBottomCollect.typeface =  resources.getFont(R.font.regular)
        }
    }

    // 界面点击重试
    override fun onLoadRetry() {
        mViewModel.getPostContent(id)
    }

    // 关闭自定义tabbar
    override fun showToolBar() = false

}