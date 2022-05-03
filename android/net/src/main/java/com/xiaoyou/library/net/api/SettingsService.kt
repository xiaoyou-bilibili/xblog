package com.xiaoyou.library.net.api

import com.xiaoyou.library.net.entity.response.AppSetting
import com.xiaoyou.library.net.entity.response.DonateSettings
import com.xiaoyou.library.net.entity.response.FriendSettings
import com.xiaoyou.library.net.entity.response.MusicDetail
import retrofit2.Call
import retrofit2.http.GET

/**
 * @description 设置板块的接口
 * @author 小游
 * @data 2021/02/14
 */
interface SettingsService {
    // 获取赞助界面设置
    @GET("settings/plugins/sponsor")
    fun getPluginsSponsor(): Call<DonateSettings>
    // 获取友链申请设置
    @GET("settings/plugins/friends")
    fun getFriend(): Call<FriendSettings>
    // 获取音乐盒设置
    @GET("settings/side/music")
    fun getMusic(): Call<List<MusicDetail>>
    // 获取APP全局设置
    @GET("settings/app")
    fun getAppSetting(): Call<AppSetting>
}