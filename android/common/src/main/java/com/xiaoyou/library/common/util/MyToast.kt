package com.xiaoyou.library.common.util

import android.content.Context
import android.widget.Toast
import com.xiaoyou.library.common.base.appContext
import es.dmoral.toasty.Toasty

/**
 * @description 自定义提示框
 * @author 小游
 * @data 2021/02/21
 */
object MyToast {
    fun success(message: String){
        Toasty.success(appContext, message, Toast.LENGTH_SHORT, true).show();
    }

    fun error(message: String){
        Toasty.error(appContext, message, Toast.LENGTH_SHORT, true).show();
    }

    fun info(message: String){
        Toasty.info(appContext, message, Toast.LENGTH_SHORT, true).show();
    }

    fun waring(message: String){
        Toasty.warning(appContext, message, Toast.LENGTH_SHORT, true).show();
    }

}