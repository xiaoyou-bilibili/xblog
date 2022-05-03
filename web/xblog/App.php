<?php

namespace xblog;

// 自动注册文件
spl_autoload_register(function($className){
    // 获取文件名
    $file = dirname(dirname(__FILE__)) . '/' . str_replace('\\', '/', $className) . '.php';
    // 加载文件
    if (is_file($file)){
        require $file;
    }
});

use xblog\route\Handle;

/**
 * Class App 核心包
 * @description APP核心包
 * @author 小游
 * @date 2021/4/30
 * @package xblog
 */
class App
{
    public function __construct()
    {
        // 首先处理请求
        $handle = new Handle();
    }
}