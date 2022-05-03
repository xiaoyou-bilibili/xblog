package com.xiaoyou.library.plugins.music

import android.content.Context
import android.content.Intent
import android.media.AudioAttributes
import android.media.AudioManager
import android.media.MediaPlayer
import android.net.Uri
import androidx.localbroadcastmanager.content.LocalBroadcastManager
import com.xiaoyou.library.common.base.appContext
import com.xiaoyou.library.common.util.MyToast
import com.xiaoyou.library.common.util.XLog
import com.xiaoyou.library.net.entity.response.MusicDetail
import java.util.*
import java.util.jar.Attributes
import kotlin.collections.ArrayList

/**
 * @description 自定义音乐播放器管理类,这里使用object来确保只初始化一次
 * @author 小游
 * @data 2021/03/10
 */
object PlayerManager :MediaPlayer.OnPreparedListener,MediaPlayer.OnCompletionListener{
    // 初始化音乐播放器
    private var player:MediaPlayer? = MediaPlayer()
    // 音乐播放列表
    private var musics:List<MusicDetail> = ArrayList()
    // 设置当前正在播放的音乐
    private var now:MusicDetail = MusicDetail()
    // 设置音乐播放列表
    fun setMusicList(data:List<MusicDetail>){
        musics = data
    }
    // 音乐播放器准备阶段，一般都是播放一首新音乐
    private fun prepare(music: MusicDetail){
        // 设置当前正在播放的音乐
        now = music
        // 首先要确保音乐播放器存在 
        if (player == null){
            new()
        }
        // 如果正在播放那么我们就关闭
        if (player!!.isPlaying){
            player!!.stop()
        }
        // 这里我们必须要reset一下，要不然会报错
        player!!.reset()
        // 设置播放源，然后我们准备一下
        player!!.setDataSource(music.url)
        player!!.setOnPreparedListener(this)
        player!!.setOnCompletionListener(this)
        player!!.prepare()
    }
    // 开始播放
    fun start() {
        // 判断当前是否有音乐在播放
        if (now.url.isEmpty()){
            // 那么我们就直接播放歌单的第一首歌
            if (musics.isNotEmpty()){
                // 我们设置音乐,然后播放
                prepare(musics[0])
            } else {
                MyToast.waring("没有歌单！")
            }
        } else {
            player?.start()
        }
    }
    // 暂停
    fun pause() = player?.pause()
    // 是否在播放
    fun isPlaying() = player?.isPlaying?:false
    // 下一首
    fun next() {
        // 获取当前音乐的位置
        val index = musics.indexOf(now)
        if (index >= musics.size - 1){
            // 说明当前是最后一首
            prepare(musics.first())
        } else {
            // 播放下一首
            prepare(musics[index+1])
        }
    }
    // 上一首
    fun previous() {
        // 获取当前音乐的位置
        val index = musics.indexOf(now)
        if (index <= 0){
            // 说明当前是第一首
            prepare(musics.last())
        } else {
            // 播放下一首
            prepare(musics[index-1])
        }
    }
    // 获取当前播放的音乐
    fun getCurrent() = now
    // 获取当前进度
    fun getProcess():Int{
        // 获取当前进度和总进度计算求出百分比
        val current= (player?.currentPosition?:0).toFloat()
        val total = (player?.duration?:0).toFloat()
        return (current/total*100).toInt()
    }
    // 设置音乐播放器进度
    fun setProcess(process:Int){
        val total = (player?.duration?:0).toFloat()
        player?.seekTo((process.toFloat()/100*total).toInt())
    }
    // 设置当播放的音乐
    fun setCurrent(music: MusicDetail){
        prepare(music)
    }

    // 新建音乐播放器
    fun new(){
        player = MediaPlayer()
    }
    // 释放音乐播放器
    fun release() = player?.release()
    // 当播放器准备好后我们就开始播放
    override fun onPrepared(mp: MediaPlayer?) {
        player?.start()
        // 这里我们启动一个定时任务，定时发送广播
        Timer().schedule(object : TimerTask(){
            override fun run() {
                LocalBroadcastManager.getInstance(appContext).sendBroadcast(Intent(PlayAction.MUSIC.action))
            }
        }, Date(),1000)
    }

    // 当播放完成播放时，我播放下一首
    override fun onCompletion(mp: MediaPlayer?) {
        next()
        // 发送完成的广播
        LocalBroadcastManager.getInstance(appContext).sendBroadcast(Intent(PlayAction.FINISH.action))
    }
}