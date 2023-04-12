import { EntitiesPostResponse } from '~/codegen/api'
import { baseURL } from '~/constants'

export function getPostDetails(slug: string): Promise<EntitiesPostResponse> {
	return $fetch(`${baseURL}/api/v1/posts/${slug}`)
}

export function usePostDetails(slug: string) {
	function fetchPostDetails() {
		return getPostDetails(slug)
	}

	return useAsyncData<EntitiesPostResponse>(
		`${slug}-details`,
		fetchPostDetails
	)
}
