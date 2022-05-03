import * as post from  "../../config/api"
import requestProcess from "../../utils/request-process"
export const state = {
  collections: []
}

// 定义一个空函数
const none = _=>{}
// 获取首页文章
export function getPostList(data){return requestProcess(post.postGetPostList,none,none,data)}
// 获取文章内容
export function getPostContent(data){return requestProcess(post.postGetContent,none,none,data)}
// 获取加密文章内容
export function getEncryption(id,data){return requestProcess(post.postGetEncryption,none,none,id,data)}
// 获取文章评论
export function getComment(data){return requestProcess(post.postGetComment,none,none,data)}
// 判断用户是否收藏文章
export function getPostCollection(id,openid){return requestProcess(post.postGetPostCollection,none,none,id,openid)}
// 用户给文章收藏或点赞
export function updatePostCollection(id,openid,data){return requestProcess(post.postUpdateCollection,none,none,id,openid,data)}
// 用户提交评论
export function commitComment(id,data){return requestProcess(post.postCommitComment,none,none,id,data)}