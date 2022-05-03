package com.xiaoyou.library.net.core

import com.google.gson.Gson
import com.xiaoyou.library.net.entity.response.ReturnError
import retrofit2.*

/**
 * @description call对象处理  https://juejin.cn/post/6844904041869213703
 * @author 小游
 * @data 2021/02/20
 */
object CallHandle {
    // 处理call请求，自己手动转换为协程的形式
     fun <T> handleCall(call: Call<T>,dataReceive: DataReceive<T?>) {
        // 使用suspendCoroutine来把进程给挂起
//        return suspendCoroutine { continuation ->
            call.enqueue(object: Callback<T>{
                // 请求成功调用该方法
                override fun onResponse(call: Call<T>, response: Response<T>) {
                    // 判断我们的请求是否成功
                    if (response.isSuccessful) {
                        // 请求成功后我们直接返回body
                        dataReceive.success(response.body())
//                        continuation.resume(response.body())
                    } else {
                        // 如果请求失败，我们就获取错误内容
                        var error = response.errorBody()?.string()
                        error = if (error != ""){
                            // 那我们解析一下错误的内容
                            Gson().fromJson(error,ReturnError::class.java).message
                        } else {
                            "未知错误"
                        }
                        // 返回我们的异常
                        dataReceive.error(error)
//                        continuation.resumeWithException(Exception(error))
                    }
                }
                // 请求失败时调用该方法
                override fun onFailure(call: Call<T>, t: Throwable) {
//                    continuation.resumeWithException(t)
                    dataReceive.error(t.message?:"")
                }

            })
    }

}