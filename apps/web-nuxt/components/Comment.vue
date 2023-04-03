<script lang="ts" setup>
import VueMarkdown from 'markdown-vue'
import { ActionIcon, Button } from 'ui-vue'
import EllipsisIcon from '~icons/mdi/ellipsis-horizontal'
import ChatIcon from '~icons/ri/chat-1-line'
import HeartIcon from '~icons/solar/heart-angle-linear'

export interface CommentProps {
	user: {
		name: string
		avatarUrl: string
	}
	updated: boolean
	like: number
	content: string
	replies: CommentProps[]
	createdAt: string
}

const { content, updated, like, replies, createdAt } =
	defineProps<CommentProps>()
</script>

<template>
	<div class="flex flex-col gap-4">
		<div class="flex gap-2">
			<div class="w-full max-w-[32px]">
				<Avatar
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
							class="flex items-center gap-2 text-neutral-500 text-sm"
						>
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
								v-if="updated"
							>
								<div
									class="h-1 w-1 rounded-full bg-neutral-500"
								/>
								Edited
							</template>
						</section>
						<ActionIcon
							color="indigo"
							variant="subtle"
						>
							<EllipsisIcon
								class="text-neutral-700"
							/>
						</ActionIcon>
					</header>
					<main>
						<VueMarkdown
							:source="content"
						/>
					</main>
				</div>
				<footer class="flex items-center">
					<Button
						color="gray"
						variant="subtle"
						size="xs"
					>
						<template #leftIcon>
							<HeartIcon
								class="text-neutral-700"
							/>
						</template>
						<span
							class="text-sm font-normal text-neutral-700"
						>
							{{ like }}
							<span
								class="hidden text-sm font-normal text-neutral-700 sm:inline"
								>reactions</span
							>
						</span>
					</Button>
					<Button
						color="gray"
						variant="subtle"
						size="xs"
					>
						<template #leftIcon>
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
				</footer>
			</div>
		</div>
	</div>
	<div class="pl-3">
		<Comment
			v-for="(comment, i) in replies"
			:key="i"
			v-bind="comment"
		/>
	</div>
</template>
