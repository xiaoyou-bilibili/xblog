package com.xiaoyou.library.common.util

import com.tencent.mmkv.MMKV
import com.xiaoyou.library.net.entity.response.AppSetting

object Setting {
    private val kv: MMKV by lazy {  MMKV.defaultMMKV() }

    private const val CHAT = "setting_chat"
    private const val FRIEND = "setting_friend"
    private const val ANIMATION = "setting_animation"
    private const val DONATE = "setting_donate"
    private const val DOU_BAN = "setting_dou_ban"
    private const val MUSIC = "setting_music"
    private const val DOC = "setting_doc"
    private const val PROJECT = "setting_project"
    private const val NAVIGATION = "setting_navigation"
    private const val LOGIN = "setting_login"

    // 设置设置信息
    fun setSetting(data :AppSetting){
        kv.encode(CHAT,data.chat)
        kv.encode(FRIEND,data.friend)
        kv.encode(ANIMATION,data.animation)
        kv.encode(DONATE,data.donate)
        kv.encode(DOU_BAN,data.dou_ban)
        kv.encode(MUSIC,data.music)
        kv.encode(DOC,data.doc)
        kv.encode(PROJECT,data.project)
        kv.encode(NAVIGATION,data.navigation)
        kv.encode(LOGIN,data.login)
    }

    // 获取设置信息
    fun getSetting() = AppSetting(
            kv.getBoolean(CHAT,true),
            kv.getBoolean(FRIEND,true),
            kv.getBoolean(ANIMATION,true),
            kv.getBoolean(DONATE,true),
            kv.getBoolean(DOU_BAN,true),
            kv.getBoolean(MUSIC,true),
            kv.getBoolean(DOC,true),
            kv.getBoolean(PROJECT,true),
            kv.getBoolean(NAVIGATION,true),
            kv.getString(LOGIN,"")?:"",
    )

}