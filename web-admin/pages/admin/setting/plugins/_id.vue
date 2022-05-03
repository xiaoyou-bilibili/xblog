<template>
  <div>
    <el-card class="box-card">
      <div v-if="pluginOptions.length>=1">
        <el-tabs type="card">
          <el-tab-pane v-for="(item,index) in pluginOptions" :key="index" :label="item.name" :name="index.toString()">
            <el-row v-if="item.type===1" :gutter="20">
              <!--功能选择区-->
              <el-col :xs="24" :sm="15" :md="7">
                <type-index :options="item.setting" />
              </el-col>
            </el-row>
            <plugins v-else-if="item.type===2" :name="item.extra" />
          </el-tab-pane>
        </el-tabs>
      </div>
    </el-card>
  </div>
</template>

<script>
import admin from '@/components/mixin/admin-seo'
import { mapGetters } from 'vuex'
import TypeIndex from '@/components/content/admin/content/setting/type/index'
import Plugins from '@/components/content/admin/content/setting/plugins'

export default {
  layout: 'admin',
  components: { TypeIndex, Plugins },
  mixins: [admin],
  computed: {
    // 把getter混入到computed对象中,我们可以直接调用vex里面的内容即可
    ...mapGetters('admin-plugins', ['pluginOptions'])
  },
  mounted () {
    this.$store.dispatch('admin-plugins/getPluginsSetting', this.$route.params.id).catch((msg) => {
      // 清空设置
      this.$store.commit('admin-plugins/setPluginsOption', [])
      this.$message.warning(msg)
    })
  }
}
</script>

<style scoped>

</style>
