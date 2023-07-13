import { useQuery } from '@tanstack/vue-query'
import { EntitiesUserResponse } from '~/codegen/api'

type UserProfile = {
	id?: number
	name?: string
	avatar?: string
	email?: string
	username?: string
	phone?: string
}

function selectUserProfile({ fullname, ...profile }: EntitiesUserResponse) {
	const userProfile: UserProfile = {
		name: fullname,
		...profile
	}
	return userProfile
}

async function getUserProfile() {
	const { $blogApi } = useNuxtApp()
	return $blogApi.auth.authMeGet().then((res) => res.data)
}

export function useUserProfile() {
	const isLoggedIn = useIsAuthenticated()
	const enabled = computed(() => !!isLoggedIn.value)

	return useQuery({
		queryFn: getUserProfile,
		queryKey: ['UserProfile'],
		select: selectUserProfile,
		enabled
	})
}

export function useOtherProfile(id: string) {
	async function getOtherProfile() {
		const { $blogApi } = useNuxtApp()
		return $blogApi.user.userIdGet(id).then((res) => res.data)
	}

	return useQuery({
		queryFn: getOtherProfile,
		queryKey: ['Profile', id]
	})
}
