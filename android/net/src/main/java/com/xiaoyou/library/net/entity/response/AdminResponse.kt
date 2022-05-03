package com.xiaoyou.library.net.entity.response

/**
 * @description 管理员板块返回的信息
 * @author 小游
 * @data 2021/03/06
 */

/**
 *  图片上传
 * @property url String 图片地址
 * @property name String 图片名字
 * @constructor
 */
data class UploadFile(
        val url: String,
        val name: String
)

/**
 * APP下载链接
 * @property version String 版本信息
 * @property dec String 描述
 * @property download String 下载链接
 * @constructor
 */
data class AppVersion(
        val version:String,
        val dec: String,
        val download: String
)