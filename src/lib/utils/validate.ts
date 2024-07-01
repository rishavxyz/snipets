import { z } from "zod";

export default function validate(data: unknown, schema: z.ZodRawShape) {
    const res = z.object(schema).safeParse(data)

    if (res.error)
        return res.error.flatten().fieldErrors

    return false
}