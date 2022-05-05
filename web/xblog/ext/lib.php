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
    echo '<script src="https://cdn.jsdelivr.net/npm/vue/dist/vue.js"></script>';
}
function lib_script_jquery(){
    echo '<script type="text/javascript" src="https://cdn.bootcss.com/jquery/3.1.1/jquery.min.js?ver=2.1.4"></script>';
}
// 自己封装的库
function lib_script_xiao_you(){
    echo '<script type="text/javascript" src="'.setting_web().'/static/js/xiaoyou.js"></script>';
}
// 代码高亮插件
function lib_script_highlight(){
    echo '<script src="//cdnjs.cloudflare.com/ajax/libs/highlight.js/9.15.10/highlight.min.js"></script>';
}
// 另外一个代码高亮插件


function lib_script_element(){
    echo '<script src="//unpkg.com/element-ui@2.15.1/lib/index.js"></script>';
}
function lib_script_layui(){
    echo '<script src="//cdnjs.cloudflare.com/ajax/libs/layui/2.6.5/layui.min.js"></script>';
}
// layui弹窗
function lib_script_layer(){
    echo '<script src="//cdn.bootcdn.net/ajax/libs/layer/3.3.0/layer.min.js"></script>';
}
function lib_script_fancybox(){
    echo '<script src="//cdn.bootcdn.net/ajax/libs/fancybox/3.5.6/jquery.fancybox.min.js"></script>';
}


function lib_css_element(){
    echo '<link rel="stylesheet" href="//unpkg.com/element-ui@2.15.1/lib/theme-chalk/index.css">';
}
function lib_css_font_awesome(){
    echo '<link href="//cdn.bootcdn.net/ajax/libs/font-awesome/5.15.3/css/all.css" rel="stylesheet">';
}
function lib_css_font_awesome4(){
    echo '<link href="//cdn.bootcdn.net/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css" rel="stylesheet">';
}
function lib_css_bootstrap(){
    echo '<link href="//cdn.jsdelivr.net/npm/bootstrap@4.6.0/dist/css/bootstrap.min.css" rel="stylesheet">';
}
function lib_css_layui(){
    echo '<link rel="stylesheet" href="//cdnjs.cloudflare.com/ajax/libs/layui/2.6.5/css/layui.min.css" />';
}
// 一款显示图片和视频的库
function lib_css_fancybox(){
    echo '<link rel="stylesheet" type="text/css" href="//cdn.bootcdn.net/ajax/libs/fancybox/3.5.6/jquery.fancybox.min.css" />';
}