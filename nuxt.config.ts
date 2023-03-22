import type { NuxtConfig } from 'nuxt/config'

const defaultNuxtConfig: NuxtConfig = {
	modules: [
		// '@nuxt/typescript-build',
		'@nuxtjs/tailwindcss',
		'nuxt-vitest',
		'@nuxt/image-edge'
		// '@nuxtjs/i18n',
		// 'nuxt-svgo'
		// '@nuxtjs/google-fonts',
		// '@nuxtjs/web-vitals',
	]
}

export default defaultNuxtConfig
