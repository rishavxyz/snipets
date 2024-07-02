// See https://kit.svelte.dev/docs/types#app

import type { UUID } from "crypto";

// for information about these interfaces
declare global {
	namespace App {
		// interface Error {}
		interface Locals {
			userid: null | UUID
			user: {
				id: UUID
				username: string
				name: string
			}
		}
		// interface PageData {}
		// interface PageState {}
		// interface Platform {}
		interface ImportMetaEnv {
			readonly _URL: string
		}
		interface ImportMeta {
			readonly env: ImportMetaEnv
		}
	}
}

export { };
