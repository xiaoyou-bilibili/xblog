package com.xiaoyou.library.widget.state

import com.kingja.loadsir.callback.Callback
import com.xiaoyou.library.widget.R

/**
 * @description 获取数据失败时的回调函数
 * @author 小游
 * @data 2021/02/20
 */
class ErrorCallback: Callback() {
    override fun onCreateView(): Int {
        return R.layout.layout_error
    }
}