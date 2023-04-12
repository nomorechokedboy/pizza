import { faker } from '@faker-js/faker'
import { expect, test } from '@playwright/test'

test('Test publish post happy case', async ({ page, isMobile }) => {
	await page.goto('http://localhost:3000/')
	if (!isMobile) {
		await page.getByRole('button', { name: 'Log in' }).click()
	} else {
		await page.goto('http://localhost:3000/login')
	}
	await page.getByLabel('Email').click()
	await page.getByLabel('Email').fill('19110073@student.hcmute.edu.vn')
	await page.getByLabel('Email').press('Tab')
	await page.getByLabel('Password').fill('lmao123')
	await page
		.getByRole('button', { name: 'Continue', exact: true })
		.click()

	await page.getByRole('button', { name: 'Create Post' }).click()
	await page.getByRole('textbox', { name: 'Post title here...' }).click()

	const title = faker.lorem.sentence()
	const content = faker.lorem.paragraphs(
		faker.datatype.number({ min: 1, max: 20 })
	)

	await page
		.getByRole('textbox', { name: 'Post title here...' })
		.fill(title)
	await page.getByPlaceholder('Post content here...').click()
	await page.getByPlaceholder('Post content here...').fill(content)
	await page.getByRole('button', { name: 'Publish' }).click()
	page.on('requestfailed', (request) => {
		expect(request.failure()).toBeNull()
	})
})
