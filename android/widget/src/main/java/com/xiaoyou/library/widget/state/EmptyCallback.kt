package com.xiaoyou.library.widget.state

import com.kingja.loadsir.callback.Callback
import com.xiaoyou.library.widget.R


/**
 * @description 没有数据时的回调，这里我们会返回一个空的状态
 * @author 小游
 * @data 2021/02/20
 */
class EmptyCallback : Callback() {
    override fun onCreateView(): Int {
        return R.layout.layout_empty
    }
}