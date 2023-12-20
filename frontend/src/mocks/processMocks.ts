const processMock: Record<string, any> = {
  total: 11,
  limit: 10,
  page: 1,
  data: [
    {
      id: 1,
      user: 'joshuaoluikpe',
      pid: 1,
      cpuUsage: 0.5,
      memoryUsage: 0.6,
      residentMemorySize: 53434,
      virtualMemorySize: 56,
      state: 'running',
      totalTime: '20.6',
      cpuTime: '670.9',
      command: 'go',
      priority: 'medium',
      createdAt: new Date('2023-09-14T12:00:00.000Z'),
      updatedAt: new Date('2023-09-14T12:00:00.000Z')
    },
    {
      id: 2,
      user: 'joshuaoluikpe',
      pid: 2,
      cpuUsage: 0.7,
      memoryUsage: 0.6,
      residentMemorySize: 53434,
      virtualMemorySize: 56,
      state: 'sleeping',
      totalTime: '20.6',
      cpuTime: '670.9',
      command: 'go',
      priority: 'medium',
      createdAt: new Date('2023-09-14T12:00:00.000Z'),
      updatedAt: new Date('2023-09-14T12:00:00.000Z')
    },
    {
      id: 3,
      user: 'joshuaoluikpe',
      pid: 3,
      cpuUsage: 0.5,
      memoryUsage: 0.6,
      residentMemorySize: 53434,
      virtualMemorySize: 56,
      state: 'running',
      totalTime: '20.6',
      cpuTime: '670.9',
      command: 'go',
      priority: 'medium',
      createdAt: new Date('2023-09-14T12:00:00.000Z'),
      updatedAt: new Date('2023-09-14T12:00:00.000Z')
    }
  ]
}

export { processMock }
