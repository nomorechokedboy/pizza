<!-- <script lang="ts" setup>
import { Codemirror as BaseCodemirror } from 'vue-codemirror'
import { markdown } from '@codemirror/lang-markdown'
import { vim } from "@replit/codemirror-vim"
import { emacs } from '@replit/codemirror-emacs'
import { ViewUpdate } from '@codemirror/view'

export interface CodemirrorProps {
    keymap: 'vim' | 'emacs'
    modelValue?: string
}
const emit = defineEmits(['update:modelValue'])
const codemirror = ref<any>(null)

function getKeyMapExt(): typeof vim | typeof emacs {
    return keymap === 'emacs' ? emacs : vim
}
function handleUpdateModelValue(value: string, viewUpdate: ViewUpdate) {
    emit('update:modelValue', value, viewUpdate)
}

const { keymap, modelValue } = defineProps<CodemirrorProps>()
const keymapExt = computed(getKeyMapExt)
const extensions = [markdown(), keymapExt.value?.()]

const view = shallowRef<import("@codemirror/view").EditorView>()
const handleReady = (payload: {
    view: import("@codemirror/view").EditorView;
    state: import("@codemirror/state").EditorState;
    container: HTMLDivElement;
}) => {
    view.value = payload.view
}
</script>

<template>
    <BaseCodemirror :model-value="modelValue" @update:model-value="handleUpdateModelValue"
        placeholder="Code goes here..." :style="{ height: '100vh', }" :autofocus="true" :indent-with-tab="true"
        :tab-size="2" :extensions="extensions" @ready="handleReady" ref="codemirror" />
</template> -->

<template>
    <div>

    </div>
</template>
