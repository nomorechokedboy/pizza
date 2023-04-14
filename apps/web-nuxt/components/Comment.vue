<script lang="ts" setup>
import { Button } from 'ui-vue'
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
	loading?: boolean
}

const { content, updated, like, replies, createdAt } =
	defineProps<CommentProps>()
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
						<!-- <ActionIcon
							color="indigo"
							variant="subtle"
						>
							<EllipsisIcon
								class="text-neutral-700"
							/>
						</ActionIcon> -->
					</header>
					<main>
						<textarea
							class="skeleton resize-none"
							v-if="loading"
						/>
						<VueMarkdown
							v-else
							:source="content"
						/>
					</main>
				</div>
				<footer
					class="flex items-center"
					:class="{ skeleton: loading }"
				>
					<div v-if="loading" class="py-[5px]">
						<br />
					</div>
					<template v-else>
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
					</template>
				</footer>
			</div>
		</div>
	</div>
	<div class="pl-3" v-if="!loading">
		<Comment
			v-for="(comment, i) in replies"
			:key="i"
			v-bind="comment"
		/>
	</div>
</template>
