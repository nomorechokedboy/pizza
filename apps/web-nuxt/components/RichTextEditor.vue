<script lang="ts" setup>
export interface RichTextEditorProps {
	modelValue?: string
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

const input = ref<HTMLTextAreaElement | null>(null)
const emit = defineEmits(['update:modelValue'])
const { modelValue } = defineProps<RichTextEditorProps>()
</script>

<template>
	<div
		class="flex-1 flex flex-col gap-2 max-w-1/2 h-full overflow-auto break-words bg-white rounded-l"
		ref="editor"
	>
		<!-- <div class="pt-8">
                    <Codemirror class="codemirror" v-model="content" keymap="vim" />
                </div> -->
		<div class="flex items-center gap-2 p-2 bg-neutral-300">
			<button @click.prevent="bold">Bold</button>
			<button @click.prevent="italic">Italic</button>
			<button @click.prevent="underline">Underline</button>
			<button @click.prevent="strikethrough">
				Strikethrough
			</button>
			<button @click.prevent="code">Code</button>
		</div>
		<textarea
			class="rounded-l-md w-full h-full focus:outline-none resize-none p-2"
			:value="modelValue"
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
