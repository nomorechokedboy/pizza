import { defineVitestConfig } from 'nuxt-vitest/config'

export default defineVitestConfig({
    environment: 'nuxt',
    test: {
        globals: true,
        exclude: [
            'node_modules',
            '.git',
            'dist',
            '.idea',
            '.cache',
            'tests',
            'tests-examples'
        ]
    }
})
