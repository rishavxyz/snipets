import type { UUID } from "crypto"
import type { BundledLanguage, BundledTheme } from "shiki"

export interface Snipet {
    id: UUID
    code: string
    title: string
    desc: string
    lang: BundledLanguage | "text"
    theme: BundledTheme
}

export interface Response<T> {
    status: number,
    error: string | null,
    data: T,
}

export interface SnipetResponse {
    snipets: Snipet[],
    length: number,
    capacity: number
}
