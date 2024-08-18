<script setup lang="ts">
import { useProcessStore } from '@/stores/ProcessStore'
import { storeToRefs } from 'pinia'

const store = useProcessStore()

// store.fetchProcessUsersReport()

const { options, series } = storeToRefs(store)
</script>

<template>
  <div class="Chart">
    <apexchart height="500" type="bar" :options="options" :series="series"></apexchart>
  </div>
</template>

<style lang="postcss" scoped>
.Chart {
  width: 100%;
}
</style>
