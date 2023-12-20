const isHttps = (): boolean => {
  return window?.location.protocol === 'https:'
}

export { isHttps }
