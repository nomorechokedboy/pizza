export default defineNuxtRouteMiddleware((to, _from) => {
	const token = to.query.token
	if (!token || token === '') {
		return navigateTo('/')
	}
})
