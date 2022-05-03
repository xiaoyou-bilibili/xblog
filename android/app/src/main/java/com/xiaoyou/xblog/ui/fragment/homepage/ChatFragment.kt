package com.xiaoyou.xblog.ui.fragment.homepage

import android.content.Intent
import android.os.Bundle
import com.stfalcon.chatkit.dialogs.DialogsListAdapter
import com.xiaoyou.library.common.base.BaseVmFragment
import com.xiaoyou.library.common.base.appContext
import com.xiaoyou.library.common.util.ImageUtil
import com.xiaoyou.xblog.R
import com.xiaoyou.xblog.data.chat.ChatDialog
import com.xiaoyou.xblog.data.chat.ChatMessage
import com.xiaoyou.xblog.data.chat.ChatUser
import com.xiaoyou.xblog.ui.activity.ChatActivity
import com.xiaoyou.xblog.viewmodel.PluginsVM
import kotlinx.android.synthetic.main.fragment_chat.*
import kotlinx.android.synthetic.main.fragment_chat.refreshLayout
import java.util.*

/**
 * @description 聊天界面视图
 * @author 小游
 * @data 2021/02/21
 */
// 参考：https://github.com/stfalcon-studio/ChatKit/blob/master/docs/COMPONENT_DIALOGS_LIST.MD
class ChatFragment(override val layoutId: Int = R.layout.fragment_chat) : BaseVmFragment<PluginsVM>() {

    // 聊天室
    private val chatRooms:MutableList<ChatDialog> = ArrayList()

    /**
     * 初始化view操作
     */
    override fun initView(savedInstanceState: Bundle?) {
        // 添加数据
        initUI()
    }

    // UI初始化
    private fun initUI() {
        // 初始化adapter
        val adapter = DialogsListAdapter<ChatDialog> { imageView, url, _ ->
            if (url != null) {
                ImageUtil.setImageByUrl(url,imageView,R.mipmap.avatar)
            }
        }
        // 刷新事件
        refreshLayout.apply {
            finishLoadMoreWithNoMoreData()
            setOnRefreshListener{
                mViewModel.getChatRoom()
            }
        }
        // 监听chatRoom变化
        mViewModel.chatRoom.observe(this){
            refreshLayout.finishRefresh()
            chatRooms.clear()
            for (item in it){
                val message = ChatMessage(item.message.id, ChatUser(1,"",""),item.message.content, Date(item.message.date))
                val dialog = ChatDialog(item.id,item.avatar,item.name,message,item.count)
                chatRooms.add(dialog)
            }
            // 设置聊天室
            adapter.setItems(chatRooms)
        }

        // 设置点击事件
        adapter.setOnDialogClickListener {
            val intent = Intent(appContext,ChatActivity::class.java)
            // 添加参数
            intent.putExtra("title",it.name)
            intent.putExtra("target",it.id)
            // 启动activity
            startActivity(intent)
        }
        chatList.setAdapter(adapter)
    }

    override fun onResume() {
        super.onResume()
        // 返回界面时自动刷新
        mViewModel.getChatRoom()
    }

    // 失败重试
    override fun onLoadRetry() {
        mViewModel.getChatRoom()
    }
}