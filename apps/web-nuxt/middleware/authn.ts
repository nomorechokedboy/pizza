export default defineNuxtRouteMiddleware(() => {
	const isAuthenticated = useIsAuthenticated()

	if (!isAuthenticated.value) {
		return navigateTo('/login')
	}
})
