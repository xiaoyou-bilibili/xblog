<?php

use xblog\request\Api;

/**
 * @description 用户板块
 * @author 小游
 * @date 2021/05/13
 */
// 登录注册界面报错信息
$errorMessage = "";
// 登录注册界面的正确信息
$okMessage = "";
// 用户信息(默认为空)
$userInfo = null;
// 获取错误信息
function get_err_msg(){
    return $GLOBALS["errorMessage"];
}
// 获取正确信息
function get_ok_msg(){
    return $GLOBALS["okMessage"];
}
// 获取用户信息
function get_user_info(){
    if (user_is_login()) {
        return $GLOBALS["userInfo"];
    } else {
        return [];
    }
}
// 判断用户是否登录
function user_is_login(): bool
{
    global $userInfo;
    if ($userInfo==null){
        try {
            $userInfo = Api::newAPi()->getInfo();
        } catch (Exception $e){
            return false;
        }
    }
    return true;
}
function user_avatar(){
    echo $GLOBALS["userInfo"]["avatar"];
}
function user_sign(){
    echo $GLOBALS["userInfo"]["sign"];
}
function user_level(){
    echo $GLOBALS["userInfo"]["level"];
}
function user_hang(){
    echo $GLOBALS["userInfo"]["hang"];
}
function user_username(){
    echo $GLOBALS["userInfo"]["username"];
}
function user_nickname(){
    echo $GLOBALS["userInfo"]["nickname"];
}
function user_email(){
    echo $GLOBALS["userInfo"]["email"];
}
function user_user_id(){
    echo $GLOBALS["userInfo"]["user_id"];
}
function user_identity(){
    echo $GLOBALS["userInfo"]["identity"];
}
function get_user_subscription(){
    return $GLOBALS["userInfo"]["subscription"];
}