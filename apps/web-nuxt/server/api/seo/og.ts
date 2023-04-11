// https://github.com/wobsoriano/unplugin-font-to-buffer
import Roboto from '@/assets/fonts/Roboto/Roboto-Regular.ttf'
import OgImage from '@/components/OgImage.ts'
import { satori } from 'v-satori'

export default eventHandler(async (event) => {
	const query = getQuery(event)

	const svg = await satori(OgImage, {
		props: {
			title: query.title,
			website: query.website
		},
		width: 1200,
		height: 630,
		fonts: [
			{
				name: 'Roboto',
				data: Roboto,
				weight: 400,
				style: 'normal'
			}
		]
	})

	setHeader(event, 'Content-Type', 'image/svg+xml')

	return svg
})
