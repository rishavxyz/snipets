import { error } from "@sveltejs/kit"
import { url } from "$lib/config"
import type { PageServerLoad } from "./$types"
import type { Response, SnipetResponse } from "$lib/types"

export const load: PageServerLoad = async ({ fetch }) => {
    try {
        const req = await fetch(url)
        const res = await req.json() as Response<SnipetResponse>

        return { ...res }

    } catch (_) {
        throw error(500, "Cannot reach out to the sever at the moment")
    }
}
