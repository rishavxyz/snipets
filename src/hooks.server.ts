import { error, redirect, type Handle } from "@sveltejs/kit"
import { sequence } from "@sveltejs/kit/hooks"

const hooks: Handle[] = [
    // auth user
    async ({ event, resolve }) => {
        if (event.url.pathname.startsWith("/app")) {
            const userid = event.cookies.get("userid")

            if (!userid) {
                event.locals.userid = null
                return redirect(302, "/signup")
            }
            else event.locals.user = JSON.parse(userid)
        }

        return await resolve(event)
    },

    // setup user
    // async ({ event, resolve }) => {
    //     // todo
    //     console.log({ path: event.url.pathname })

    //     return resolve(event)
    // },

    // playground is in progress
    async ({ event, resolve }) => {
        if (event.url.pathname.includes("/playground")) {
            return error(501, "Still working on this")
        }

        return resolve(event)
    }
]

export const handle: Handle = sequence(...hooks)
