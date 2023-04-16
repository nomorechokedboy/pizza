import { useInfiniteQuery } from '@tanstack/vue-query'

export function usePostComments() {
	const { $blogApi } = useNuxtApp()
	const pageSize = 30
	const slug = inject<string>('slug', 'no slug')
	const { data: postDetails } = usePostDetails(slug)
	const postId = computed(() => postDetails.value?.id)
	const enabled = computed(() => !!postDetails.value?.id)

	async function getPostComments({ pageParam: page = 1 }) {
		return $blogApi.comment
			.commentsGet(
				undefined,
				postId.value,
				undefined,
				page,
				pageSize,
				'desc',
				'id'
			)
			.then((resp) => resp.data)
	}

	return useInfiniteQuery({
		queryKey: ['comments', postId],
		queryFn: getPostComments,
		enabled,
		getNextPageParam: (lastPage) =>
			lastPage.page && lastPage.items?.length === pageSize
				? lastPage.page + 1
				: undefined
	})
}
