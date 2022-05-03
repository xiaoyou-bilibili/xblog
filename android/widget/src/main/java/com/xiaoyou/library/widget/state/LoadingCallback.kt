package com.xiaoyou.library.widget.state

import com.kingja.loadsir.callback.Callback
import com.xiaoyou.library.widget.R

/**
 * @description 显示加载中的界面
 * @author 小游
 * @data 2021/02/20
 */
class LoadingCallback : Callback() {
    override fun onCreateView(): Int {
        return R.layout.layout_error
    }
}