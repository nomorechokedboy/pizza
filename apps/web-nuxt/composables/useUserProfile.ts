import { useQuery } from '@tanstack/vue-query'
import { EntitiesUserResponse } from '~/codegen/api'

type UserProfile = {
	id?: number
	name?: string
	avatar?: string
	email?: string
	username?: string
}

function selectUserProfile({
	fullname,
	phone: _phone,
	...profile
}: EntitiesUserResponse) {
	const userProfile: UserProfile = {
		name: fullname,
		...profile
	}
	return userProfile
}

export function useUserProfile() {
	const { $blogApi } = useNuxtApp()
	const isLoggedIn = useIsAuthenticated()
	const enabled = computed(() => !!isLoggedIn.value)
	async function getUserProfile() {
		return $blogApi.auth.authMeGet().then((res) => res.data)
	}

	return useQuery({
		queryFn: getUserProfile,
		queryKey: ['UserProfile'],
		select: selectUserProfile,
		enabled
	})
}
