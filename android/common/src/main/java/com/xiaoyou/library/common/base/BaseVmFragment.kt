package com.xiaoyou.library.common.base

import android.content.Context
import android.os.Bundle
import android.util.Log
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.TextView
import androidx.appcompat.app.AppCompatActivity
import androidx.lifecycle.Lifecycle
import androidx.lifecycle.ViewModelProvider
import com.kingja.loadsir.core.LoadService
import com.kingja.loadsir.core.LoadSir
import com.xiaoyou.library.common.R
import com.xiaoyou.library.common.ext.dismissLoading
import com.xiaoyou.library.common.ext.getVmClazz
import com.xiaoyou.library.common.ext.logD
import com.xiaoyou.library.common.ext.showLoading
import com.xiaoyou.library.common.util.MyToast
import com.xiaoyou.library.common.util.XLog
import com.xiaoyou.library.net.entity.base.LoadStatusEntity
import com.xiaoyou.library.net.entity.base.LoadingType
import com.xiaoyou.library.widget.state.EmptyCallback
import com.xiaoyou.library.widget.state.ErrorCallback
import com.xiaoyou.library.widget.state.LoadingCallback

/**
 * @description 带ViewModel的fragment
 * @author 小游
 * @data 2021/02/20
 */
abstract class BaseVmFragment<VM : BaseViewModel> : BaseFragment(), BaseIView {

    //界面状态管理者
    lateinit var uiStatusManger: LoadService<*>

    //是否第一次加载
    private var isFirst: Boolean = true

    //当前Fragment绑定的泛型类ViewModel
    lateinit var mViewModel: VM

    //父类activity
    lateinit var mActivity: AppCompatActivity
    // 父类的baseVMActivity(因为有些东西需要使用到)
    lateinit var mVmActivity : BaseVmActivity<*>

    // 当数据视图创建的时候
    override fun onCreateView(
        inflater: LayoutInflater,
        container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View? {
        isFirst = true
        javaClass.simpleName.logD()
        // 判断是否绑定过内容，如果绑定过就直接返回
        val rootView = if (dataBindView == null) {
            inflater.inflate(layoutId, container, false)
        } else {
            dataBindView
        }
        // 这里我们传递View来显示状态界面，如果为空就默认使用整个activity
        return if(getLoadingView()==null){
            // 这里我们设置布局管理的界面
            uiStatusManger = LoadSir.getDefault().register(rootView){
                onLoadRetry()
            }
            container?.removeView(uiStatusManger.loadLayout)
            uiStatusManger.loadLayout
        }else{
            rootView
        }
    }

    //  自动获取父activity供fragment使用
    override fun onAttach(context: Context) {
        super.onAttach(context)
        mActivity = context as AppCompatActivity
        mVmActivity = context as BaseVmActivity<*>
    }

    // 这进行布局初始化操作
    override fun onViewCreated(view: View, savedInstanceState: Bundle?) {
        super.onViewCreated(view, savedInstanceState)
        mViewModel = createViewModel()
        initStatusView(view, savedInstanceState)
        initLoadingUiChange()
        initObserver()
        onRequestSuccess()
        onBindViewClick()
    }

    // 初始化状态界面，这里我们传入view
    private fun initStatusView(view: View, savedInstanceState: Bundle?) {
        getLoadingView()?.let {
            uiStatusManger = LoadSir.getDefault().register(it){
                onLoadRetry()
            }
        }
        view.post {
            initView(savedInstanceState)
        }
    }

    /**
     * 创建viewModel
     */
    private fun createViewModel(): VM {
        return ViewModelProvider(this).get(getVmClazz(this))
    }

    /**
     * 初始化view操作
     */
    abstract fun initView(savedInstanceState: Bundle?)

    /**
     * 懒加载
     */
    open fun lazyLoadData() {}

    /**
     * 创建观察者
     */
    open fun initObserver() {}

    override fun onResume() {
        super.onResume()
        onVisible()
    }

    /**
     * 是否需要懒加载
     */
    private fun onVisible() {
        if (lifecycle.currentState == Lifecycle.State.STARTED && isFirst) {
            view?.post {
                lazyLoadData()
                isFirst = false
            }
        }
    }

    /**
     * 子类可传入需要被包裹的View，做状态显示-空、错误、加载
     * 如果子类不覆盖该方法 那么会将整个当前Fragment界面都当做View包裹
     */
    open fun getLoadingView(): View? {
        return null
    }

    /**
     * 点击事件方法 配合setOnclick()拓展函数调用，做到黄油刀类似的点击事件
     */
    open fun onBindViewClick() {}

    /**
     * 注册 UI 事件
     */
    private fun initLoadingUiChange() {
        mViewModel.loadingChange.run {
            loading.observeInFragment(this@BaseVmFragment) {
                if (it.loadingType == LoadingType.LOADING_DIALOG) {
                    if (it.isShow) {
                        showLoading(it.loadingMessage)
                    } else {
                        dismissLoading()
                    }
                }
            }
            showEmpty.observeInFragment(this@BaseVmFragment) {
                onRequestEmpty(it)
            }
            showError.observeInFragment(this@BaseVmFragment) {
                // 判断错误显示类型，如果需要显示错误布局，我们就进行显示，否则只显示错误弹框
                if (it.errorType == LoadingType.LOADING_XML) {
                    showErrorUi(it.errorMessage)
                } else if (it.errorType == LoadingType.LOADING_NULL) {
                    // 默认显示错误信息
                    MyToast.error(it.errorMessage)
                }
                // 接口回调
                onRequestError(it)
            }
            showSuccess.observeInFragment(this@BaseVmFragment) {
                showSuccessUi()
            }
        }
    }

    /**
     * 请求列表数据为空时 回调
     * @param loadStatus LoadStatusEntity
     */
    override fun onRequestEmpty(loadStatus: LoadStatusEntity) {
        showEmptyUi()
    }

    /**
     * 请求接口失败回调，如果界面有请求接口，需要处理错误业务，请实现它
     * @param loadStatus LoadStatusEntity
     */
    override fun onRequestError(loadStatus: LoadStatusEntity) {}

    /**
     * 请求成功的回调放在这里面 没干啥就是取了个名字，到时候好找
     */
    override fun onRequestSuccess() {

    }

    /**
     * 空界面，错误界面 点击重试时触发的方法，如果有使用 状态布局的话，一般子类都要实现
     */
    override fun onLoadRetry() {}

    /**
     * 显示 成功状态界面
     */
    override fun showSuccessUi() {
        uiStatusManger.showSuccess()
    }

    /**
     * 显示 错误 状态界面
     */
    override fun showErrorUi(errMessage: String) {
        uiStatusManger.showCallback(ErrorCallback::class.java)
        // 使用自带的Callback修改内容显示
        uiStatusManger.setCallBack(ErrorCallback::class.java) {
                _, view -> view.findViewById<TextView>(R.id.state_error_text).text = errMessage
        }
    }


    /**
     * 显示 错误 状态界面
     */
    override fun showEmptyUi() {
        uiStatusManger.showCallback(EmptyCallback::class.java)
    }

    /**
     * 显示 loading 状态界面
     */
    override fun showLoadingUi() {
        uiStatusManger.showCallback(LoadingCallback::class.java)
    }

}