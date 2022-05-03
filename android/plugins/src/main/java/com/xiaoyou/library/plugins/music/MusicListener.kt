package com.xiaoyou.library.plugins.music

import com.xiaoyou.library.net.entity.response.MusicDetail

/**
 * @description 自定义音乐播放器回接口
 * @author 小游
 * @data 2021/03/15
 */
interface MusicListener {
    /**
     * 播放音乐
     * @param musicDetail MusicDetail
     */
    fun onPlay(musicDetail: MusicDetail)
}