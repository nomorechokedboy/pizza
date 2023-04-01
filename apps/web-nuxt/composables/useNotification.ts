import { AlertProps } from '@@/components/Alert.vue'

export const useNotification = () =>
	useState<AlertProps[]>('notification', () => [])
