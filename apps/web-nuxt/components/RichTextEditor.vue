<script lang="ts" setup>
import { ActionIcon } from 'ui-vue'
import CodeBlockIcon from '~icons/material-symbols/code-blocks-rounded'
import CodeIcon from '~icons/mdi/code-tags'
import BoldIcon from '~icons/mdi/format-bold'
import ItalicIcon from '~icons/mdi/format-italic'
import StrikethroughIcon from '~icons/mdi/format-strikethrough'
import UnderlineIcon from '~icons/mdi/format-underline'

export interface RichTextEditorProps {
	modelValue?: string
	reversed?: boolean
}

/* function handleScroll(event: any) {
    console.log({ event })
    const target = event.target
    const other: any =
        target === editor.value ? preview.value : editor.value
    other.scrollTop = target.scrollTop
} */

function wrapSelection(startTag: string, endTag: string = startTag) {
	if (!input.value) {
		return
	}

	const start = input.value.selectionStart
	const end = input.value.selectionEnd
	const text = input.value.value
	const newText =
		text.slice(0, start) +
		startTag +
		text.slice(start, end) +
		endTag +
		text.slice(end)

	emit('update:modelValue', newText)
	input.value.setSelectionRange(
		start + startTag.length,
		end + startTag.length
	)
	input.value.focus()
}

function bold() {
	wrapSelection('**')
}

function italic() {
	wrapSelection('_')
}

function underline() {
	wrapSelection('<u>', '</u>')
}

function strikethrough() {
	wrapSelection('~~')
}

function code() {
	wrapSelection('`')
}

function codeBlock() {
	wrapSelection('```\n', '\n```')
}

const input = ref<HTMLTextAreaElement | null>(null)
const emit = defineEmits(['update:modelValue'])
const { modelValue } = defineProps<RichTextEditorProps>()
const icons = [
	{ icon: BoldIcon, handler: bold },
	{ icon: ItalicIcon, handler: italic },
	{ icon: UnderlineIcon, handler: underline },
	{ icon: StrikethroughIcon, handler: strikethrough },
	{ icon: CodeIcon, handler: code },
	{ icon: CodeBlockIcon, handler: codeBlock }
]
</script>

<template>
	<div
		class="flex-1 flex gap-2 max-w-1/2 overflow-auto break-words bg-white rounded-l"
		:class="{ 'flex-col': !reversed, 'flex-col-reverse': reversed }"
		ref="editor"
	>
		<div class="flex items-center gap-3 p-2 bg-neutral-300">
			<ActionIcon
				v-for="{ handler, icon } of icons"
				@click.prevent="handler"
				color="gray"
				variant="subtle"
			>
				<span class="text-black">
					<component :is="icon" />
				</span>
			</ActionIcon>
		</div>
		<textarea
			class="rounded-l-md flex-1 w-full h-full focus:outline-none resize-none p-2"
			:value="modelValue"
			@keydown.ctrl.b.prevent="bold"
			@keydown.ctrl.i.prevent="italic"
			@keydown.ctrl.u.prevent="underline"
			@input="
				$emit(
					'update:modelValue',
					($event.target as HTMLInputElement)
						.value
				)
			"
			placeholder="Post content here..."
			ref="input"
		/>
	</div>
</template>

<style scoped>
.toolbar {
	display: flex;
	gap: 10px;
	margin-bottom: 10px;
}

.markdown {
	overflow-x: auto;
	overflow-wrap: break-word;
}

.markdown > p {
	width: 800px;
	word-wrap: break-word;
}

.preview::-webkit-scrollbar {
	display: none;
}

.preview {
	-ms-overflow-style: none;
	scrollbar-width: none;
}
</style>
