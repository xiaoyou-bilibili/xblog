package com.xiaoyou.library.net.core

import android.content.Context
import com.xiaoyou.library.net.R
import com.xiaoyou.library.net.interception.MyHeadInterceptor
import okhttp3.Dispatcher
import okhttp3.OkHttpClient
import retrofit2.Retrofit
import retrofit2.converter.gson.GsonConverterFactory
import java.util.concurrent.TimeUnit

/**
 * @description retrofit核心库，这里负责封装请求
 * @author 小游
 * @data 2021/02/13
 */
// 参考 https://github.com/kongxiaoan/Network-Demo
object ServiceCreator {
    // api链接
    private val SERVER = R.string.API
    // 封装httpClient请求
    private val okHttpClient by lazy { OkHttpClient().newBuilder() }
    //  retrofit对象
    private fun retrofit(context: Context) :Retrofit  {
        val builder = Retrofit.Builder()
            .baseUrl("${context.getString(SERVER)}/api/v3/")
            // 添加转换工厂，这里我们使用gson来进行转换
            .addConverterFactory(GsonConverterFactory.create())
        // Dispatcher是负责对okhttp所有的请求进行调度管理的类，可以获取并取消所有请求
        val dispatcher = Dispatcher()
        dispatcher.maxRequests = 1
        // 建立okhttp链接,最后我们设置头部拦截来自动给头部加上token等参数
        okHttpClient
            .connectTimeout(10, TimeUnit.SECONDS)
            .writeTimeout(10, TimeUnit.SECONDS)
            .readTimeout(10, TimeUnit.SECONDS)
            .dispatcher(dispatcher)
            .addInterceptor(MyHeadInterceptor())
        return builder.client(okHttpClient.build()).build()
    }

    fun <T> create(clazz: Class<T>,context: Context): T = retrofit(context).create(clazz)

    // 通过interface接口来创建服务对象,这里需要我们传入一个context对象用于获取地址信息
    inline fun <reified T> createService(clazz: Class<T>,context: Context): T = create(clazz,context)
}