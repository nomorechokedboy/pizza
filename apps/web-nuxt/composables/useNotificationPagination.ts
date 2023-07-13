import { useInfiniteQuery } from '@tanstack/vue-query'

export function useNotificationPagination() {
	const { $blogApi } = useNuxtApp()
	const PAGE_SIZE = 10

	async function fetchNotifications({ pageParam = 0 }) {
		return $blogApi.notification
			.getNotifications(pageParam, PAGE_SIZE)
			.then((res) => res.data)
	}

	return useInfiniteQuery({
		queryKey: ['notifications'],
		queryFn: fetchNotifications,
		getNextPageParam: (lastPage) =>
			lastPage.data.length === PAGE_SIZE
				? lastPage.page + 1
				: undefined
	})
}
