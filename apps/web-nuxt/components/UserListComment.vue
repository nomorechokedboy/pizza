<script lang="ts" setup>
import { useInfiniteQuery } from '@tanstack/vue-query'

async function fetchUserComments({ pageParam = 1 }) {
	return $blogApi.comment
		.commentsGet(
			userProfile.value?.id,
			undefined,
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
	data: userCommentData,
	isFetching,
	isFetchingNextPage,
	isLoading,
	fetchNextPage,
	hasNextPage
} = useInfiniteQuery({
	queryKey: ['user-comments', userProfile.value?.id],
	queryFn: fetchUserComments,
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

watchEffect(() => {
	const loading =
		isFetching.value || isLoading.value || isFetchingNextPage.value
	if (targetIsVisible.value && hasNextPage?.value && !loading) {
		fetchNextPage()
	}
})
</script>

<template>
	<div class="flex-grow overflow-y-auto flex flex-col border-t">
		<div class="overflow-y-auto max-h-full">
			<div class="flex flex-col gap-2">
				<template
					v-for="page in userCommentData?.pages"
				>
					<RecentComment
						v-for="c in page.items"
						:created-at="
							timeFromNow(c.createdAt)
						"
						:comment-content="
							c.content || 'Oops'
						"
						:post-title="
							c.post?.title || ''
						"
					/>
				</template>
			</div>
		</div>
	</div>
</template>
