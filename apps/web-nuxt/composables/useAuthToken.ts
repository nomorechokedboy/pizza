export type AuthToken = {
	accessToken?: string
	refreshToken?: string
}

export function useAuthToken() {
	return useCookie<AuthToken>('token', {
		default: () => ({
			accessToken: undefined,
			refreshToken: undefined
		}),
		sameSite: 'strict',
		maxAge: 60 * 60
	})
}

export function removeToken() {
	const token = useAuthToken()
	token.value.refreshToken = undefined
	token.value.accessToken = undefined
}
