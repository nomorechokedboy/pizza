<script setup lang="ts">
<<<<<<< HEAD
import { ActionIcon, Button } from 'ui-vue'
import { reactive } from 'vue'
import IconBell from '~icons/ph/bell-simple'
import Avatar from './Avatar.vue'
import Dropdown from './Dropdown.vue'
import { Button } from 'ui-vue'

const token = useAuthToken()
const isLoggedIn = computed(
    () => token.value.accessToken && token.value.refreshToken
)
const { $blogApi } = useNuxtApp()
const toggle = reactive({
    open: false,
    onChange() {
        this.open = !this.open
    }
})
watchEffect(() => {
    if (isLoggedIn.value) {
        $blogApi.auth.authMeGet()
    }
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
					class="bg-sky-500 h-10 lg:w-full lg:max-w-md"
				></div>
			</div>
			<div class="flex flex-row gap-2">
				<NuxtLink v-if="isLoggedIn" to="/new">
					<Button color="indigo"
						>Create Post</Button
					>
				</NuxtLink>
				<!-- <ActionIcon
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
						class="focus:ring-4 focus:outline-none focus:ring-gray-300"
					>
						<Avatar
							class="md:w-8"
							width="24"
							:src="'https://avatars.githubusercontent.com/u/42694704?v=4'"
							@click="
								toggle.onChange()
							"
						/>
					</ActionIcon>
					<Dropdown
						:open="toggle.open"
						:user="{
							name: 'Đỗ Viên',
							username: 'Cpea2506'
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
