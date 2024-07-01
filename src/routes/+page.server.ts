import { error } from "@sveltejs/kit"
import type { PageServerLoad } from "./$types"
import type { Snipet, SnipetResponse } from "$lib/types"
import type { Response } from "$lib/types"

export const load: PageServerLoad = async ({ fetch }) => {
    try {
        const req = await fetch("http://localhost:8000")
        const res = await req.json() as Response<SnipetResponse>

        return { ...res }

    } catch (_) {
        throw error(500, "Cannot reach out to the sever at the moment")
    }
}
