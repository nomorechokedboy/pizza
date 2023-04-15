<script setup lang="ts">
import { ActionIcon, Button } from 'ui-vue'
import { dicebearMedia } from '~/constants'
import IconBell from '~icons/ph/bell-simple'

const isLoggedIn = useIsAuthenticated()
const toggle = ref(false)
const { data: userProfile } = useUserProfile()
const userAvatar = computed(
	() =>
		userProfile.value?.avatar ||
		`${dicebearMedia}${
			userProfile.value?.name ||
			'A6Blog&backgroundColor=000000'
		}`
)

function handleToggle() {
	toggle.value = !toggle.value
}

async function logout() {
	removeToken()
	navigateTo('/login')
}
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
					<Dropdown :open="toggle">
						<div
							class="pb-2 mb-2 hover:underline hover:rounded-lg px-4 py-3 dark:text-gray-200 hover:bg-gray-100 dark:hover:bg-gray-600 dark:hover:text-white"
						>
							<a href="#">
								<div>
									<span
										class="fw-medium block"
									>
										{{
											userProfile?.name ||
											'No username'
										}}
									</span>
									<small
										class="text-gray-400"
										>@{{
											userProfile?.username ||
											'No username'
										}}
									</small>
								</div>
							</a>
						</div>
						<ul
							class="py-2 text-gray-700 dark:text-gray-200"
							aria-labelledby="dropdownDefaultButton"
						>
							<!-- <li>
                <a href="#"
                    class="block px-4 py-2 hover:underline hover:rounded-lg hover:bg-gray-100 dark:hover:bg-gray-600 dark:hover:text-white">Dashboard</a>
            </li> -->
							<li>
								<NuxtLink
									to="/new"
									class="block px-4 py-2 hover:underline hover:rounded-lg hover:bg-gray-100 dark:hover:bg-gray-600 dark:hover:text-white"
								>
									Create
									Post</NuxtLink
								>
							</li>
							<!-- <li>
                <a href="#"
                    class="block px-4 py-2 hover:underline hover:rounded-lg hover:bg-gray-100 dark:hover:bg-gray-600 dark:hover:text-white">Reading
                    list</a>
            </li>
            <li>
                <a href="#"
                    class="block px-4 py-2 hover:underline hover:rounded-lg hover:bg-gray-100 dark:hover:bg-gray-600 dark:hover:text-white">Settings</a>
            </li> -->
						</ul>
						<div
							class="pt-2 dark:text-gray-200"
						>
							<div
								class="cursor-pointer px-4 py-2 hover:underline hover:rounded-lg hover:bg-gray-100 dark:hover:bg-gray-600 dark:hover:text-white"
								@click="logout"
							>
								Sign out
							</div>
						</div>
					</Dropdown>
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
