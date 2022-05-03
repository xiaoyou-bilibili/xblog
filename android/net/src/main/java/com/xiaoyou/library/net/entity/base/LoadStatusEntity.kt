package com.xiaoyou.library.net.entity.base

/**
 * @description 数据加载状态显示
 * @author 小游
 * @data 2021/02/21
 */
data class LoadStatusEntity(
    // 失败的异常类
    var throwable: Throwable,
    // 错误消息
    var errorMessage: String,
    // 是否是列表类型刷新
    var isRefresh: Boolean = false,
    // 加载错误时，显示的错误类型
    @LoadingType var errorType: Int = LoadingType.LOADING_NULL, // loading类型
    // 如果是请求失败，重新请求的参数
    var intentData: Any? = null
)