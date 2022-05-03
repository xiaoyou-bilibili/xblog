package com.xiaoyou.library.common.components

import android.content.Context
import android.util.AttributeSet
import android.webkit.WebView

/**
 *作者:created by HP on 2021/2/14 21:43
 *邮箱:sakurajimamai2020@qq.com
 */
class NoScrollWebView : WebView {
    constructor(context: Context?, attrs: AttributeSet?, defStyleAttr: Int, defStyleRes: Int) : super(context!!, attrs, defStyleAttr, defStyleRes) {}
    constructor(context: Context?, attrs: AttributeSet?, defStyleAttr: Int) : super(context!!, attrs, defStyleAttr) {}
    constructor(context: Context?, attrs: AttributeSet?) : super(context!!, attrs) {}
    constructor(context: Context?) : super(context!!) {}

    override fun onMeasure(widthMeasureSpec: Int, heightMeasureSpec: Int) {
        val mExpandSpec = MeasureSpec.makeMeasureSpec(Int.MAX_VALUE shr 2, MeasureSpec.AT_MOST)
        super.onMeasure(widthMeasureSpec, mExpandSpec)
    }
}