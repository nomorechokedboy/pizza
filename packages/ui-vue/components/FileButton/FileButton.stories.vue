<script lang="ts" setup>
import Button from '$/Button/Button.vue'
import { reactive, ref } from 'vue'
import FileButton, { FileButtonProps } from './FileButton.vue'

const file = ref<File | null>(null)
const files = ref<File[] | null>(null)
const resetFile = ref<File | null>(null)
const state = reactive<
	Omit<FileButtonProps, 'onChange'> & { disabled?: boolean }
>({
	disabled: false,
	multiple: false
})
const fileButtonRef = ref<InstanceType<typeof FileButton> | null>(null)
const handleFile = (e: File | File[]) => {
	if (Array.isArray(e)) {
		file.value = e[0]
	} else {
		file.value = e
	}
}
const handleFiles = (e: File | File[]) => {
	if (!Array.isArray(e)) return

	files.value = e
}
const clearFile = () => {
	resetFile.value = null
	fileButtonRef.value?.reset()
}
</script>

<template>
	<Story>
		<Variant title="Usage" auto-props-disabled>
			<template #controls>
				<HstCheckbox
					v-model="state.disabled"
					title="Disabled"
				/>
				<HstText
					v-model="state.accept"
					title="Accept"
				/>
			</template>
			<template #default>
				<FileButton
					:on-change="handleFile"
					:disabled="state.disabled"
					:accept="state.accept"
					v-slot="slotProps"
				>
					<Button :="slotProps"
						>Upload image</Button
					>
					<div v-if="file" class="text-dark-0">
						Picked file: {{ file?.name }}
					</div>
				</FileButton>
			</template>
		</Variant>
		<Variant title="Multiple" auto-props-disabled>
			<template #default>
				<FileButton
					:on-change="handleFiles"
					v-slot="slotProps"
				>
					<Button :="slotProps"
						>Upload multiple</Button
					>
				</FileButton>
				<template v-if="files">
					<div
						v-for="file in files"
						class="text-dark-0"
					>
						Picked file: {{ file?.name }}
					</div>
				</template>
			</template>
		</Variant>
		<Variant title="Reset" auto-props-disabled>
			<template #default>
				<FileButton
					:on-change="
						(e) => {
							if (Array.isArray(e))
								return
							resetFile = e
						}
					"
					v-slot="slotProps"
				>
					<Button :="slotProps"
						>Upload file</Button
					>
					<Button
						:disabled="!resetFile"
						color="red"
						@click="clearFile"
					>
						Reset
					</Button>
				</FileButton>
				<div v-if="resetFile" class="text-dark-0">
					Picked file: {{ resetFile?.name }}
				</div>
			</template>
		</Variant>
	</Story>
</template>
