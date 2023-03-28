<script setup lang="ts">
import { ActionIcon, Button } from 'ui-vue'
import IconBell from '~icons/ph/bell-simple'

const token = useAuthToken()
const isLoggedIn = computed(
	() => token.value.accessToken && token.value.refreshToken
)
const { $blogApi } = useNuxtApp()

watchEffect(() => {
	if (isLoggedIn.value) {
		$blogApi.auth.authMeGet()
	}
})
</script>

<template>
	<header
		class="grid place-items-center p-2.5 bg-white shadow-md sticky top-0 z-20 lg:px-8"
	>
		<div
			class="max-w-7xl w-full flex flex-row items-center justify-end px-2"
		>
			<div class="flex flex-row gap-2">
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
				<NuxtLink class="hidden md:inline" to="/login">
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
	</header>
</template>
