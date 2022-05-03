// 管理员设置接口

import request from '@/utils/requests-v3'

// 基本路径
const base = process.env.SERVER + 'api/v3/admin/settings'
const row = process.env.SERVER + 'api/v3/'

// 获取站点设置的设置信息
export function getSiteOption (data) { return request(base + '/site', data, 'get') }
// 保存站点的设置
export function updateOption (data) { return request(base, data, 'put') }
// 获取壁纸设置
export function getBackgroundOption (data) { return request(base + '/background', data, 'get') }
// 获取侧边栏设置信息
export function getSideOption (data) { return request(base + '/side', data, 'get') }
// 更新侧边栏设置
export function updateSideOption (data) { return request(base + '/side', data, 'put') }
// 获取导航栏设置
export function getNavOption (data) { return request(base + '/nav', data, 'get') }
// 更新导航栏设置
export function updateNavOption (data) { return request(base + '/nav', data, 'put') }
// 获取微信小程序的设置信息
export function getWechatOption (data) { return request(base + '/other/mini_program', data, 'get') }
// 获取安卓APP的设置信息
export function getAPPOption (data) { return request(base + '/other/app', data, 'get') }
// 原始请求
export function rowRequest (data) { return request(row + data.url, data.data, data.type) }
