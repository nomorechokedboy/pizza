import { mount } from '@vue/test-utils'
import Loader from './Loader.vue'

describe('PizzaUI Loader', () => {
	it('should set loader size base on size prop', () => {
		const wrapper = mount(Loader, {
			props: {
				size: 'lg'
			}
		})
		expect(wrapper.element).toHaveAttribute('data-size')
	})
})
