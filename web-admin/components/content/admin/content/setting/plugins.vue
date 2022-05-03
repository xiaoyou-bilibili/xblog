<template>
  <setting-friend v-if="name==='friend'" />
  <setting-project v-else-if="name==='project'" />
  <setting-donate v-else-if="name==='donate'" />
  <setting-navigation v-else-if="name==='navigation'" />
  <setting-barrages v-else-if="name==='barrages'" />
</template>

<script>
export default {
  components: {
    SettingFriend: () => import('@/components/plugins/friend'),
    SettingProject: () => import('@/components/plugins/project'),
    SettingDonate: () => import('@/components/plugins/donate'),
    SettingNavigation: () => import('@/components/plugins/navigation'),
    SettingBarrages: () => import('@/components/plugins/barrages')
  },
  props: { name: { type: String, default: '' } }
}
</script>

<style scoped>

</style>
