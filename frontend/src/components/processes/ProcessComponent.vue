<script setup lang="ts">
import { useProcessStore } from '@/stores/ProcessStore'
import { ref, watchEffect } from 'vue'
import TableComponent from '@/components/table/TableComponent.vue'
import InputComponent from '@/components/input/InputComponent.vue'
import SelectComponent from '@/components/select/SelectComponent.vue'
import { FunnelIcon } from '@heroicons/vue/24/outline'
import Card from '@/components/card/Card.vue'
import Button from '@/components/button/Button.vue'
import { debounce } from '@/utils/debounce'
import { storeToRefs } from 'pinia'
import ProcessChart from './ProcessChart.vue'

const store = useProcessStore()

store.fetchProcesses({})
store.fetchProcessUsers()
store.fetchProcessInfo()

const { processes, totalCount, stateOptions, userOptions, processInfo } = storeToRefs(store)

console.log('processes', processes)

const columns = ref([
  {
    header: 'PID',
    accessor: 'pid'
  },
  {
    header: 'User',
    accessor: 'user'
  },

  {
    header: 'Command',
    accessor: 'command'
  },
  {
    header: 'CPU(%)',
    accessor: 'cpu_usage'
  },
  {
    header: 'Memory(%)',
    accessor: 'memory_usage'
  },
  {
    header: 'Resident Memory Size(MB)',
    accessor: 'resident_memorySize'
  },
  {
    header: 'Virtual Memory Size(MB)',
    accessor: 'virtual_memory_size'
  },
  {
    header: 'State',
    accessor: 'state',
    isPill: true
  },
  {
    header: 'Priority',
    accessor: 'priority',
    isPill: true
  },
  {
    header: 'Process Time',
    accessor: 'total_time'
  },
  {
    header: 'CPU Time',
    accessor: 'cpu_time'
  }
])

const searchValue = ref('')
const page = ref(0)
const state = ref('')
const user = ref('')

function handlePageChange(value: any) {
  page.value = value.page
  store.fetchProcesses(
    {
      state: state.value,
      user: user.value
    },
    value?.limit,
    value?.page
  )
}

function handleFilter() {
  if (state.value !== '' || user.value !== '') {
    store.fetchProcesses({
      state: state.value,
      user: user.value
    }),
      500
  }
}

watchEffect(async () => {
  if (searchValue.value != '') {
    debounce(store.fetchProcesses({ search: searchValue.value }), 500)
  }
})
</script>

<template>
  <div class="Cards">
    <div v-for="(info, i) in processInfo" :key="i">
      <Card :title="info.title" :value="info.value" />
    </div>
  </div>
  <ProcessChart />
  <div class="TableHeader">
    <div class="flex md:w-[30%] w-full">
      <div class="FilterIcon">
        <Popper placement="bottom-end">
          <FunnelIcon class="h-6 w-6 text-[#4F4999]" />
          <template #content>
            <div class="FilterContent">
              <p>Filter</p>
              <div class="FilterOptions">
                <label>Status</label>
                <SelectComponent :options="stateOptions" v-model:modelValue="state" />
              </div>

              <div class="FilterOptions">
                <label>Users</label>
                <SelectComponent :options="userOptions" v-model:modelValue="user" />
              </div>

              <Button label="Filter" @clicked="handleFilter"></Button>
            </div>
          </template>
        </Popper>
      </div>
      <InputComponent v-model:modelValue="searchValue" :isSearch="true" />
    </div>
  </div>

  <TableComponent
    :columns="columns"
    :rows="processes"
    testId="process-collection"
    :isLoading="false"
    :isPaginated="true"
    :totalcount="totalCount"
    @page-change="handlePageChange"
  />
</template>

<style lang="postcss" scoped>
.Cards {
  @apply md:grid gap-6 md:grid-cols-4 mb-20;
}
.TableHeader {
  @apply md:flex justify-between w-full mb-3 mt-4;

  .FilterIcon {
    @apply h-[40px] w-[40px] flex items-center justify-between cursor-pointer;
    background-color: white;
    &:last-child {
      margin-bottom: 30px;
    }
  }
}

.FilterContent {
  @apply p-6 rounded-2xl;
  background-color: white;
  box-shadow: 0px 6px 32px 0px #1529521f;
  width: 300px;
  & > p {
    @apply mb-4;
    font-size: 16px;
  }
  .FilterOptions {
    @apply mb-4;

    & > label {
      @apply text-sm mb-2 block;
    }
  }
}
</style>
