import { dev } from "$app/environment";
import { env } from "$env/dynamic/private";

export const url = (!dev ? env.URL! : "http://localhost:8000") + "/snipets"