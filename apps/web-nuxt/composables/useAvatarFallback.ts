import { dicebearMedia } from '~~/constants'
import { useUserProfile } from './useUserProfile'

export function useAvatarFallback(fallbackName = 'A6blog') {
	const { data: userProfile } = useUserProfile()

	return computed(
		() =>
			userProfile.value?.avatar ||
			`${dicebearMedia}${
				userProfile.value?.name || fallbackName
			}`
	)
}
