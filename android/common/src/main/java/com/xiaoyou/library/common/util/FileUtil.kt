package com.xiaoyou.library.common.util

import android.os.Environment
import com.xiaoyou.library.common.base.appContext
import java.io.File
import java.math.BigDecimal


/**
 * @description 和文件有关的工具类
 * @author 小游
 * @data 2021/03/06
 */
object FileUtil {

    /**
     *  获取缓存大小
     * @return String 缓存大小
     */
    fun getTotalCacheSize():String{
        return try {
            var cacheSize = getFolderSize(appContext.cacheDir)
            if (Environment.getExternalStorageState() == Environment.MEDIA_MOUNTED) {
                appContext.externalCacheDir?.let { cacheSize += getFolderSize(it) }
            }
            getFormatSize(cacheSize.toDouble())
        } catch (e: java.lang.Exception) {
            "读取失败!"
        }
    }


    /**
     * 清除缓存
     */
    fun clearAllCache() {
        deleteDir(appContext.cacheDir)
        if (Environment.getExternalStorageState().equals(Environment.MEDIA_MOUNTED)) {
            deleteDir(appContext.externalCacheDir)
        }
    }

    /**
     * 删除文件夹
     * @param dir 当前文件夹
     * @return 是否删除
     */
    private fun deleteDir(dir: File?): Boolean {
        if (dir != null && dir.isDirectory) {
            val children = dir.list()?:return false
            for (i in children.indices) {
                val success = deleteDir(File(dir, children[i]))
                if (!success) {
                    return false
                }
            }
        }
        if (dir != null) {
            return dir.delete()
        }
        return false
    }


    /**
     * 获取文件夹大小
     * @param file 文件
     * @return 返回文件夹大小
     * @throws Exception 抛出异常
     */
    @Throws(Exception::class)
    fun getFolderSize(file: File): Long {
        var size: Long = 0
        try {
            val fileList = file.listFiles() ?: return 0
            for (i in fileList.indices) {
                // 如果下面还有文件
                size += if (fileList[i].isDirectory) {
                    getFolderSize(fileList[i])
                } else {
                    fileList[i].length()
                }
            }
        } catch (e: Exception) {
            e.printStackTrace()
        }
        return size
    }


    /**
     * 格式化单位
     * @param size 文件大小
     * @return 返回大小
     */
    fun getFormatSize(size: Double): String {
        val kiloByte = size / 1024
        if (kiloByte < 1) {
            return "0K"
        }
        val megaByte = kiloByte / 1024
        if (megaByte < 1) {
            val result1 = BigDecimal(kiloByte.toString())
            return result1.setScale(2, BigDecimal.ROUND_HALF_UP)
                    .toPlainString().toString() + "K"
        }
        val gigaByte = megaByte / 1024
        if (gigaByte < 1) {
            val result2 = BigDecimal(megaByte.toString())
            return result2.setScale(2, BigDecimal.ROUND_HALF_UP)
                    .toPlainString().toString() + "M"
        }
        val teraBytes = gigaByte / 1024
        if (teraBytes < 1) {
            val result3 = BigDecimal(gigaByte.toString())
            return result3.setScale(2, BigDecimal.ROUND_HALF_UP)
                    .toPlainString().toString() + "GB"
        }
        val result4 = BigDecimal(teraBytes)
        return result4.setScale(2, BigDecimal.ROUND_HALF_UP).toPlainString()
                .toString() + "TB"
    }

}