package com.xiaoyou.library.common.base

import androidx.lifecycle.ViewModel
import com.kunminx.architecture.ui.callback.UnPeekLiveData
import com.xiaoyou.library.net.entity.base.LoadStatusEntity
import com.xiaoyou.library.net.entity.base.LoadingDialogEntity

/**
 * @description 基础viewmodel，包含几个请求的状态
 * @author 小游
 * @data 2021/02/20
 */
open class BaseViewModel : ViewModel() {


    val loadingChange: UiLoadingChange by lazy { UiLoadingChange() }

    // 当前网络的加载状态，activity中可以通过监听这个来确定当前网络的请求状态
    inner class UiLoadingChange {
        //请求时loading
        val loading by lazy { UnPeekLiveData<LoadingDialogEntity>() }

        //界面显示空布局
        val showEmpty by lazy { UnPeekLiveData<LoadStatusEntity>() }

        //界面显示错误布局
        val showError by lazy { UnPeekLiveData<LoadStatusEntity>() }

        //界面显示成功布局
        val showSuccess by lazy { UnPeekLiveData<Boolean>() }
    }
}