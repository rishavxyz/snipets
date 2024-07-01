import escape from "escape-html"
import validate from "$lib/utils/validate"

import { Redirect_1, error, fail, redirect, type Actions } from "@sveltejs/kit"
import { z } from "zod"

const schema = {
	title: z.string().min(2,
		{
			message: "Must be 2 or more letters"
		}
	).max(60,
		{
			message: "Must be less than 60 letters"
		}
	),

	code: z.string().min(5,
		{
			message: "Snipet must be 5 or more characters long"
		}
	).max(5000,
		{
			message: "Must be less than 5000 characters"
		}
	),

	desc: z.string().nullable(),
}

type Response = {
	status: number,
	error: string | null,
	data: {
		title?: string
		code?: string
	}
}

function response(props: Response) {
	return { ...props }
}

export const actions: Actions = {
	new: async ({ request, fetch }) => {
		const fd = await request.formData()

		const dirtyFields = {
			title: fd.get("title"),
			desc: fd.get("desc"),
			code: fd.get("code")
		}

		const err = validate(dirtyFields, schema)

		if (err) return fail(400, response(
			{
				status: 400,
				error: "input validation error",
				data: err
			}
		))

		type Field = "title" | "desc" | "code"

		for (const key in dirtyFields) {
			const cleanStr = escape(
				dirtyFields[key as Field]?.toString()
			)
			fd.set(key, cleanStr)
		}

		const req = await fetch("http://localhost:8000/new", {
			method: "post",
			body: fd
		})
		const res = await req.json() as Response

		if (res.status != 200)
			return fail(res.status, res)

		throw redirect(302, "http://localhost:5173")
	}
}
