import { AlertProps } from '@@/components/Alert.vue'

export const useNotification = () =>
	useState<AlertProps[]>('notification', () => [])

export const notify = (notification: AlertProps) => {
	const notificationList = useNotification()
	console.log({ notification })

	notificationList.value = [...notificationList.value, notification]
}

export const removeNotification = (k: string) => {
	const notificationList = useNotification()
	notificationList.value = notificationList.value.filter((n) => n.k !== k)
}
