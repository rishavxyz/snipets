export default function preventDefault(fn: Function) {
    return (event: Event) => fn.call(fn, event)
}