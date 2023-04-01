import { AlertProps } from '@@/components/Alert.vue'

export function useNotification() {
	return useState<AlertProps[]>('notification', () => [])
}

export function notify(notification: AlertProps) {
	const notificationList = useNotification()

	notificationList.value = [...notificationList.value, notification]
}

export function removeNotification(k: string) {
	const notificationList = useNotification()
	notificationList.value = notificationList.value.filter((n) => n.k !== k)
}
