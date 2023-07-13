<script lang="ts" setup>
import { Button } from 'ui-vue'

export interface CommentFormProps {
	loading?: boolean
	onDismiss: () => void
	onSubmit: (content: string) => void
}

function handlePreview() {
	previewMode.value = !previewMode.value
}

defineProps<CommentFormProps>()
defineEmits(['update:modelValue'])
const formData = reactive({ content: '' })
const previewMode = ref(false)
</script>

<template>
	<form class="flex flex-col gap-3 flex-1">
		<div
			class="text-neutral-900 text-base shadow rounded-md flex-1 flex flex-col gap-4 px-3 pt-2 pb-4"
		>
			<main class="min-h-[134px]">
				<Markdown
					v-if="previewMode"
					:source="formData.content"
					:key="formData.content"
					class="markdown"
				/>
				<RichTextEditor
					v-else
					v-model="formData.content"
					placeholder="Reply..."
					reversed
				/>
			</main>
		</div>
		<footer
			class="flex items-center"
			:class="{ skeleton: loading }"
		>
			<div class="flex items-center gap-3">
				<Button
					@click.prevent="
						onSubmit(formData.content)
					"
					:disabled="
						loading ||
						formData.content.length === 0
					"
					type="submit"
					>Reply</Button
				>
				<Button
					@click.prevent="handlePreview"
					color="gray"
					:disabled="
						loading ||
						formData.content.length === 0
					"
					>{{
						previewMode
							? 'Continue'
							: 'Preview'
					}}</Button
				>
				<Button
					@click="onDismiss"
					color="gray"
					:disabled="loading"
					variant="subtle"
				>
					<span class="text-black">Dismiss</span>
				</Button>
			</div>
		</footer>
	</form>
</template>
