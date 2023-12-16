<script setup lang="ts">
import { useProcessStore } from '@/stores/ProcessStore'

const store = useProcessStore()

store.fetchProcesses()
</script>

<template>
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

<style scoped>
h1 {
  font-weight: 500;
  font-size: 2.6rem;
  position: relative;
  top: -10px;
}

h3 {
  font-size: 1.2rem;
}

.greetings h1,
.greetings h3 {
  text-align: center;
}

@media (min-width: 1024px) {
  .greetings h1,
  .greetings h3 {
    text-align: left;
  }
}
</style>
