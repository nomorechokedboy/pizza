import { mount } from '@vue/test-utils'
import { itShouldRenderChildren } from '../tests'
import Overlay from './Overlay.vue'

describe('PizzaUI Overlay', () => {
	itShouldRenderChildren(Overlay)

	it('should set overlay center base on center props', () => {
		const wrapper = mount(Overlay, {
			props: {
				center: false
			}
		})
		expect(wrapper.element).toHaveAttribute('data-center')
	})

	it('should set overlay radius base on radius props', () => {
		const wrapper = mount(Overlay, {
			props: {
				radius: 'lg'
			}
		})
		expect(wrapper.element).toHaveAttribute('data-radius')
	})

	it('should set overlay fixed base on fixed props', () => {
		const wrapper = mount(Overlay, {
			props: {
				fixed: false
			}
		})
		expect(wrapper.element).toHaveAttribute('data-fixed')
	})
})
