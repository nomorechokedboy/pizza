<script lang="ts" setup>
import dayjs from 'dayjs'

export interface NotificationItemProps {
	avatar?: string
	actionType: string
	userName: string
	fullName?: string | null
	loading?: boolean
	readAt?: string | null
	createdAt: string
	id: number
}

async function handleReadNotification() {
	if (loading !== false && readAt !== null) {
		return
	}

	try {
		await $blogApi.notification.readAt(id)
		refetch()
	} catch (e) {
		console.error({ e })
	}
}

const { id, loading, readAt } = defineProps<NotificationItemProps>()
const { $blogApi } = useNuxtApp()
const appConfig = useRuntimeConfig()
const { refetch } = useNotificationPagination()
</script>

<template>
	<button
		class="flex items-center gap-3 p-2 relative rounded-lg hover:bg-blue-50"
		:class="{ 'bg-blue-50': readAt === null && loading == false }"
		@click="handleReadNotification"
	>
		<div :class="{ 'h-8': loading }">
			<div class="skeleton rounded-full" v-if="loading">
				<div class="w-8 h-8" />
			</div>
			<Avatar
				v-else
				width="32"
				:src="
					avatar ||
					`${appConfig.public.dicebearMedia}${
						fullName || userName
					}`
				"
				:alt="`${fullName || userName} avatar`"
			/>
		</div>
		<div class="flex flex-col flex-1 gap-1 pr-2 text-start">
			<div :class="{ skeleton: loading }">
				<br v-if="loading" />
				<template v-else>
					<span
						class="font-bold text-black capitalize"
					>
						{{ fullName || userName }}
					</span>
					<span>{{ ' ' + actionType }}</span>
				</template>
			</div>
			<span class="text-sm text-slate-500">{{
				dayjs(createdAt).fromNow()
			}}</span>
		</div>
	</button>
</template>
