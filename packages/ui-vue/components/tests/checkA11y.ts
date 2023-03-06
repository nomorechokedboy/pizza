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

interface RenderComponent {
	component: any
	props: any
}

export const checkA11y = (components: RenderComponent[]) => {
	it('should has no a11y violation', async () => {
		for (const { component, props } of components) {
			const wrapper = mount(component, {
				slots: {
					default: 'Pizza Component'
				},
				props: props
			})
			const res = await axe(wrapper.element, config)
			expect(res).toHaveNoViolations()
		}
	})
}
