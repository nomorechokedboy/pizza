<script setup lang="ts">
import { Button } from 'ui-vue'

const config = useRuntimeConfig()
const token = useAuthToken()
const refreshToken = useRefreshToken()
const isLoggedIn = computed(() => !!token.value)
const { $blogApi } = useNuxtApp()
const notificationEventSource = useNotificationEventSource()

watchEffect((onStop) => {
	if (
		!isLoggedIn.value ||
		notificationEventSource.value ||
		!process.client
	) {
		return
	}

	notificationEventSource.value = new EventSource(
		config.public.notifyUrl,
		{ withCredentials: true }
	)
	notificationEventSource.value?.addEventListener('notification', (e) => {
		console.debug(e.data)
	})
	notificationEventSource.value.onerror = () => {
		if (!token.value || !refreshToken.value) {
			cleanupNotificationEventSource()
			return
		}

		try {
			const isExpired = isTokenExpired(token.value)
			if (isExpired) {
				$blogApi.auth
					.authRefreshTokenPost({
						refresh_token:
							refreshToken.value
					})
					.then((resp) => {
						onRefreshToken(resp)
					})
					.catch(cleanupNotificationEventSource)
			}
		} catch (err) {
			console.error({ err })
		}
	}

	onStop(cleanupNotificationEventSource)
})
</script>

<template>
	<div class="grid place-items-center p-2.5 bg-white shadow-md lg:px-8">
		<div
			class="max-w-7xl w-full flex flex-row items-center justify-between px-2"
		>
			<div class="flex items-center gap-4 lg:w-1/2">
				<NuxtLink to="/">
					<nuxt-img
						alt="Accessiblog logo"
						src="/logo.png"
						width="50"
						height="40"
					/>
				</NuxtLink>
				<div
					class="bg-transparent h-10 lg:w-full lg:max-w-md"
				></div>
			</div>
			<div class="flex flex-row gap-2 md:gap-5">
				<NuxtLink v-if="isLoggedIn" to="/new">
					<Button color="indigo"
						>Create Post</Button
					>
				</NuxtLink>
				<NotificationPopover v-if="isLoggedIn" />
				<UserPopover v-if="isLoggedIn" />
				<NuxtLink
					v-if="!isLoggedIn"
					class="hidden md:inline"
					to="/login"
				>
					<Button
						variant="subtle"
						color="indigo"
						class="text-black group"
					>
						<span
							class="text-neutral-700 font-normal text-base group-hover:underline group-hover:text-indigo-500"
							>Log in</span
						>
					</Button>
				</NuxtLink>
				<NuxtLink v-if="!isLoggedIn" to="/signup">
					<Button color="indigo"
						>Create account</Button
					>
				</NuxtLink>
			</div>
		</div>
	</div>
</template>
