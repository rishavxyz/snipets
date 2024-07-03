<script lang="ts">
	import Seo from "$lib/components/seo.svelte"
	import { fade, fly, slide } from "svelte/transition"
	import { generateUsername } from "unique-username-generator"

	let username = $state("")
	let value = $state("")
	let tid: NodeJS.Timeout

	function genUsername(e: Event) {
		clearTimeout(tid)

		tid = setTimeout(() => {
			const t = e.target as HTMLInputElement

			if (t.value.length > 2) {
				const name = t.value.trim().split(" ")[0]?.toLowerCase()
				const n = generateUsername("", 0, 5)
				username = n + "_" + name + Math.floor(Math.random() * 999)
			}
		}, 600)
	}

	function setUsername(e: Event) {
		e.preventDefault()

		value = username
		username = ""
	}
</script>

<article>
	<hgroup>
		<h2>Sign up to Snipets</h2>
		<h6>Already have an account? <a href="/signin">Sign in.</a></h6>
	</hgroup>

	<form action="?/signup" method="post">
		<fieldset>
			<label>
				Display name <small class="tip">(Optional)</small>
				<input
					type="text"
					name="name"
					autocomplete="off"
					oninput={genUsername}
				/>
			</label>
			<label>
				Username <small class="tip">(Space not allowed)</small>
				<input type="text" name="username" autocomplete="username" bind:value />
				{#if username}
					<small transition:slide>
						How about
						{#key username}
							<button class="act" in:fly|local={{ y: 8 }} onclick={setUsername}
								>{username}</button
							>? Click to set.
						{/key}
					</small>
				{/if}
			</label>
		</fieldset>

		<small
			>If Display name is left blank, your username will be used as display
			name.</small
		>

		<input type="submit" value="Sign up" />
	</form>
</article>

<Seo title="Sign up" />

<style>
	hgroup h6 {
		margin-top: 10px;
	}

	button.act {
		all: unset;
		color: var(--pico-primary);
		cursor: pointer;

		&:focus {
			border: none;
		}
	}
</style>
