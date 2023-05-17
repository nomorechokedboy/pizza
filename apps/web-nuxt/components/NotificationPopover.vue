<script setup lang="ts">
import {
	EventStreamContentType,
	fetchEventSource
} from '@microsoft/fetch-event-source'
import { onClickOutside } from '@vueuse/core'
import { ActionIcon, Button } from 'ui-vue'
import IconBell from '~icons/ph/bell-simple'

function handleToggle() {
	toggle.value = !toggle.value
}

const toggle = ref(false)
const { data: userProfile } = useUserProfile()
const target = ref<HTMLDivElement | null>(null)
const { $blogApi } = useNuxtApp()
const {
	data: notifications,
	isFetching,
	isFetchingNextPage,
	isLoading,
	fetchNextPage,
	hasNextPage,
	refetch
} = useNotificationPagination()
const unread = computed(() => {
	let count = 0
	if (!notifications.value?.pages) {
		return count
	}

	for (let i = 0; i < notifications.value.pages.length; i++) {
		const notification = notifications.value.pages.at(0)?.data[i]
		if (notification?.notifications?.[0].readAt === null) {
			count++
		}
	}

	return count
})
const config = useRuntimeConfig()
const token = useAuthToken()
const refreshToken = useRefreshToken()
const isLoggedIn = useIsAuthenticated()
const isEmpty = computed(() => notifications.value?.pages[0].data.length === 0)
const notifyController = inject<AbortController>('notifyController')

onClickOutside(target, () => (toggle.value = false))
watchEffect(() => {
	if (!isLoggedIn.value) {
		return
	}

	fetchEventSource(config.public.notifyUrl, {
		signal: notifyController?.signal,
		openWhenHidden: true,
		headers: {
			Authorization: `Bearer ${token.value}`
		},
		async onopen(response) {
			if (
				response.ok &&
				response.headers.get('content-type') ===
					EventStreamContentType
			) {
				return // everything's good
			} else if (response.status === 401) {
				if (!token.value || !refreshToken.value) {
					throw Error('Fatal')
				}

				try {
					const isExpired = isTokenExpired(
						token.value
					)
					if (isExpired) {
						$blogApi.auth
							.authRefreshTokenPost({
								refresh_token:
									refreshToken.value
							})
							.then((resp) => {
								onRefreshToken(
									resp
								)
								throw Error(
									'Retry'
								)
							})
							.catch(() => {
								throw Error(
									'Fatal'
								)
							})
					}
				} catch (err) {
					console.error({ err })
					throw err
				}
			} else if (
				response.status !== 401 &&
				response.status >= 400 &&
				response.status < 500 &&
				response.status !== 429
			) {
				// client-side errors are usually non-retriable:
				throw Error('Fatal')
			} else {
				throw Error('Retry')
			}
		},
		onmessage(msg) {
			// if the server emits an error message, throw an exception
			// so it gets handled by the onerror callback below:
			if (msg.event === 'notification') {
				refetch()
			}
		},
		onclose() {
			// if the server closes the connection unexpectedly, retry:
			throw Error('Retry')
		},
		onerror(err) {
			if (err === Error('Fatal')) {
				throw err // rethrow to stop the operation
			} else {
				// do nothing to automatically retry. You can also
				// return a specific retry interval here.
			}
		}
	})
})
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
			class="max-h-[90vh] overflow-y-auto px-2 flex flex-col gap-2"
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
									userProfile?.id
							)?.readAt
						"
						:id="notification.id"
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
				:id="n"
				actionType=""
				createdAt=""
				:loading="true"
				userName=""
			/>
			<template v-if="isEmpty">
				<div>You have no notification!</div>
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
