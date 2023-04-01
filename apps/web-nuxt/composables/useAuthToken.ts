export type AuthToken = {
	accessToken?: string
	refreshToken?: string
}

export const useAuthToken = () =>
	useCookie<AuthToken>('token', {
		default: () => ({
			accessToken: undefined,
			refreshToken: undefined
		}),
		sameSite: 'strict',
		maxAge: 60 * 60
	})
