<script setup lang="ts">
import { ref } from 'vue'
import RowContent from './RowContent.vue'
interface ColumnProps {
  header: string
  accessor: string
  cell?: (columnValue: string | number) => any
  styles?: Record<string, string>
}
interface TableProps {
  columns: ColumnProps[]
  rows: Record<string, any>[]
  isLoading?: boolean
  testId?: string
}

const props = defineProps<TableProps>()
</script>

<template>
  <div class="TableContainer">
    <table class="TableWrapper">
      <thead>
        <tr>
          <th
            class="Th"
            v-for="(column, index) in props.columns"
            :key="index"
            :style="column.styles"
          >
            {{ column.header }}
          </th>
        </tr>
      </thead>

      <!-- <tbody>
          {rows.map((row: Record<string, string | number>, i: number) => (
            <Tr key={i} data-testid={`${testId}-row`}>
              {columns.map((column: ColumnProps, i: number) => (
                <Td key={i} style={column.styles}>
                  <RowContent isLoading={isLoading} row={row} column={column} />
                </Td>
              ))}
            </Tr>
          ))}
        </tbody> -->

      <tbody>
        <tr
          class="Tr"
          v-for="(row, index) in props.rows"
          :key="index"
          :data-testid="`${testId}-row`"
        >
          <!-- {columns.map((column: ColumnProps, i: number) => ( -->
          <td class="Td" v-for="(column, i) in props.columns" :key="i" :style="column.styles">
            Hello
            <RowContent :isLoading="isLoading" :column="column" :row="row" />
            <!-- <RowContent isLoading="{isLoading}" row="{row}" column="{column}" /> -->
          </td>
          <!-- ))} -->
        </tr>
      </tbody>
    </table>
  </div>
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
