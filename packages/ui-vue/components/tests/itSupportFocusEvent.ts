import { mount } from '@vue/test-utils'
import { MountParams } from './types'

export const itSupportFocusEvent = (
	component: any,
	selector: keyof HTMLElementTagNameMap,
	options?: MountParams[1]
) => {
	it('should support focus event', () => {
		const onFocusMock = vi.fn()
		const onBlurMock = vi.fn()
		const wrapper = mount(component, {
			...options,
			props: {
				...options?.props,
				onfocus: onFocusMock,
				onblur: onBlurMock
			}
		})

		wrapper.find(selector).trigger('focus')
		expect(onFocusMock).toHaveBeenCalled()

		wrapper.find(selector).trigger('blur')
		expect(onBlurMock).toHaveBeenCalled()
	})
}
