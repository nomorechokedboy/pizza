export default defineNuxtRouteMiddleware(() => {
	const token = useAuthToken()
	const refreshToken = useRefreshToken()
	const isAuthenticated = !!token.value && !!refreshToken.value
	if (isAuthenticated) {
		return navigateTo('/')
	}
})
