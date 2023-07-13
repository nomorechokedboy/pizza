<script lang="ts" setup>
import { Button } from 'ui-vue'

const appConfig = useRuntimeConfig()
const { data: userProfile, isLoading: isUserProfileLoading } = useUserProfile()
const userAvatar = computed(() =>
	userProfile.value?.avatar
		? `${appConfig.public.mediaUrl}${userProfile.value?.avatar}`
		: `${appConfig.public.dicebearMedia}${userProfile.value?.name}`
)
</script>

<template>
	<header
		class="w-full flex items-center justify-between rounded shadow p-5 bg-white flex-shrink-0"
	>
		<div class="flex items-center gap-5 w-1/2">
			<Avatar
				class="w-14 md:!w-full max-w-[80px]"
				:src="userAvatar"
				width="80"
				v-if="!isUserProfileLoading"
			/>
			<div class="skeleton !rounded-full max-w-[80px]" v-else>
				<div class="w-14 h-14 md:!w-full md:!h-20" />
			</div>
			<div class="flex flex-col gap-2 w-full">
				<p
					class="text-2xl font-bold text-neutral-700 md:text-4xl"
					:class="{
						skeleton: isUserProfileLoading
					}"
				>
					<template v-if="!isUserProfileLoading">
						{{
							userProfile?.name ||
							userProfile?.username ||
							'Lmao error'
						}}
					</template>
					<br v-else />
				</p>
				<p
					class="text-xs text-neutral-600 md:text-base"
					:class="{
						skeleton: isUserProfileLoading
					}"
				>
					<template v-if="!isUserProfileLoading">
						{{
							userProfile?.email
						}}</template
					>
					<br v-else />
				</p>
			</div>
		</div>
		<NuxtLink :to="`/${$route.params.users}/edit`">
			<Button
				:loading="isUserProfileLoading"
				loader-position="center"
				>Edit Profile</Button
			></NuxtLink
		>
	</header>
</template>
