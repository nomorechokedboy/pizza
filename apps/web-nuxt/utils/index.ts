import { AxiosError } from 'axios'
import { notify } from '~~/composables/useNotification'
import { AlertProps } from '~~/components/Alert.vue'

export function notifyError(e: unknown) {
	if (e instanceof AxiosError) {
		const notification: AlertProps = {
			content: '',
			k: crypto.randomUUID(),
			type: 'error'
		}

		if (e.response?.data) {
			notification.content = e.response.data
		} else if (e.message) {
			notification.content = e.message
		} else {
			notification.content = 'Unknown error'
		}

		notify(notification)
	}
}
