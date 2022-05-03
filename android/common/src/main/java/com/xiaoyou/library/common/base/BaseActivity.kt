package com.xiaoyou.library.common.base

import android.view.View
import com.jaredrummler.cyanea.app.CyaneaAppCompatActivity

/**
 * @description 基类
 * @author 小游
 * @data 2021/02/20
 */
abstract class BaseActivity : CyaneaAppCompatActivity(){
    abstract val layoutId: Int

    var dataBindView :View? = null


}