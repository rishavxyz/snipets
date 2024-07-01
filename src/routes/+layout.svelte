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

<header>
	<nav class="container">
		<ul>
			<li><h2><a href="/" class="contrast">Snipets 🐶</a></h2></li>
		</ul>
		<ul>
			<li><a href="/" class="contrast">Profile</a></li>
			<li><a href="/new" class="secondary">New</a></li>
		</ul>
	</nav>
</header>

<main class="container">
	{@render children()}
</main>

<style>
	nav {
		align-items: baseline;
	}
</style>
