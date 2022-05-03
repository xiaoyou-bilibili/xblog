<?php
    /**
     * @description 主函数入口
     * @author 小游
     * @date 2021/4/30
     */
    // 博客系统主页配置
    // 注意这个两个地址不能一样，要不然替换会出错。。。
    const SERVER = "http://192.168.123.119";
    const API = "http://127.0.0.1:2333";
    // 引入启动文件
    require __DIR__ . '/xblog/App.php';
    // 启动服务
    $handle = new xblog\App();