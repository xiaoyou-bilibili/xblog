<template>
  <div>
    <not-found v-if="error.statusCode === 404" />
    <server-error v-if="error.statusCode === 502" />
  </div>
</template>

<script>
import ServerError from '@/components/common/server-error'
import NotFound from '@/components/common/not-fond'

export default {
  components: { NotFound, ServerError },
  props: ['error']
}
</script>

<style scoped>

</style>
