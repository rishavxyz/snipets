export default function selectText(event: Event, start = 0, end?: number) {
    const target = event.target as HTMLInputElement
    end = end || target.value.length

    target.setSelectionRange(start, end)
}