export default function once(fn: Function | null) {
    return (event: Event) => {
        if (fn) fn.call(fn, event)
        fn = null
    }
}