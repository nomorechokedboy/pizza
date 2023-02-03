import { defineVitestConfig } from 'nuxt-vitest/config'

export default defineVitestConfig({
	environment: 'nuxt',
	test: {
		globals: true
	}
})
