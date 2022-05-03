package com.xiaoyou.library.common.ext

import com.xiaoyou.library.widget.R
import com.xiaoyou.library.widget.toolbar.CustomToolBar

/**
 * @description 自定义顶部标题栏
 * @author 小游
 * @data 2021/02/21
 */

/**
 * 初始化有返回键的toolbar
 */
fun CustomToolBar.initBack(
    titleStr: String = "标题",
    backImg: Int = R.drawable.ic_back,
    onBack: (toolbar: CustomToolBar) -> Unit
): CustomToolBar {
    this.setCenterTitle(titleStr)
    this.getBaseToolBar().setNavigationIcon(backImg)
    this.getBaseToolBar().setNavigationOnClickListener { onBack.invoke(this) }
    return this
}



