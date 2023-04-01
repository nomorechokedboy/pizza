<script setup lang="ts">
import { ActionIcon, ActionIconProps } from 'ui-vue'

export interface AlertProps {
	type: 'success' | 'error' | 'warning'
	content: string
	key: string
}

const notificationList = useNotification()
const { content, type, key } = defineProps<AlertProps>()
const iconColor: ActionIconProps['color'] =
	type === 'error' ? 'red' : type === 'warning' ? 'yellow' : 'green'
const timeout = ref<NodeJS.Timeout>()
function handleClose() {
	notificationList.value = notificationList.value.filter(
		(n) => n.key !== key
	)
}
function handleAutoClose() {
	timeout.value = setTimeout(handleClose, 5000)
}
function cancelDelay() {
	if (timeout.value) {
		clearTimeout(timeout.value)
	}
}

onMounted(handleAutoClose)
onUnmounted(cancelDelay)
</script>

<template>
	<div
		@mouseover="cancelDelay"
		@mouseleave="handleAutoClose"
		id="alert-2"
		class="flex items-center gap-3 p-4 rounded-lg bg-red-50 dark:bg-gray-800"
		:class="[
			{ 'text-red-800 dark:text-red-400': type === 'error' },
			{
				'text-green-800 dark:text-green-400':
					type === 'success'
			},
			{
				'text-yellow-800 dark:bg-yellow-400':
					type === 'warning'
			}
		]"
		role="alert"
	>
		<svg
			aria-hidden="true"
			class="flex-shrink-0 w-5 h-5"
			fill="currentColor"
			viewBox="0 0 20 20"
			xmlns="http://www.w3.org/2000/svg"
		>
			<path
				fill-rule="evenodd"
				d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7-4a1 1 0 11-2 0 1 1 0 012 0zM9 9a1 1 0 000 2v3a1 1 0 001 1h1a1 1 0 100-2v-3a1 1 0 00-1-1H9z"
				clip-rule="evenodd"
			></path>
		</svg>
		<span class="sr-only">Info</span>
		<div class="text-sm font-medium">
			{{ content }}
		</div>
		<ActionIcon
			variant="transparent"
			:color="iconColor"
			data-dismiss-target="#alert-2"
			aria-label="Close"
			@click="handleClose"
		>
			<span class="sr-only">Close</span>
			<svg
				class="w-5 h-5"
				fill="currentColor"
				viewBox="0 0 20 20"
				xmlns="http://www.w3.org/2000/svg"
			>
				<path
					fill-rule="evenodd"
					d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z"
					clip-rule="evenodd"
				></path>
			</svg>
		</ActionIcon>
	</div>
</template>
