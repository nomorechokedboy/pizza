export default defineNuxtRouteMiddleware(() => {
	const token = useAuthToken()
	const isAuthenticated =
		token.value.accessToken && token.value.refreshToken

	if (!isAuthenticated) {
		return navigateTo('/login')
	}
})
