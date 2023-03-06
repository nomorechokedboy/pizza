<script lang="ts" setup>
import { computed, ref } from 'vue'

export interface CopyButtonProps {
	value: string
	timeout?: number
}

const { value, timeout } = defineProps<CopyButtonProps>()
const copied = ref(false)
const copyToClipBoard = computed(() => () => {
	navigator.clipboard.writeText(value)
	copied.value = true
	setTimeout(() => {
		copied.value = false
	}, timeout ?? 500)
})
</script>

<template>
	<slot :copied="copied" :copy="copyToClipBoard" />
</template>
