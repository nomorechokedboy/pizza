<script lang="ts" setup>
import { useElementVisibility } from '@vueuse/core'
import { ActionIcon, Button } from 'ui-vue'
import { EntitiesComment } from '~/codegen/api'
import ChevronIcon from '~icons/mdi/chevron-up-down'
import { CommentProps } from './Comment.vue'

const appConfig = useRuntimeConfig()
const { $blogApi } = useNuxtApp()
const { data: userProfile, isLoading: isUserProfileLoading } = useUserProfile()
const isAuth = useIsAuthenticated()
const userAvatar = computed(computeUserAvatar)
const formData = reactive({ content: '' })
const isFormValid = computed(computeFormValidity)
const slug = inject<string>('slug', 'no slug')
const { data: postDetails } = usePostDetails(slug)
const {
	data: commentData,
	isFetching,
	isFetchingNextPage,
	isLoading,
	fetchNextPage,
	hasNextPage,
	refetch: refetchComments
} = usePostComments()
const comments = computed(computeComments)
const loading = ref(false)
const blankCommentProps: Omit<CommentProps, 'id'> = {
	like: 0,
	loading: true,
	createdAt: '',
	updated: false,
	replies: [],
	user: { avatarUrl: '', name: '', id: 1 },
	content: ''
}
const target = ref(null)
const targetIsVisible = useElementVisibility(target)

function computeUserAvatar() {
	return userProfile.value?.avatar
		? `${appConfig.public.mediaUrl}${userProfile.value?.avatar}`
		: `${appConfig.public.dicebearMedia}${
				userProfile.value?.name ||
				'A6Blog&backgroundColor=000000'
		  }`
}

function computeFormValidity() {
	return (
		formData.content.length !== 0 &&
		postDetails.value?.id !== undefined
	)
}

function commentToProps({
	id,
	content,
	user,
	createdAt,
	updatedAt,
	replies
}: EntitiesComment): CommentProps {
	const cProps: CommentProps = {
		id: id || 0,
		content: content || 'Error: No content',
		user: {
			id: user?.id!,
			name: user?.fullName || 'Error no fullName',
			avatarUrl:
				!user?.avatar || user.avatar === ''
					? `${appConfig.public.dicebearMedia}${
							user?.fullName ||
							'A6Blog'
					  }`
					: `${appConfig.public.mediaUrl}${user.avatar}`
		},
		like: 0,
		replies: replies?.map(commentToProps) || [],
		updated: createdAt !== updatedAt,
		createdAt: getCalendarTime(createdAt)
	}
	return cProps
}

function computeComments() {
	const comments: CommentProps[] = []
	commentData.value?.pages.forEach((page) =>
		page?.items?.forEach((c) => {
			comments.push(commentToProps(c))
		})
	)

	return comments
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
					Top comments ({{
						postDetails?.commentCount || 0
					}})
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
						v-if="!isUserProfileLoading"
						class="md:w-8"
						width="24"
						:src="userAvatar"
					/>
					<div
						class="skeleton !rounded-full"
						v-else
					>
						<div
							class="w-6 h-6 md:!w-8 md:h-8"
						/>
					</div>
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
								placeholder="Your comment..."
								reversed
							/>
						</main>
					</div>
					<div
						class="flex items-center gap-3 justify-start"
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
			v-for="comment in comments"
			v-bind="comment"
			:key="`${comment.id}${comment.content}`"
		/>
		<template v-if="isFetching || isFetchingNextPage || isLoading">
			<Comment
				v-for="n in 5"
				v-bind="blankCommentProps"
				:id="n"
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
