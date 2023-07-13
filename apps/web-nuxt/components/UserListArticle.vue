<script lang="ts" setup>
import { useInfiniteQuery } from '@tanstack/vue-query'
import dayjs from 'dayjs'
import relativeTime from 'dayjs/plugin/relativeTime'

dayjs.extend(relativeTime)

async function fetchUserPosts({ pageParam = 1 }) {
	return $blogApi.post
		.postsGet(
			userProfile.value?.id,
			undefined,
			pageParam,
			PAGE_SIZE,
			'desc'
		)
		.then((res) => res.data)
}

const PAGE_SIZE = 10
const { data: userProfile } = useUserProfile()
const { $blogApi } = useNuxtApp()
const enabled = computed(() => !!userProfile.value?.id)
const {
	data: userPostData,
	isFetching,
	isFetchingNextPage,
	isLoading,
	fetchNextPage,
	hasNextPage
} = useInfiniteQuery({
	queryKey: ['user-posts', userProfile.value?.id],
	queryFn: fetchUserPosts,
	getNextPageParam: (lastPage) =>
		lastPage.page && lastPage.items?.length === PAGE_SIZE
			? lastPage.page + 1
			: undefined,
	enabled
})
const target = ref(null)
const targetIsVisible = ref(false)
useIntersectionObserver(target, ([{ isIntersecting }], _) => {
	targetIsVisible.value = isIntersecting
})
const userPosts = computed(() => flattenPostData(userPostData))

watchEffect(() => {
	const loading =
		isFetching.value || isLoading.value || isFetchingNextPage.value
	if (targetIsVisible.value && hasNextPage?.value && !loading) {
		fetchNextPage()
	}
})
</script>

<template>
	<div class="flex-grow overflow-y-auto flex flex-col">
		<div class="overflow-y-auto max-h-full">
			<div class="flex flex-col gap-2">
				<ListArticle
					:loading="
						isFetching ||
						isFetchingNextPage ||
						isLoading
					"
					:data="userPosts"
				/>
				<div ref="target" />
			</div>
		</div>
	</div>
</template>
