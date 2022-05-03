package com.xiaoyou.library.plugins.music

import android.annotation.SuppressLint
import android.content.BroadcastReceiver
import android.content.Context
import android.content.Intent
import android.content.IntentFilter
import android.media.MediaPlayer
import android.net.Uri
import android.os.Bundle
import android.widget.SeekBar
import android.widget.TextView
import androidx.core.app.JobIntentService.enqueueWork
import androidx.localbroadcastmanager.content.LocalBroadcastManager
import androidx.recyclerview.widget.StaggeredGridLayoutManager
import com.lxj.xpopup.XPopup
import com.lxj.xpopup.enums.PopupAnimation
import com.xiaoyou.library.common.base.BaseDbActivity
import com.xiaoyou.library.common.base.appContext
import com.xiaoyou.library.common.util.XLog
import com.xiaoyou.library.net.core.DataReceive
import com.xiaoyou.library.net.entity.response.MusicDetail
import com.xiaoyou.library.plugins.R
import com.xiaoyou.library.plugins.databinding.ActivityMusicBinding
import com.xiaoyou.library.plugins.friend.FriendAdapter
import com.xiaoyou.library.plugins.viewmodel.PluginsVM
import com.xiaoyou.library.plugins.viewmodel.SettingVM

/**
 * @description 我的播放器的界面
 * @author 小游
 * @data 2021/03/09
 */
class MusicActivity :BaseDbActivity<SettingVM,ActivityMusicBinding>() {
    // 音乐播放列表
    private val musics:MutableList<MusicDetail> = ArrayList()
    val adapter = MusicAdapter(musics)
    // 当前音乐播放状态
    private val playing = false

    override fun initView(savedInstanceState: Bundle?) {
        // 注册广播接受者接收数据
        // 初始化多个intentFilter
        val filter = IntentFilter()
        filter.addAction(PlayAction.PLAY.action)
        filter.addAction(PlayAction.PREVIOUS.action)
        filter.addAction(PlayAction.NEXT.action)
        filter.addAction(PlayAction.FINISH.action)
        // 这里我们监听音乐播放器的播放，上一首，下一首,播放完成事件
        LocalBroadcastManager.getInstance(this).registerReceiver(object : BroadcastReceiver() {
            override fun onReceive(context: Context?, intent: Intent?) {
                // 设置封面
                setPlayer(PlayerManager.getCurrent())
                // 设置为播放按钮
                setStatus(true)
            }
        }, filter)
        // 监听播放器进度
        LocalBroadcastManager.getInstance(this).registerReceiver(object : BroadcastReceiver() {
            override fun onReceive(context: Context?, intent: Intent?) {
                // 设置当前进度
                mDataBind.currentprocess.progress = PlayerManager.getProcess()
            }
        }, IntentFilter(PlayAction.MUSIC.action))
        // 监听播放器拖动事件
        mDataBind.currentprocess.setOnSeekBarChangeListener(object: SeekBar.OnSeekBarChangeListener{
            override fun onProgressChanged(seekBar: SeekBar?, progress: Int, fromUser: Boolean) {
            }
            override fun onStartTrackingTouch(seekBar: SeekBar?) {
            }
            override fun onStopTrackingTouch(seekBar: SeekBar?) {
                // 停止移动时我们设置播放器进度
                PlayerManager.setProcess(mDataBind.currentprocess.progress)
            }
        })
        // 初始化音乐播放器adapter
        mDataBind.musicList.layoutManager = StaggeredGridLayoutManager(
            1,
            StaggeredGridLayoutManager.VERTICAL
        )
        mDataBind.musicList.adapter = adapter
        // 获取音乐播放数据
        mViewModel.getMusic(object :DataReceive<List<MusicDetail>?>(){
            override fun success(data: List<MusicDetail>?) {
                if (data is List<MusicDetail>){
                    // 设置数据
                    adapter.setList(data)
                    // 这里我们给音乐播放器设置数据
                    PlayerManager.setMusicList(data)
                    // 先判断一下当前是否有音乐
                    if (PlayerManager.getCurrent().url!=""){
                        setPlayer(PlayerManager.getCurrent())
                        setStatus(PlayerManager.isPlaying())
                    } else {
                        setPlayer(data[0])
                    }
                }
            }
        })
        // 获取播放状态
        setStatus(PlayerManager.isPlaying())
        // 播放点击播放按钮
        mDataBind.play.setOnClickListener{
            if (PlayerManager.isPlaying()){
                setCommand(PlayAction.PAUSE)
                setStatus(false)
            } else {
                setCommand(PlayAction.PLAY)
                setStatus(true)
            }
        }
        // 下一首
        mDataBind.next.setOnClickListener{
            setCommand(PlayAction.NEXT)
            setStatus(true)
        }
        // 上一首
        mDataBind.previous.setOnClickListener{
            setCommand(PlayAction.PREVIOUS)
            setStatus(true)
        }
        // 设置播放列表点击事件
        adapter.setOnItemClickListener{ _, _, position ->
            val data = adapter.getItem(position)
            PlayerManager.setCurrent(data)
            setPlayer(data)
            // 设置为播放按钮
            setStatus(true)
        }
    }


    /**
     * 设置播放命令
     * @param action PlayAction
     */
    private fun setCommand(action: PlayAction){
        // 启动服务
        Intent(appContext,MusicService::class.java).also {
            it.action = action.action
            startService(it)
        }
    }

    /**
     *  设置音乐播放器状态
     * @param play Boolean 是否在播放
     */
    @SuppressLint("UseCompatLoadingForDrawables")
    private fun setStatus(play: Boolean){
        if (play){
            mDataBind.play.background = getDrawable(R.drawable.ic_pause)
        } else {
            mDataBind.play.background = getDrawable(R.drawable.ic_play)
        }
    }

    /**
     *  设置播放器内容
     * @param now MusicDetail 播放器详细数据
     */
    @SuppressLint("ResourceAsColor")
    private fun setPlayer(now: MusicDetail){
        // 设置列表为选中状态
        val current = musics.indexOf(now)
        for (i:Int in musics.indices){
            // 获取当前正在播放的音乐
            val musicText:TextView? = adapter.getViewByPosition(i,R.layout.item_music)?.findViewById(R.id.musicName)
            if (musicText!=null){
                if (i==current){
                    musicText.text = "测试"
                    musicText.setTextColor(R.color.theme_pink)
                } else {
                    musicText.text = "不测试"
                    musicText.setTextColor(R.color.text)
                }
            }
        }
        // 设置音乐名字和作家
        mDataBind.musicName.text = now.name
        mDataBind.musicArtist.text = now.artist
        // 设置封面
        var cover = now.cover
        if (cover.indexOf("http")==-1){
            cover = "https:${cover}"
        }
        mDataBind.musicCover.setImageURI(cover)
    }
}