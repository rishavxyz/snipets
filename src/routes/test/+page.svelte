<script lang="ts">
	import {
		codeToHtml as shiki,
		type BundledLanguage,
		type BundledTheme
	} from "shiki"

	import { langs } from "$lib/langs.json"
	import { themes } from "$lib/themes.json"
	import { selectText } from "$lib/utils"
	import { once } from "$lib/utils/event-modifiers"

	type Props = {
		defaultCode?: string
		defaultLang?: BundledLanguage | "text"
		defaultTheme?: BundledTheme
		on_code_change?: (value: string) => void
		on_lang_change?: (theme: string) => void
		on_theme_change?: (theme: string) => void
	}

	let {
		defaultCode = "",
		defaultLang = "javascript",
		defaultTheme = "ayu-dark",
		...props
	}: Props = $props()

	let code = $state("")
	let value = $state(defaultCode.trim())
	let selectedLang = $state<BundledLanguage | "text">(defaultLang)
	let selectedTheme = $state<BundledTheme>(defaultTheme)

	const TAB_SPACE = "  "

	function updateText(e: KeyboardEvent) {
		type CursorPosition = number

		const t = e.target as HTMLTextAreaElement
		const pairKeys = ["()", "{}", "[]"]

		const getCursorPos = () => t.selectionStart
		const setCursorPos = (pos: CursorPosition) => t.setSelectionRange(pos, pos)
		const getTextBeforeCursor = (pos: CursorPosition) =>
			t.value.substring(0, pos)
		const getTextAfterCursor = (pos: CursorPosition) => t.value.substring(pos)

		function setTextAfterCursor(text: string): CursorPosition {
			const pos = getCursorPos()
			const textBeforeCursor = t.value.substring(0, pos)
			const textAfterCursor = t.value.substring(pos)

			t.value = textBeforeCursor + text + textAfterCursor
			value = t.value

			return pos + text.length
		}

		switch (e.key) {
			case "Tab":
				{
					e.preventDefault()
					const pos = setTextAfterCursor(TAB_SPACE)
					setCursorPos(pos)
				}
				break

			case "Enter":
				{
					e.preventDefault()

					const space = getTextBeforeCursor(getCursorPos())
						.split(/\n/)
						.slice(-1)[0]
						.match(/^\s+/g)?.[0]

					const pos = setTextAfterCursor("\n" + (space ?? ""))
					setCursorPos(pos)
				}
				break

			case "{":
			case "[":
			case "(":
				{
					e.preventDefault()
					const key = pairKeys.find((key) => key[0] == e.key)
					const pos = setTextAfterCursor(key!)
					setCursorPos(pos - 1)
				}
				break

			case "}":
			case "]":
			case ")":
				{
					const pos = getCursorPos()
					const charBeforeCursor = t.value[pos - 1]
					const charAfterCursor = t.value[pos]

					pairKeys.forEach((key) => {
						if (charBeforeCursor == key[0] && charAfterCursor == key[1]) {
							e.preventDefault()
							setCursorPos(pos + 1)
						}
					})
				}
				break

			case "Backspace":
				{
					pairKeys.forEach((key) => {
						const pos = getCursorPos()
						const charBeforeCursor = t.value[pos - 1]
						const charAfterCursor = t.value[pos]

						if (charBeforeCursor == key[0] && charAfterCursor == key[1]) {
							e.preventDefault()
							t.value =
								getTextBeforeCursor(pos - 1) + getTextAfterCursor(pos + 1)
							value = t.value

							setCursorPos(pos - 1)
						}
					})
				}
				break
		}
	}

	function setCaretColor() {
		const pre = document.querySelector("pre.shiki") as HTMLPreElement | null
		const textarea = document.querySelector(
			".editor > textarea"
		) as HTMLTextAreaElement

		if (!pre) return

		const { color } = getComputedStyle(pre)

		textarea.style.caretColor = color
	}

	function matchScrollPos(e: Event) {
		const textarea = e.target as HTMLTextAreaElement
		const code = document.querySelector("pre.shiki") as HTMLDivElement
		const scrollPos = textarea.scrollTop
		console.log({ scrollPos })

		code.style.setProperty("--posY", -scrollPos + "px")
	}

	async function codeToHtml() {
		code = await shiki(value, {
			lang: selectedLang,
			theme: selectedTheme
		})
	}

	$effect(() => {
		codeToHtml()
		if (typeof props.on_code_change != "undefined") props.on_code_change(value)

		if (typeof props.on_lang_change != "undefined")
			props.on_lang_change(selectedLang)

		if (typeof props.on_theme_change != "undefined")
			props.on_theme_change(selectedTheme)
	})
</script>

{#snippet options(arr)}
	{#each arr as obj}
		<option value={obj.id}>{obj.name}</option>
	{/each}
{/snippet}

{#snippet themeOptions(optionLabel, themes)}
	<optgroup label={optionLabel}>
		{@render options(themes)}
	</optgroup>
{/snippet}

<div class="grid">
	<label>
		Language
		<select name="langSelect" bind:value={selectedLang}>
			<option value="text" selected>Plain text</option>
			{@render options(langs)}
		</select>
	</label>

	<label>
		Theme
		<select
			name="themeSelect"
			bind:value={selectedTheme}
			onchange={setCaretColor}
		>
			{@render themeOptions("Dark themes", themes.dark)}
			{@render themeOptions("Light themes", themes.light)}
		</select>
	</label>
</div>

<div class="editor">
	<textarea
		spellcheck="false"
		autocomplete="off"
		autocapitalize="off"
		autocorrect="off"
		onkeydown={updateText}
		onfocus={once(selectText)}
		onscroll={matchScrollPos}
		bind:value
	></textarea>

	{#if code}{@html code}{/if}
</div>

<style>
	.editor {
		overflow: hidden;
		position: relative;
		height: 360px;

		& > textarea {
			background-color: transparent;
			caret-color: white;
			color: transparent;
			z-index: 1;
			resize: none;
		}
	}

	textarea,
	:global(pre.shiki) {
		position: absolute;
		inset: 0;
		margin: 0 !important;
		padding: 0.785rem !important;
	}

	textarea,
	:global(pre.shiki),
	:global(pre.shiki code) {
		font-family: monospace !important;
		font-weight: 400 !important;
		font-size: var(--text-sm) !important;
		line-height: 1.25 !important;
		white-space: pre !important;
	}

	:global(pre.shiki) {
		--posY: 0;
		transform: translate3d(0, var(--posY), 0);
	}

	:global(pre.shiki code) {
		padding: 0;
		margin: 0;
	}
</style>
