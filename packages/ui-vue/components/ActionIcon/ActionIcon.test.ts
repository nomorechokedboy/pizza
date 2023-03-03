import { mount } from '@vue/test-utils'
import Loader from '../Loader/Loader.vue'
import { checkA11y, itShouldRenderChildren } from '../tests'
import { itSupportFocusEvent } from '../tests/itSupportFocusEvent'
import ActionIcon from './ActionIcon.vue'

describe('PizzaUI ActionIcon', () => {
	itShouldRenderChildren(ActionIcon)
	itSupportFocusEvent(ActionIcon, 'button')
	checkA11y(ActionIcon)

	it('should passes type to button component', () => {
		const wrapper = mount(ActionIcon, {
			props: {
				type: 'submit'
			}
		})
		expect(wrapper.element.firstChild).toHaveAttribute(
			'type',
			'submit'
		)
	})

	it('should set disabled attribute', () => {
		const wrapper = mount(ActionIcon, {
			props: {
				disabled: true
			}
		})
		expect(wrapper.element.firstChild).toBeDisabled()
	})

	it('should be disabled when inside disabled fieldset', () => {
		const DisabledFieldset = {
			components: { ActionIconVue: ActionIcon },
			template: `<fieldset disabled>
					<ActionIconVue type="submit" />
				</fieldset>`
		}
		const wrapper = mount(DisabledFieldset)
		expect(wrapper.get('button').element).toBeDisabled()
	})

	it('should set loading attribute', () => {
		const wrapper = mount(ActionIcon, {
			props: {
				loading: true
			}
		})
		expect(wrapper.element.firstChild).toHaveAttribute(
			'data-loading'
		)
	})

	it('should replace children with Loader when loading is set to true', () => {
		const defaultSlot = {
			slots: {
				default: 'Lmao'
			}
		}
		const loadingWrapper = mount(ActionIcon, {
			props: {
				loading: true
			},
			...defaultSlot
		})
		const defaultWrapper = mount(ActionIcon, {
			props: {
				loading: false
			},
			...defaultSlot
		})
		expect(
			loadingWrapper.findComponent(Loader).exists()
		).toBeTruthy()
		expect(loadingWrapper.text()).toBe('')
		expect(
			defaultWrapper.findComponent(Loader).exists()
		).toBeFalsy()
		expect(defaultWrapper.text()).toContain('Lmao')
	})
})
