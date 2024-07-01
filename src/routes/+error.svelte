<script lang="ts">
	import { page } from "$app/stores"
</script>

<article>
	<figure>
		<img src="//http.cat/{$page.status}" alt="Error cat" />
	</figure>

	<footer>
		<hgroup>
			<h3>Oh no, that's a {$page.status}</h3>
			<h6><small>{$page.error?.message}</small></h6>
		</hgroup>

		<div class="grid">
			<a href="/" role="button" class="contrast">Take me to home</a>
			<a href="/playground" role="button" class="secondary">Playground</a>
		</div>
	</footer>
</article>

<style>
	article {
		padding: 0;
		overflow: clip;

		& > footer {
			margin: 0;
		}
	}

	figure {
		display: grid;
		place-content: center;
		& > img {
			width: 100%;
			max-width: 640px;
		}
	}
</style>
