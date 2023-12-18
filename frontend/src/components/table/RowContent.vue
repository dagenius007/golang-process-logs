export defalut { props: ['isLoading', 'row', 'column'], setup() { console.log({ props }) const count
= ref(props.row[props.column?.accessor] || '-') // return { count } }, template: ` ` // Can also
target an in-DOM template: // template: '#my-template-element' }

<script setup lang="ts">
import { h, ref } from 'vue'
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

let v = ''
if (props.column.cell) {
  v = props.column.cell(props.row[props.column?.accessor])
} else {
  v = props.row[props.column?.accessor] || '-'
}

const value = ref(v)
</script>

<template>
  <h1 v-if="props.column.cell" v-html="value" />
  <h1 v-else>{{ value }}</h1>
</template>

<style lang="postcss" scoped>
.TableContainer {
  @apply overflow-scroll;
  .TableWrapper {
    @apply w-full whitespace-nowrap border-collapse;
  }

  .Th {
    @apply text-left font-medium;
    font-size: 13px;
    padding: 15px;
    color: #6f767e;
  }

  .Td {
    @apply text-sm font-medium;
    color: #1a1d1f;
    padding: 15px;
  }

  .Tr {
    @apply cursor-pointer;
    &:nth-child(odd) {
      background: rgba(244, 244, 244, 0.5);
    }
  }
}
</style>
