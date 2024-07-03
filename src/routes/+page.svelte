<script>
	import Seo from "$lib/components/seo.svelte"
</script>

<section class="hero">
	<div>
		<h1>
			A place for developers
			<br />
			what you can call it a home
		</h1>
		<p>
			Snipets is a code <i>snippet</i> sharing platform where you can share your
			code snippet like a boss.
		</p>
	</div>

	<div class="grid cta">
		<a href="/app/new" role="button">Share your snipet</a>
		<a href="/app" role="button" class="contrast">Explore snipets</a>
	</div>
</section>

<Seo title="Share code snippets" />

<style>
	.hero {
		margin-top: 3rem;
		margin-inline: auto;
		text-align: center;
		max-width: 60ch;

		& h1 {
			text-wrap: balance;
		}

		.cta {
			max-width: 69%;
			margin-inline: auto;
		}
	}

	@media (width <= 768px) {
		.hero {
			margin-top: auto;
			max-width: inherit;

			.cta {
				max-width: inherit;
			}
		}
	}
</style>
