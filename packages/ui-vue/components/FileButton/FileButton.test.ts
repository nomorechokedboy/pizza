import { checkA11y } from '$tests'
import { mount } from '@vue/test-utils'
import FileButton from './FileButton.vue'

describe('PizzaUI FileButton', () => {
	const onChange = (_payload: unknown) => {}
	checkA11y([{ component: FileButton, props: { onChange } }])
	it('should set input name', async () => {
		const wrapper = mount(FileButton, {
			props: {
				name: 'test name',
				onChange
			},
			slots: {
				default: `<button @click="params.onClick">Upload file</button>`
			}
		})
		expect(
			wrapper.get('input[type="file"]').element
		).toHaveAttribute('name', 'test name')
	})

	it('should set input accept', () => {
		const wrapper = mount(FileButton, {
			props: {
				accept: 'image/png,image/jpeg',
				onChange
			},
			slots: {
				default: `<button @click="params.onClick">Upload file</button>`
			}
		})
		expect(
			wrapper.get('input[type="file"]').element
		).toHaveAttribute('accept', 'image/png,image/jpeg')
	})
})
