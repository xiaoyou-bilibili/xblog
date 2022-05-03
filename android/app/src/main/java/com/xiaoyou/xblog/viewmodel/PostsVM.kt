package com.xiaoyou.xblog.viewmodel

import androidx.lifecycle.MutableLiveData
import com.xiaoyou.library.common.base.BaseViewModel
import com.xiaoyou.library.common.base.appContext
import com.xiaoyou.library.common.ext.handleCall
import com.xiaoyou.library.net.api.PostsService
import com.xiaoyou.library.net.core.DataReceive
import com.xiaoyou.library.net.core.Repository
import com.xiaoyou.library.net.entity.base.LoadingType
import com.xiaoyou.library.net.entity.param.EncryptContent
import com.xiaoyou.library.net.entity.param.PostComment
import com.xiaoyou.library.net.entity.param.PostsStatus
import com.xiaoyou.library.net.entity.response.CommentDetail
import com.xiaoyou.library.net.entity.response.PostContentDetail
import com.xiaoyou.library.net.entity.response.PostDetail
import com.xiaoyou.library.net.entity.response.ReturnList

/**
 * @description 文章vm
 * @author 小游
 * @data 2021/02/20
 */
class PostsVM(private val service: PostsService = Repository.postsService(appContext)) : BaseViewModel() {

    // 文章列表
    val posts = MutableLiveData<ReturnList<PostDetail>>()
    // 文章内容
    val postContent = MutableLiveData<PostContentDetail>()
    // 文章评论
    val postComment = MutableLiveData<List<CommentDetail>>()

    /**
     *  获取文章列表
     * @param page Int 页数
     * @param q String 关键词
     */
    fun getPosts(page: Int, q: String = "") = handleCall<ReturnList<PostDetail>> {
        call = service.getPosts(page,q,"","")
        onSuccess = { posts.postValue(it) }
        errorType = if (page==1) LoadingType.LOADING_XML else LoadingType.LOADING_NULL
    }

    /**
     *  获取文章内容
     * @param id Int 文章id
     * @return LiveData<Resource<PostContentDetail>>
     */
    fun getPostContent(id: Int) = handleCall<PostContentDetail> {
        call = service.getPostsContent(id)
        onSuccess = { postContent.postValue(it) }
        loadingType = LoadingType.LOADING_DIALOG
        loadingMessage = "正在获取文章数据..."
        errorType = LoadingType.LOADING_XML
    }

    /**
     *  获取文章评论数据
     * @param id Int 文章评论id
     */
    fun getComment(id: Int) = handleCall<List<CommentDetail>> {
        call = service.getComments(id)
        onSuccess = { postComment.postValue(it) }
        loadingType = LoadingType.LOADING_DIALOG
        loadingMessage = "正在获取文章数据..."
        errorType = LoadingType.LOADING_XML
    }

    // 获取文章状态
    fun getPostStatus(id: Int, receive: DataReceive<PostsStatus?>) = handleCall<PostsStatus> {
        call = service.getStatus(id)
        onSuccess = { receive.success(it) }
        errorType = LoadingType.LOADING_ERROR_NULL
    }

    /**
     *  更新文章状态
     * @param id Int 文章id
     * @param status PostsStatus 文章状态
     * @param receive OnDataReceive<PostsStatus?> 回调函数
     */
    fun updatePostStatus(id: Int, status :PostsStatus, receive: DataReceive<PostsStatus?>) = handleCall<PostsStatus> {
        call = service.updateStatus(id,status)
        onSuccess = { receive.success(it) }
    }

    /**
     *  文章发布评论
     * @param id Int 文章id
     * @param comment PostComment 评论参数
     * @param receive OnDataReceive<PostComment?>
     */
    fun postComment(id: Int, comment :PostComment, receive: DataReceive<PostComment?>) = handleCall<PostComment> {
        call = service.postComments(id,comment)
        onSuccess = {receive.success(it)}
        loadingType = LoadingType.LOADING_DIALOG
        loadingMessage = "正在发表..."
        onError = {receive.error(it.message?:"")}
    }

    /**
     *  获取加密文章内容
     * @param id Int 文章id
     * @param password String 文章密码
     * @param receive DataReceive<EncryptContent?> 回调函数
     */
    fun getEncryptContent(id:Int,password:String, receive: DataReceive<EncryptContent?>) = handleCall<EncryptContent> {
        call = service.getEncryptContent(id,password)
        onSuccess = {receive.success(it)}
    }
}