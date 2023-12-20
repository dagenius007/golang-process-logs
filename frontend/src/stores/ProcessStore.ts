import { ref } from 'vue'
import { defineStore } from 'pinia'
import api from '@/internals/api'
import { isHttps } from '@/utils/helper'

interface Process {
  key: number
  command: string
}

const handleWebsocketMessage = (data: string) => {
  let processes = []
  if (data) {
    try {
      processes = JSON.parse(data)
    } catch (e) {
      console.error(e)
    }
  }

  return processes
}
const useProcessStore = defineStore('process', () => {
  const processes = ref<Process[]>([])

  const userOptions = ref<{ label: string; value: string }[]>([])
  const processInfo = ref<Record<string, string>[]>([])

  const totalCount = ref(0)

  const stateOptions = ref([
    {
      label: 'Sleeping',
      value: 'sleeping'
    },
    {
      label: 'Running',
      value: 'running'
    },
    {
      label: 'Stopped Child Process',
      value: 'stopped_child'
    },
    {
      label: 'Stopped',
      value: 'stopped'
    },
    {
      label: 'Idle',
      value: 'idle'
    }
  ])

  // let dataInterval
  async function fetchProcesses(params: Record<string, string | number>, limit = 10, page = 1) {
    try {
      //build query

      let query = `limit=${limit}&page=${page}`

      Object.keys(params).forEach((key: string) => {
        if (params[key]) {
          const value = params[key] as string
          query += `&${key}=${value}`
        }
      })
      const {
        data: { data }
      } = await api.get(`/data?${query}`)

      processes.value = data.processes

      totalCount.value = data.total

      const apiHost = import.meta.env.VITE_API_HOST

      const socket = new WebSocket(`${isHttps() ? 'wss:' : 'ws:'}//${apiHost}/ws?${query}`)

      socket.onmessage = (event: MessageEvent) => {
        const _processes = handleWebsocketMessage(event.data)
        if (_processes.length > 0) {
          processes.value = _processes
        }
        // }, 300)
      }
    } catch (e) {
      console.error({ e })
    }
  }

  async function fetchProcessUsers() {
    try {
      const {
        data: { data }
      } = await api.get(`/users`)

      userOptions.value = data.map((_data: string) => ({ label: _data, value: _data }))
    } catch (e) {
      console.error({ e })
    }
  }

  async function fetchProcessInfo() {
    try {
      const {
        data: { data }
      } = await api.get(`/counts`)

      Object.keys(data).map((key: string) => {
        switch (true) {
          case key == 'processCount':
            processInfo.value.push({ title: 'Total process', value: data[key].toString() })
            break
          case key == 'usersCount':
            processInfo.value.push({ title: 'Total users', value: data[key].toString() })
            break
        }
      })
    } catch (e) {
      console.log({ e })
    }
  }
  return {
    processes,
    fetchProcesses,
    totalCount,
    stateOptions,
    userOptions,
    fetchProcessUsers,
    fetchProcessInfo,
    processInfo
  }
})

export { useProcessStore }
