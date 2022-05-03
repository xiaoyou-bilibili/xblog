package com.xiaoyou.library.net.api

import com.xiaoyou.library.net.entity.param.SubmitAdvice
import com.xiaoyou.library.net.entity.response.AppVersion
import com.xiaoyou.library.net.entity.response.UploadFile
import retrofit2.Call
import retrofit2.http.*

/**
 * @description 工具板块
 * @author 小游
 * @data 2021/02/19
 */
interface ToolsService {

    // 意见反馈
    @POST("tools/advice")
    fun postAdvice(
            @Body submitAdvice: SubmitAdvice
    ): Call<SubmitAdvice>

    // 图片上传结构
    @POST("tools/file/images/base64")
    @FormUrlEncoded
    fun uploadImage(@Field("data") data: String): Call<UploadFile>

    // 检查应用更新
    @GET("tools/app/version")
    fun getAppVersion(): Call<AppVersion>
}