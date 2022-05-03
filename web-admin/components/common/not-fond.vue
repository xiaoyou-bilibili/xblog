<template>
  <div id="main" class="panel-cover">
    <span class="mobile btn-mobile-menu">
      <i class="fa fa-list btn-mobile-menu__icon" />
      <i class="fa fa-angle-up btn-mobile-close__icon hidden" />
    </span>
    <header class="panel-cover" style="background-image: url(/static/images/error/bg.jpg)">
      <div class="panel-main">
        <div class="panel-main__inner panel-inverted">
          <div class="panel-main__content">
            <a href="/" title="回到首页"><img src="/static/images/error/few.jpg" width="150" alt="Noisky logo" class="panel-cover__logo logo"></a>
            <h1 class="panel-cover__title panel-title">
              <a href="/" title="link to homepage for noisky">404</a>
            </h1>

            <hr class="panel-cover__divider">
            <p class="panel-cover__description">
              对不起，您要找的页面不存在!
            </p>
            <hr class="panel-cover__divider panel-cover__divider--secondary">

            <!--            <p class="panel-cover__description">-->
            <!--              欢迎来到我的个人世界!-->
            <!--            </p>-->

            <p class="footer">
              © 2021  |  <a style="TEXT-DECORATION: none" href="/" rel="nofollow" target="_blank"><span style="color: #b3b3b3" /></a>
            </p>

            <!--            <div class="navigation-wrapper">-->
            <!--              <div>-->
            <!--                <nav class="cover-navigation cover-navigation&#45;&#45;primary">-->
            <!--                  <ul class="navigation">-->
            <!--                    <li class="navigation__item">-->
            <!--                      <a href="" title="访问博客">博客</a>-->
            <!--                    </li>-->
            <!--                    <li class="navigation__item">-->
            <!--                      <a href="" target="_blank" title="</a>-->
            <!--                    </li>-->
            <!--                    <li class="navigation__item">-->
            <!--                      <a href="" title="">百度么</a>-->
            <!--                    </li>-->
            <!--                    <li class="navigation__item">-->
            <!--                      <a href="" title="">旗下网站</a>-->
            <!--                    </li>-->
            <!--                    <li class="navigation__item">-->
            <!--                      <a href="" title="">关于</a>-->
            <!--                    </li>-->
            <!--                  </ul>-->
            <!--                </nav>-->
            <!--              </div>-->
            <!--            </div>-->
          </div>
        </div>
        <div class="panel-cover--overlay cover-slate" />
      </div>
    </header>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'

export default {
  name: 'NotFound',
  computed: {
    // 计算属性映射
  }
}
</script>

<style src="~/assets/css/error.css" scoped />
