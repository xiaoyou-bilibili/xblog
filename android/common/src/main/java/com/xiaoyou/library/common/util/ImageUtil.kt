package com.xiaoyou.library.common.util

import android.content.ContentValues
import android.content.Context
import android.graphics.Bitmap
import android.graphics.BitmapFactory
import android.net.Uri
import android.os.Build
import android.os.Environment
import android.provider.MediaStore
import android.util.Base64
import android.widget.ImageView
import androidx.annotation.RequiresApi
import com.bumptech.glide.Glide
import com.xiaoyou.library.common.R
import com.xiaoyou.library.common.base.appContext
import java.io.ByteArrayOutputStream
import java.io.File
import java.io.FileOutputStream
import java.io.OutputStream

/**
 * @description 和图片有关的工具类
 * @author 小游
 * @data 2021/02/25
 */
object ImageUtil {
    /**
     *  手动给imageView设置图像
     * @param url String 图片地址
     * @param view ImageView ImageView
     * @param err Int 加载错误显示的内容
     */
    fun setImageByUrl(url: String, view : ImageView, err :Int = R.mipmap.avatar){
        Glide.with(appContext)
            .load(url)
            .error(err)
            .centerCrop()
            .into(view)
    }

    /**
     *  读取本地的图片文件并转bitmap
     * @param path String 路径
     * @return Bitmap? bitmap对象
     */
    // https://www.jianshu.com/p/dce7fdcac03b
    fun getBitMapFromPath(path: String):Bitmap?{
        // 先判断图片是否存在
        val file = File(path)
        if (!file.exists()){
            return null
        }
        // 读取文件
        val bytes = file.readBytes()
        // 转bitmap offset我们设置为0即可
        return BitmapFactory.decodeByteArray(bytes,0,bytes.size)
    }

    /**
     *  bitmap转base64编码
     * @param bitmap Bitmap bitmap对象
     * @return String base64字符串
     */
    fun bitmap2Base64(bitmap: Bitmap):String{
        // 初始化outputString
        val ous = ByteArrayOutputStream()
        // 输出bitmap，这里我们设置quality为100 就是不进行压缩
        bitmap.compress(Bitmap.CompressFormat.JPEG,100,ous)
        // 返回base64编码
        return Base64.encodeToString(ous.toByteArray(),Base64.DEFAULT)
    }

    /**
     *  保存图片到手机相册
     * @param bitmap Bitmap bitmap对象
     * @param folderName String 文件夹名字
     * @return Boolean
     */
    fun saveImage(bitmap: Bitmap, folderName: String): Boolean{
        if (android.os.Build.VERSION.SDK_INT >= 29) {
            val values = contentValues()
            values.put(MediaStore.Images.Media.RELATIVE_PATH, "Pictures/" + folderName)
            values.put(MediaStore.Images.Media.IS_PENDING, true)
            // RELATIVE_PATH and IS_PENDING are introduced in API 29.
            val uri: Uri? = appContext.contentResolver.insert(MediaStore.Images.Media.EXTERNAL_CONTENT_URI, values)
            if (uri != null) {
                saveImageToStream(bitmap, appContext.contentResolver.openOutputStream(uri))
                values.put(MediaStore.Images.Media.IS_PENDING, false)
                appContext.contentResolver.update(uri, values, null, null)
                return true
            }
            return false
        } else {
            val directory = File(Environment.getExternalStorageDirectory().toString() + File.separator + folderName)
            // getExternalStorageDirectory is deprecated in API 29
            if (!directory.exists()) {
                directory.mkdirs()
            }
            val fileName = System.currentTimeMillis().toString() + ".png"
            val file = File(directory, fileName)
            saveImageToStream(bitmap, FileOutputStream(file))
            val values = contentValues()
            values.put(MediaStore.Images.Media.DATA, file.absolutePath)
            // .DATA is deprecated in API 29
            appContext.contentResolver.insert(MediaStore.Images.Media.EXTERNAL_CONTENT_URI, values)
            return true
        }
    }

    /**
     *  获取contentValues对象
     * @return ContentValues
     */
    private fun contentValues() : ContentValues {
        val values = ContentValues()
        values.put(MediaStore.Images.Media.MIME_TYPE, "image/png")
        values.put(MediaStore.Images.Media.DATE_ADDED, System.currentTimeMillis() / 1000);
        values.put(MediaStore.Images.Media.DATE_TAKEN, System.currentTimeMillis());
        return values
    }

    /**
     *  保存bitmap为stream流
     * @param bitmap Bitmap
     * @param outputStream OutputStream?
     */
    private fun saveImageToStream(bitmap: Bitmap, outputStream: OutputStream?) {
        if (outputStream != null) {
            try {
                bitmap.compress(Bitmap.CompressFormat.PNG, 100, outputStream)
                outputStream.close()
            } catch (e: Exception) {
                e.printStackTrace()
            }
        }
    }
}