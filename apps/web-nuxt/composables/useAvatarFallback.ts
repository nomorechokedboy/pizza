import { useUserProfile } from './useUserProfile'

export function useAvatarFallback(fallbackName = 'A6blog') {
	const appConfig = useRuntimeConfig()
	const { data: userProfile } = useUserProfile()

	return computed(
		() =>
			userProfile.value?.avatar ||
			`${appConfig.public.dicebearMedia}${
				userProfile.value?.name || fallbackName
			}`
	)
}
