import { error } from "@sveltejs/kit"
import type { PageServerLoad } from "./$types"
import type { Snipet } from "$lib/types"

export const load: PageServerLoad = async ({ fetch }) => {
	try {
		const req = await fetch("http://localhost:8000")
		const res = await req.json()

		return { res } as {
			res: {
				data: {
					snipets: Snipet[]
					length: number
					capacity: number
				}
			}
		}
	} catch (_) {
		throw error(500)
	}
}
