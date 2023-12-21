<script setup lang="ts">
import RowContent from './RowContent.vue'
import { ref } from 'vue'
interface ColumnProps {
  header: string
  accessor: string
  cell?: (columnValue: string | number) => any
  styles?: Record<string, string>
}
interface TableProps {
  columns: ColumnProps[]
  rows: Record<string, any>[]
  isLoading: boolean
  testId?: string
  isPaginated?: boolean
  isPill?: boolean
  totalcount: number
}

const props = defineProps<TableProps>()
const currentPage = ref(1)

const emit = defineEmits(['page-change'])

function onPageChange() {
  emit('page-change', { page: currentPage.value, limit: 10 })
}
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

      <tbody>
        <tr
          class="Tr"
          v-for="(row, index) in props.rows"
          :key="index"
          :data-testid="`${testId}-row`"
        >
          <td class="Td" v-for="(column, i) in props.columns" :key="i" :style="column.styles">
            <RowContent :isLoading="isLoading" :column="column" :row="row" />
          </td>
        </tr>
      </tbody>
    </table>
  </div>
  <div class="Pagination">
    <vue-awesome-paginate
      :total-items="totalcount"
      :items-per-page="10"
      :max-pages-shown="10"
      v-model="currentPage"
      :on-click="onPageChange"
    />
  </div>
</template>

<style lang="postcss">
.TableContainer {
  @apply overflow-scroll;
  .TableWrapper {
    @apply w-full whitespace-nowrap border-collapse;
  }

  .Th {
    @apply text-left font-medium;
    font-size: 13px;
    padding: 20px;
    color: #6f767e;
  }

  .Td {
    @apply text-xs;
    color: #5c5c80;
    padding: 20px;
  }

  .Tr {
    @apply cursor-pointer;
    &:nth-child(odd) {
      background: rgba(244, 244, 244, 0.5);
    }
  }
}

.Pagination {
  @apply mt-5 flex justify-center;
}

.pagination-container {
  display: flex;
  column-gap: 10px;
}
.paginate-buttons {
  height: 40px;
  width: 40px;
  cursor: pointer;
  background-color: transparent;
  color: black;
  border: 0;
}
.paginate-buttons:hover {
  background-color: #d8d8d8;
  border-radius: 50%;
}
.active-page {
  background-color: #b188e6;
  border: 1px solid #b188e6;
  color: white;
  border-radius: 50%;
}
.active-page:hover {
  background-color: #b188e6;
}
</style>
