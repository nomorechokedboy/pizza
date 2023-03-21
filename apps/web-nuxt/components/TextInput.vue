<script setup lang="ts">
export interface TextInputProps {
	id?: string
	modelValue?: string
}

const { id } = defineProps<TextInputProps>()
defineEmits(['update:modelValue'])
</script>

<template>
	<div class="flex flex-col gap-2">
		<label :for="id">
			<slot name="label" />
		</label>
		<input
			class="border rounded-md py-[6.5px] px-2 focus-visible:outline-none focus:border-2 focus:border-blue-700"
			:="$attrs"
			:value="modelValue"
			@input="
				$emit(
					'update:modelValue',
					($event.target as HTMLInputElement)
						.value
				)
			"
			:id="id"
		/>
	</div>
</template>
