<script lang="ts" setup>
import { EntitiesPostRes } from '~~/codegen/api'
import { baseURL, dicebearMedia } from '~~/constants'

function getPostDetails(): Promise<EntitiesPostRes> {
	return $fetch(`${baseURL}/api/v1/posts/${slug}`)
}

function computePostedOn() {
	return `${
		postDetails.value?.published ? 'Posted on' : 'Created on'
	} ${convertDate(
		postDetails.value?.publishedAt ?? postDetails.value?.createdAt
	)}`
}

const route = useRoute()
// const tags = ['tag1', 'tag2']
let slug: string = ''
if (typeof route.params.slug === 'object') {
	slug = route.params.slug.pop() || ''
} else {
	slug = route.params.slug
}

const { data: postDetails } = await useAsyncData<EntitiesPostRes>(
	`${slug}-details`,
	getPostDetails
)
const nuxtApp = useNuxtApp()
const host = nuxtApp.ssrContext?.event.node.req.headers.host
const postedOn = computed(computePostedOn)
useSeoMeta({
	title: postDetails.value?.title,
	description: postDetails.value?.content,
	ogImage: `${host}/api/seo/og?title=${postDetails.value?.title}`
})
</script>

<template>
	<SidebarLayout class="md:p-4 lg:max-w-7xl">
		<template #sidebar>
			<LeftSidebar class="lg:max-w-[64px]" />
		</template>
		<SidebarLayout :reverse="true">
			<template #sidebar>
				<LeftSidebar
					class="hidden lg:block lg:max-w-xs"
				/>
			</template>
			<div class="w-full max-w-[832px]">
				<article
					class="bg-white border border-neutral-200"
				>
					<nuxt-img
						v-if="postDetails?.image"
						sizes="sm:100vw md:680px, 806px"
						:src="postDetails.image"
					/>
					<section
						class="flex flex-col gap-3 p-4"
					>
						<NoClue
							:alt="`${postDetails?.user?.fullName} avatar`"
							:description="postedOn"
							:src="
								postDetails
									?.user
									?.avatar ||
								`${dicebearMedia}${postDetails?.user?.fullName}`
							"
							:title="
								postDetails
									?.user
									?.fullName ||
								'User name'
							"
						/>
						<h1
							class="text-3xl font-bold text-neutral-900"
						>
							{{ postDetails?.title }}
						</h1>
						<!-- <Tags :tags="tags" /> -->
						<main class="max-w-full">
							<VueMarkdown
								class="markdown"
								:key="
									postDetails?.content
								"
								:source="
									postDetails?.content ||
									''
								"
							/>
						</main>
					</section>
					<CommentSection />
				</article>
			</div>
		</SidebarLayout>
	</SidebarLayout>
</template>

<style lang="postcss" scoped>
.markdown > * {
	@apply overflow-x-auto break-words;
}
</style>
