package com.xiaoyou.library.plugins.douban

import android.content.Context
import android.content.Intent
import android.net.Uri
import android.text.method.ScrollingMovementMethod
import com.xiaoyou.library.common.base.BaseDialog
import com.xiaoyou.library.common.base.appContext
import com.xiaoyou.library.common.util.Common
import com.xiaoyou.library.net.entity.response.DouBanDetail
import com.xiaoyou.library.plugins.databinding.DialogDouBanBinding

/**
 * @description 自定义豆瓣弹窗
 * @author 小游
 * @data 2021/03/08
 */
class DouBanDialog(context: Context,douBanDetail: DouBanDetail) :BaseDialog<DialogDouBanBinding>(context) {
    init {
        // 赋值
        mDataBind.item = douBanDetail
        mDataBind.douBanImg.setImageURI(douBanDetail.picture)
        // 设置textView滚动
        mDataBind.douBanComment.movementMethod = ScrollingMovementMethod.getInstance()
        // 点击打开网址
        mDataBind.openLink.setOnClickListener{
            Common.openUrl(douBanDetail.url)
        }
    }
}