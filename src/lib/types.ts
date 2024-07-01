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