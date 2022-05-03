package com.xiaoyou.xblog.viewmodel

import androidx.lifecycle.MutableLiveData
import com.xiaoyou.library.common.base.BaseViewModel
import com.xiaoyou.library.common.base.appContext
import com.xiaoyou.library.common.util.XLog
import com.xiaoyou.library.net.core.Ws
import com.xiaoyou.library.net.core.WsListener
import com.xiaoyou.library.net.entity.response.ChatInfo
import com.xiaoyou.library.net.entity.response.DiaryDetail
import com.xiaoyou.library.net.entity.response.ReturnList
import com.xiaoyou.xblog.data.chat.ChatMessage
import com.xiaoyou.xblog.data.chat.ChatUser
import okhttp3.WebSocket
import java.util.*

/**
 * @description 聊天界面viewModel对象
 * @author 小游
 * @data 2021/03/01
 */
class ChatVM  : BaseViewModel(){
    // 新消息对象
    val newMessage = MutableLiveData<ChatInfo>()
    // 获取到的消息
    val receiveMessage = MutableLiveData<List<ChatInfo>>()
    // 获取到的历史消息
    val historyMessage = MutableLiveData<List<ChatInfo>>()

    /**
     * 建立webSocket连接，因为okHttp会自动把websocket的请求放入后台
     * 所以我们不能通过回调来获取数据，我们必须使用liveData来监听数据变化
     * @param wsListener WsListener
     */
    fun startChatSocket(wsListener: WsListener) = Ws.createChatSocket(object : WsListener(){
        // 当收到消息
        override fun onMessageReceive(message: List<ChatInfo>?) = receiveMessage.postValue(message)
        // 当有新消息时
        override fun onNewMessage(message: ChatInfo) = newMessage.postValue(message)
        // 当获取到历史消息时
        override fun onHistoryReceive(message: List<ChatInfo>?) = historyMessage.postValue(message)
        // 有消息时
        override fun onMessage(webSocket: WebSocket, message: String)  = wsListener.onMessage(webSocket,message)
        // 建立连接
        override fun onOpen(webSocket: WebSocket) = wsListener.onOpen(webSocket)
        // 当有错误发生时
        override fun onError(err: String) = wsListener.onError(err)
    }, appContext)


}