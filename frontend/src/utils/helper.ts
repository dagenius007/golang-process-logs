export const isHttps = (): boolean => {
  return window?.location.protocol === 'https:'
}

export const stripHttp = (value: string): string => {
  return value.replace(/^https?:\/\//, '')
}
