<script lang="ts" setup>
import { onClickOutside } from '@vueuse/core'
import { ActionIcon, Button } from 'ui-vue'
import EllipsisIcon from '~icons/mdi/ellipsis-horizontal'
import ChatIcon from '~icons/ri/chat-1-line'
import HeartIcon from '~icons/solar/heart-angle-linear'

export interface CommentProps {
	user: {
		name: string
		avatarUrl: string
		id: number
	}
	updated: boolean
	like: number
	content: string
	replies: CommentProps[]
	createdAt: string
	loading?: boolean
	id: number
}

function handleReply() {
	replyMode.value = true
}

function handleDismiss() {
	replyMode.value = false
}

function handlePreview() {
	previewMode.value = !previewMode.value
}

async function handleSubmitReply(content: string) {
	try {
		await $blogApi.comment.commentsPost({
			content,
			parentId: id,
			postId: postDetails.value?.id
		})
		handleDismiss()
		await refetchComments()
	} catch (e) {
		console.error('Error submit reply: ', e)
	}
}

const isAuth = useIsAuthenticated()
const { $blogApi } = useNuxtApp()
const { content, updated, like, replies, createdAt, id } =
	defineProps<CommentProps>()
const opened = ref(false)
const target = ref<HTMLDivElement | null>(null)
const editMode = ref(false)
const replyMode = ref(false)
const previewMode = ref(false)
const isFormValid = computed(computeFormValidity)
const formData = reactive({ content })
const formLoading = ref(false)
const slug = inject<string>('slug', 'no slug')
const { data: postDetails } = usePostDetails(slug)
const { refetch: refetchComments } = usePostComments()
const { data: userProfile } = useUserProfile()

function handleToggleDropdown() {
	opened.value = !opened.value
}

function handleCloseDropdown() {
	opened.value = false
}

function toEditMode() {
	editMode.value = true
}

function toReadonlyMode() {
	editMode.value = false
	opened.value = false
	previewMode.value = false
}

async function handleSubmitEditForm() {
	if (!isFormValid) {
		return
	}

	formLoading.value = true
	try {
		await $blogApi.comment.commentsIdPut(id, {
			content: formData.content
		})
		await refetchComments()
		formData.content = ''
	} catch (e) {
		notifyError(e)
	} finally {
		toReadonlyMode()
		formLoading.value = false
	}
}

function computeFormValidity() {
	return (
		formData.content.length !== 0 &&
		postDetails.value?.id !== undefined
	)
}

async function handleDeleteComment() {
	try {
		await $blogApi.comment.commentsIdDelete(id)
		refetchComments()
	} catch (e) {
		notifyError(e)
	}
}

onClickOutside(target, handleCloseDropdown)
</script>

<template>
	<div class="flex flex-col gap-4">
		<div class="flex gap-2">
			<div>
				<div class="skeleton" v-if="loading">
					<div class="w-6 h-6 md:!w-8 md:h-8" />
				</div>
				<Avatar
					v-else
					class="md:w-8"
					width="24"
					:src="user.avatarUrl"
				/>
			</div>
			<div class="flex flex-col gap-1 flex-1">
				<div
					class="text-neutral-900 text-base shadow rounded-md flex-1 flex flex-col gap-4 px-3 pt-2 pb-4"
				>
					<header
						class="flex items-center justify-between"
					>
						<section
							class="w-full flex items-center gap-2 text-neutral-500 text-sm"
						>
							<div
								v-if="loading"
								class="skeleton"
							>
								<br />
							</div>
							<template v-else>
								<span>{{
									user.name
								}}</span>
								<div
									class="h-1 w-1 rounded-full bg-[rgb(189,189,189)] mx-1"
								/>
								<span>{{
									createdAt
								}}</span>
								<template
									v-if="
										updated
									"
								>
									<div
										class="h-1 w-1 rounded-full bg-neutral-500"
									/>
									Edited
								</template>
							</template>
						</section>
						<div
							class="relative"
							ref="target"
							v-if="
								!editMode &&
								isAuth &&
								user.id ===
									userProfile?.id
							"
						>
							<ActionIcon
								color="gray"
								variant="subtle"
								size="md"
								v-if="!loading"
								@click="
									handleToggleDropdown
								"
							>
								<EllipsisIcon
									class="text-neutral-700"
								/>
							</ActionIcon>
							<Dropdown
								:open="opened"
							>
								<Button
									@click="
										toEditMode
									"
									color="blue"
									variant="subtle"
									block
								>
									Edit
								</Button>
								<Button
									@click="
										handleDeleteComment
									"
									color="red"
									variant="subtle"
									block
								>
									Delete
								</Button>
							</Dropdown>
						</div>
					</header>
					<main>
						<RichTextEditor
							v-if="
								editMode &&
								!previewMode
							"
							v-model="
								formData.content
							"
							reversed
						/>
						<VueMarkdown
							v-if="
								editMode &&
								previewMode
							"
							:source="
								formData.content
							"
						/>
						<template v-if="!editMode">
							<textarea
								class="skeleton resize-none"
								v-if="loading"
							/>
							<VueMarkdown
								v-else
								:source="
									content
								"
							/>
						</template>
					</main>
				</div>
				<footer
					class="flex items-center"
					:class="{ skeleton: loading }"
					v-if="!replyMode"
				>
					<div
						class="flex items-center gap-3"
						v-if="editMode"
					>
						<Button
							@click="
								handleSubmitEditForm
							"
							>Submit</Button
						>
						<Button
							@click="handlePreview"
							color="gray"
							>{{
								previewMode
									? 'Continue'
									: 'Preview'
							}}</Button
						>
						<Button
							@click="toReadonlyMode"
							color="gray"
							variant="subtle"
						>
							<span class="text-black"
								>Dismiss</span
							>
						</Button>
					</div>
					<template v-else>
						<div
							v-if="loading"
							class="py-[5px]"
						>
							<br />
						</div>
						<template v-else>
							<Button
								color="gray"
								variant="subtle"
								size="xs"
							>
								<template
									#leftIcon
								>
									<HeartIcon
										class="text-neutral-700"
									/>
								</template>
								<span
									class="text-sm font-normal text-neutral-700"
								>
									{{
										like
									}}
									<span
										class="hidden text-sm font-normal text-neutral-700 sm:inline"
										>reactions</span
									>
								</span>
							</Button>
							<Button
								@click="
									handleReply
								"
								color="gray"
								variant="subtle"
								size="xs"
							>
								<template
									#leftIcon
								>
									<ChatIcon
										class="text-neutral-800"
									/>
								</template>
								<span
									class="text-neutral-800 font-normal hidden md:inline"
								>
									Reply
								</span>
							</Button>
						</template>
					</template>
				</footer>
				<CommentForm
					class="mt-3"
					v-else
					:onDismiss="handleDismiss"
					:onSubmit="handleSubmitReply"
				/>
			</div>
		</div>
	</div>
	<div class="pl-3 flex flex-col gap-6" v-if="!loading">
		<Comment
			v-for="(comment, i) in replies"
			:key="i"
			v-bind="comment"
		/>
	</div>
</template>
