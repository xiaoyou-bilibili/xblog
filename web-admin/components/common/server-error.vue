<template>
  <div class="auto">
    <div class="container">
      <div class="settings">
        <span class="icon">502</span>
        <h4>很抱歉,博客系统好像出了一点问题!</h4>
        <p>你可以清空浏览器缓存后再试，或者使用ctrl+f5强制刷新</p>
        <div>
          <a href="/" target="_blank">回到主页</a>
          <a href="javascript:window.location.reload()">刷新</a>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'ServerError'
}
</script>

<style scoped>
  body {
    background-color: #fff;
    font: 12px/1.5 'Microsoft Yahei', Tahoma, Helvetica, 'SimSun', sans-serif;
  }
  .auto {
    width: 1000px;
    margin: 130px auto;
  }
  .container {
    background: url(/static/images/error/erbg.png) no-repeat 560px top;
  }
  .settings {
    padding-left: 100px;
    padding-bottom: 50px;
  }
  .settings .icon {
    display: block;
    width: 242px;
    height: 106px;
    font-size: 90px;
  }
  .settings h4 {
    margin: 30px 0 15px 0;
    font-size: 18px;
    color: #2cb7fd;
  }
  .settings p {
    font-size: 14px;
    color: #999;
  }
  .settings > div {
    margin-top: 40px;
    font-size: 0;
  }
  .settings > div a {
    display: inline-block;
    padding: 10px 40px;
    border: 1px solid #2cb7fd;
    font-size: 15px;
    color: #2cb7fd;
    text-decoration: none;
  }
  .settings > div a:first-child {
    margin-right: 20px;
    color: #fff;
    background-color: #2cb7fd;
  }
</style>
