<script lang="ts" setup>
import { Button } from 'ui-vue'
import ChatIcon from '~icons/ri/chat-1-line'
import HeartIcon from '~icons/solar/heart-angle-bold'

export interface ArticleFooterProps {
	like: number
	comment: string
}

const { comment, like } = defineProps<ArticleFooterProps>()
</script>

<template>
	<div class="flex items-center justify-between">
		<div class="flex items-center gap-7 py-2">
			<Button v-if="like !== 0" color="gray" variant="subtle">
				<template #leftIcon>
					<HeartIcon />
					<span
						class="text-sm font-normal text-neutral-700"
					>
						{{ like }}
						<span
							class="hidden text-sm font-normal text-neutral-700 sm:inline"
							>reactions</span
						>
					</span>
				</template>
			</Button>
			<Button
				class="!text-neutral-800"
				color="gray"
				variant="subtle"
			>
				<ChatIcon />
				<span
					class="hidden text-sm font-normal text-neutral-700 sm:inline"
				>
					{{ comment }}
				</span>
			</Button>
		</div>
		<!-- <div class="flex items-center gap-2">
			<p class="text-xs text-neutral-600">3 min read</p>
			<ActionIcon
				class="!text-neutral-800"
				color="indigo"
				variant="subtle"
				size="xl"
			>
				<BookmarkIcon width="24" height="24" />
			</ActionIcon>
		</div> -->
	</div>
</template>
