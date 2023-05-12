export function useNotificationEventSource() {
	return useState<EventSource | undefined>(
		'notificationEventSource',
		() => undefined
	)
}

export function cleanupNotificationEventSource() {
	const notificationEventSource = useNotificationEventSource()
	notificationEventSource.value?.close()
	notificationEventSource.value = undefined
}
