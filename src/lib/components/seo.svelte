<script lang="ts">
	import { page } from "$app/stores"

	type Seo = {
		title?: string
		description?: string
		author?: string
		url?: string
		keywords?: string
		meta?: [name: string, content: string][]
	}

	type Props = {} & Seo

	let { ...seo }: Props = $props()

	const brand = "Snipets"
	const _title_ = [seo.title, "•", brand]
	const title = seo.title
		? $page.url.pathname == "/"
			? _title_.reverse().join(" ")
			: _title_.join(" ")
		: brand
	const desc = seo.description || "Share code snippets and be a chad programmer"
	const url = seo.url || $page.url.href
</script>

<svelte:head>
	<title>{title}</title>
	<link rel="canonical" href={url} />
	<link rel="icon" href="//fav.farm/🐶" />

	<meta name="description" content={desc} />
	<meta name="keywords" content={seo.keywords} />
	<meta name="author" content="Rishav Mandal" />

	<meta property="og:type" content="website" />
	<meta property="og:url" content={url} />
	<meta property="og:title" content={title} />
	<meta property="og:description" content={desc} />
	<meta property="og:locale" content="en_US" />

	<meta property="twitter:card" content="summary_large_image" />
	<meta property="twitter:url" content={url} />
	<meta property="twitter:title" content={title} />
	<meta property="twitter:description" content={desc} />

	{#if seo.meta}
		{#each seo.meta as [name, content]}
			{#if name.startsWith("prop ")}
				<meta property={name.slice(5)} {content} />
			{:else}
				<meta {name} {content} />
			{/if}
		{/each}
	{/if}
</svelte:head>
