const debounce = (fn: any, delay: number) => {
  let timeout: any

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
