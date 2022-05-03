package com.xiaoyou.library.common.base

import android.app.Application
import android.view.View
import android.widget.ImageView
import androidx.lifecycle.ViewModelProvider
import androidx.lifecycle.ViewModelStore
import androidx.lifecycle.ViewModelStoreOwner
import com.bumptech.glide.Glide
import com.effective.android.anchors.AnchorsManager
import com.effective.android.anchors.BuildConfig
import com.effective.android.anchors.Project
import com.tencent.mmkv.MMKV
import com.xiaoyou.library.common.R
import com.xiaoyou.library.common.util.*
import com.xiaoyou.library.net.core.ServiceCreator


/**
 * @description 自己定义的baseApplication用于完成控件初始化
 * @author 小游
 * @data 2021/02/20
 */
//全局上下文
val appContext: BaseApplication by lazy { BaseApplication.instance }
// api和web的链接
val SERVER by lazy { appContext.getString(R.string.API) }
val WEB by lazy { appContext.getString(R.string.WEB) }
// 全局kv对象,注意需要延迟加载，只有当我们需要的时候才调用
val kv: MMKV by lazy {  MMKV.defaultMMKV() }
// 自定义application
open class BaseApplication : Application(), ViewModelStoreOwner {

    companion object {
        lateinit var instance: BaseApplication
    }

    private lateinit var mAppViewModelStore: ViewModelStore

    private var mFactory: ViewModelProvider.Factory? = null

    override fun onCreate() {
        super.onCreate()
        instance = this
        // 存储当前的viewmodel
        mAppViewModelStore = ViewModelStore()
        // 主进程初始化
        onMainProcessInit()
    }

    /**
     *  主进程初始化
     */
    open fun onMainProcessInit() {
        // 因为我们需要初始化一些东西，我们的初始化代码需要异步执行，所以使用AnchorsManager
        AnchorsManager.getInstance()
            .debuggable(BuildConfig.DEBUG)
            // 设置不同任务的id并依次来执行这些初始化任务
            .addAnchor(
                InitUtils.TASK_ID,
                InitComm.TASK_ID,
                InitAppLifecycle.TASK_ID,
                InitToast.TASK_ID
            ).start(
                Project.Builder("app", AppTaskFactory())
                    .add(InitComm.TASK_ID)
                    .add(InitUtils.TASK_ID)
                    .add(InitToast.TASK_ID)
                    .add(InitAppLifecycle.TASK_ID)
                    .build()
            )
    }


    /**
     * 获取一个全局的ViewModel
     */
    fun getAppViewModelProvider(): ViewModelProvider {
        return ViewModelProvider(this, this.getAppFactory())
    }

    // 通过ViewModelProvider.Factory 来创建viewmodel对象
    private fun getAppFactory(): ViewModelProvider.Factory {
        if (mFactory == null) {
            mFactory = ViewModelProvider.AndroidViewModelFactory.getInstance(this)
        }
        return mFactory as ViewModelProvider.Factory
    }

    override fun getViewModelStore(): ViewModelStore {
        return mAppViewModelStore
    }
}
