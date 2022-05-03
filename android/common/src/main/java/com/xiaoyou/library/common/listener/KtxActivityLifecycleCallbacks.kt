package com.xiaoyou.library.common.listener

import android.app.Activity
import android.app.Application
import android.os.Bundle
import com.xiaoyou.library.common.ext.addActivity
import com.xiaoyou.library.common.util.XLog


/**
 * @description 生命周期监听
 * @author 小游
 * @data 2021/02/21
 */
class KtxActivityLifecycleCallbacks : Application.ActivityLifecycleCallbacks {
    override fun onActivityPaused(p0: Activity) {

    }

    override fun onActivityStarted(p0: Activity) {

    }

    override fun onActivityDestroyed(activity: Activity) {
//        removeActivity(activity)
    }

    override fun onActivitySaveInstanceState(p0: Activity, p1: Bundle) {
    }

    override fun onActivityStopped(p0: Activity) {
    }

    override fun onActivityCreated(activity: Activity, p1: Bundle?) {
        XLog.d(activity.javaClass.simpleName)
//        addActivity(activity)
    }

    override fun onActivityResumed(p0: Activity) {
    }

}