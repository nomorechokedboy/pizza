import { mount } from '@vue/test-utils'
import { axe } from 'vitest-axe'

const config = {
	rules: {
		region: {
			enabled: false
		},

		'autocomplete-valid': {
			enabled: false
		}
	}
}

export const checkA11y = (component: any) => {
	it('should has no a11y violation', async () => {
		const wrapper = mount(component, {
			slots: {
				default: 'Pizza Component'
			}
		})
		const res = await axe(wrapper.element, config)
		expect(res).toHaveNoViolations()
	})
}
