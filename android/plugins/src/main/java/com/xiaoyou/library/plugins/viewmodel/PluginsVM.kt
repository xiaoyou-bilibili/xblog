package com.xiaoyou.library.plugins.viewmodel

import androidx.lifecycle.MutableLiveData
import com.xiaoyou.library.common.base.BaseViewModel
import com.xiaoyou.library.common.base.appContext
import com.xiaoyou.library.common.ext.handleCall
import com.xiaoyou.library.net.api.PluginsService
import com.xiaoyou.library.net.api.PostsService
import com.xiaoyou.library.net.core.DataReceive
import com.xiaoyou.library.net.core.Repository
import com.xiaoyou.library.net.entity.base.LoadingType
import com.xiaoyou.library.net.entity.param.SubmitFriend
import com.xiaoyou.library.net.entity.response.*

/**
 * @description 插件部分ViewModel
 * @author 小游
 * @data 2021/02/26
 */
class PluginsVM (private val service: PluginsService = Repository.pluginsService(appContext)) : BaseViewModel(){

    // 追番数据
    val animations = MutableLiveData<Animations>()
    // 豆瓣数据
    val douBans = MutableLiveData<ReturnList<DouBanDetail>>()

    /**
     * 获取友链数据
     * @param handle OnDataReceive<List<FriendDetail>?>
     */
    fun getFriend(handle: DataReceive<List<FriendDetail>?>) = handleCall<List<FriendDetail>> {
        call = service.getFriends()
        onSuccess = { handle.success(it) }
        loadingType = LoadingType.LOADING_DIALOG
        loadingMessage = "正在获取友链数据..."
        errorType = LoadingType.LOADING_XML
    }

    /**
     * 获取追番数据
     * @param page Int 当前页数
     */
    fun getAnimation(page :Int) = handleCall<Animations> {
        call = service.getAnimations(page)
        onSuccess = {animations.postValue(it)}
        loadingType = LoadingType.LOADING_DIALOG
        loadingMessage = "正在获取追番数据..."
        errorType = if (page==1) LoadingType.LOADING_XML else LoadingType.LOADING_NULL
    }

    /**
     *  获取赞助数据
     * @param handle OnDataReceive<List<Sponsors>?>
     */
    fun getSponsor(handle: DataReceive<List<Sponsors>?>) = handleCall<List<Sponsors>> {
        call = service.getSponsors()
        onSuccess = {handle.success(it)}
        loadingType = LoadingType.LOADING_DIALOG
        loadingMessage = "正在获取赞助记录..."
        errorType = LoadingType.LOADING_XML
    }

    /**
     *  获取豆瓣记录
     * @param page Int 第几页
     * @param type String 类型
     */
    fun getDouBan(page: Int,type:String) = handleCall<ReturnList<DouBanDetail>> {
        call = service.getDouBan(type,page)
        onSuccess = {douBans.postValue(it)}
        loadingType = LoadingType.LOADING_DIALOG
        loadingMessage = "正在获取豆瓣记录..."
        errorType = if (page==1) LoadingType.LOADING_XML else LoadingType.LOADING_NULL
    }

    /**
     *  提交友链数据
     * @param param SubmitFriend 参数
     * @param handle DataReceive<SubmitFriend?> 回调函数
     */
    fun submitFriend(param: SubmitFriend,handle: DataReceive<SubmitFriend?>) = handleCall<SubmitFriend> {
        call = service.submitFriend(param)
        onSuccess = {handle.success(it)}
    }

    /**
     * 获取文档列表
     * @param handle DataReceive<List<DocList>?>
     */
    fun getDocs(handle: DataReceive<List<DocList>?>) = handleCall<List<DocList>> {
        call = service.getDocs()
        onSuccess = {handle.success(it)}
    }

    /**
     *  获取文档内容
     * @param id Int 文档id
     * @param handle DataReceive<DocContent?> 回调函数
     */
    fun getDocContent(id :Int,handle: DataReceive<DocContent?>) = handleCall<DocContent> {
        call = service.getDocContent(id)
        onSuccess = {handle.success(it)}
    }

    /**
     * 获取我的项目
     * @param handle DataReceive<Project?> 回调函数
     */
    fun getProject(handle: DataReceive<Project?>) = handleCall<Project> {
        call = service.getProjects()
        onSuccess = {handle.success(it)}
        loadingType = LoadingType.LOADING_DIALOG
        loadingMessage = "获取个人项目..."
        errorType = LoadingType.LOADING_XML
    }

    /**
     * 获取个人导航
     * @param handle DataReceive<Project?> 回调函数
     */
    fun getNavigationLink(handle: DataReceive<List<Link>?>) = handleCall<List<Link>> {
        call = service.getNavigationLink()
        onSuccess = {handle.success(it)}
        loadingType = LoadingType.LOADING_DIALOG
        loadingMessage = "获取个人导航..."
        errorType = LoadingType.LOADING_XML
    }

}