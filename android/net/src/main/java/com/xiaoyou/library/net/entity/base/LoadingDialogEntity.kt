package com.xiaoyou.library.net.entity.base


/**
 * @description 加载中样式
 * @author 小游
 * @data 2021/02/21
 */
data class LoadingDialogEntity(
    // 加载中的样式类型
    @LoadingType var loadingType: Int = LoadingType.LOADING_NULL,
    // 显示的文字
    var loadingMessage: String = "",
    // 是否显示
    var isShow: Boolean = false
)