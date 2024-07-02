<script lang="ts">
	import { codeToHtml } from "$lib/utils/codeToHtml"

	type Props = {
		data: import("./$types").PageServerData
	}

	let { data }: Props = $props()

	const { snipets } = data.data
</script>

<section>
	<h1>Share code, Share love</h1>
</section>

<section>
	<p>Popular snipets</p>
	{#each snipets as snipet}
		{@const param = {
			code: snipet.code,
			lang: snipet.lang,
			theme: snipet.theme
		}}
		<article>
			<hgroup>
				<h5>{snipet.title}</h5>
				<h6><small>{snipet.desc}</small></h6>
			</hgroup>

			{#await codeToHtml(param)}
				<pre><code id="code">Getting ready</code></pre>
			{:then code}
				<div class="snipet">
					<code>{snipet.lang}</code>
					{@html code}
				</div>
			{/await}
		</article>
	{/each}
</section>

<style>
	code#code {
		background-color: var(--pico-background-color);
	}

	.snipet {
		position: relative;
		& > code {
			position: absolute;
			z-index: 999;
			isolation: isolate;
			top: 0.25rem;
			right: 0.2rem;
			font-size: var(--text-xs);
		}
	}
</style>
