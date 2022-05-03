<template>
  <div>
    <el-row :gutter="20">
      <!--功能选择区-->
      <el-col :xs="24" :sm="15" :md="7">
        <type-index :options="wechatOption" />
      </el-col>
    </el-row>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import TypeIndex from '@/components/content/admin/content/setting/type/index'
export default {
  name: 'SettingWechat',
  components: { TypeIndex },
  computed: {
    // 把getter混入到computed对象中,我们可以直接调用vex里面的内容即可
    ...mapGetters('admin-settings', ['wechatOption'])
  }
}
</script>

<style scoped>

</style>
