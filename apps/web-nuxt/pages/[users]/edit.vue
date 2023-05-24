<script lang="ts" setup>
import useVuelidate from '@vuelidate/core'
import {
	email,
	helpers,
	maxLength,
	minLength,
	numeric,
	required
} from '@vuelidate/validators'
import { Button } from 'ui-vue'

function formDataRules() {
	return {
		email: {
			required: helpers.withMessage(
				'The email field is required',
				required
			),
			email: helpers.withMessage(
				'Invalid email format',
				email
			)
		},
		username: {
			required: helpers.withMessage(
				'The user name field is required',
				required
			),
			minLength: minLength(3),
			maxLength: maxLength(30)
		},
		fullname: {
			required: helpers.withMessage(
				'The full name field is required',
				required
			),
			minLength: minLength(3),
			maxLength: maxLength(30)
		},
		phone: {
			numeric: helpers.withMessage('Number only', numeric),
			minLength: minLength(10),
			maxLength: maxLength(20)
		}
	}
}

async function handleUpdateProfile() {
	const isFormValid = await v$.value.$validate()
	if (!isFormValid) {
		return
	}

	try {
		if (currentFile.value) {
			const { data: imageData } = await $blogApi.media.image(
				currentFile.value
			)

			formData.value.avatar = imageData as unknown as string
		} else {
			formData.value.avatar = userProfile.value?.avatar as any
		}

		const { avatar, email, fullname, phone, username } =
			formData.value
		await $blogApi.user.userUpdatePut({
			avatar,
			email,
			fullname,
			phonenumber: phone,
			username
		})
		notify({
			k: crypto.randomUUID(),
			content: 'Update profile success',
			type: 'success'
		})
		await refetch()
	} catch (e) {
		notifyError(e)
	}
}

function handleBrowseImage() {
	avatarInput.value?.click()
}

const appConfig = useRuntimeConfig()
const { $blogApi } = useNuxtApp()
const { data: userProfile, refetch } = useUserProfile()
const loading = ref(false)
const currentFile = ref<File | undefined>()
const userAvatar = computed(() =>
	userProfile.value?.avatar
		? `${appConfig.public.mediaUrl}${userProfile.value?.avatar}`
		: `${appConfig.public.dicebearMedia}${userProfile.value?.name}`
)
const formData = computed(() =>
	reactive({
		avatar: userAvatar.value,
		email: userProfile.value?.email,
		fullname: userProfile.value?.name,
		phone: userProfile.value?.phone,
		username: userProfile.value?.username
	})
)
const rules = computed(formDataRules)
const v$ = useVuelidate(rules, formData)
const avatarInput = ref<HTMLInputElement | null>(null)

definePageMeta({ middleware: ['authn'] })
</script>

<template>
	<main class="min-w-full w-full flex-grow py-10">
		<div
			class="w-full md:max-w-5xl flex flex-col mx-auto bg-white rounded p-5 gap-3"
		>
			<h3 class="text-neutral-800 font-bold text-2xl">
				User
			</h3>
			<form class="flex flex-col">
				<TextInput
					:error="v$.fullname.$error"
					:errors="v$.fullname.$errors"
					id="fullname"
					name="fullname"
					type="fullname"
					v-model="formData.fullname"
					@blur="v$.fullname.$touch"
				>
					<template #label>Name</template>
				</TextInput>
				<TextInput
					:error="v$.email.$error"
					:errors="v$.email.$errors"
					id="email"
					name="email"
					type="email"
					v-model="formData.email"
					@blur="v$.email.$touch"
				>
					<template #label>Email</template>
				</TextInput>
				<TextInput
					:error="v$.username.$error"
					:errors="v$.username.$errors"
					id="username"
					name="username"
					type="username"
					v-model="formData.username"
					@blur="v$.username.$touch"
				>
					<template #label>Username</template>
				</TextInput>
				<div class="flex flex-col gap-2 pb-4">
					<label for="avatar">Avatar</label>
					<div class="flex items-center gap-3">
						<Avatar
							:src="userAvatar"
							width="48"
						/>
						<Button
							@click.prevent="
								handleBrowseImage
							"
							>Browse</Button
						>
						<input
							type="file"
							name="avatar"
							id="avatar"
							class="hidden"
							accept="image/*"
							@change="
								currentFile =
									$event
										.target
										.files?.[0]
							"
							ref="avatarInput"
						/>
					</div>
				</div>
				<Button
					color="indigo"
					size="md"
					class="!py-3"
					@click.prevent="handleUpdateProfile"
					block
					:loading="loading"
					>Save profile</Button
				>
			</form>
		</div>
	</main>
</template>
