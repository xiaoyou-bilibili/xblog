package com.xiaoyou.library.common.base

import android.os.Bundle
import android.view.LayoutInflater
import android.view.View
import android.widget.TextView
import androidx.lifecycle.ViewModelProvider
import com.kingja.loadsir.core.LoadService
import com.kingja.loadsir.core.LoadSir
import com.xiaoyou.library.widget.toolbar.CustomToolBar
import com.xiaoyou.library.common.R
import com.xiaoyou.library.common.ext.*
import com.xiaoyou.library.common.util.MyToast
import com.xiaoyou.library.net.entity.base.LoadStatusEntity
import com.xiaoyou.library.net.entity.base.LoadingType
import com.xiaoyou.library.widget.state.EmptyCallback
import com.xiaoyou.library.widget.state.ErrorCallback
import com.xiaoyou.library.widget.state.LoadingCallback
import kotlinx.android.synthetic.main.activity_base.*


/**
 * @description 带viewModel的activity
 * @author 小游
 * @data 2021/02/20
 */
abstract class BaseVmActivity<VM : BaseViewModel> : BaseActivity(), BaseIView {

    // 界面状态管理，用于显示各种界面
    private lateinit var uiStatusManger: LoadService<*>

    //当前Activity绑定的 ViewModel
    lateinit var mViewModel: VM

    // 自定义toobar，内容在widget那个library里
    lateinit var mToolbar: CustomToolBar

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        // 首先设置界面为空
        setContentView(R.layout.activity_base)
        // 生成ViewModel
        mViewModel = createViewModel()
        // 初始化 status View
        initStatusView(savedInstanceState)
        //注册界面响应事件
        initLoadingUiChange()
        //初始化绑定observer
        initObserver()
        //初始化请求成功方法
        onRequestSuccess()
        //初始化绑定点击方法
        onBindViewClick()
    }

    /**
     *  初始化statusView，用于显示自定义状态栏等内容
     * @param savedInstanceState Bundle?
     */
    private fun initStatusView(savedInstanceState: Bundle?) {
        val title = intent.getStringExtra("title")
        mToolbar = baseToolBar
        if (title.isNullOrEmpty()){
            mToolbar.visibleOrGone(showToolBar())
        }else{
            mToolbar.visibleOrGone(true)
            mToolbar.initBack(title) { finish() }
        }
        // 显示界面
        baseContentView.addView(if (dataBindView == null) LayoutInflater.from(this).inflate(layoutId, null) else dataBindView)
        // 初始化界面状态管理
        uiStatusManger = LoadSir.getDefault().register(if (getLoadingView() == null) baseContentView else getLoadingView()!!){
            onLoadRetry()
        }
        // 初始化fragment
        baseContentView.post {
            initView(savedInstanceState)
        }
    }

    /**
     * 初始化view
     */
    abstract fun initView(savedInstanceState: Bundle?)

    /**
     * 创建观察者
     */
    open fun initObserver() {}

    /**
     * 创建viewModel
     */
    private fun createViewModel(): VM {
        return ViewModelProvider(this).get(getVmClazz(this))
    }

    /**
     * 是否隐藏 标题栏 默认显示
     */
    open fun showToolBar(): Boolean {
        return true
    }


    /**
     * 点击事件方法 配合setOnclick()拓展函数调用，做到黄油刀类似的点击事件
     */
    open fun onBindViewClick() {}

    /**
     * 注册 UI 事件，用于显示网络请求
     */
    private fun initLoadingUiChange() {
        // 这里我们监听我们的网络请求，对不同的状态显示不同的ui界面
        mViewModel.loadingChange.run {
            loading.observeInActivity(this@BaseVmActivity) {
                if (it.loadingType == LoadingType.LOADING_DIALOG) {
                    if (it.isShow) {
                        showLoading(it.loadingMessage)
                    } else {
                        dismissLoading()
                    }
                }
            }
            // 观察空数据，并显示
            showEmpty.observeInActivity(this@BaseVmActivity) {
                onRequestEmpty(it)
            }
            showError.observeInActivity(this@BaseVmActivity) {
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
            showSuccess.observeInActivity(this@BaseVmActivity) {
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
    override fun onRequestSuccess() {}

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
        // 设置错误界面
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

    /**
     * 子类可传入需要被包裹的View，做状态显示-空、错误、加载
     * 如果子类不覆盖该方法 那么会将整个当前Activity界面（除封装的头部Toolbar）都当做View包裹
     */
    open fun getLoadingView(): View? {
        return null
    }
}