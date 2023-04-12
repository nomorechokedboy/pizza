type UserProfile = {
	id?: number
	name?: string
	avatar?: string
	email?: string
	username?: string
}

export function useUserProfile() {
	return useState<UserProfile>('userProfile', () => ({
		email: undefined,
		avatar: undefined,
		name: undefined,
		id: undefined,
		username: undefined
	}))
}

export function setUserProfile(profile: UserProfile) {
	const userProfile = useUserProfile()
	userProfile.value = profile
}
