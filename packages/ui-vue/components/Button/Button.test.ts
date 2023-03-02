import { mount } from '@vue/test-utils'
import { checkA11y, itShouldRenderChildren } from '../tests'
import { itSupportFocusEvent } from '../tests/itSupportFocusEvent'
import Button from './Button.vue'

describe('PizzaUI Button', () => {
	itShouldRenderChildren(Button)
	itSupportFocusEvent(Button, 'button')
	checkA11y(Button)

	it('should passes type to button component', () => {
		const wrapper = mount(Button, {
			props: {
				type: 'submit'
			}
		})
		expect(wrapper.element).toHaveAttribute('type', 'submit')
	})

	it('should set disabled attribute', () => {
		const wrapper = mount(Button, {
			props: {
				disabled: true
			}
		})
		expect(wrapper.element).toBeDisabled()
	})

	it('should be disabled when inside disabled fieldset', () => {
		const DisabledFieldset = {
			components: { Button },
			template: `<fieldset disabled>
					<Button type="submit" />
				</fieldset>`
		}
		const wrapper = mount(DisabledFieldset)
		expect(wrapper.get('button').element).toBeDisabled()
	})

	it('should set loading attribute', () => {
		const wrapper = mount(Button, {
			props: {
				loading: true
			}
		})
		expect(wrapper.element).toHaveAttribute('data-loading')
	})

	it('should render left and right icons if they are provided', () => {
		const wrapper = mount(Button, {
			slots: {
				leftIcon: '<div>LeftIcon</div>',
				rightIcon: '<div>RightIcon</div>'
			}
		})
		const text = wrapper.text()
		expect(text).toContain('LeftIcon')
		expect(text).toContain('RightIcon')
	})
})
