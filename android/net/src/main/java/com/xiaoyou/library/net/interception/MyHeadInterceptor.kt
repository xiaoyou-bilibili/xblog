package com.xiaoyou.library.net.interception

import android.util.Log
import com.tencent.mmkv.MMKV
import com.xiaoyou.library.net.core.Token
import okhttp3.Interceptor
import okhttp3.Response
import java.io.IOException
import kotlin.jvm.Throws

/**
 * @description 自定义头部拦截器,自动添加token信息
 * @author 小游
 * @data 2021/02/20
 */
class MyHeadInterceptor : Interceptor {

    @Throws(IOException::class)
    override fun intercept(chain: Interceptor.Chain): Response {
        val builder = chain.request().newBuilder()
        // 首先我们获取token信息
        val token = Token.getToken()
        if (token.userId != 0){
            // 这里我们手动添加user_id和token这两个参数
            builder.addHeader(Token.USER_ID, "${token.userId}").build()
            builder.addHeader(Token.TOKEN, token.token).build()
        }
        return chain.proceed(builder.build())
    }
}