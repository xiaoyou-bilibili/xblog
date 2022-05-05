<?php
/**
 * @description 文章处理板块
 * @author 小游
 * @date 2021/04/30
 */

use xblog\ext\Load;
use xblog\request\Api;

// 文章信息
$post = [];
// 当前文章数据
$current = [];

// 文章板块提供的相关函数
function get_post_no(){
    return $GLOBALS['post']["no"];
}
function post_title(){
    echo $GLOBALS['post']["title"];
}
function post_summer(){
    echo $GLOBALS['post']["content"];
}
function post_content(){
    echo $GLOBALS['post']["content"];
}
function post_id(){
    echo $GLOBALS['post']["id"];
}
function get_post_id(){
    return $GLOBALS['post']["id"];
}
function post_date(){
    echo $GLOBALS['post']["date"];
}
function post_good(){
    echo $GLOBALS['post']["good"];
}
function post_view(){
    echo $GLOBALS['post']["view"];
}
function post_image(){
    echo $GLOBALS['post']["image"];
}
function post_comment(){
    echo $GLOBALS['post']["comment"];
}
function post_modify(){
    echo $GLOBALS['post']["modify"];
}
function post_alipay(){
    echo $GLOBALS['post']["alipay"];
}
function post_wechat(){
    echo $GLOBALS['post']["wechat"];
}
function get_post_encrypt(){
    return $GLOBALS['post']["encrypt"];
}
function get_post_category(){
    // 判断一下大小
    if (count($GLOBALS['post']["category"])<=0){
        return [
            ["name"=>"无分类","link"=>"#"]
        ];
    } else {
        return $GLOBALS['post']["category"];
    }
}
function get_post_tags(){
    return $GLOBALS['post']["tag"];
}
function get_post_is_top(){
    return $GLOBALS['post']["is_top"];
}
function get_post_encryption(){
    return $GLOBALS['post']["encryption"];
}
// 总页数
function current_total(){
    echo $GLOBALS['current']["total"];
}
// 数据总数
function current_count(){
    echo $GLOBALS['current']["total"]*count($GLOBALS['current']["contents"]);
}
function get_current_count(){
    return $GLOBALS['current']["total"]*count($GLOBALS['current']["contents"]);
}
// 第几页
function current_current(){
    echo $GLOBALS['current']["current"];
}
function get_current_current(){
    return $GLOBALS['current']["current"];
}
// 页面大小
function current_size(){
    echo count($GLOBALS['current']["contents"]);
}

// 评论数据
$comments = [];
function get_post_comments(){
    return $GLOBALS['comments'];
}

// 获取文章评论
function xy_comments($id){
    global $comments;
    try {
        $comments = Api::newAPi()->getPostComment($id);
        Load::loadPage("/content/comments.php");
    } catch (Exception $e) {

    }
}

// 获取文章分类
function get_index_category(): array
{
    $category = [
        "parent"=>[],
        "child"=>[]
    ];
    try {
        return Api::newAPi()->getPostCategory();
    } catch (Exception $e){
        echo $e;
        return $category;
    }
}
