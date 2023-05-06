<script setup lang="ts">
import { fetchEventSource } from '@microsoft/fetch-event-source'
import { ActionIcon, Button } from 'ui-vue'
import { dicebearMedia, notifyUrl } from '~/constants'
import IconBell from '~icons/ph/bell-simple'
import Avatar from './Avatar.vue'
import Dropdown from './Dropdown.vue'

function handleToggle() {
	toggle.value = !toggle.value
}

function handleCleanupSSE() {
	eventSource.value?.close()
}

const token = useAuthToken()
const isLoggedIn = computed(
	() => token.value.accessToken && token.value.refreshToken
)
const { $blogApi } = useNuxtApp()
const toggle = ref(false)
const userProfile = useUserProfile()
const userAvatar = computed(
	() =>
		userProfile.value?.avatar ||
		`${dicebearMedia}${userProfile.value.name}`
)
const isSSEConnected = ref(false)
watchEffect((onStop) => {
	if (!isLoggedIn.value) {
		return
	}

	$blogApi.auth.authMeGet().then(({ data }) => {
		if (data) {
			const { avatar, username, id, email, fullname } = data
			setUserProfile({
				name: fullname,
				email,
				id,
				username,
				avatar
			})
		}
	})

	if (process.client && userProfile.value.id && !isSSEConnected.value) {
		fetchEventSource(`${notifyUrl}/${userProfile.value.id}`, {
			headers: {
				Authorization: `Bearer ${token.value.accessToken}`
			},
			onmessage: (ev: unknown) => {
				console.log(ev)
			},
			onopen: async () => {
				isSSEConnected.value = true
			}
		})
	}

	onStop(handleCleanupSSE)
})
</script>

<template>
	<div class="grid place-items-center p-2.5 bg-white shadow-md lg:px-8">
		<div
			class="max-w-7xl w-full flex flex-row items-center justify-between px-2"
		>
			<div class="flex items-center gap-4 w-1/2">
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
				<ActionIcon
					color="indigo"
					class="group"
					size="lg"
					variant="subtle"
					v-if="isLoggedIn"
				>
					<span
						class="group-hover:text-indigo-500 text-black"
					>
						<IconBell />
					</span>
				</ActionIcon>
				<div
					class="relative inline-block"
					v-if="isLoggedIn"
				>
					<ActionIcon
						radius="xl"
						size="lg"
						variant="subtle"
						@click="handleToggle"
						class="focus:ring-4 focus:outline-none focus:ring-gray-300"
					>
						<Avatar
							width="32"
							:src="userAvatar"
						/>
					</ActionIcon>
					<Dropdown
						:open="toggle"
						:user="{
							name:
								userProfile?.name ||
								'No username',
							username:
								userProfile?.username ||
								'No username'
						}"
					/>
				</div>
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
