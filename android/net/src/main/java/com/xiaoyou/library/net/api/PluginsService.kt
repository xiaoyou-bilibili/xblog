package com.xiaoyou.library.net.api

import com.xiaoyou.library.net.entity.param.SubmitFriend
import com.xiaoyou.library.net.entity.response.*
import retrofit2.Call
import retrofit2.http.*

/**
 * @description 插件板块的API调用
 * @author 小游
 * @data 2021/02/13
 */
interface PluginsService {
    // 使用suspend 协程可以直接返回结果，不返回call对象
    // 获取日记列表
    @GET("plugins/diary")
    fun getDiaryList(@Query("page")page:Int, @Query("filter") filter:Int): Call<ReturnList<DiaryDetail>>

    // 获取友链数据
    @GET("plugins/friends")
    fun getFriends():  Call<List<FriendDetail>>

    // 获取追番数据
    @GET("plugins/animations")
    fun getAnimations(@Query("page")page:Int): Call<Animations>

    // 获取赞助博主数据
    @GET("plugins/sponsors")
    fun getSponsors(): Call<List<Sponsors>>

    // 获取我的豆瓣记录
    @GET("plugins/dou_ban/{type}")
    fun getDouBan(@Path("type") type:String,@Query("page")page:Int): Call<ReturnList<DouBanDetail>>

    // 获取聊天室消息
    @GET("plugins/chatRoom")
    fun getChatRoom(): Call<List<ChatRoom>>

    // 提交友链申请
    @POST("plugins/friends")
    fun submitFriend(@Body friend: SubmitFriend): Call<SubmitFriend>

    // 获取所有文档
    @GET("plugins/docs")
    fun getDocs(): Call<List<DocList>>

    // 获取文档内容
    @GET("plugins/docs/{id}")
    fun getDocContent(@Path("id") id:Int): Call<DocContent>

    // 获取我的项目
    @GET("plugins/projects")
    fun getProjects(): Call<Project>

    // 获取个人导航
    @GET("plugins/navigation/links")
    fun getNavigationLink(): Call<List<Link>>

}