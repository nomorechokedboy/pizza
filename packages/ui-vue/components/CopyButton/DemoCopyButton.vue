<script lang="ts" setup>
import Button from '../Button/Button.vue'

export interface DemoCopyButtonProps {
	copied: boolean
	copy: () => void
}

const { copied, copy } = defineProps<DemoCopyButtonProps>()
</script>

<template>
	<Button :color="copied ? 'teal' : 'blue'" @click="copy">
		{{ copied ? 'Copied url' : 'Copy url' }}
	</Button>
</template>
