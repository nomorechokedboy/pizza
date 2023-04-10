import { NuxtConfig } from '@nuxt/schema'
import { internalIpV4 } from 'internal-ip'
import defaultNuxtConfig from '../../nuxt.config'

// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig(async () => {
	const host = await internalIpV4()
	const config: NuxtConfig = {
		app: {
			head: {
				meta: [
					{
						name: 'viewport',
						content: 'width=device-width, initial-scale=1'
					},
					{
						name: 'description',
						content: "Accessiblog is a web application that focuses on creating accessible and inclusive blog content. With a range of tools and features, including a rich text editor and image descriptions, Accessiblog makes it easy for bloggers to create content that can be enjoyed by all users. Join our community of bloggers today and start creating content that's accessible to everyone."
					},
					{
						name: 'apple-mobile-web-app-title',
						content: 'Accessiblog'
					}
				],
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
				]
			}
		},
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
		image: {
			dir: 'assets/'
		},
		ssr: process.env.TAURI_ENV === undefined
	}

	return config
})
