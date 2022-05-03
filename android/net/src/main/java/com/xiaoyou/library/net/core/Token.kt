package com.xiaoyou.library.net.core

import com.tencent.mmkv.MMKV
import com.xiaoyou.library.net.entity.response.TokenDetail

/**
 * @description 用户认证相关的信息
 * @author 小游
 * @data 2021/02/25
 */
object Token {
    private val kv: MMKV by lazy {  MMKV.defaultMMKV() }
    // 用户id
    const val USER_ID = "user_id"
    // 用户token信息
    const val TOKEN = "token"
    // 获取token信息
    fun getToken() = TokenDetail(kv.getInt(USER_ID,0),kv.getString(TOKEN,"")?:"")
    // 删除token信息
    fun clearToken() {
        kv.removeValueForKey(USER_ID)
        kv.removeValueForKey(TOKEN)
    }
}