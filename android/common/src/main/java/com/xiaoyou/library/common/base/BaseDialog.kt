package com.xiaoyou.library.common.base

import android.app.Dialog
import android.content.Context
import android.util.DisplayMetrics
import android.view.*
import androidx.databinding.DataBindingUtil
import androidx.databinding.ViewDataBinding
import com.xiaoyou.library.common.R
import java.lang.reflect.ParameterizedType

/**
 * @description 自定义的基础dialog
 * @author 小游
 * @data 2021/03/06
 */
abstract class BaseDialog<DB: ViewDataBinding>(context: Context,width: Float = 0.99f): Dialog(context){
    // dataBind对象
    lateinit var mDataBind: DB
    // dataBindView视图
    private val dataBindView : View  by lazy { mDataBind.root }

    init {
        //利用反射 根据泛型得到 ViewDataBinding
        val superClass = javaClass.genericSuperclass
        val aClass = (superClass as ParameterizedType).actualTypeArguments[0] as Class<*>
        val method = aClass.getDeclaredMethod("inflate",LayoutInflater::class.java)
        mDataBind =  method.invoke(null,layoutInflater) as DB
        // 视图绑定
        setContentView(dataBindView)
        // 设置dataBind
        mDataBind = DataBindingUtil.bind(dataBindView)!!
        // 设置动画样式
        val dialogWindow: Window = window!!
        dialogWindow.setWindowAnimations(R.style.AnimationAnimal)
        // 这里我们设置一下弹框的宽度
        val lp: WindowManager.LayoutParams = dialogWindow.attributes
        val d: DisplayMetrics = context.resources.displayMetrics
        lp.width = (d.widthPixels * width).toInt()
        dialogWindow.attributes = lp
    }

}