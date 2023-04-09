<script lang="ts" setup>
import { EntitiesPostRes } from '~~/codegen/api'

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
	() =>
		$fetch(
			`https://api-blog-dev-nomorechokedboy.cloud.okteto.net/api/v1/posts/${slug}`
		)
)
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
						sizes="sm:100vw md:680px, 806px"
						src="https://res.cloudinary.com/practicaldev/image/fetch/s--h0iBOiYw--/c_imagga_scale,f_auto,fl_progressive,h_420,q_auto,w_1000/https://dev-to-uploads.s3.amazonaws.com/uploads/articles/qxfnu4dptk7iazjv6mjg.png"
					/>
					<section
						class="flex flex-col gap-3 p-4"
					>
						<NoClue
							description="Lmao"
							src="https://api.dicebear.com/6.x/icons/svg?seed=CheeseRaa"
							title="owner.name"
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
