<?php
/**
 * @description 自定义页面的点击事件
 * @author 小游
 * @date 2021/04/30
 */

use xblog\request\Api;

// 页面信息
$more = [];

// 初始化head标签
function more_init_head(){
    global $head,$more;
    $head = $more["head_meta"];
}
function more_css(){
    foreach ($GLOBALS["more"]["css"] as $css){
        echo "<link rel='stylesheet' href='".$css."' />";
    }
}
function more_script(){
    foreach ($GLOBALS["more"]["script"] as $script){
        echo "<script src='".$script."'></script>";
    }
}
function more_background(){
    echo $GLOBALS["more"]["background"];
}
function more_is_full(){
    return $GLOBALS["more"]["full"];
}
function more_has_side(){
    return $GLOBALS["more"]["side"];
}
function more_content(){
    echo $GLOBALS["more"]["content"];
}

// 获取友链信息
function more_get_friends(): array
{
    try {
        // 获取友人帐
        return Api::newAPi()->getFriends();
    } catch (Exception $e) {
        return [];
    }
}
