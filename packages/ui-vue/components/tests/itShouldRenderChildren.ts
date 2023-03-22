import { mount } from '@vue/test-utils'
import { MountParams } from './types'

export const itShouldRenderChildren = (
	component: any,
	options?: MountParams[1]
) => {
	it('should render children', () => {
		const defaultSlot = 'Test-Children'
		const wrapper = mount(component, {
			...options,
			slots: {
				default: defaultSlot
			}
		})

		expect(wrapper.text()).toContain(defaultSlot)
	})
}
