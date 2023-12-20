const debounce = (fn: any, delay: number) => {
  let timeout: NodeJS.Timeout

  return (...args: any[]) => {
    if (timeout) {
      clearTimeout(timeout)
    }

    timeout = setTimeout(() => {
      fn(...args)
    }, delay)
  }
}

export { debounce }
