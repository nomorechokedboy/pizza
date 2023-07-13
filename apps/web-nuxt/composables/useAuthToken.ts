import { AxiosResponse } from 'axios'
import { EntitiesAuth } from '~/codegen/api'

export function useAuthToken() {
	const config = useRuntimeConfig()
	return useCookie<string | undefined>('token', {
		default: undefined,
		sameSite: 'strict',
		secure: true,
		maxAge: config.public.tokenExpTime
	})
}

export function useRefreshToken() {
	const config = useRuntimeConfig()
	return useCookie<string | undefined>('refreshToken', {
		default: undefined,
		sameSite: 'strict',
		secure: true,
		maxAge: config.public.tokenExpTime
	})
}

export function onRefreshToken(resp: AxiosResponse<EntitiesAuth, any>) {
	const token = useAuthToken()
	const refreshToken = useRefreshToken()

	refreshToken.value = resp.data.refresh_token
	token.value = resp.data.token
}

export function removeToken() {
	const token = useAuthToken()
	const refreshToken = useRefreshToken()

	refreshToken.value = undefined
	token.value = undefined
}

export function onRefreshTokenError() {
	removeToken()
	cleanupNotificationEventSource()
	notify({
		content: 'Session expired!',
		k: crypto.randomUUID(),
		type: 'error'
	})
}

export function useIsAuthenticated() {
	const token = useAuthToken()
	const refreshToken = useRefreshToken()

	function computeIsAuthenticated() {
		return token.value && refreshToken.value
	}

	return computed(computeIsAuthenticated)
}
