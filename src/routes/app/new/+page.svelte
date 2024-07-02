<script lang="ts">
	import Editor from "$lib/components/editor.svelte"
	import Seo from "$lib/components/seo.svelte"
	import { selectText } from "$lib/utils"
	import { once } from "$lib/utils/event-modifiers"

	type Props = {
		form: import("./$types").ActionData
	}

	let title = $state("")
	let code = $state("")
	let lang = $state("")
	let theme = $state("")

	let { form }: Props = $props()
	const err = form?.error as { code: string[]; title: string[] } | undefined
</script>

<article>
	<h2>Share your snipet</h2>

	<label>
		<Editor
			defaultCode={form?.data.code ? form.data.code.toString() : ""}
			on_code_change={(value) => (code = value)}
			on_lang_change={(value) => (lang = value)}
			on_theme_change={(value) => (theme = value)}
		/>
		{#if err}
			<small class="error">{err.code}</small>
		{/if}
	</label>

	<form method="post" action="?/new">
		<fieldset>
			<label>
				Title
				<input
					type="text"
					name="title"
					value={err ? form?.data.title : title}
					onfocus={once(selectText)}
					placeholder="My awesome gleam"
				/>
				{#if err}
					<small class="error">{err.title}</small>
				{/if}
			</label>

			<details>
				<!-- svelte-ignore a11y_no_redundant_roles -->
				<summary role="button" class="outline contrast"
					>Add a description?</summary
				>
				<label>
					Description
					<span style="color:var(--pico-secondary)">
						<small>&lpar;Optional&rpar;</small>
					</span>
					<textarea name="desc" rows="8" style="resize:none"></textarea>
				</label>
			</details>
		</fieldset>

		<div class="hidden">
			<textarea name="code" hidden>{code}</textarea>
			<input name="lang" value={lang} hidden />
			<input name="theme" value={theme} hidden />
		</div>

		<div class="grid">
			<input type="submit" value="Share" />
			<input type="reset" value="Cancel" class="outline" />
		</div>
	</form>
</article>

<Seo title="Create a new Snipet" />

<style>
	small.error {
		color: salmon;
	}
</style>
