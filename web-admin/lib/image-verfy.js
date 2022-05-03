// 图片验证码生成插件
let checkCode
function createCode (el) {
  checkCode = ''
  const codeLength = 4// 验证码的长度，可变
  const canvas = document.getElementById(el)// 获取画布
  const selectChar = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z']// 所有候选组成验证码的字符
  for (let i = 0; i < codeLength; i++) {
    const charIndex = Math.floor(Math.random() * 36)
    checkCode += selectChar[charIndex]
  }
  if (canvas) {
    const ctx = canvas.getContext('2d')
    ctx.fillStyle = '#FFFFFF'
    ctx.fillRect(0, 0, 70, 27)
    ctx.font = '20px arial'
    // 创建渐变
    const gradient = ctx.createLinearGradient(0, 0, canvas.width, 0)
    gradient.addColorStop(0, 'magenta')
    gradient.addColorStop(0.5, 'blue')
    gradient.addColorStop(1.0, 'red')
    // 用渐变填色
    ctx.strokeStyle = gradient
    ctx.strokeText(checkCode, 5, 20)// 画布上添加验证码
  }
}

export default {
  install (Vue, option) {
    // 生成验证码
    Vue.prototype.$createCode = (el) => {
      createCode(el)
    }
    // 判断验证码是否正确
    Vue.prototype.$verify = (code) => {
      return checkCode.toLowerCase() === code.toLowerCase()
    }
  }
}
