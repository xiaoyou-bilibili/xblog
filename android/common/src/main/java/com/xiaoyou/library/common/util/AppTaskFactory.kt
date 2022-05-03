package com.xiaoyou.library.common.util

import com.effective.android.anchors.Project
import com.effective.android.anchors.Task
import com.effective.android.anchors.TaskCreator
import com.facebook.drawee.backends.pipeline.Fresco
import com.jaredrummler.cyanea.Cyanea
import com.kingja.loadsir.callback.SuccessCallback
import com.kingja.loadsir.core.LoadSir
import com.lxj.xpopup.XPopup
import com.scwang.smart.refresh.footer.ClassicsFooter
import com.scwang.smart.refresh.header.ClassicsHeader
import com.scwang.smart.refresh.layout.SmartRefreshLayout
import com.tencent.mmkv.MMKV
import com.xiaoyou.library.common.base.appContext
import com.xiaoyou.library.common.core.ThemeManage
import com.xiaoyou.library.widget.state.EmptyCallback
import com.xiaoyou.library.widget.state.ErrorCallback
import com.xiaoyou.library.widget.state.LoadingCallback
import com.xiaoyou.library.common.listener.KtxActivityLifecycleCallbacks
import es.dmoral.toasty.Toasty
import java.util.*

/**
 * @description 自定义任务处理工厂
 * @author 小游
 * @data 2021/02/21
 */
object TaskCreator : TaskCreator {
    override fun createTask(taskName: String): Task {
        return when (taskName) {
            InitComm.TASK_ID -> InitComm()
            InitUtils.TASK_ID -> InitUtils()
            InitToast.TASK_ID -> InitToast()
            InitAppLifecycle.TASK_ID -> InitAppLifecycle()
            else -> InitDefault()
        }
    }
}

class InitDefault : Task(TASK_ID, true) {
    companion object {
        const val TASK_ID = "0"
    }

    override fun run(name: String) {

    }
}


//初始化常用控件类
class InitComm : Task(TASK_ID, true) {
    companion object {
        const val TASK_ID = "2"
    }

    override fun run(name: String) {
        // 注册界面状态管理
        LoadSir.beginBuilder()
        .addCallback(ErrorCallback())
        .addCallback(EmptyCallback())
        .addCallback(LoadingCallback())
        .setDefaultCallback(SuccessCallback::class.java)
        .commit()
        // 初始化主题引擎
        Cyanea.init(appContext, appContext.resources)
        // 初始化下拉加载框架
        // 初始化设置刷新组件的设置
        // 全局设置刷新方式 下面这种方式是kotlin lambda的用法
        // 设置全局加载方式
        SmartRefreshLayout.setDefaultRefreshInitializer { _, layout ->
            // 设置自动刷新
            layout.autoRefresh()
        }
        //设置上拉刷新
        SmartRefreshLayout.setDefaultRefreshHeaderCreator { context, _ ->
            // 设置雷达头部作为默认上拉刷新样式
            ClassicsHeader(context)
        }
        // 设置下拉加载
        SmartRefreshLayout.setDefaultRefreshFooterCreator{ context, _ ->
            // 设置下拉加载样式
            ClassicsFooter(context)
        }
        // 图片加载库初始化
        Fresco.initialize(appContext)
        // 初始化弹窗主题
        XPopup.setPrimaryColor(appContext.getColor(ThemeManage.getTheme().primary))
    }
}

//初始化Utils
class InitUtils : Task(TASK_ID, true) {
    companion object {
        const val TASK_ID = "3"
    }
    override fun run(name: String) {
        //初始化Log打印,我们显示日志
        XLog.init(true)
        //初始化MMKV
        MMKV.initialize(appContext.filesDir.absolutePath + "/mmkv")
    }
}

//初始化Utils
class InitToast : Task(TASK_ID, false) {
    companion object {
        const val TASK_ID = "4"
    }

    override fun run(name: String) {
        //初始化吐司
        Toasty.Config.getInstance()
            .setTextSize(15) // optional
            .allowQueue(true) // optional (prevents several Toastys from queuing)
            .apply(); // required
    }
}

//初始化Utils
class InitAppLifecycle : Task(TASK_ID, true) {
    companion object {
        const val TASK_ID = "5"
    }
    override fun run(name: String) {
        //注册全局的Activity生命周期管理
        appContext.registerActivityLifecycleCallbacks(KtxActivityLifecycleCallbacks())
    }
}




class AppTaskFactory : Project.TaskFactory(TaskCreator)

/**
 * 模拟初始化SDK
 * @param millis Long
 */
fun doJob(millis: Long) {
    val nowTime = System.currentTimeMillis()
    while (System.currentTimeMillis() < nowTime + millis) {
        //程序阻塞指定时间
        val min = 10
        val max = 99
        val random = Random()
        val num = random.nextInt(max) % (max - min + 1) + min
    }
}