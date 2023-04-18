import { QueryFunctionContext, useQuery } from '@tanstack/vue-query'

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
