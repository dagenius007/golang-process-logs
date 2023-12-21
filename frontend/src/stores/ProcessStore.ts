import { ref } from 'vue'
import { defineStore } from 'pinia'
import api from '@/internals/api'
import { isHttps } from '@/utils/helper'

interface IProcess {
  key: number
  command: string
}

interface IOptions {
  chart: {
    id: string
  }
  xaxis: {
    categories: string[]
  }
  plotOptions?: {
    bar: {
      horizontal: boolean
      columnWidth: string
      endingShape: string
    }
  }
  legend?: {
    show?: boolean
  }
  dataLabels: {
    enabled: boolean
  }
}

interface ISeries {
  name: string
  data: number[]
}

const handleWebsocketMessage = (data: string) => {
  let _data = {
    processes: [],
    reports: []
  }
  if (data) {
    try {
      _data = JSON.parse(data)
    } catch (e) {
      console.error(e)
    }
  }

  return _data
}

const createSerisAndOptions = (data: any) => {
  const series = [
    {
      name: 'Total CPU Usage(%)',
      data: [] as number[]
    },
    {
      name: 'Total Memory Usage(%)',
      data: [] as number[]
    },
    {
      name: 'Total Processes',
      data: [] as number[]
    }
  ]

  const xAxixs = [] as string[]

  if (data.length > 0) {
    for (const value of data) {
      Object.keys(value).forEach((key: string) => {
        if (key === 'user') {
          xAxixs.push(value[key])
        }

        if (key === 'totalCpuUsage') {
          series[0].data.push(value[key])
        }

        if (key === 'totalMemoryUsage') {
          series[1].data.push(value[key])
        }

        if (key === 'totalProcesses') {
          series[2].data.push(value[key])
        }
      })
    }
  }

  return {
    series,
    xAxixs
  }
}
const useProcessStore = defineStore('process', () => {
  const processes = ref<IProcess[]>([])

  const userOptions = ref<{ label: string; value: string }[]>([])
  const processInfo = ref<Record<string, string>[]>([])
  const webSocketData = ref<IProcess[]>([])

  const options = ref<IOptions>({
    chart: {
      id: 'process-user-chart'
    },
    xaxis: {
      categories: []
    },
    legend: {
      show: false
    },
    dataLabels: {
      enabled: false
    }
    //   plotOptions: {
    //     bar: {
    //       horizontal: false,
    //       columnWidth: '100%',
    //       endingShape: 'rounded'
    //     }
    //   }
  })

  const series = ref<ISeries[]>([])

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
        const _data = handleWebsocketMessage(event.data)
        if (_data.processes.length > 0) {
          processes.value = _data.processes
        }

        if (_data.processes.length > 0) {
          const { series: _series, xAxixs } = createSerisAndOptions(_data.reports)

          options.value = {
            ...options.value,
            xaxis: {
              categories: xAxixs
            }
          }

          series.value = _series
        }
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

  async function fetchProcessUsersReport() {
    try {
      const {
        data: { data }
      } = await api.get(`/reports`)

      const { series: _series, xAxixs } = createSerisAndOptions(data)

      options.value = {
        ...options.value,
        xaxis: {
          categories: xAxixs
        }
      }

      series.value = _series
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
    processInfo,
    fetchProcessUsersReport,
    options,
    series,
    webSocketData
  }
})

export { useProcessStore }
