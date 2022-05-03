package com.xiaoyou.library.common.util

import java.time.Instant
import java.time.LocalDateTime
import java.time.ZoneId
import java.time.format.DateTimeFormatter

/**
 * @description 和时间有关的函数
 * @author 小游
 * @data 2021/03/01
 */
object TimeUtil {
    fun formatTimeFromUnix(format: String,unix :Long = Instant.now().toEpochMilli()):String{
        val formatter =  DateTimeFormatter.ofPattern(format)
        // 设置时间
        return formatter.format(LocalDateTime.ofInstant(Instant.ofEpochMilli(unix), ZoneId.systemDefault()))
    }
}
