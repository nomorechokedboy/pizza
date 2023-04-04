<script setup lang="ts">
import useVuelidate from '@vuelidate/core'
import { helpers, minLength, required, sameAs } from '@vuelidate/validators'
import { Button } from 'ui-vue'
import TextInput from '~~/components/TextInput.vue'

definePageMeta({
	middleware: ['authn', 'reset-password-guard']
})

const { $blogApi } = useNuxtApp()
const route = useRoute()
const token = route.query.token?.toString()
const formData = reactive({ password: '', confirmPassword: '' })
const passwordInput = ref<InstanceType<typeof TextInput> | null>(null)
function formDataRules() {
	return {
		password: {
			required: helpers.withMessage(
				'The password field is required',
				required
			),
			minLength: minLength(8)
		},
		confirmPassword: {
			required: helpers.withMessage(
				'The password confirmation field is required',
				required
			),
			sameAs: helpers.withMessage(
				"Passwords don't match",
				sameAs(formData.password)
			)
		}
	}
}
const rules = computed(formDataRules)
const v$ = useVuelidate(rules, formData)
async function handleSubmit() {
	const isFormValid = await v$.value.$validate()
	if (!isFormValid) {
		return
	}

	try {
		await $blogApi.auth.authResetPasswordPut({
			password: formData.password,
			token
		})
		notify({
			content: 'Reset password success!',
			k: crypto.randomUUID(),
			type: 'success'
		})
	} catch (e) {
		notifyError(e)
	}
}
</script>

<template>
	<div class="pt-12 pb-32 bg-[rgb(245,245,245)]">
		<div
			class="flex flex-col gap-4 p-8 pt-3 border bg-white lg:max-w-[490px] m-auto"
		>
			<h1 class="font-bold text-[27px]">
				Reset your password
			</h1>
			<form class="flex flex-col gap-3">
				<TextInput
					:error="v$.password.$error"
					:errors="v$.password.$errors"
					id="password"
					name="password"
					type="password"
					v-model="formData.password"
					ref="passwordInput"
					@blur="v$.password.$touch"
				>
					<template #label>Password</template>
				</TextInput>
				<TextInput
					:errors="v$.confirmPassword.$errors"
					:error="v$.confirmPassword.$error"
					name="confirmPassword"
					type="password"
					v-model="formData.confirmPassword"
					@blur="v$.confirmPassword.$touch"
				>
					<template #label
						>Confirm Password</template
					>
				</TextInput>
				<Button
					color="indigo"
					size="md"
					class="!font-medium"
					@click.prevent="handleSubmit"
					block
				>
					Reset
				</Button>
			</form>
		</div>
	</div>
</template>
