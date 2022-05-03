package com.xiaoyou.library.net.core

import com.xiaoyou.library.net.entity.response.ChatInfo
import okhttp3.WebSocket

/**
 * @description webSocket获取数据的接口
 * @author 小游
 * @data 2021/03/01
 */
interface WsListenerInterface {
    fun onMessage(webSocket: WebSocket, message: String)
    fun onNewMessage(message: ChatInfo)
    fun onMessageReceive(message: List<ChatInfo>?)
    fun onHistoryReceive(message: List<ChatInfo>?)
    fun onOpen(webSocket: WebSocket)
    fun onError(err: String)
    fun onSuccess(info: String)
}