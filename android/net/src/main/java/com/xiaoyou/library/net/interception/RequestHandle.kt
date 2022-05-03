package com.xiaoyou.library.net.interception

import android.util.Log
import okhttp3.Interceptor
import okhttp3.Response
import okio.Buffer
import java.nio.charset.Charset
import java.util.concurrent.TimeUnit

/**
 * @description 自定义网络错误处理
 * @author 小游
 * @data 2021/02/20
 */
class RequestHandle : Interceptor {
    private val TAG = RequestHandle::class.java.simpleName
    private val UTF8 = Charset.forName("UTF-8")

    override fun intercept(chain: Interceptor.Chain): Response {
        val request = chain.request()
        val requestBody = request.body
        var body: String? = null
        requestBody?.let {
            val buffer = Buffer()
            requestBody.writeTo(buffer)
            var charset: Charset? = UTF8
            val contentType = requestBody.contentType()
            contentType?.let {
                charset = contentType.charset(UTF8)
            }
            body = buffer.readString(charset!!)
        }
        Log.i(TAG,
            " send request: \n method = ${request.method}"
                    + " \n url = ${request.url}"
                    + " \n request header = ${request.headers}"
                    + " \n request params = $body"
        )
        val startNs = System.nanoTime()
        val response = chain.proceed(request)
        val tookMs = TimeUnit.NANOSECONDS.toMillis(System.nanoTime() - startNs)

        val responseBody = response.body
        val rBody: String

        val source = responseBody!!.source()
        source.request(java.lang.Long.MAX_VALUE)
        val buffer = source.buffer()

        var charset: Charset? = UTF8
        val contentType = responseBody.contentType()
        contentType?.let {
            charset = contentType.charset(UTF8)
        }
        rBody = buffer.clone().readString(charset!!)

        Log.i(TAG,
            "received : code = ${response.code}"
                    + "\n url = ${response.request.url}"
                    + "\n body = $body"
                    + "\n response $rBody ")

        return response
    }
}