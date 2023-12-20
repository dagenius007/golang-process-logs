export defalut { props: ['isLoading', 'row', 'column'], setup() { console.log({ props }) const count
= ref(props.row[props.column?.accessor] || '-') // return { count } }, template: ` ` // Can also
target an in-DOM template: // template: '#my-template-element' }

<script setup lang="ts">
import { watchEffect, ref } from 'vue'
interface ColumnProps {
  header: string
  accessor: string
  cell?: (columnValue: string | number) => any
  styles?: Record<string, string>
}

const props = defineProps<{
  row: Record<string, any>
  column: ColumnProps
  isLoading: boolean
}>()

const value = ref('')

watchEffect(async () => {
  let _v = ''
  if (props.column.cell) {
    _v = props.column.cell(props.row[props.column?.accessor])
  } else {
    _v = props.row[props.column?.accessor] ?? '-'
  }

  value.value = _v
})
</script>

<template>
  <p v-if="props.column.cell" v-html="value" />
  <p v-else>{{ value }}</p>
</template>
