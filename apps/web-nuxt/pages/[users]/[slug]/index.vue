<script lang="ts" setup>
import { Button } from 'ui-vue'
import { dicebearMedia } from '~~/constants'

function computePostedOn() {
	return `${
		postDetails.value?.published ? 'Posted on' : 'Created on'
	} ${convertDate(
		postDetails.value?.publishedAt ?? postDetails.value?.createdAt
	)}`
}

function handleOpenModal() {
	openedModal.value = true
}

function handleCloseModal() {
	openedModal.value = false
}

async function handleDeletePost() {
	if (!postDetails.value?.id) {
		return
	}

	loading.value = true
	try {
		await $blogApi.post.postsIdDelete(postDetails.value?.id)
		notify({
			type: 'success',
			k: crypto.randomUUID(),
			content: 'Delete post success!'
		})
		openedModal.value = false
		await navigateTo('/')
	} catch (e) {
		notifyError(e)
	} finally {
		loading.value = false
	}
}

const route = useRoute()
// const tags = ['tag1', 'tag2']
let slug: string = ''
if (typeof route.params.slug === 'object') {
	slug = route.params.slug.pop() || ''
} else {
	slug = route.params.slug
}

const { data: postDetails } = usePostDetails(slug)
const url =
	postDetails.value?.image ??
	`/api/seo/og?title=${postDetails.value?.title}`
const postedOn = computed(computePostedOn)
const { $blogApi } = useNuxtApp()
const openedModal = ref(false)
const loading = ref(false)
const userProfile = useUserProfile()
const editUrl = computed(
	() => `/${postDetails.value?.user?.id}/${postDetails.value?.slug}/edit`
)
useSeoMeta({
	title: postDetails.value?.title,
	description: postDetails.value?.content,
	ogImage: url,
	ogImageSecureUrl: url,
	ogImageWidth: 1200,
	ogImageHeight: 600,
	twitterImage: url
})
</script>

<template>
	<div
		id="popup-modal"
		tabindex="-1"
		v-if="openedModal"
		class="fixed top-0 left-0 right-0 z-50 p-4 overflow-x-hidden overflow-y-auto md:inset-0 h-[calc(100%-1rem)] max-h-full"
	>
		<div class="relative w-full max-w-md max-h-full m-auto">
			<div
				class="relative bg-white rounded-lg shadow dark:bg-gray-700"
			>
				<button
					type="button"
					:disabled="loading"
					class="absolute top-3 right-2.5 text-gray-400 bg-transparent hover:bg-gray-200 hover:text-gray-900 rounded-lg text-sm p-1.5 ml-auto inline-flex items-center dark:hover:bg-gray-800 dark:hover:text-white"
					@click="handleCloseModal"
					data-modal-hide="popup-modal"
				>
					<svg
						aria-hidden="true"
						class="w-5 h-5"
						fill="currentColor"
						viewBox="0 0 20 20"
						xmlns="http://www.w3.org/2000/svg"
					>
						<path
							fill-rule="evenodd"
							d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z"
							clip-rule="evenodd"
						></path>
					</svg>
					<span class="sr-only">Close modal</span>
				</button>
				<div class="p-6 text-center">
					<svg
						aria-hidden="true"
						class="mx-auto mb-4 text-gray-400 w-14 h-14 dark:text-gray-200"
						fill="none"
						stroke="currentColor"
						viewBox="0 0 24 24"
						xmlns="http://www.w3.org/2000/svg"
					>
						<path
							stroke-linecap="round"
							stroke-linejoin="round"
							stroke-width="2"
							d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"
						></path>
					</svg>
					<h3
						class="mb-5 text-lg font-normal text-gray-500 dark:text-gray-400"
					>
						Are you sure you want to delete
						this product?
					</h3>
					<div
						class="flex items-center justify-center gap-3"
					>
						<Button
							color="red"
							@click="
								handleDeletePost
							"
							:loading="loading"
						>
							Yes, I'm sure
						</Button>
						<Button
							@click="
								handleCloseModal
							"
							:disabled="loading"
						>
							No, cancel</Button
						>
					</div>
				</div>
			</div>
		</div>
	</div>
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
						<div
							class="flex items-center justify-between"
						>
							<div class="w-1/2">
								<NoClue
									:alt="`${postDetails?.user?.fullName} avatar`"
									:description="
										postedOn
									"
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
							</div>
							<div
								class="flex items-center gap-3"
								v-if="
									postDetails
										?.user
										?.id ===
									userProfile?.id
								"
							>
								<Button
									color="red"
									@click="
										handleOpenModal
									"
									>Delete</Button
								>
								<NuxtLink
									:to="
										editUrl
									"
								>
									<Button
										>Edit</Button
									>
								</NuxtLink>
							</div>
						</div>
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
