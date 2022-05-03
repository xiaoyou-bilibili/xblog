package com.xiaoyou.library.plugins.music

/**
 * @description 播放器的几个命令
 * @author 小游
 * @data 2021/03/10
 */
enum class PlayAction(val action :String) {
    // 发送音乐的广播
    MUSIC("MUSIC"),
    // 播放完毕的广播
    FINISH("FINISH"),
    PLAY("PLAY"),
    NEXT("NEXT"),
    PREVIOUS("PREVIOUS"),
    PAUSE("PAUSE")
}