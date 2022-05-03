package com.xiaoyou.library.plugins.friend

import android.annotation.SuppressLint
import android.content.Context
import android.view.animation.AccelerateInterpolator
import android.view.animation.Animation
import android.view.animation.DecelerateInterpolator
import com.xiaoyou.library.common.base.BaseDialog
import com.xiaoyou.library.common.ext.visibleOrGone
import com.xiaoyou.library.common.listener.MyClickListener
import com.xiaoyou.library.common.util.*
import com.xiaoyou.library.net.entity.param.SubmitFriend
import com.xiaoyou.library.net.entity.response.FriendSettings
import com.xiaoyou.library.plugins.R
import com.xiaoyou.library.plugins.databinding.DialogFriendBinding
import java.util.*


/**
 * @description 友链申请dialog
 * @author 小游
 * @data 2021/03/07
 */
class FriendDialog(context: Context): BaseDialog<DialogFriendBinding>(context,0.88f){


    /**翻转动画 */
    private var isOpen = false
    private var centerX = 0f
    private var centerY = 0f
    private val depthZ = 700f
    /**修改此处可以改变距离来达到你满意的效果 */
    private val duration = 300
    /**动画时间 */
    private var openAnimation: Rotate3dAnimation? = null
    private var closeAnimation: Rotate3dAnimation? = null
    /** 友链数据 */
    private lateinit var friend :FriendSettings
    private var param = SubmitFriend()


    /**
     * 设置dialog的内容
     * @param friend 友链设置
     * @return 返回本身
     */
    @SuppressLint("SetTextI18n")
    fun setContent(friend: FriendSettings): FriendDialog {
        this.friend = friend
        // 设置头像
        ImageUtil.setImageByUrl(friend.avatar,mDataBind.userAvatar)
        // 设置显示文字
        mDataBind.dialogContent.text = """
        本站信息
        名字:${friend.name}
        描述:${friend.dec}
        网址:${friend.link}
        头像:${friend.avatar}
        友链申请要求
        1.申请时请先加上本站的连接(〃'▽'〃)
        2.原创博客，非采集站，全站 HTTPS 优先
        * 图标仅支持 png / jpg /gif 等格式，请勿提交 ico 或 分辨率小于 100x100 的图标
        """.trimIndent()
        return this
    }

    /**
     *  设置点击提交的监听事件
     * @param listener MyClickListener<SubmitFriend> 回调函数
     * @return FriendDialog
     */
    fun setOnSubmitListener(listener: MyClickListener<SubmitFriend>): FriendDialog {
        // 点击提交
        mDataBind.dialogBtnSubmit.setOnClickListener{
            // 验证参数
            var message = ""
            when{
                param.name.isEmpty() -> message = "请输入名字"
                param.site.isEmpty() -> message = "请输入站点地址"
                !Validator.isEmailOk(param.email) -> message = "邮箱格式不正确"
            }
            if (message==""){
                // 触发回调事件
                listener.onClick(param)
            } else {
                MyToast.waring(message)
            }
        }
        return this
    }

    /**
     * 控件的初始化
     */
    init {
        //清除背景颜色
        Objects.requireNonNull(window)?.setBackgroundDrawableResource(R.color.transparent)
        // 绑定数据
        mDataBind.item = param
        // 点击申请
        mDataBind.dialogStart.setOnClickListener{
            startAnimation()
        }
        // 剪贴板复制
        mDataBind.dialogCopy.setOnClickListener{
            if(Common.setClipboard("""
             名字:${friend.name}
             描述:${friend.dec}
             网址:${friend.link}
             头像:${friend.avatar}
             """.trimIndent())){
                MyToast.success("复制成功！")
            }else{
                MyToast.error("复制失败!")
            }
        }
        // 点击取消
        mDataBind.dialogBtnCancel.setOnClickListener{
            onBackPressed()
        }
        mDataBind.dialogCancel.setOnClickListener{
            onBackPressed()
        }
    }


    /**
     * 开启旋转动画效果
     */
    private fun startAnimation() {
        //接口回调传递参数
        centerX = (mDataBind.dialogContainer.width / 2).toFloat()
        centerY = (mDataBind.dialogContainer.height / 2).toFloat()
        if (openAnimation == null) {
            initOpenAnim()
            initCloseAnim()
        }
        //用作判断当前点击事件发生时动画是否正在执行
        if (openAnimation!!.hasStarted() && !openAnimation!!.hasEnded()) {
            return
        }
        if (closeAnimation!!.hasStarted() && !closeAnimation!!.hasEnded()) {
            return
        }
        //判断动画执行
        if (isOpen) {
            mDataBind.dialogContainer.startAnimation(openAnimation)
        } else {
            mDataBind.dialogContainer.startAnimation(closeAnimation)
        }
        isOpen = !isOpen
    }


    /**
     * 注意旋转角度
     */
    private fun initOpenAnim() {
        //从0到90度，顺时针旋转视图，此时reverse参数为true，达到90度时动画结束时视图变得不可见，
        openAnimation = Rotate3dAnimation(0f, 90f, centerX, centerY, depthZ, true)
        openAnimation?.duration = duration.toLong()
        openAnimation?.fillAfter = true
        openAnimation?.interpolator = AccelerateInterpolator()
        openAnimation?.setAnimationListener(object : Animation.AnimationListener {
            override fun onAnimationStart(animation: Animation) {}
            override fun onAnimationRepeat(animation: Animation) {}
            override fun onAnimationEnd(animation: Animation) {
                mDataBind.dialogInfo.visibleOrGone(false)
                mDataBind.dialogInput.visibleOrGone(true)
                mDataBind.userAvatar.visibleOrGone(false)
                //从270到360度，顺时针旋转视图，此时reverse参数为false，达到360度动画结束时视图变得可见
                val rotateAnimation = Rotate3dAnimation(270f, 360f, centerX, centerY, depthZ, false)
                rotateAnimation.duration = duration.toLong()
                rotateAnimation.fillAfter = true
                rotateAnimation.interpolator = DecelerateInterpolator()
                mDataBind.dialogContainer.startAnimation(rotateAnimation)
            }
        })
    }


    private fun initCloseAnim() {
        closeAnimation = Rotate3dAnimation(360f, 270f, centerX, centerY, depthZ, true)
        closeAnimation?.duration = duration.toLong()
        closeAnimation?.fillAfter = true
        closeAnimation?.interpolator = AccelerateInterpolator()
        closeAnimation?.setAnimationListener(object : Animation.AnimationListener {
            override fun onAnimationStart(animation: Animation) {}
            override fun onAnimationRepeat(animation: Animation) {}
            override fun onAnimationEnd(animation: Animation) {
                mDataBind.dialogInput.visibleOrGone(true)
                mDataBind.dialogInfo.visibleOrGone(false)
                mDataBind.userAvatar.visibleOrGone(true)
                val rotateAnimation = Rotate3dAnimation(90f, 0f, centerX, centerY, depthZ, false)
                rotateAnimation.duration = duration.toLong()
                rotateAnimation.fillAfter = true
                rotateAnimation.interpolator = DecelerateInterpolator()
                mDataBind.dialogContainer.startAnimation(rotateAnimation)
            }
        })
    }
}