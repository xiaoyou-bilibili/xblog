package com.xiaoyou.xblog.ui.activity

import android.annotation.SuppressLint
import android.os.Bundle
import androidx.lifecycle.ViewModelProvider
import com.stfalcon.chatkit.messages.MessageHolders
import com.stfalcon.chatkit.messages.MessagesListAdapter
import com.stfalcon.chatkit.utils.DateFormatter
import com.xiaoyou.library.common.base.BaseDbActivity
import com.xiaoyou.library.common.ext.isVisible
import com.xiaoyou.library.common.ext.visibleOrGone
import com.xiaoyou.library.common.util.ImageUtil
import com.xiaoyou.library.common.util.MyToast
import com.xiaoyou.library.net.core.Token
import com.xiaoyou.library.net.core.Ws
import com.xiaoyou.library.net.core.WsListener
import com.xiaoyou.library.net.entity.response.ChatInfo
import com.xiaoyou.library.net.entity.response.UserDetail
import com.xiaoyou.xblog.R
import com.xiaoyou.xblog.data.chat.ChatMessage
import com.xiaoyou.xblog.data.chat.ChatUser
import com.xiaoyou.xblog.databinding.ActivityChatBinding
import com.xiaoyou.xblog.ui.adapter.chat.MyInHolder
import com.xiaoyou.xblog.ui.adapter.chat.MyOutHolder
import com.xiaoyou.xblog.viewmodel.ChatVM
import com.xiaoyou.xblog.viewmodel.UserVM
import kotlinx.android.synthetic.main.activity_chat.*
import okhttp3.WebSocket
import java.util.*


/**
 * @description 聊天界面
 * @author 小游
 * @data 2021/02/21
 */
// 参考：https://github.com/stfalcon-studio/ChatKit/blob/master/docs/COMPONENT_MESSAGES_LIST.md
class ChatActivity(override val layoutId: Int = R.layout.activity_chat) : BaseDbActivity<ChatVM, ActivityChatBinding>() {

    private var target = 0
    // 用户是否登录
    private var isLogin = false
    // 用户信息
    private lateinit var info : UserDetail
    // 信息
    private val messages:MutableList<ChatMessage> = ArrayList()
    // 用户信息的viewModel
    private val userViewModel by lazy { ViewModelProvider(this).get(UserVM::class.java) }

    override fun initView(savedInstanceState: Bundle?) {
        // 初始化UI
        initUI()
    }

    /**
     *  UI界面初始化
     */
    @SuppressLint("SetTextI18n")
    private fun initUI() {
        // 获取用户信息
        if (Token.getToken().userId > 0){
            userViewModel.getUserInfo()
            // 监听用户信息变化
            userViewModel.info.observe(this){
                isLogin = true
                info = it
                ImageUtil.setImageByUrl(it.avatar,mDataBind.chatAvatar)
            }
        }
        // 设置自己的布局
        val holdersConfig = MessageHolders()
        holdersConfig.setIncomingTextLayout(R.layout.item_chat)
        holdersConfig.setIncomingTextHolder(MyInHolder::class.java)
        holdersConfig.setOutcomingTextHolder(MyOutHolder::class.java)
        // 初始化adapter并设置图片加载
        val adapter = MessagesListAdapter<ChatMessage>(Token.getToken().userId.toString(),holdersConfig) { imageView, url, _ ->
            ImageUtil.setImageByUrl(url ?: "", imageView)
        }
        // 格式化时间显示
        adapter.setDateHeadersFormatter { DateFormatter.format(it, "M-d HH:mm") }
        messagesList.setAdapter(adapter)
        // 设置下拉刷新事件
        mDataBind.refreshLayout.apply {
            finishRefresh()
            setEnableAutoLoadMore(false)
            setEnableLoadMore(false)
            setOnRefreshListener {
                if (messages.isNullOrEmpty()){
                    MyToast.waring("没有更多数据了")
                    finishRefresh()
                    return@setOnRefreshListener
                }
                // 上拉刷新事件
                Ws.getContent(target,messages.last().date)
            }
        }
        // 当获取到数据时
        mViewModel.receiveMessage.observe(this){
            if (it is List<ChatInfo>){
                messages.clear()
                for (item in it){
                    messages.add(
                        ChatMessage(
                            item.id,
                            ChatUser(
                                item.user_id,
                                item.nickname,
                                item.avatar
                            ),
                            item.content,
                            Date(item.date),
                            item.date
                        )
                    )
                }
                adapter.clear()
                adapter.addToEnd(messages, true)
            }
        }
        // 当获取到历史记录时
        mViewModel.historyMessage.observe(this){
            mDataBind.refreshLayout.finishRefresh()
            if (it is List<ChatInfo>){
                messages.clear()
                for (item in it){
                    messages.add(
                        ChatMessage(
                            item.id,
                            ChatUser(
                                item.user_id,
                                item.nickname,
                                item.avatar
                            ),
                            item.content,
                            Date(item.date),
                            item.date
                        )
                    )
                }
                adapter.addToEnd(messages, true)
            }
        }
        // 当收到新消息时
        mViewModel.newMessage.observe(this){
            if (it is ChatInfo){
                val message = ChatMessage(
                    it.id,
                    ChatUser(
                        it.user_id,
                        it.nickname,
                        it.avatar
                    ),
                    it.content,
                    Date(it.date),
                    it.date
                )
                adapter.addToStart(message, true)
                adapter.notifyDataSetChanged()
            }
        }

        // 表情点击事件
        mDataBind.chatSmile.setOnClickListener{
            if (mDataBind.faceContent.isVisible()) {
                // 隐藏表情包
                mDataBind.faceContent.visibleOrGone(false)
                mDataBind.chatSmile.setTextColor(getColor(R.color.text))
            } else {
                // 显示表情包
                mDataBind.faceContent.visibleOrGone(true)
                mDataBind.chatSmile.setTextColor(getColor(R.color.cyanea_primary))
            }
        }
        // 表情部分回调
        mDataBind.faceContent.setListener {
            // 获取表情数据
            val face = it.tag.toString()
            // 设置编辑框内容，因为表情数据包含单引号，所以需要去掉
            mDataBind.chatEdit.setText(mDataBind.chatEdit.text.toString() + face.substring(1, face.length - 1))
        }
        // 点击发送按钮
        mDataBind.chatSend.setOnClickListener{
            if (isLogin){
                if (!mDataBind.chatEdit.text.isNullOrEmpty()){
                    // 提交留言
                    Ws.commitMessage(target,mDataBind.chatEdit.text.toString(),0)
                    mDataBind.chatEdit.setText("")
                } else{
                    MyToast.waring("请输入内容！")
                }
            } else {
                MyToast.waring("请先登录！")
            }
        }
    }

    // 界面恢复时重新建立连接
    override fun onResume() {
        super.onResume()
        // 获取发送目标
        target = intent.getIntExtra("target",0)
        // 这里我们直接调用接口来获取数据
        mViewModel.startChatSocket(object : WsListener() {
            override fun onOpen(webSocket: WebSocket) {
                Ws.getContent(target)
                // 定时发送心跳包,避免链接断开
                Timer().schedule(object : TimerTask(){
                    override fun run() {
                        Ws.mWebSocket?.send("ping")
                    }
                }, Date(),5000)
            }
        })
    }

    // 界面暂停时关闭
    override fun onPause() {
        super.onPause()
        Ws.mWebSocket?.close(1000,"")
    }

    override fun showToolBar() = true
}