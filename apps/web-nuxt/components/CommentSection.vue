<script lang="ts" setup>
import { useInfiniteQuery } from '@tanstack/vue-query'
import { useElementVisibility } from '@vueuse/core'
import { ActionIcon, Button } from 'ui-vue'
import { dicebearMedia } from '~/constants'
import ChevronIcon from '~icons/mdi/chevron-up-down'
import { CommentProps } from './Comment.vue'

export interface CommentSectionProps {
	slug: string
}

const { $blogApi } = useNuxtApp()
const { slug } = defineProps<CommentSectionProps>()
const userProfile = useUserProfile()
const isAuth = useIsAuthenticated()
const userAvatar = computed(computeUserAvatar)
const { data: postDetails } = usePostDetails(slug)
const formData = reactive({ content: '' })
const isFormValid = computed(computeFormValidity)
const postId = computed(() => postDetails.value?.id)
const enabled = computed(() => !!postDetails.value?.id)
const {
	data: commentData,
	isFetching,
	isFetchingNextPage,
	isLoading,
	fetchNextPage,
	hasNextPage,
	refetch: refetchComments
} = useInfiniteQuery({
	queryKey: ['comments', postId],
	queryFn: getPostComments,
	enabled,
	getNextPageParam: (lastPage) =>
		lastPage.page && lastPage.items?.length === pageSize
			? lastPage.page + 1
			: undefined
})
const comments = computed(computeComments)
const loading = ref(false)
const blankCommentProps: CommentProps = {
	like: 0,
	loading: true,
	createdAt: '',
	updated: false,
	replies: [],
	user: { avatarUrl: '', name: '' },
	content: ''
}
const pageSize = 30
const target = ref(null)
const targetIsVisible = useElementVisibility(target)

function computeUserAvatar() {
	return (
		userProfile.value?.avatar ||
		`${dicebearMedia}${
			userProfile.value.name ||
			'A6Blog&backgroundColor=000000'
		}`
	)
}

function computeFormValidity() {
	return (
		formData.content.length !== 0 &&
		postDetails.value?.id !== undefined
	)
}

function computeComments() {
	const comments: (CommentProps & { id: number })[] = []
	commentData.value?.pages.forEach((page) =>
		page?.items?.forEach(
			({ content, id, user, createdAt, updatedAt }) => {
				comments.push({
					id,
					content,
					user: {
						name: user?.fullName,
						avatarUrl:
							user?.avatar === ''
								? `${dicebearMedia}${
										user?.fullName ||
										'A6Blog'
								  }`
								: user?.avatar
					},
					like: 0,
					replies: [],
					updated: createdAt !== updatedAt,
					createdAt: getCalendarTime(createdAt)
				} as CommentProps & { id: number })
			}
		)
	)

	return comments
}

async function getPostComments({ pageParam: page = 1 }) {
	return $blogApi.comment
		.commentsGet(
			undefined,
			postId.value,
			undefined,
			page,
			pageSize,
			'desc',
			'id'
		)
		.then((resp) => resp.data)
}

async function handleSubmit() {
	if (!isFormValid) {
		return
	}

	loading.value = true
	try {
		await $blogApi.comment.commentsPost({
			content: formData.content,
			postId: postDetails.value!.id
		})
		refetchComments()
		formData.content = ''
	} catch (e) {
		notifyError(e)
	} finally {
		loading.value = false
	}
}
watchEffect(() => {
	const loading =
		isFetching.value || isLoading.value || isFetchingNextPage.value
	if (targetIsVisible.value && hasNextPage?.value && !loading) {
		fetchNextPage()
	}
})
</script>

<template>
	<section
		id="comments"
		class="p-4 border-t border-t-neutral-200 flex flex-col gap-6"
	>
		<header class="flex items-center justify-between">
			<div class="flex items-center">
				<h2 class="text-xl text-neutral-800 font-bold">
					Top comments (0)
				</h2>
				<ActionIcon
					color="indigo"
					variant="subtle"
					size="xl"
				>
					<ChevronIcon
						class="text-neutral-800"
						width="24"
					/>
				</ActionIcon>
			</div>
			<!-- <Button>Subscribe</Button> -->
		</header>
		<div class="flex flex-col gap-4">
			<div class="flex gap-2">
				<div class="w-full max-w-[32px]">
					<Avatar
						class="md:w-8"
						width="24"
						:src="userAvatar"
					/>
				</div>
				<form class="flex flex-col gap-4 flex-1">
					<div
						class="text-neutral-900 text-base shadow rounded-md flex-1 flex flex-col gap-4 px-3 pt-2 pb-4"
					>
						<main>
							<RichTextEditor
								v-model="
									formData.content
								"
								reversed
							/>
						</main>
					</div>
					<div
						class="flex items-center gap-3 justify-end"
						v-if="isAuth"
					>
						<Button
							type="submit"
							@click.prevent="
								handleSubmit
							"
							:disabled="!isFormValid"
							:loading="loading"
							>Comment</Button
						>
						<Button
							color="indigo"
							variant="subtle"
							:disabled="
								!isFormValid ||
								loading
							"
							>Preview</Button
						>
					</div>
				</form>
			</div>
		</div>
		<Comment
			v-for="{ id, ...comment } in comments"
			v-bind="comment"
			:key="id"
		/>
		<template v-if="isFetching || isFetchingNextPage || isLoading">
			<Comment
				v-for="n in 5"
				v-bind="blankCommentProps"
				:key="n"
			/>
		</template>
		<div ref="target" />
		<!-- <nav class="flex items-center justify-center text-neutral-500 text-sm">
            <p>Code of conduct</p>
            <div class="py-2 px-3">
                <div class="h-1 w-1 rounded-full bg-[rgb(189,189,189)]" />
            </div>
            <p>Report abuse</p>
        </nav> -->
	</section>
</template>
