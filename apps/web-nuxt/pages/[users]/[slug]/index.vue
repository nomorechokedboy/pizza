<script lang="ts" setup>
import { ActionIcon, Button } from 'ui-vue'
import { useAVWaveform } from 'vue-audio-visual'
import RoundPauseCircleIcon from '~icons/ic/round-pause-circle'
import RoundPlayCircleIcon from '~icons/ic/round-play-circle'

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

function handleAudioButton() {
	if (player.value?.paused) {
		player.value?.play()
	} else {
		player.value?.pause()
	}
}

function updatePauseState() {
	if (!player.value) {
		return
	}

	isPause.value = player.value.paused
}

const appConfig = useRuntimeConfig()
const route = useRoute()
// const tags = ['tag1', 'tag2']
let slug: string = ''
if (typeof route.params.slug === 'object') {
	slug = route.params.slug.pop() || ''
} else {
	slug = route.params.slug
}

const { data: postDetails } = usePostDetailsSSR(slug)
const url =
	postDetails.value?.image ??
	`/api/seo/og?title=${postDetails.value?.title}`
const postedOn = computed(computePostedOn)
const { $blogApi } = useNuxtApp()
const openedModal = ref(false)
const loading = ref(false)
const { data: userProfile } = useUserProfile()
const editUrl = computed(
	() => `/${postDetails.value?.user?.id}/${postDetails.value?.slug}/edit`
)
const isAuthenticated = useIsAuthenticated()
const audioUrl = computed(
	() =>
		`${appConfig.public.apiUrl}/api/v1/posts/t2s/${postDetails.value?.slug}`
)
const player = ref<HTMLAudioElement | null>(null)
const canvas = ref<HTMLCanvasElement | null>(null)
const isPause = ref(true)

if (postDetails.value) {
	useSeoMeta({
		title: postDetails.value?.title,
		description: postDetails.value?.content,
		ogImage: url,
		ogImageSecureUrl: url,
		ogImageWidth: 1200,
		ogImageHeight: 600,
		twitterImage: url
	})
}
if (process.client) {
	useAVWaveform(player, canvas, {
		src: audioUrl,
		canvHeight: 56,
		canvWidth: 768,
		barColor: 'lime'
	})
}

onMounted(() => {
	player.value?.addEventListener('play', updatePauseState)
	player.value?.addEventListener('pause', updatePauseState)
})
onBeforeUnmount(() => {
	player.value?.removeEventListener('play', updatePauseState)
	player.value?.removeEventListener('pause', updatePauseState)
})
provide('slug', slug)
</script>

<template>
	<Modal :opened="openedModal" :onClose="handleCloseModal">
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
				But sir, are you sure you want to
				<br />
				delete this post?
			</h3>
			<div class="flex items-center justify-center gap-3">
				<Button
					color="red"
					@click="handleDeletePost"
					:loading="loading"
				>
					Yes, just do it
				</Button>
				<Button
					@click="handleCloseModal"
					:disabled="loading"
				>
					Wait, I changed my mind</Button
				>
			</div>
		</div>
	</Modal>
	<SidebarLayout class="md:p-4 lg:max-w-7xl">
		<template #sidebar>
			<LeftSidebar class="lg:max-w-[64px]" />
		</template>
		<SidebarLayout reverse>
			<template #sidebar>
				<AuthorSidebar />
			</template>
			<div class="w-full max-w-[832px]">
				<article
					class="bg-white border border-neutral-200 rounded-lg pt-4"
				>
					<nuxt-img
						class="h-80 md:w-full"
						v-if="postDetails?.image"
						sizes="sm:100vw md:680px, 806px"
						height="160"
						:alt="`${postDetails.title} image`"
						:src="postDetails.image"
						:modifiers="{
							default: '/image_not_available.png'
						}"
					/>
					<div class="px-2">
						<div
							aria-label="Audio"
							class="flex items-center border border-green-500 rounded"
						>
							<ActionIcon
								aria-label="Play audio"
								@click="
									handleAudioButton
								"
								color="green"
								size="xl"
								variant="transparent"
							>
								<RoundPlayCircleIcon
									v-if="
										isPause
									"
								/>
								<RoundPauseCircleIcon
									v-else
								/>
							</ActionIcon>
							<canvas ref="canvas" />
						</div>
						<audio
							class="hidden"
							ref="player"
							v-if="
								postDetails?.content !==
								undefined
							"
							:src="audioUrl"
							controls
						/>
					</div>
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
											?.avatar
											? `${appConfig.public.mediaUrl}${postDetails?.user?.avatar}`
											: `${appConfig.public.dicebearMedia}${postDetails?.user?.fullName}`
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
									isAuthenticated &&
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
