<script setup lang="ts">
import { useInfiniteQuery } from '@tanstack/vue-query'
import dayjs from 'dayjs'
import relativeTime from 'dayjs/plugin/relativeTime'

dayjs.extend(relativeTime)

type HeaderNav = {
	content: string
	match: string
}

async function fetchPosts({ pageParam: page = 1 }) {
	return $blogApi.post
		.postsGet(undefined, undefined, page, pageSize, 'desc')
		.then((resp) => resp.data)
}

const pageSize = 50
const { $blogApi } = useNuxtApp()
const headerNav: HeaderNav[] = [
	{ content: 'Relevant', match: '/' }
	/* { content: 'Latest', match: '/latest' },
    { content: 'Top', match: '/top' } */
]
const route = useRoute()
const {
	data: postData,
	isFetching,
	isFetchingNextPage,
	isLoading,
	fetchNextPage,
	hasNextPage
} = useInfiniteQuery({
	queryKey: ['posts'],
	queryFn: fetchPosts,
	getNextPageParam: (lastPage) =>
		lastPage.page && lastPage.items?.length === pageSize
			? lastPage.page + 1
			: undefined
})
const target = ref(null)
const targetIsVisible = useElementVisibility(target)
const posts = computed(() => flattenPostData(postData))

watchEffect(() => {
	const loading =
		isFetching.value || isLoading.value || isFetchingNextPage.value
	if (targetIsVisible.value && hasNextPage?.value && !loading) {
		fetchNextPage()
	}
})

useSeoMeta({
	title: 'Accessiblog',
	description:
		"Accessiblog is a web application that focuses on creating accessible and inclusive blog content. With a range of tools and features, including a rich text editor and image descriptions, Accessiblog makes it easy for bloggers to create content that can be enjoyed by all users. Join our community of bloggers today and start creating content that's accessible to everyone.",
	ogImage: '/logo.png'
})
</script>

<template>
	<SidebarLayout class="md:p-4 lg:max-w-7xl">
		<template #sidebar>
			<LeftSidebar />
		</template>
		<SidebarLayout :reverse="true">
			<template #sidebar>
				<LeftSidebar
					class="hidden lg:block lg:max-w-xs"
				/>
			</template>
			<div class="w-full">
				<header class="md:p-0 md:px-0 md:mb-2 p-2 px-3">
					<h1 class="sr-only">Posts</h1>
					<nav
						class="m:mx-0 s:flex -mx-3 items-center justify-between"
					>
						<ul
							class="flex items-center text-lg text-[rgb(87,87,87)]"
						>
							<li
								v-for="{
									content,
									match
								} in headerNav"
								:key="content"
								class="py-2 px-3 capitalize"
								:class="{
									'font-bold text-neutral-900':
										route.path ===
										match
								}"
							>
								<nuxt-link
									:to="
										match
									"
								>
									{{
										content
									}}
								</nuxt-link>
							</li>
						</ul>
					</nav>
				</header>
				<div class="flex flex-col gap-2">
					<ListArticle
						:data="posts"
						:loading="
							isFetching ||
							isFetchingNextPage ||
							isLoading
						"
					/>
					<div ref="target" />
				</div>
			</div>
		</SidebarLayout>
	</SidebarLayout>
</template>
