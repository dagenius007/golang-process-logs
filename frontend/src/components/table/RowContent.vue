<script setup lang="ts">
import { watchEffect, ref } from 'vue'

const statusMapping: Record<string, string> = {
  running: '#60CA57',
  sleeping: '#6f767e66',
  stopped: '#6f767e66',
  stopped_child: '#6f767e66',
  medium: '#F1872D',
  low: '#6f767e66',
  high: 'red'
}

interface ColumnProps {
  header: string
  accessor: string
  cell?: (columnValue: string | number) => any
  styles?: Record<string, string>
  isPill?: boolean
}

const props = defineProps<{
  row: Record<string, any>
  column: ColumnProps
  isLoading: boolean
}>()

const value = ref('')
const color = ref({})
const styleObject = ref({})

watchEffect(async () => {
  let _v = ''
  if (props.column.cell) {
    _v = props.column.cell(props.row[props.column?.accessor])
  } else {
    _v = props.row[props.column?.accessor] ?? '-'
  }

  value.value = _v

  if (props.column.isPill) {
    const _value = props.row[props.column?.accessor]
    color.value = statusMapping[_value]
    styleObject.value = {
      backgroundColor: statusMapping[_value]
    }
  }
})
</script>

<template>
  <p v-if="props.column.cell" v-html="value" />
  <p v-else :class="[props.column.isPill ? 'isPill' : null]" :style="styleObject">
    {{ value }}
  </p>
</template>

<style scoped>
.isPill {
  width: fit-content;
  height: 28px;
  padding: 0 15px;
  border-radius: 15px;
  display: flex;
  align-items: center;
  color: white;
}
</style>
