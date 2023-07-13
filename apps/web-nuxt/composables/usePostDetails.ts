import { QueryFunctionContext, useQuery } from '@tanstack/vue-query'
import { EntitiesPostResponse } from '~/codegen/api'

export async function getPostDetails({
	queryKey
}: QueryFunctionContext<string[], any>) {
	const { $blogApi } = useNuxtApp()
	return $blogApi.post.postsSlugGet(queryKey[1]).then((resp) => resp.data)
}

export function usePostDetails(slug: string) {
	return useQuery({
		queryFn: getPostDetails,
		queryKey: ['postDetails', slug]
	})
}

export function getPostDetailsSSR(slug: string): Promise<EntitiesPostResponse> {
	const appConfig = useRuntimeConfig()
	return $fetch(`${appConfig.public.apiUrl}/api/v1/posts/${slug}`)
}

export function usePostDetailsSSR(slug: string) {
	if (slug === '[object Object]') {
		return { data: ref(null) }
	}

	function fetchPostDetails() {
		return getPostDetailsSSR(slug)
	}

	return useAsyncData<EntitiesPostResponse>(
		`${slug}-details`,
		fetchPostDetails
	)
}
