import { NuxtConfig } from '@nuxt/schema'
import { internalIpV4 } from 'internal-ip'
import defaultNuxtConfig from '../../nuxt.config'

// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig(async () => {
	const host = await internalIpV4()
	const config: NuxtConfig = {
		modules: [
			...defaultNuxtConfig.modules,
			['unplugin-icons/nuxt', { scale: 1.5 }]
		],
		css: ['ui-vue/dist/style.css'],
		vite: {
			clearScreen: false,
			envPrefix: ['VITE_', 'TAURI_'],
			build: {
				target: 'es2022',
				// don't minify for debug builds
				minify: !process.env.TAURI_DEBUG
					? 'esbuild'
					: false,
				// produce sourcemaps for debug builds
				sourcemap: !!process.env.TAURI_DEBUG
			},
			server: {
				strictPort: true,
				hmr: {
					protocol: 'ws',
					host,
					port: 5183
				}
			}
		},
		image: {}
	}

	return config
})
