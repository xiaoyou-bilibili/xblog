<?php
/**
 * @description 主页文件相关的函数
 * @author 小游
 * @date 2021/4/30
 */

use xblog\ext\Load;
use xblog\request\Api;

/**
 *  获取头部信息
 */
function get_header($page=''){
    // 获取主页设置
    setting_index($page);
    // 判断是否需要执行其他函数
    if ($page=='plugin'){
        more_init_head();
    }
//    if ($action!=""){
//        $action();
//    }
    Load::loadPage("/common/header.php");
}

/**
 * 获取底部信息
 */
function get_footer(){
    Load::loadPage("/common/footer.php");
}

/**
 *  加载文章
 */
function xy_posts(){
    global $post,$current;
    try {
        // 获取文章列表
        $data = Api::newAPi()->getPostList($_GET);
        // 设置当前数据
        $current = $data;
        $no = 0;
        foreach ($data["contents"] as $item){
            $post = $item;
            $post["no"] = $no;
            Load::loadPage("/content/post.php");
            $no++;
        }
    } catch (Exception $e) {

    }
}

/**
 *  加载左侧边栏
 */
function xy_left_side(){
    global $widget,$index;
    foreach ($index["left_side"] as $item){
        $widget = $item;
        Load::loadPage("/common/side.php");
    }
}

/**
 *  加载右侧边栏
 */
function xy_right_side(){
    global $widget,$index;
    foreach ($index["right_side"] as $item){
        $widget = $item;
        Load::loadPage("/common/side.php");
    }
}