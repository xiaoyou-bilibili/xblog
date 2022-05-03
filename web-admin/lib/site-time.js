// 网站运行时间统计
function createTime (time, el) {
  const now = new Date()
  const grt = new Date(time)
  now.setTime(now.getTime() + 250)
  const days = (now - grt) / 1000 / 60 / 60 / 24
  const dnum = Math.floor(days)
  const hours = (now - grt) / 1000 / 60 / 60 - (24 * dnum)
  let hnum = Math.floor(hours)
  if (String(hnum).length === 1) {
    hnum = '0' + hnum
  }
  const minutes = (now - grt) / 1000 / 60 - (24 * 60 * dnum) - (60 * hnum)
  let mnum = Math.floor(minutes)
  if (String(mnum).length === 1) {
    mnum = '0' + mnum
  }
  const seconds = (now - grt) / 1000 - (24 * 60 * 60 * dnum) - (60 * 60 * hnum) - (60 * mnum)
  let snum = Math.round(seconds)
  if (String(snum).length === 1) {
    snum = '0' + snum
  }
  document.getElementById(el).innerHTML = dnum + '天' + hnum + '小时' + mnum + '分' + snum + '秒'
}
export default {
  install (Vue, option) {
    // 定义一个方法用来触发时间统计
    Vue.prototype.$countTime = (build, el) => {
      createTime(build, el)
    }
  }
}
