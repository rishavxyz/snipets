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
		defaultLang = "gleam",
		defaultTheme = "poimandres",
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
		const textarea = document.querySelector("textarea")!

		if (!pre) return

		const { color } = getComputedStyle(pre)

		textarea.style.caretColor = color
	}

	function matchScrollPos(e: Event) {
		const textarea = e.target as HTMLTextAreaElement
		const code = document.querySelector(".code") as HTMLDivElement
		const scrollPos = textarea.scrollTop

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
			onload={setCaretColor}
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
		onscroll={matchScrollPos}
		onfocus={once(selectText)}
		bind:value
	></textarea>

	<div class="code">
		{#if code}
			{@html code}
		{/if}
	</div>
</div>

<style>
	.editor {
		overflow: hidden;
		position: relative;
		height: 250px;
	}

	textarea,
	.code {
		position: absolute;
		inset: 0;
		margin: 0 !important;
		display: block !important;
		height: 100% !important;
	}
	textarea {
		background-color: transparent;
		color: transparent;
		caret-color: white;
		z-index: 2;
		isolation: isolate;
		height: 100px;
		resize: none !important;
	}
	.code {
		--posY: 0px;
		left: 2px;
		transform: translate3d(0, var(--posY), 0);
	}
</style>
