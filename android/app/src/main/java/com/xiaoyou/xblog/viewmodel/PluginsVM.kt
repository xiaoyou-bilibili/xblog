package com.xiaoyou.xblog.viewmodel

import androidx.lifecycle.MutableLiveData
import com.xiaoyou.library.common.base.BaseViewModel
import com.xiaoyou.library.common.base.appContext
import com.xiaoyou.library.common.ext.handleCall
import com.xiaoyou.library.net.api.PluginsService
import com.xiaoyou.library.net.core.Repository
import com.xiaoyou.library.net.entity.base.LoadingType
import com.xiaoyou.library.net.entity.response.ChatRoom
import com.xiaoyou.library.net.entity.response.DiaryDetail
import com.xiaoyou.library.net.entity.response.ReturnList

/**
 * @description 插件部分ViewModel
 * @author 小游
 * @data 2021/02/24
 */
class PluginsVM (private val service: PluginsService = Repository.pluginsService(appContext)) : BaseViewModel(){
    // 日记列表
    val diary = MutableLiveData<ReturnList<DiaryDetail>>()
    // 聊天室
    val chatRoom = MutableLiveData<List<ChatRoom>>()
    // 获取日记列表
    fun getDiary(page: Int) = handleCall<ReturnList<DiaryDetail>> {
        call = service.getDiaryList(page,1)
        onSuccess = { diary.postValue(it) }
        errorType = if (page==1) LoadingType.LOADING_XML else LoadingType.LOADING_NULL
    }

    // 获取聊天室数据
    fun getChatRoom() = handleCall<List<ChatRoom>>{
        call = service.getChatRoom()
        onSuccess = { chatRoom.postValue(it) }
        loadingType = LoadingType.LOADING_DIALOG
        loadingMessage = "正在获取最新消息..."
        errorType = LoadingType.LOADING_XML
    }
}