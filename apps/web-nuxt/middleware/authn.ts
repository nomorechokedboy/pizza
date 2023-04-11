export default defineNuxtRouteMiddleware((to, _from) => {
	const token = useAuthToken()

	// if((!token.value.accessToken  || !token.value.refreshToken ) && to.path === '')
	if (
		token.value.accessToken &&
		token.value.refreshToken &&
		(to.path === '/login' || to.path === '/signup')
	) {
		return navigateTo('/')
	}
})
