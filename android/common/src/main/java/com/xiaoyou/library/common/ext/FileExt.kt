package com.xiaoyou.library.common.ext

import android.content.Context
import java.io.BufferedReader
import java.io.IOException
import java.io.InputStreamReader

/**
 * @description 和文件读取有关的函数
 * @author 小游
 * @data 2021/02/25
 */
// 读取文件
fun readAssets(context: Context, src: String?): String {
    try {
        context.assets.open(src!!).use { `is` ->
            val reader = InputStreamReader(`is`)
            val bufferedReader = BufferedReader(reader)
            val buffer = StringBuilder("")
            var str: String?
            while (bufferedReader.readLine().also { str = it } != null) {
                buffer.append(str)
                buffer.append("\n")
            }
            return buffer.toString()
        }
    } catch (e: IOException) {
        return ""
    }
}