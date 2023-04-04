<script setup lang="ts">
import { ErrorObject } from '@vuelidate/core'

export interface TextInputProps {
	id?: string
	modelValue?: string
	errors?: ErrorObject[]
	error?: boolean
}

const inputRef = ref<HTMLInputElement | null>(null)
const focus = () => {
	inputRef.value?.focus()
}
const { id, errors, error } = defineProps<TextInputProps>()
defineEmits(['update:modelValue'])
defineExpose({
	focus
})
</script>

<template>
	<div class="flex flex-col gap-1">
		<div class="flex flex-col gap-2">
			<label :for="id">
				<slot name="label" />
			</label>
			<input
				:class="{
					'border-red-500 focus:border-red-500':
						error
				}"
				class="border rounded-md py-[6.5px] px-2 focus-visible:outline-none focus:border-2 focus:border-blue-700"
				:="$attrs"
				:value="modelValue"
				@input="
					$emit(
						'update:modelValue',
						(
							$event.target as HTMLInputElement
						).value
					)
				"
				:id="id"
				ref="inputRef"
			/>
		</div>
		<div class="text-xs text-red-500">
			<br v-if="!error" />
			<template v-for="error of errors" :key="error.$uid">
				{{ error.$message }}
			</template>
		</div>
	</div>
</template>
