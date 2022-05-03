package com.xiaoyou.library.net.api

import com.xiaoyou.library.net.entity.param.*
import com.xiaoyou.library.net.entity.response.PostDetail
import com.xiaoyou.library.net.entity.response.TokenDetail
import com.xiaoyou.library.net.entity.response.UserDetail
import retrofit2.Call
import retrofit2.http.*

/**
 * @description 用户板块接口
 * @author 小游
 * @data 2021/02/14
 */
interface UserService {
    // 用户登录接口
    @POST("user/token")
    fun postToken(@Body loginParam: UserLoginParam): Call<TokenDetail>

    // 获取用户个人信息
    @GET("user")
    fun getUser(): Call<UserDetail>

    // 获取用户个人收藏
    @GET("user/collections")
    fun getCollection(): Call<List<PostDetail>>

    // 用户注册函数
    @POST("user/app")
    fun postRegister(@Body userRegisterParam: UserRegisterParam): Call<UserRegisterParam>

    // 获取验证码
    @POST("user/code")
    fun postCode(@Body userGetCodeParam: UserGetCodeParam): Call<UserGetCodeParam>

    // 判断用户名或邮箱是否注册
    @GET("user/username")
    fun getUsername(
        @Query("user") user: String,
        @Query("email") email: String
    ): Call<UserGetUsername>

    // 重置密码
    @PUT("user/app/password")
    fun resetPassword(@Body resetPassword: ResetPassword): Call<ResetPassword>

    //  更新用户信息
    @PUT("user")
    fun updateInfo(@Body updateUserInfo: UpdateUserInfo): Call<UpdateUserInfo>

}

