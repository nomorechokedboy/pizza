<script lang="ts" setup>
import { computed, InputHTMLAttributes, ref } from 'vue'

export interface FileButtonProps<Multiple extends boolean = false> {
	onChange: (payload: Multiple extends true ? File[] : File) => void
	accept?: string
	disabled?: boolean
	multiple?: Multiple
	name?: string
	inputProps?: Omit<
		InputHTMLAttributes,
		'accept' | 'multiple' | 'name' | 'disabled'
	>
	form?: string
	capture?: boolean | 'user' | 'environment'
}

const {
	onChange,
	accept,
	multiple,
	name,
	inputProps,
	disabled,
	capture,
	form
} = defineProps<FileButtonProps<boolean>>()
const input = ref<HTMLInputElement | null>(null)
const handleClick = () => {
	!disabled && input.value?.click()
}
const handleChange = computed<InputHTMLAttributes['onChange']>(() => (e) => {
	const target = e.currentTarget as HTMLInputElement
	if (target && target.files) {
		if (multiple || (multiple as unknown as string) === '') {
			onChange(Array.from(target.files) as any)
		} else {
			onChange(target.files[0])
		}
	}
})
const reset = () => {
	if (!input.value) {
		return
	}

	input.value.value = ''
}
defineExpose({ reset })
</script>

<template>
	<input
		:multiple="multiple"
		:accept="accept"
		:name="name"
		:capture="capture"
		:form="form"
		@change="handleChange"
		v-show="false"
		type="file"
		ref="input"
		v-bind="inputProps"
	/>
	<slot :onClick="handleClick" />
</template>
