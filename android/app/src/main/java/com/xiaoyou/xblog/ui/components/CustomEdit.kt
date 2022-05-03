package com.xiaoyou.xblog.ui.components

import android.view.View
import android.widget.EditText
import androidx.core.content.ContextCompat
import com.xiaoyou.library.common.base.appContext
import com.xiaoyou.xblog.R

/**
 * @description 自定义编辑框组件
 * @author 小游
 * @data 2021/03/04
 */
object CustomEdit {
    /**
     * 用于修复EditText背景切换的Bug
     * 即如果没有这个函数,变为红色边框的EditText失去Focus时依旧是红色边框,这不符合失去焦点时无边框背景的bug
     * @receiver EditText
     */
    fun setEditTextBackground(editText: EditText) {
        editText.onFocusChangeListener = View.OnFocusChangeListener { v, hasFocus ->
            when (hasFocus) {
                false -> editText.background = v
                        ?.context
                        ?.let { ContextCompat.getDrawable(it, R.drawable.sign_activity_edit_text_selector) }
            }
        }
    }
    /**
     *  设置错误提示
     * @param editText EditText
     */
    fun setError(editText: EditText){
        editText.background = ContextCompat.getDrawable(appContext, R.drawable.sign_activity_edit_text_error_bk)
    }
}