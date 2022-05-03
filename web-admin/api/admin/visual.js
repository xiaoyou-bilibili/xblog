// 数可视化接口

import request from '@/utils/requests-v3'

// 基本路径
const base = process.env.SERVER + 'api/v3/admin/visualization'

// 获取文章数，用户数，浏览数，评论数
export function getTotal (data) { return request(base, data, 'get') }
// 获取文章分布情况
export function getPostDistributed (data) { return request(base + '/posts/distributed', data, 'get') }
// 获取文章的详细数据
export function getPostDetail (data) { return request(base + '/posts/detail', data, 'get') }
