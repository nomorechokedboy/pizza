<script setup lang="ts">
import { CommonBasePaginationResponseEntitiesPostRes } from '~~/codegen/api'

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
		`https://api-blog-dev-nomorechokedboy.cloud.okteto.net/api/v1/posts?${queryParams}`
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
const { data: posts } =
	await useAsyncData<CommonBasePaginationResponseEntitiesPostRes>(
		`posts-${baseQuery.page}-${baseQuery.pageSize}`,
		getPosts
	)
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
					<!-- <div v-if="loading">Loading...</div> -->
					<Article
						v-for="(
							{
								user,
								id,
								publishedAt,
								title,
								slug
							},
							i
						) in posts?.items || []"
						src="https://res.cloudinary.com/practicaldev/image/fetch/s--iJh8Y2cI--/c_imagga_scale,f_auto,fl_progressive,h_420,q_auto,w_1000/https://dev-to-uploads.s3.amazonaws.com/uploads/articles/qqi78z7yx9q6koq6dmrc.png"
						:owner="{
							src:
								user?.avatar ||
								'https://res.cloudinary.com/practicaldev/image/fetch/s--5VEqFAA8--/c_fill,f_auto,fl_progressive,h_90,q_66,w_90/https://dev-to-uploads.s3.amazonaws.com/uploads/user/profile_image/909049/9a19683f-1e9f-4933-bdba-e7ea2fe5e71c.gif',
							name:
								user?.username ||
								'User name'
						}"
						:slug="slug || ''"
						:tags="['tags']"
						:publishedAt="publishedAt || ''"
						:title="title || ''"
						:showImage="i === 0"
						:comments="0"
						:like="0"
						:key="id"
					/>
				</div>
			</div>
		</SidebarLayout>
	</SidebarLayout>
</template>
