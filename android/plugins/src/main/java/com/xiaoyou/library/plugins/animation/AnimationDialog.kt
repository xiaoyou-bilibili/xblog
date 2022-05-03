package com.xiaoyou.library.plugins.animation

import android.content.Context
import android.content.Intent
import android.net.Uri
import android.text.method.ScrollingMovementMethod
import com.xiaoyou.library.common.base.BaseDialog
import com.xiaoyou.library.common.base.appContext
import com.xiaoyou.library.common.util.Common.openUrl
import com.xiaoyou.library.net.entity.response.AnimationDetail
import com.xiaoyou.library.plugins.databinding.DialogAnimationBinding


/**
 * @description 我的追番界面的dialog
 * @author 小游
 * @data 2021/03/06
 */
class AnimationDialog(context: Context, val animation: AnimationDetail) : BaseDialog<DialogAnimationBinding>(context) {

    init {
        // 设置textView可以滚动
        mDataBind.animationDec.movementMethod = ScrollingMovementMethod.getInstance()
        // 绑定视图
        mDataBind.item = animation
        // 设置图片
        mDataBind.animationCover.setImageURI(animation.cover)
        // 点击播放
        mDataBind.playAnimation.setOnClickListener{
            openUrl(animation.url)
        }
    }

}