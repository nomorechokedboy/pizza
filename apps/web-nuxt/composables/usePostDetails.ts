import { QueryFunctionContext, useQuery } from '@tanstack/vue-query'
import { EntitiesPostDetailResponse } from '~/codegen/api'

export async function getPostDetails({
	queryKey
}: QueryFunctionContext<string[], any>): Promise<EntitiesPostDetailResponse> {
	const { $blogApi } = useNuxtApp()
	return $blogApi.post.postsSlugGet(queryKey[1]).then((resp) => resp.data)
}

export function usePostDetails(slug: string) {
	return useQuery({
		queryFn: getPostDetails,
		queryKey: ['postDetails', slug]
	})
}
