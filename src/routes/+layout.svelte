<script lang="ts">
	import "../pico.css"
	import "$lib/css/font.css"
	import "$lib/css/main.css"

	import type { Snippet } from "svelte"

	type Props = {
		children: Snippet
	}

	let { children }: Props = $props()
</script>

<svelte:head>
	<link
		rel="preload"
		href="/Lotion-Regular.ttf"
		as="font"
		type="font/ttf"
		crossorigin=""
	/>
	<link
		rel="preload"
		href="/Lotion-RegularItalic.ttf"
		as="font"
		type="font/ttf"
		crossorigin=""
	/>
	<link
		rel="preload"
		href="/Lotion-SemiBold.ttf"
		as="font"
		type="font/ttf"
		crossorigin=""
	/>
</svelte:head>

<header>
	<nav class="container">
		<ul>
			<li><h2><a href="/" class="contrast">Snipets 🐶</a></h2></li>
		</ul>
		<ul>
			<li><a href="/app/new" class="secondary">New</a></li>
		</ul>
	</nav>
</header>

<main class="container">
	{@render children()}
</main>

<style>
	nav {
		align-items: baseline;

		& li > h2 a:hover {
			text-decoration: none;
		}
	}
</style>
