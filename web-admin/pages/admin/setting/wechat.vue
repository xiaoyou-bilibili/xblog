<template>
  <div>
    <el-card class="box-card">
      <div slot="header" class="clearfix">
        <span>小程序/APP设置</span>
      </div>
      <el-tabs v-model="active" type="card" @tab-click="change">
        <el-tab-pane label="小程序设置" name="01">
          <setting-wechat />
        </el-tab-pane>
        <el-tab-pane label="APP设置" name="02">
          <setting-a-p-p />
        </el-tab-pane>
      </el-tabs>
    </el-card>
  </div>
</template>

<script>
import SettingWechat from '@/components/content/admin/content/setting/wechat'
import admin from '@/components/mixin/admin-seo'
import SettingAPP from '@/components/content/admin/content/setting/app'
export default {
  components: { SettingAPP, SettingWechat },
  layout: 'admin',
  mixins: [admin],
  data () {
    return {
      active: '01'
    }
  },
  mounted () {
    // 首先获取网站设置
    this.change({ name: '01' })
  },
  methods: {
    // 当导航栏改变的时候自动获取设置信息
    change (item) {
      switch (item.name) {
        case '01':
          this.$store.dispatch('admin-settings/getWechatOption')
          break
        case '02':
          this.$store.dispatch('admin-settings/getAPPOption')
          break
      }
    }
  }
}
</script>

<style scoped>
</style>
