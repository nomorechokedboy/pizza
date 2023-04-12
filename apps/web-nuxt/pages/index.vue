<script setup lang="ts">
import { useInfiniteQuery } from '@tanstack/vue-query'
import { useElementVisibility } from '@vueuse/core'
import dayjs from 'dayjs'
import relativeTime from 'dayjs/plugin/relativeTime'
import { dicebearMedia } from '~~/constants'

dayjs.extend(relativeTime)

type HeaderNav = {
	content: string
	match: string
}

async function fetchPosts({ pageParam = 1 }) {
	return $blogApi.post
		.postsGet(undefined, undefined, pageParam, undefined, 'desc')
		.then((resp) => resp.data)
}

const { $blogApi } = useNuxtApp()
const headerNav: HeaderNav[] = [
	{ content: 'Relevant', match: '/' }
	/* { content: 'Latest', match: '/latest' },
    { content: 'Top', match: '/top' } */
]
const route = useRoute()
const {
	data: testPosts,
	isFetching,
	isFetchingNextPage,
	isLoading,
	fetchNextPage,
	hasNextPage
} = useInfiniteQuery({
	queryKey: ['posts'],
	queryFn: fetchPosts,
	getNextPageParam: (lastPage) =>
		lastPage.page && lastPage.items?.length === 10
			? lastPage.page + 1
			: undefined
})
const target = ref(null)
const targetIsVisible = useElementVisibility(target)
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
					<template
						v-for="page in testPosts?.pages"
					>
						<Article
							v-for="{
								user,
								id,
								publishedAt,
								title,
								slug,
								commentCount,
								image
							} in page.items"
							:owner="{
								src:
									user?.avatar ||
									`${dicebearMedia}${user?.fullName}`,
								name:
									user?.fullName ||
									'User name',
								id: id ?? 0
							}"
							:slug="slug || ''"
							:tags="['tags']"
							:publishedAt="`${convertDate(
								publishedAt
							)} (${dayjs(
								publishedAt
							).toNow(true)} ago)`"
							:title="title || ''"
							:src="image"
							:comments="
								commentCount ||
								0
							"
							:like="0"
							:key="id"
						/>
					</template>
					<template
						v-if="
							isFetching ||
							isFetchingNextPage ||
							isLoading
						"
					>
						<Article
							v-for="n in 3"
							:owner="{
								id: n,
								name: 'lmao',
								src: 'https://res.cloudinary.com/practicaldev/image/fetch/s--5VEqFAA8--/c_fill,f_auto,fl_progressive,h_90,q_66,w_90/https://dev-to-uploads.s3.amazonaws.com/uploads/user/profile_image/909049/9a19683f-1e9f-4933-bdba-e7ea2fe5e71c.gif'
							}"
							slug="test"
							publishedAt="Lmao"
							title="test"
							:comments="0"
							:like="0"
							src="https://res.cloudinary.com/practicaldev/image/fetch/s--xow1lZzw--/c_imagga_scale,f_auto,fl_progressive,h_420,q_auto,w_1000/https://dev-to-uploads.s3.amazonaws.com/uploads/articles/ndp8r87ccsoa5yzw6zqh.png"
							:key="n"
							showImage
							loading
						/>
					</template>
					<div ref="target" />
				</div>
			</div>
		</SidebarLayout>
	</SidebarLayout>
</template>
