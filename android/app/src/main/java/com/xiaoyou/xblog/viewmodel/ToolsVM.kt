package com.xiaoyou.xblog.viewmodel

import androidx.lifecycle.MutableLiveData
import com.xiaoyou.library.common.base.BaseViewModel
import com.xiaoyou.library.common.base.appContext
import com.xiaoyou.library.common.ext.handleCall
import com.xiaoyou.library.net.api.PluginsService
import com.xiaoyou.library.net.api.ToolsService
import com.xiaoyou.library.net.core.DataReceive
import com.xiaoyou.library.net.core.Repository
import com.xiaoyou.library.net.entity.base.LoadingType
import com.xiaoyou.library.net.entity.param.SubmitAdvice
import com.xiaoyou.library.net.entity.response.DiaryDetail
import com.xiaoyou.library.net.entity.response.ReturnList

/**
 * @description
 * @author 小游
 * @data 2021/02/25
 */
class ToolsVM (private val service: ToolsService = Repository.toolsService(appContext)) : BaseViewModel(){
    // 文章列表
//    val diary = MutableLiveData<ReturnList<DiaryDetail>>()

    /**
     *  用户提交意见反馈
     * @param content String 反馈的内容
     * @param handle DataReceive<SubmitAdvice?> 回调函数
     */
    fun submitAdvice(content: String,handle: DataReceive<SubmitAdvice?>) = handleCall<SubmitAdvice> {
        call = service.postAdvice(SubmitAdvice("",content))
        onSuccess = { handle.success(it) }
    }
}