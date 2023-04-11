<script setup lang="ts">
import dayjs from 'dayjs'
import relativeTime from 'dayjs/plugin/relativeTime'
import { CommonBasePaginationResponseEntitiesPostRes } from '~~/codegen/api'
import { baseURL, dicebearMedia } from '~~/constants'

dayjs.extend(relativeTime)

type HeaderNav = {
	content: string
	match: string
}

type BaseQuery = {
	page?: number
	pageSize?: number
	q?: string
	sort?: 'asc' | 'desc'
	sortBy?: string
}

async function getPosts() {
	let queryParams = ''
	for (const [key, val] of Object.entries(baseQuery)) {
		if (val !== undefined) {
			queryParams += `${key}=${val}&`
		}
	}

	return $fetch<CommonBasePaginationResponseEntitiesPostRes>(
		`${baseURL}/api/v1/posts?${queryParams}`
	)
}

const baseQuery = reactive<BaseQuery>({
	sort: 'desc',
	sortBy: 'id',
	q: undefined,
	pageSize: 20,
	page: 0
})
const headerNav: HeaderNav[] = [
	{ content: 'Relevant', match: '/' }
	/* { content: 'Latest', match: '/latest' },
    { content: 'Top', match: '/top' } */
]
const route = useRoute()
const { data: posts, pending: isPostsPending } =
	await useAsyncData<CommonBasePaginationResponseEntitiesPostRes>(
		`posts-${baseQuery.page}-${baseQuery.pageSize}`,
		getPosts
	)
useSeoMeta({
	title: 'Accessiblog',
	description:
		"Accessiblog is a web application that focuses on creating accessible and inclusive blog content. With a range of tools and features, including a rich text editor and image descriptions, Accessiblog makes it easy for bloggers to create content that can be enjoyed by all users. Join our community of bloggers today and start creating content that's accessible to everyone."
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
					<div v-if="isPostsPending">
						Loading...
					</div>
					<Article
						v-for="(
							{
								user,
								id,
								publishedAt,
								title,
								slug,
								comments
							},
							i
						) in posts?.items || []"
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
						)} (${dayjs(publishedAt).toNow(
							true
						)} ago)`"
						:title="title || ''"
						:showImage="i === 0"
						:comments="
							comments?.length || 0
						"
						:like="0"
						:key="id"
					/>
				</div>
			</div>
		</SidebarLayout>
	</SidebarLayout>
</template>
