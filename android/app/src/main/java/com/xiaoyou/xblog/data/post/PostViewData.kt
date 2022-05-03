package com.xiaoyou.xblog.data.post


/**
 * Post浏览数据模型
 * @property date String 日期
 * @property view Int 访问量
 * @property comment Int 评论数
 * @property good Int 点赞数
 * @constructor
 */
data class PostViewData(val date: String,
                               val view: String,
                               val comment: String,
                               val good: String)