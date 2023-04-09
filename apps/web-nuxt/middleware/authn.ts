export default defineNuxtRouteMiddleware((to, _from) => {
	const token = useAuthToken()

	const isProtectedRoute = to.path === '/new'
	if (
		!token.value.accessToken &&
		!token.value.refreshToken &&
		isProtectedRoute
	) {
		return navigateTo('/login')
	}

	if (
		token.value.accessToken &&
		token.value.refreshToken &&
		(to.path === '/login' ||
			to.path === '/signup' ||
			to.path === '/forgot-password')
	) {
		return navigateTo('/')
	}
})
