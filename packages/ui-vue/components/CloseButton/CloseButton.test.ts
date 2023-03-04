import { mount } from '@vue/test-utils'
import { checkA11y, itShouldRenderChildren } from '../tests'
import { itSupportFocusEvent } from '../tests/itSupportFocusEvent'
import CloseButton from './CloseButton.vue'

describe('PizzaUI CloseButton', () => {
	itShouldRenderChildren(CloseButton)
	itSupportFocusEvent(CloseButton, 'button')
	checkA11y([
		{ component: CloseButton, props: { title: 'Close Button' } },
		{
			component: CloseButton,
			props: { 'aria-label': 'Close Button' }
		}
	])

	it('should set XMarkIcon width and height based on iconSize prop', () => {
		const wrapper = mount(CloseButton, {
			props: {
				iconSize: 'lg'
			}
		})
		const svg = wrapper.get('svg')
		expect(svg.attributes('width')).toBe('1.75rem')
		expect(svg.attributes('height')).toBe('1.75rem')
	})
})
