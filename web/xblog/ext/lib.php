<?php
/**
 * @description 常用开源库
 * @author 小游
 * @date 2021/05/01
 */
// 几个CDN加速库
// https://www.bootcdn.cn/
// https://www.jsdelivr.com/?docs=gh

function lib_script_vue(){
    echo '<script src="//cdn.bootcdn.net/ajax/libs/vue/3.2.33/vue.cjs.js"></script>';
}
function lib_script_jquery(){
    echo '<script type="text/javascript" src="//cdn.bootcdn.net/ajax/libs/jquery/3.1.1/jquery.min.js?ver=2.1.4"></script>';
}
// 自己封装的库
function lib_script_xiao_you(){
    echo '<script type="text/javascript" src="'.setting_web().'/static/js/xiaoyou.js"></script>';
}
// 代码高亮插件
function lib_script_highlight(){
    echo '<script src="//cdn.bootcdn.net/ajax/libs/highlight.js/9.15.10/highlight.min.js"></script>';
}
// 另外一个代码高亮插件


function lib_script_element(){
    echo '<script src="//cdn.bootcdn.net/ajax/libs/element-ui/2.15.8/index.js"></script>';
}
function lib_script_layui(){
    echo '<script src="//cdn.bootcdn.net/ajax/libs/layui/2.6.5/layui.min.js"></script>';
}
// layui弹窗
function lib_script_layer(){
    echo '<script src="//cdn.bootcdn.net/ajax/libs/layer/3.3.0/layer.min.js"></script>';
}
function lib_script_fancybox(){
    echo '<script src="//cdn.bootcdn.net/ajax/libs/fancybox/3.5.6/jquery.fancybox.min.js"></script>';
}


function lib_css_element(){
    echo '<link rel="stylesheet" href="//cdn.bootcdn.net/ajax/libs/element-ui/2.15.8/theme-chalk/index.min.css">';
}
function lib_css_font_awesome(){
    echo '<link href="//cdn.bootcdn.net/ajax/libs/font-awesome/5.15.3/css/all.css" rel="stylesheet">';
}
function lib_css_font_awesome4(){
    echo '<link href="//cdn.bootcdn.net/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css" rel="stylesheet">';
}
function lib_css_bootstrap(){
    echo '<link href="//cdn.bootcdn.net/ajax/libs/twitter-bootstrap/4.6.0/css/bootstrap.min.css" rel="stylesheet">';
}
function lib_css_layui(){
    echo '<link rel="stylesheet" href="//cdn.bootcdn.net/ajax/libs/layui/2.6.5/layui.min.js" />';
}
// 一款显示图片和视频的库
function lib_css_fancybox(){
    echo '<link rel="stylesheet" type="text/css" href="//cdn.bootcdn.net/ajax/libs/fancybox/3.5.6/jquery.fancybox.min.css" />';
}