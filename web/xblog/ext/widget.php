<?php
/**
 * @description 侧边栏工具
 * @author 小游
 * @date 2021/04/30
 */

$widget = [];

function widget_title(){
    echo $GLOBALS['widget']["title"];
}
function get_widget_title(){
    return $GLOBALS['widget']["title"];
}
function widget_html(){
    echo $GLOBALS['widget']["html"];
}