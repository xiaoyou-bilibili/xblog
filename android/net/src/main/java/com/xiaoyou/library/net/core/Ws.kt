package com.xiaoyou.library.net.core

import android.content.Context
import com.google.gson.Gson
import com.google.gson.reflect.TypeToken
import com.xiaoyou.library.net.R
import com.xiaoyou.library.net.entity.param.SendMessage
import com.xiaoyou.library.net.entity.param.UserGet
import com.xiaoyou.library.net.entity.param.UserSend
import com.xiaoyou.library.net.entity.response.ChatInfo
import com.xiaoyou.library.net.entity.response.WsReturn
import okhttp3.*
import java.lang.Exception

/**
 * @description webSocket请求核心库
 * @author 小游
 * @data 2021/03/01
 */
object Ws {
    // api链接
    val WS = R.string.WS
    // 当前链接的webSocket对象
    var mWebSocket : WebSocket? = null
    // client对象
    private  var client : OkHttpClient? = null

    // 创建webSocket链接
    private fun webSocketCreate(address:String,listener:WebSocketListener,context: Context){
        client = OkHttpClient.Builder().build()
        val request = Request.Builder().url(context.getString(WS)+address).build()
        client?.newWebSocket(request,listener)
    }

    // 创建一个聊天的webSocket链接
    fun createChatSocket(wsListener: WsListener,context: Context){
        val token = Token.getToken()
        webSocketCreate("/ws/v1/chat?id=${token.userId}&token=${token.token}",object :WebSocketListener(){
            override fun onOpen(webSocket: WebSocket, response: Response) {
                // 设置当前webSocket
                mWebSocket = webSocket
                // 回调
                wsListener.onOpen(webSocket)
            }

            override fun onMessage(webSocket: WebSocket, text: String) {
                wsListener.onMessage(webSocket,text)
                try {
                    // 尝试解析数据
                    val data = parseData<WsReturn<Any?>>(text)
//                    val row = data.data.toString()
//                    Log.e("xiaoyou",text)
//                    Log.e("xiaoyou",parseData<WsReturn<List<ChatInfo>>>(text).data.toString())

                    when(data.code){
                        // 0表示错误
                        0 -> wsListener.onError(data.message)
                        // 1表示成功
                        1 -> wsListener.onSuccess(data.message)
                        // 2表示有新的消息
                        2 -> wsListener.onNewMessage(parseData<WsReturn<ChatInfo>>(text).data)
                        // 3表示获取到了数据
                        3 -> wsListener.onMessageReceive(parseData<WsReturn<List<ChatInfo>>>(text).data)
                        // 4表示获取历史数据
                        4 -> wsListener.onHistoryReceive(parseData<WsReturn<List<ChatInfo>>>(text).data)
                    }
                }catch (e: Exception){
//                    Log.e("xiaoyou",e.message)
                    wsListener.onError("数据解析异常")
                }
            }
            override fun onFailure(webSocket: WebSocket, t: Throwable, response: Response?) {
//                Log.e("xiaoyou","发送失败")
                wsListener.onError(t.message?:"")
            }

            override fun onClosing(webSocket: WebSocket, code: Int, reason: String) {
//                super.onClosing(webSocket, code, reason)
//                Log.e("xiaoyou","连接已关闭")
                // 关闭时webSocket置为空
                mWebSocket = null
                // client也置为空
                client = null
                wsListener.onError("连接已关闭")
            }
        },context)
    }

    /**
     *  用户提交信息
     * @param to Int 发送的对象
     * @param content String 发送的内容
     * @param messageType Int 发送消息的类型
     */
    fun commitMessage(to :Int,content:String,messageType:Int = 1){
        // 获取用户token信息
        val token=Token.getToken()
        val message = SendMessage(
            token.userId,
            token.token,
            to,
            "send",
            UserSend(messageType,content)
        )
        mWebSocket?.send(decodeDara(message))
    }

    /**
     *  用户获取内容
     * @param to Int 发送的对象
     * @param date Long 历史记录
     * @param size Int 大小
     */
    fun getContent(to: Int,date:Long = 0,size :Int = 10){
        // 获取用户token信息
        val token=Token.getToken()
        val message = SendMessage(
            token.userId,
            token.token,
            to,
            "get",
            UserGet(date,size)
        )
        mWebSocket?.send(decodeDara(message))
    }

    /**
     *  解析数据
     * @param row String 原始数据
     * @return WsReturn<T> 解析后的数据
     */
    inline fun <reified T> parseData(row :String): T{
       return Gson().fromJson(row, object: TypeToken<T>(){}.type)
    }

    /**
     *  把数据转换为json编码
     * @param data Any 数据
     * @return String 返回的字符串
     */
    private fun decodeDara(data: Any): String = Gson().toJson(data)
}