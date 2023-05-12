<script setup lang="ts">
import { useInfiniteQuery } from '@tanstack/vue-query'
import { onClickOutside } from '@vueuse/core'
import { ActionIcon, Button } from 'ui-vue'
import IconBell from '~icons/ph/bell-simple'

const PAGE_SIZE = 10

function handleToggle() {
	toggle.value = !toggle.value
}

async function fetchNotifications({ pageParam = 1 }) {
	return $blogApi.notification
		.getNotifications(pageParam, PAGE_SIZE)
		.then((res) => res.data)
}

const toggle = ref(false)
const userProfile = useUserProfile()
const target = ref<HTMLDivElement | null>(null)
const { $blogApi } = useNuxtApp()
const {
	data: notifications,
	isFetching,
	isFetchingNextPage,
	isLoading,
	fetchNextPage,
	hasNextPage
} = useInfiniteQuery({
	queryKey: ['notifications'],
	queryFn: fetchNotifications,
	getNextPageParam: (lastPage) =>
		lastPage.page && lastPage.data.length === PAGE_SIZE
			? lastPage.page + 1
			: undefined
})
const unread = computed(() => {
	let count = 0
	if (!notifications.value) {
		return count
	}

	for (let i = 0; i < notifications.value.pages.length; i++) {
		const notification = notifications.value.pages[0].data[i]
		if (notification.notifications[0].readAt === null) {
			count++
		}
	}

	return count
})

onClickOutside(target, () => (toggle.value = false))
</script>

<template>
	<div class="relative inline-block" ref="target">
		<span class="sr-only">Notifications</span>
		<span
			class="absolute top-0 right-0 flex w-3 h-3 bg-red-500 rounded-full"
			v-if="unread > 0"
		></span>
		<ActionIcon
			color="indigo"
			class="group"
			size="lg"
			variant="subtle"
			@click="handleToggle"
		>
			<span class="group-hover:text-indigo-500 text-black">
				<IconBell />
			</span>
		</ActionIcon>
		<Dropdown
			class="h-[90vh] overflow-y-auto px-2 flex flex-col gap-2"
			:open="toggle"
		>
			<div class="py-2 px-4">
				<h3 class="font-bold text-black text-2xl">
					Notifications
				</h3>
			</div>
			<template v-for="page in notifications?.pages">
				<template
					v-for="notification in page.data"
					:key="notification.id"
				>
					<NotificationItem
						:avatar="
							notification
								.notificationChange
								.actor.avatar
						"
						:actionType="
							notification.actionType
						"
						:createdAt="
							notification.createdAt
						"
						:fullName="
							notification
								.notificationChange
								.actor.fullName
						"
						:userName="
							notification
								.notificationChange
								.actor.userName
						"
						:readAt="
							notification.notifications.find(
								(n) =>
									n
										.notifier
										.id ===
									userProfile.id
							)?.readAt
						"
					/>
				</template>
			</template>
			<NotificationItem
				v-if="
					isFetching ||
					isFetchingNextPage ||
					isLoading
				"
				v-for="n in 3"
				actionType=""
				createdAt=""
				:loading="true"
				userName=""
			/>
			<template v-if="notifications?.pages.length === 0">
				<div>You have no notification</div>
			</template>
			<div class="pt-3 w-full grid place-items-center">
				<Button
					v-if="hasNextPage"
					:disabled="
						isFetching ||
						isFetchingNextPage ||
						isLoading
					"
					@click="fetchNextPage"
					>Load more</Button
				>
			</div>
		</Dropdown>
	</div>
</template>
