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
