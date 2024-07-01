import type { Response } from "$lib/types"

export function response<T>(props: Response<T>) {
    return { ...props }
}
