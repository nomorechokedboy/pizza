import defaultNuxtConfig from '../../nuxt.config'

// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
	...defaultNuxtConfig,
	image: {
		cloudinary: {
			baseURL: 'https://res.cloudinary.com/nuxt/image/upload/'
		}
	}
})
