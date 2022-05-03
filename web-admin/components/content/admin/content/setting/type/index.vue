<template>
  <div>
    <div v-for="(item,index) in options" :key="index">
      <el-divider v-if="item.type==='divider'" content-position="left">
        {{ item.title }}
      </el-divider>
      <type-input v-else-if="item.type==='input'" :title="item.title" :field="item.key" :model.sync="item.value" />
      <type-number v-else-if="item.type==='number'" :title="item.title" :field="item.key" :model.sync="item.value" />
      <type-upload v-else-if="item.type==='upload'" :title="item.title" :field="item.key" :model.sync="item.value" />
      <type-text v-else-if="item.type==='text'" :title="item.title" :field="item.key" :model.sync="item.value" />
      <type-switch v-else-if="item.type==='switch'" :title="item.title" :field="item.key" :model.sync="item.value" />
      <type-row-request v-else-if="item.type==='row'" :title="item.title" :url="item.value" />
    </div>
  </div>
</template>

<script>
import TypeText from '@/components/content/admin/content/setting/type/text'
import TypeUpload from '@/components/content/admin/content/setting/type/upload'
import TypeNumber from '@/components/content/admin/content/setting/type/number'
import TypeInput from '@/components/content/admin/content/setting/type/input'
import TypeSwitch from '@/components/content/admin/content/setting/type/switch'
import TypeRowRequest from '@/components/content/admin/content/setting/type/row-request'

export default {
  name: 'TypeIndex',
  components: { TypeRowRequest, TypeSwitch, TypeText, TypeUpload, TypeNumber, TypeInput },
  props: { options: { type: Array, default: () => ([]) } }
}
</script>

<style scoped>

</style>
