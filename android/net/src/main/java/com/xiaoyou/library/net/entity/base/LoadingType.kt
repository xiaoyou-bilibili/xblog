package com.xiaoyou.library.net.entity.base

import androidx.annotation.IntDef

/**
 * @description 加载类型，这里包括加载中，加载失败的两个类型
 * @author 小游
 * @data 2021/02/21
 */
@IntDef(LoadingType.LOADING_NULL, LoadingType.LOADING_DIALOG, LoadingType.LOADING_XML)
@Retention(AnnotationRetention.SOURCE)
annotation class LoadingType {
    companion object {
        // 请求时 不需要Loading
        const val LOADING_NULL = 0
        // 请求时 弹出 Dialog弹窗Loading
        const val LOADING_DIALOG = 1
        // 请求错误时显示一个错误界面
        const val LOADING_XML = 2
        // 什么错误也不显示
        const val LOADING_ERROR_NULL = 3
    }
}