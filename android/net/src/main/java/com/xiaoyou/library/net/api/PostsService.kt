package com.xiaoyou.library.net.api

import com.xiaoyou.library.net.entity.param.EncryptContent
import com.xiaoyou.library.net.entity.param.PostComment
import com.xiaoyou.library.net.entity.param.PostsStatus
import com.xiaoyou.library.net.entity.response.*
import retrofit2.Call
import retrofit2.http.*


/**
 * @description posts板块的api调用
 * @author 小游
 * @data 2021/02/13
 */
interface PostsService {
    // 获取文章列表
    @GET("posts")
    fun getPosts(
        @Query("page") page: Int,
        @Query("q") q: String,
        @Query("category") category: String,
        @Query("tag") tag: String
    ): Call<ReturnList<PostDetail>>

    // 获取文章内容
    @GET("posts/{id}")
    fun getPostsContent(
        @Path("id") Id: Int
    ): Call<PostContentDetail>

    // 获取文章评论
    @GET("posts/{id}/comments")
    fun getComments(@Path("id") Id: Int): Call<List<CommentDetail>>

    // 获取文章状态
    @GET("posts/{id}/status")
    fun getStatus(@Path("id") Id: Int): Call<PostsStatus>

    // 用户更新文章状态
    @PUT("posts/{id}/status")
    fun updateStatus(@Path("id") Id: Int, @Body status: PostsStatus): Call<PostsStatus>

    // 用户发表评论
    @POST("posts/{id}/comments")
    fun postComments(@Path("id") Id: Int, @Body comment: PostComment): Call<PostComment>

    // 获取加密文章内容
    @GET("posts/{id}/encryption")
    fun getEncryptContent(@Path("id") Id: Int, @Query("password") password: String): Call<EncryptContent>

}