import { ref, computed } from 'vue'
import { defineStore } from 'pinia'
import api from '@/internals/api'

interface Process {
  key: number
  command: string
}
const useProcessStore = defineStore('process', () => {
  const count = ref(0)
  const processes = ref<Process[]>([])
  async function fetchProcesses() {
    try {
      const { data } = await api.get('/data')

      console.log({ data })

      processes.value = data
    } catch (e) {
      console.log({ e })
    }
  }

  return { count, processes, fetchProcesses }
})

export { useProcessStore }
