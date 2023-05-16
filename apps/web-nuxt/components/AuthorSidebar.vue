<script lang="ts" setup>
import { Button } from 'ui-vue'

function computeUserAvatar() {
	return (
		postDetails.value?.user?.avatar ||
		`${appConfig.public.dicebearMedia}${
			postDetails.value?.user?.fullName ||
			postDetails.value?.user?.userName ||
			'A6Blog&backgroundColor=000000'
		}`
	)
}

const appConfig = useRuntimeConfig()
const slug = inject<string>('slug', 'no slug')
const { data: postDetails, isLoading: isPostDetailsLoading } =
	usePostDetails(slug)
const { data: userProfile } = useUserProfile()
const isUser = computed(
	() => postDetails.value?.user?.id === userProfile.value?.id
)
const userAvatar = computed(computeUserAvatar)
</script>

<template>
	<aside
		class="hidden w-full bg-black/50 md:relative md:max-w-xs md:z-10 md:block md:bg-transparent md:h-screen"
	>
		<div class="w-full bg-transparent flex flex-col md:h-full">
			<header
				class="bg-white flex flex-col gap-4 md:border-t-[32px] rounded-lg md:border-black px-4 pb-4"
			>
				<div class="flex items-center gap-2 -mt-4">
					<div>
						<Avatar
							v-if="
								!isPostDetailsLoading
							"
							class="w-10 md:!w-12"
							:src="userAvatar"
							:alt="`${
								postDetails
									?.user
									?.fullName ||
								postDetails
									?.user
									?.userName
							} avatar`"
							width="48"
						/>
						<div
							class="skeleton !rounded-full"
							v-else
						>
							<div
								class="w-10 h-10 md:!w-12 md:!h-12"
							/>
						</div>
					</div>
					<span
						class="text-neutral-700 text-xl font-bold self-end"
						:class="{
							skeleton: isPostDetailsLoading
						}"
					>
						<template
							v-if="
								!isPostDetailsLoading
							"
							>{{
								postDetails
									?.user
									?.fullName ||
								postDetails
									?.user
									?.userName
							}}
						</template>
						<br v-else />
					</span>
				</div>
				<div>
					<Button
						block
						:loading="isPostDetailsLoading"
						:loader-position="'center'"
					>
						{{
							isUser
								? 'Edit profile'
								: 'Follow'
						}}
					</Button>
				</div>
				<div>
					<h3
						class="uppercase text-xs text-neutral-600 font-bold"
						:class="{
							skeleton: isPostDetailsLoading
						}"
					>
						<template
							v-if="
								!isPostDetailsLoading
							"
							>Joined</template
						>
						<br v-else />
					</h3>
					<time
						class="text-neutral-700 text-base"
						:class="{
							skeleton: isPostDetailsLoading
						}"
					>
						<template
							v-if="
								!isPostDetailsLoading
							"
						>
							{{
								convertDate(
									postDetails
										?.user
										?.createdAt,
									'MMM D, YYYY'
								)
							}}
						</template>
						<br v-else />
					</time>
				</div>
			</header>
			<nav class="hidden h-full bg-transparent p-2"></nav>
		</div>
	</aside>
</template>
