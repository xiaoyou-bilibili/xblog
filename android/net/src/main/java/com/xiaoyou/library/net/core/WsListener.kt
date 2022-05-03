package com.xiaoyou.library.net.core

import com.xiaoyou.library.net.entity.response.ChatInfo
import com.xiaoyou.library.net.entity.response.WsReturn
import okhttp3.WebSocket

/**
 * @description 自定义webSocket回调事件
 * @author 小游
 * @data 2021/03/01
 */
abstract class WsListener : WsListenerInterface{
    /**
     *  当有消息的时候,这个是通用的，只要有数据就会调用
     * @param webSocket WebSocket webSocket对象
     * @param message String 发送的消息
     */
    override fun onMessage(webSocket: WebSocket, message: String){}

    /**
     *  当有新消息的时候通知
     * @param message WsReturn<ChatInfo>
     */
    override fun onNewMessage(message: ChatInfo){}


    /**
     *  当我们获取到数据时
     * @param message WsReturn<List<ChatInfo>>
     */
    override fun onMessageReceive(message: List<ChatInfo>?){}

    /**
     *  当获取到历史数据时
     * @param message WsReturn<List<ChatInfo>>
     */
    override fun onHistoryReceive(message: List<ChatInfo>?){}

    /**
     *  当链接建立的时候
     * @param webSocket WebSocket
     */
    override fun onOpen(webSocket: WebSocket){}

    /**
     *  当出现错误时
     * @param err String 错误信息
     */
    override fun onError(err: String){}

    /**
     *  当我们的请求通过的时候
     * @param info String
     */
    override fun onSuccess(info: String) {}
}