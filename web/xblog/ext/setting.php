<?php
/**
 * @description 设置有关的函数
 * @author 小游
 * @date 2021/4/30
 */

use xblog\request\Api;
use xblog\util\File;
// 一些路径
const THEME_DIR = __DIR__ . "/../../themes/";

// 设置信息
$index = null;
// 头部信息
$head = [];
// 登录界面设置
$access = null;

// 获取主页设置
function setting_index($page=''){
    global $index,$head;
    // 如果index为空，那么就获取
    if ($index==null){
        try {
            $index = Api::newAPi()->getSettingIndex($page);
            $head = $index["head_meta"];
        } catch (Exception $e) {

        }
    }
}

// 获取登录注册界面设置
function setting_access(){
    global $access,$head;
    // 如果index为空，那么就获取
    if ($access==null){
        try {
            $access = Api::newAPi()->getSettingLogin();
            $head = $access["head_meta"];
        } catch (Exception $e) {

        }
    }
}


// 获取head信息
function setting_head_title(){
    echo $GLOBALS["head"]["title"];
}
function setting_head_keyword(){
    echo $GLOBALS["head"]["keyword"];
}
function setting_head_description(){
    echo $GLOBALS["head"]["description"];
}
function setting_head_url(){
    echo $GLOBALS["head"]["url"];
}
function setting_head_image(){
    echo $GLOBALS["head"]["image"];
}
function setting_head_icon(){
    echo $GLOBALS["head"]["icon"];
}

// 获取主页信息
function setting_index_head_nav(){
    return $GLOBALS["index"]["head_nav"];
}
function setting_index_background(){
    echo $GLOBALS["index"]["background"];
}
function setting_index_navigation_background(){
    echo $GLOBALS["index"]["navigation_background"];
}
function setting_index_site_name(){
    echo $GLOBALS["index"]["site_name"];
}
function setting_index_site_url(){
    echo $GLOBALS["index"]["site_url"];
}
function setting_index_description(){
    echo $GLOBALS["index"]["description"];
}
function setting_index_build_time(){
    echo $GLOBALS["index"]["build_time"];
}
function setting_index_bei_an(){
    echo $GLOBALS["index"]["bei_an"];
}
function setting_index_gov_bei_an(){
    echo $GLOBALS["index"]["gov_bei_an"];
}
function get_setting_index_gov_bei_an(){
    return $GLOBALS["index"]["gov_bei_an"];
}
function setting_index_left_side(){
    echo $GLOBALS["index"]["left_side"];
}
function setting_index_right_side(){
    echo $GLOBALS["index"]["right_side"];
}

// 登录界面设置
function setting_access_background(){
    echo $GLOBALS["access"]["background"];
}
function setting_access_logo(){
    echo $GLOBALS["access"]["logo"];
}
function setting_access_web_text(){
    echo $GLOBALS["access"]["web_text"];
}
function setting_access_site_name(){
    echo $GLOBALS["access"]["site_name"];
}

// 服务地址
function setting_web(): string
{
    return SERVER;
}
function setting_api(): string
{
    return API;
}
// 主题地址
function setting_template(): string
{
    return SERVER."/themes/".THEME;
}

/**
 * 获取网站设置信息
 * @param string $key
 * @return mixed
 */
function xy_option(string $key)
{
    $value = null;
    try {
        $value = Api::newAPi()->getSettingThemes($key);
    } catch (Exception $e) {
    }
    // 判断是否有设置，如果没有就自己读取setting.json文件
    if ($value==null || !isset($value["value"])){
        // 获取切换的主题名字
        $filename = THEME_DIR.THEME."/setting.json";
        try{
           $res = File::readFile($filename,true);
           foreach ($res as $tmp){
               foreach ($tmp["items"] as $tmp2) {
                   if ($tmp2["key"]==$key){
                       return $tmp2["default"];
                   }
               }
           }
        }catch (Exception $e){
        }
    } else {
        return $value["value"];
    }
    return "";
}