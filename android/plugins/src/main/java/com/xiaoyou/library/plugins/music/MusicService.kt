package com.xiaoyou.library.plugins.music

import android.app.IntentService
import android.content.Intent
import androidx.localbroadcastmanager.content.LocalBroadcastManager
import com.xiaoyou.library.common.base.appContext
import com.xiaoyou.library.common.util.XLog


/**
 * @description 自定义音乐播放服务
 *
 * @author 小游
 * @data 2021/03/10
 */
// 服务参考 https://developer.android.com/guide/components/services?hl=zh-cn
class MusicService: IntentService("musicService"){

    // 这个就是发送命令
    override fun onStartCommand(intent: Intent?, flags: Int, startId: Int): Int {
        // 获取action信息
        when(intent?.action){
            PlayAction.PLAY.action -> PlayerManager.start()
            PlayAction.PAUSE.action -> PlayerManager.pause()
            PlayAction.PREVIOUS.action -> PlayerManager.previous()
            PlayAction.NEXT.action -> PlayerManager.next()
        }
        // 发送广播
        sendBroadCast(intent?.action)
        return super.onStartCommand(intent, flags, startId)
    }

    // 发送广播
    private fun sendBroadCast(action:String?){
        // 发送广播
        LocalBroadcastManager.getInstance(appContext).sendBroadcast(Intent(action))
    }

    override fun onHandleIntent(intent: Intent?) {
    }

}