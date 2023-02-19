import { mount } from '@vue/test-utils'
import Button from './Button.vue'

describe('PizzaUI Button', () => {
	it('should render children', () => {
		const buttonText = 'Click me'
		const wrapper = mount(Button, {
			slots: {
				default: buttonText
			}
		})

		expect(wrapper.text()).toContain(buttonText)
	})
})
