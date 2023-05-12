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
}

defineProps<NotificationItemProps>()
const appConfig = useRuntimeConfig()
</script>

<template>
	<div
		class="flex items-center gap-3 p-2 relative rounded-lg hover:bg-blue-50"
		:class="{ 'bg-blue-50': readAt === null && loading == false }"
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
		<div class="flex flex-col gap-1 pr-2">
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
	</div>
</template>
