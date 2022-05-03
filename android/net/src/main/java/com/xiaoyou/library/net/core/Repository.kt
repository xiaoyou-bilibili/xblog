package com.xiaoyou.library.net.core

import android.content.Context
import com.xiaoyou.library.net.api.*
import com.xiaoyou.library.net.core.ServiceCreator.createService


/**
 * @description 请求类
 * @author 小游
 * @data 2021/02/13
 */
object Repository {
    // 文章板块服务类
    fun postsService(context: Context) = createService(PostsService::class.java,context)
    // 插件板块服务类
    fun pluginsService(context: Context) = createService(PluginsService::class.java,context)
    // 设置板块服务类
    fun settingService(context: Context) = createService(SettingsService::class.java,context)
    // 用户板块服务类
    fun userService(context: Context) = createService(UserService::class.java,context)
    // 文章板块服务类
    fun toolsService(context: Context) = createService(ToolsService::class.java,context)
}