package com.xiaoyou.library.net.entity.response

import com.google.gson.annotations.SerializedName

/**
 * @description 设置板块返回的数据类型
 * @author 小游
 * @data 2021/02/14
 */
/**
 * @property donateHeadMeta DonateHeadMeta
 * @property aliPay String 	支付宝收款码
 * @property weChat String 	微信收款码
 * @property background String 赞助界面背景图片
 * @constructor
 */
data class DonateSettings(
        @SerializedName("head_meta")
        val donateHeadMeta: DonateHeadMeta,
        @SerializedName("alipay")
        val aliPay:String,
        @SerializedName("wechat")
        val weChat: String,
        val background :String
)

/**
 *  友链设置
 * @property donateHeadMeta DonateHeadMeta 头部标签
 * @property name String 名字
 * @property dec String 描述
 * @property link String 链接
 * @property avatar String 头像
 * @property background String 背景
 * @constructor
 */
data class FriendSettings(
        @SerializedName("head_meta")
        val donateHeadMeta: DonateHeadMeta,
        val name:String,
        val dec: String,
        val link :String,
        val avatar :String,
        val background :String
)

/**
 * @property title String 网站标题
 * @property keyword String 网站关键词
 * @property description String 描述信息
 * @property url String 网页链接
 * @property image String 代表性的图片
 * @property icon String 网站图标
 * @constructor
 */
data class DonateHeadMeta(
        val title:String,
        val keyword:String,
        val description:String,
        val url:String,
        val image:String,
        val icon:String
)

/**
 *  音乐播放盒数据
 * @property artist String 作家
 * @property cover String 封面名字
 * @property lrc String 歌词地址
 * @property name String 名字
 * @property url String 地址
 * @constructor
 */
data class MusicDetail(
        var artist:String="",
        var cover:String="",
        var lrc:String="",
        var name:String="",
        var url:String=""
)

/**
 *  APP的全局设置
 * @property chat Boolean 聊天功能
 * @property friend Boolean 友链功能
 * @property animation Boolean 显示我的追番
 * @property donate Boolean 显示赞助
 * @property dou_ban Boolean 显示豆瓣
 * @property music Boolean 显示音乐
 * @property doc Boolean 显示文档
 * @property project Boolean 显示项目
 * @property navigation Boolean 显示导航
 * @property login Boolean 登录界面图片
 * @constructor
 */
data class AppSetting(
        var chat:Boolean,
        var friend:Boolean,
        var animation:Boolean,
        var donate:Boolean,
        var dou_ban:Boolean,
        var music:Boolean,
        var doc:Boolean,
        var project:Boolean,
        var navigation:Boolean,
        var login: String,
)