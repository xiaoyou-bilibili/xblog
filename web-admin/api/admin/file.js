// 上传文件相关的接口

import request from '@/utils/requests-v3'

// 基本路径
const base = process.env.SERVER + 'api/v3/tools/file'

// 上传图片
export function uploadImage (data) { return request(base + '/images', data, 'post') }
