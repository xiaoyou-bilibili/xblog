package com.xiaoyou.library.common.ext

import com.google.gson.Gson
import com.xiaoyou.library.common.base.BaseViewModel
import com.xiaoyou.library.net.entity.base.LoadStatusEntity
import com.xiaoyou.library.net.entity.base.LoadingDialogEntity
import com.xiaoyou.library.net.entity.base.LoadingType
import com.xiaoyou.library.net.entity.response.ReturnError
import retrofit2.Call
import retrofit2.Callback
import retrofit2.Response
import java.lang.Exception

/**
 * @description 网络处理的handle
 * @author 小游
 * @data 2021/02/21
 */
fun <T> BaseViewModel.handleCall(request: HttpRequest<T>.() -> Unit) {
    // 获取传入的对象，并初始化
    val httpRequest = HttpRequest<T>()
    request(httpRequest)
    // 设置状态为加载中
    loadingChange.loading.value = LoadingDialogEntity(httpRequest.loadingType, httpRequest.loadingMessage, true)
    // 加载结束
    val loadingStop = LoadingDialogEntity(httpRequest.loadingType, httpRequest.loadingMessage, false)
    // 如果call对象不为空，那么我们就执行异步请求
    httpRequest.call?.enqueue(object: Callback<T> {
        // 请求成功调用该方法
        override fun onResponse(call: Call<T>, response: Response<T>) {
            // 结束加载状态
            loadingChange.loading.value = loadingStop
            // 判断我们的请求是否成功
            if (response.isSuccessful) {
                // 执行回调函数，返回请求成功的对象
                httpRequest.onSuccess(response.body())
                // 显示成功,避免错误布局一直显示
                loadingChange.showSuccess.value = true
            } else {
                // 如果请求失败，我们就获取错误内容
                var error = response.errorBody()?.string()
                // 对错误进行解析
                error = try {
                    Gson().fromJson(error, ReturnError::class.java).message
                } catch (e: Exception){
                    "未知错误"
                }
                // 判断错误是否需要自己处理
                if (httpRequest.onError == null) {
                    // 返回我们的异常
                    loadingChange.showError.value = LoadStatusEntity(
                        Exception(error),
                        error ?: "无错误信息",
                        httpRequest.isRefreshRequest,
                        httpRequest.errorType,
                        httpRequest.intentData
                    )
                } else {
                    httpRequest.onError?.let { it(Exception(error)) }
                }
            }
        }

        // 请求失败时调用该方法
        override fun onFailure(call: Call<T>, t: Throwable) {
            // 结束加载状态
            loadingChange.loading.value = loadingStop
            // 判断错误是否需要自己处理
            if (httpRequest.onError == null){
                // 返回我们的错误内容
                loadingChange.showError.value = LoadStatusEntity(
                    t,
                    "请求错误，请检查网络连接",
                    httpRequest.isRefreshRequest,
                    httpRequest.errorType,
                    httpRequest.intentData
                )
            } else {
                httpRequest.onError?.let { it(t) }
            }
        }
    })
}

// http请求对象
class HttpRequest<T> {
    // call对象，用于请求数据
    var call: Call<T>? = null
    // 回调对象，当请求成功时，会通过这个来进行回调显示
    var onSuccess:  (T?) -> Unit = {}
    // 加载中的类型
    @LoadingType var loadingType = LoadingType.LOADING_NULL
    // 如果加载中显示弹窗，设置弹窗里的文字 loadingType == LOADING_DIALOG 的时候才有用 不是的话都不用传他
    var loadingMessage: String = "请求网络中..."
    // 请求失败的类型
    @LoadingType var errorType = LoadingType.LOADING_NULL
    // 错误回调，如果想自己处理错误那么就可以调用这个
    var onError: ((Throwable) -> Unit)? = null
    //是否是刷新请求 做列表分页功能使用 一般请求用不到
    var isRefreshRequest: Boolean = false
    // 如果请求失败，设置重新发起请求的参数
    var intentData: Any? = null
}
