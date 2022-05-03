package com.xiaoyou.library.common.util

import java.util.regex.Pattern

/**
 * @description 各种参数验证
 * @author 小游
 * @data 2021/03/03
 */
object Validator {
    private const val emailPattern = "^[a-zA-Z0-9][\\w\\.-]*[a-zA-Z0-9]@[a-zA-Z0-9][\\w\\.-]*[a-zA-Z0-9]\\.[a-zA-Z][a-zA-Z\\.]*[a-zA-Z]$"

    /**
     *  验证邮箱是否合法
     * @param email String 邮箱
     * @return Boolean 是否合法
     */
    fun isEmailOk(email:String) = if (email.isEmpty()){
        false
    } else {
        Pattern.compile(emailPattern).matcher(email).matches()
    }

}