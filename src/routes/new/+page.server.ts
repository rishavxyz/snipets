import escape from "escape-html"
import { error, type Actions } from "@sveltejs/kit"

export const actions: Actions = {
	new: async ({ request, fetch }) => {
		const fd = await request.formData()
		const inputs = [
			{
				name: "title",
				value: fd.get("title")
			},
			{
				name: "desc",
				value: fd.get("desc")
			},
			{
				name: "lang",
				value: fd.get("lang")
			},
			{
				name: "theme",
				value: fd.get("theme")
			},
			{
				name: "code",
				value: fd.get("code")
			}
		]

		inputs.forEach((entry) => {
			if (entry.value == null) return
			fd.set(entry.name, escape(entry.value.toString()))
		})

		try {
			const req = await fetch("http://localhost:8000/new", {
				method: "post",
				body: fd
			})
			const res = await req.json()

			console.dir(res, { depth: 10 })
		} catch (_) {
			throw error(500, "OOPs")
		}
	}
}
