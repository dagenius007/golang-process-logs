<script setup lang="ts">
import { useProcessStore } from '@/stores/ProcessStore'
import { ref } from 'vue'
import TableComponent from '@/components/table/TableComponent.vue'
import InputComponent from '@/components/input/InputComponent.vue'
import SelectComponent from '@/components/select/SelectComponent.vue'
import { FunnelIcon } from '@heroicons/vue/24/outline'

const data: any[] = ref([
  {
    firstName: 'tanner',
    lastName: 'linsley'
  },
  {
    firstName: 'tandy',
    lastName: 'miller'
  },
  {
    firstName: 'joe',
    lastName: 'dirte'
  }
])

const columns = ref([
  {
    header: 'First name',
    accessor: 'firstName',
    styles: {
      width: '30%'
    },
    // cell: (value: string | number) => `<p>${value}</p>`
    cell: (value: string | number) => `<p>${value}</p>`
  },
  {
    header: 'Last name',
    accessor: 'lastName',
    styles: {
      width: '20%'
    }
  }
])

const store = useProcessStore()

const searchValue = ref('')

const options = ref([
  {
    label: 'Simple bale ',
    value: 'label'
  }
])

const optionValue = ref('')

console.log({ searchValue })

store.fetchProcesses()
</script>

<template>
  <div class="TableHeader">
    <div class="flex w-[30%]">
      <div class="FilterIcon">
        <Popper placement="bottom-end">
          <FunnelIcon class="h-6 w-6 text-black-500" />
          <template #content>
            <div class="FilterContent">
              <p>Filter</p>
              <div class="FilterOptions">
                <label>Status</label>
                <SelectComponent :options="options" v-model:modelValue="optionValue" />
              </div>

              <!-- <div class="FilterOptions">
                <label>Storage</label>
                <SelectComponent :options="options" v-model:modelValue="optionValue" />
              </div> -->
            </div>
          </template>
        </Popper>
      </div>
      <InputComponent v-model:modelValue="searchValue" />
    </div>
    <p>Hello</p>
  </div>
  <TableComponent :columns="columns" :rows="data" testId="payout-collection" :isLoading="false" />
  <div
    class="row mb-2 border-top border-bottom"
    v-for="process in store.processes"
    :key="process.key"
  >
    <div class="col-sm-2 mt-4">
      {{ process?.command }}
    </div>
    <div class="col-md offset-md-2 mt-4">price</div>
    <div class="col-sm-2 mt-4">
      <button class="btn btn-primary" type="submit">Add</button>
    </div>
  </div>
</template>

<style lang="postcss" scoped>
.TableHeader {
  @apply flex justify-between w-full;

  .FilterIcon {
    @apply h-[40px] w-[40px] flex items-center justify-between cursor-pointer;
    margin-right: 10px;
    background-color: white;
  }
}

.FilterContent {
  @apply p-4;
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
