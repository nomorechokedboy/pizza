import { mount } from '@vue/test-utils'
import LoadingOverlay from './LoadingOverlay.vue'

describe('PizzaUI LoadingOverlay', () => {
	it('should render base on visible props', () => {
		const visibleWrapper = mount(LoadingOverlay, {
			props: {
				visible: true
			}
		})
		const invisibleWrapper = mount(LoadingOverlay, {
			props: {
				visible: false
			}
		})
		expect(visibleWrapper.isVisible()).toBeTruthy()
		expect(invisibleWrapper.isVisible()).toBeFalsy()
	})
})
