import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';
import { browserslistToTargets } from 'lightningcss';
import browserlist from 'browserslist';

export default defineConfig({
	css: {
		transformer: 'lightningcss',
		lightningcss: {
			targets: browserslistToTargets(browserlist('>= 0.25%'))
		}
	},
	build: {
		cssMinify: 'lightningcss'
	},
	plugins: [sveltekit()]
});
