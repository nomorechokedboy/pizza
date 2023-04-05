<script setup lang="ts">
import { ArticleProps } from '~~/components/Article.vue'

type HeaderNav = {
	content: string
	match: string
}

const headerNav: HeaderNav[] = [
	{ content: 'Relevant', match: '/' },
	{ content: 'Latest', match: '/latest' },
	{ content: 'Top', match: '/top' }
]
const route = useRoute()
const articles: ArticleProps[] = [
	{
		comments: 0,
		publishedAt: 'Apr 1',
		like: 13,
		owner: {
			name: 'Duong Le',
			src: 'https://api.dicebear.com/6.x/icons/svg?seed=CheeseRaa'
		},
		tags: ['tags 1', 'tags 2'],
		title: 'Learn How to Use Artifactory as a Docker Registry',
		src: 'https://res.cloudinary.com/practicaldev/image/fetch/s--h0iBOiYw--/c_imagga_scale,f_auto,fl_progressive,h_420,q_auto,w_1000/https://dev-to-uploads.s3.amazonaws.com/uploads/articles/qxfnu4dptk7iazjv6mjg.png',
		slug: 'test-1'
	},
	{
		comments: 13,
		publishedAt: 'Apr 1',
		like: 0,
		owner: {
			name: 'Duong Le',
			src: 'https://api.dicebear.com/6.x/icons/svg?seed=CheeseRaa'
		},
		tags: [],
		title: 'Learn How to Use Artifactory as a Docker Registry',
		src: 'https://res.cloudinary.com/practicaldev/image/fetch/s--h0iBOiYw--/c_imagga_scale,f_auto,fl_progressive,h_420,q_auto,w_1000/https://dev-to-uploads.s3.amazonaws.com/uploads/articles/qxfnu4dptk7iazjv6mjg.png',
		slug: 'test-2'
	},
	{
		comments: 0,
		publishedAt: 'Apr 1',
		like: 13,
		owner: {
			name: 'Duong Le',
			src: 'https://api.dicebear.com/6.x/icons/svg?seed=CheeseRaa'
		},
		tags: ['tags 1', 'tags 2'],
		title: 'Learn How to Use Artifactory as a Docker Registry',
		src: 'https://res.cloudinary.com/practicaldev/image/fetch/s--h0iBOiYw--/c_imagga_scale,f_auto,fl_progressive,h_420,q_auto,w_1000/https://dev-to-uploads.s3.amazonaws.com/uploads/articles/qxfnu4dptk7iazjv6mjg.png',
		slug: 'test-3'
	},
	{
		comments: 0,
		publishedAt: 'Apr 1',
		like: 13,
		owner: {
			name: 'Duong Le',
			src: 'https://api.dicebear.com/6.x/icons/svg?seed=CheeseRaa'
		},
		tags: ['tags 1', 'tags 2'],
		title: 'Learn How to Use Artifactory as a Docker Registry',
		src: 'https://res.cloudinary.com/practicaldev/image/fetch/s--h0iBOiYw--/c_imagga_scale,f_auto,fl_progressive,h_420,q_auto,w_1000/https://dev-to-uploads.s3.amazonaws.com/uploads/articles/qxfnu4dptk7iazjv6mjg.png',
		slug: 'test-3'
	},
	{
		comments: 0,
		publishedAt: 'Apr 1',
		like: 13,
		owner: {
			name: 'Duong Le',
			src: 'https://api.dicebear.com/6.x/icons/svg?seed=CheeseRaa'
		},
		tags: ['tags 1', 'tags 2'],
		title: 'Learn How to Use Artifactory as a Docker Registry',
		src: 'https://res.cloudinary.com/practicaldev/image/fetch/s--h0iBOiYw--/c_imagga_scale,f_auto,fl_progressive,h_420,q_auto,w_1000/https://dev-to-uploads.s3.amazonaws.com/uploads/articles/qxfnu4dptk7iazjv6mjg.png',
		slug: 'test-3'
	},
	{
		comments: 0,
		publishedAt: 'Apr 1',
		like: 13,
		owner: {
			name: 'Duong Le',
			src: 'https://api.dicebear.com/6.x/icons/svg?seed=CheeseRaa'
		},
		tags: ['tags 1', 'tags 2'],
		title: 'Learn How to Use Artifactory as a Docker Registry',
		src: 'https://res.cloudinary.com/practicaldev/image/fetch/s--h0iBOiYw--/c_imagga_scale,f_auto,fl_progressive,h_420,q_auto,w_1000/https://dev-to-uploads.s3.amazonaws.com/uploads/articles/qxfnu4dptk7iazjv6mjg.png',
		slug: 'test-3'
	}
]
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
					<Article
						v-for="(props, i) in articles"
						:="props"
						:showImage="i === 0"
						:key="i"
					/>
				</div>
			</div>
		</SidebarLayout>
	</SidebarLayout>
</template>
