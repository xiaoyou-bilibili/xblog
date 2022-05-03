// 自定义弹幕发送插件，用来发送弹幕
import $ from 'jquery'
// 发送弹幕
function barrage (barrage) {
  barrage = $.extend({
    close: true,
    bottom: 0,
    max: 10,
    speed: 6,
    color: '#fff',
    old_ie_color: '#000000'
  }, barrage || {})

  const time = new Date().getTime()
  const barrageId = 'barrage_' + time
  const id = '#' + barrageId
  const divBarrage = $("<div class='barrage' id='" + barrageId + "'></div>").appendTo($('#main'))
  const windowHeight = $(window).height() - 100
  const bottom = (barrage.bottom === 0) ? Math.floor(Math.random() * windowHeight + 40) : barrage.bottom
  divBarrage.css('bottom', bottom + 'px')
  const divBarrageBox = $("<div class='barrage_box cl'></div>").appendTo(divBarrage)
  if (barrage.avatar) {
    const img = $("<img src='' title='' >").appendTo(id + ' .barrage_box')
    img.attr('src', barrage.avatar)
    img.attr('title', barrage.nickname)
  }
  divBarrageBox.append(" <div class='z p'></div>")
  const content = $('<p></p>').appendTo(id + ' .barrage_box .p')
  content.empty().append(barrage.content)

  if (navigator.userAgent.indexOf('MSIE 6.0') > 0 || navigator.userAgent.indexOf('MSIE 7.0') > 0 || navigator.userAgent.indexOf('MSIE 8.0') > 0) {
    content.css('color', barrage.old_ie_color)
  } else {
    content.css('color', barrage.color)
  }

  let i = 0
  divBarrage.css('margin-right', i)
  let looper = setInterval(barrager, barrage.speed)

  function barrager () {
    const windowWidth = $(window).width() + 500
    if (i < windowWidth) {
      i += 1
      $(id).css('margin-right', i)
    } else {
      $(id).remove()
      return false
    }
  }
  divBarrageBox.mouseover(function () {
    clearInterval(looper)
  })
  divBarrageBox.mouseout(function () {
    looper = setInterval(barrager, barrage.speed)
  })
}

export default {
  install (Vue, option) {
    // 我们直接混入一个方法
    Vue.mixin({
      methods: {
        startBarrage () {
          // 获取所有的弹幕
          this.$store.dispatch('more/getBarrage').then((res) => {
            // 获取到了弹幕
            const barrages = res
            const time = new Date().getTime()
            const barrageId = 'barrage_' + time
            const id = '#' + barrageId
            $('#main').append('<div id="' + barrageId + '" class="danmu-first"  style="position:fixed;"><img src="/static/images/c61.png"/><p>' + barrages.length + '条弹幕已填充完毕，准备发射!</p></div>')
            let i = 0
            let j = 0
            // 不断遍历输出内容
            // 显示弹幕
            const show = setInterval(function () {
              if (i < barrages.length) {
                barrage(barrages[i])
              } else {
                // 清除定时器
                clearInterval(show)
              }
              i++
            }, 1000)

            // 显示弹幕数
            setInterval(function () {
              const windowWidth = $(window).width() + 500
              if (j < windowWidth) {
                j += 2
                $(id).css('right', j)
              } else {
                $(id).remove()
                return false
              }
            }, 20)
          }).catch(() => {
            this.$message.error('获取弹幕失败')
          })
        }
      }
    })
  }
}
