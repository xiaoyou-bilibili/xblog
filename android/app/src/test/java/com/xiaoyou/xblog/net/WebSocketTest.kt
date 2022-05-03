package com.xiaoyou.xblog.net

import android.util.Log
import com.xiaoyou.library.common.base.appContext
import com.xiaoyou.library.net.core.Ws
import com.xiaoyou.library.net.core.WsListener
import com.xiaoyou.library.net.entity.response.ChatInfo
import com.xiaoyou.library.net.entity.response.WsReturn
import kotlinx.coroutines.delay
import okhttp3.*
import okio.ByteString
import org.junit.Assert
import org.junit.Test
import java.lang.Thread.sleep
import java.nio.charset.Charset
import java.util.*
import kotlin.concurrent.thread

/**
 * @description
 * @author 小游
 * @data 2021/03/01
 */
class WebSocketTest {
    @Test
    fun webSocketTest() {
        Ws.createChatSocket(object :WsListener(){
            override fun onError(err: String) {
                println(err)
            }
            override fun onMessage(webSocket: WebSocket, message: String) {
                println(message)
            }

            override fun onMessageReceive(message: List<ChatInfo>?) {
                println("获取数据")
                println(message?.get(0)?.content)
            }

            override fun onHistoryReceive(message: List<ChatInfo>?) {
                println("获取历史数据")
                println(message?.get(0)?.content)
            }

            override fun onNewMessage(message: ChatInfo) {
                println("有新消息")
                println(message.content)
            }

        }, appContext)
        sleep(2000L*5)
        println("${Ws.WS}ws/v1/chat")
    }


}