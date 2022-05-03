package com.xiaoyou.library.plugins.viewmodel

import com.xiaoyou.library.common.base.BaseViewModel
import com.xiaoyou.library.common.base.appContext
import com.xiaoyou.library.common.ext.handleCall
import com.xiaoyou.library.net.api.SettingsService
import com.xiaoyou.library.net.core.DataReceive
import com.xiaoyou.library.net.core.Repository
import com.xiaoyou.library.net.entity.response.AppSetting
import com.xiaoyou.library.net.entity.response.DonateSettings
import com.xiaoyou.library.net.entity.response.FriendSettings
import com.xiaoyou.library.net.entity.response.MusicDetail

/**
 * @description
 * @author 小游
 * @data 2021/02/26
 */
class SettingVM(private val service: SettingsService = Repository.settingService(appContext)) : BaseViewModel() {
    /**
     *  获取赞助界面设置
     * @param handle DataReceive<DonateSettings?>
     */
    fun getSponsor(handle: DataReceive<DonateSettings?>) = handleCall<DonateSettings> {
        call = service.getPluginsSponsor()
        onSuccess = {handle.success(it)}
    }

    /**
     *  获取友链界面设置
     * @param handle DataReceive<FriendSettings?> 回调函数
     */
    fun getFriend(handle: DataReceive<FriendSettings?>) = handleCall<FriendSettings>{
        call = service.getFriend()
        onSuccess = {handle.success(it)}
    }

    /**
     *  获取音乐盒设置
     * @param handle DataReceive<List<MusicDetail>?> 回调函数
     */
    fun getMusic(handle: DataReceive<List<MusicDetail>?>) = handleCall<List<MusicDetail>> {
        call = service.getMusic()
        onSuccess = {handle.success(it)}
    }

    /**
     *  获取全局APP设置
     * @param handle DataReceive<AppSetting?> APP设置
     */
    fun getAppSetting(handle: DataReceive<AppSetting?>) = handleCall<AppSetting> {
        call = service.getAppSetting()
        onSuccess = {handle.success(it)}
    }


}