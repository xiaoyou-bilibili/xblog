package com.xiaoyou.library.net.entity.response

/**
 * @description 工具板块返回的参数
 * @author 小游
 * @data 2021/02/19
 */

/**
 * 表情内容
 */
data class Face(var type: String? = null,
                var container: List<FaceDetail>? = null)

/**
 * 表情详情
 */
data class FaceDetail(
    var desc: String? = null,
    var icon: String? = null,
    var text: String? = null
)