import {
    codeToHtml as shiki,
    type BundledLanguage,
    type BundledTheme
} from "shiki"

export type ShikiParams = {
    code: string
    lang: BundledLanguage | "text"
    theme: BundledTheme
}

export async function codeToHtml(
    { code, lang, theme }: ShikiParams
) {
    return await shiki(code, { lang, theme })
}