import { inject } from '@vercel/analytics'

export default defineNuxtPlugin(vercel)

function vercel() {
	inject()
}
