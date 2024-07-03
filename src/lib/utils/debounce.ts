export default function debounce(fn: (args: any) => void, delay = 1000) {
    let timeoutId: NodeJS.Timeout

    return (...args: any) => {
        if (timeoutId) return clearTimeout(timeoutId)
        timeoutId = setTimeout(() => fn(args), delay)
    }
}
