import { NuxtConfig } from '@nuxt/schema'
import { internalIpV4 } from 'internal-ip'
import defaultNuxtConfig from '../../nuxt.config'

// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig(async () => {
	const host = await internalIpV4()
	const config: NuxtConfig = {
		app: {
			head: {
				link: [
					{
						rel: 'icon',
						type: 'image/png',
						sizes: '32x32',
						href: '/favicon-32x32.png'
					},
					{
						rel: 'icon',
						type: 'image/png',
						sizes: '16x16',
						href: '/favicon-16x16.png'
					},
					{
						rel: 'apple-touch-icon',
						sizes: '180x180',
						href: '/apple-touch-icon.png'
					}
				],
				noscript: [
					{ children: 'JavaScript is required' }
				],
				htmlAttrs: { lang: 'en' }
			}
		},
		modules: [
			...defaultNuxtConfig.modules,
			['unplugin-icons/nuxt', { scale: 1.5 }],
			'v-satori/nuxt',
			'unplugin-font-to-buffer/nuxt',
			'@vueuse/nuxt',
			'nuxt-simple-sitemap'
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
		image: {
			dir: 'assets/'
		},
		ssr: process.env.TAURI_ENV === undefined,
		plugins: [{ src: '~/plugins/vercel.ts', mode: 'client' }],
		webVitals: {
			debug: false
		},
		sitemap: {
			siteUrl: 'https://pizza-web-nuxt.vercel.app'
		},
		runtimeConfig: {
			public: {
				tokenExpTime: 29 * 60 * 60 * 24,
				apiUrl: 'https://api-blog-dev-nomorechokedboy.cloud.okteto.net',
				mediaUrl: 'https://api-blog-dev-nomorechokedboy.cloud.okteto.net/api/v1/media/',
				dicebearMedia:
					'https://api.dicebear.com/6.x/initials/svg?seed=',
				notificationUrl: 'http://localhost:5000',
				notifyUrl: 'http://localhost:5000/api/v1/notify'
			}
		}
	}

	return config
})
