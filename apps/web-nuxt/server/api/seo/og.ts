// https://github.com/wobsoriano/unplugin-font-to-buffer
import Roboto from '@/assets/fonts/Roboto/Roboto-Regular.ttf'
import OgImage from '@/components/OgImage.ts'
import { Resvg } from '@resvg/resvg-js'
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

	const resvg = new Resvg(svg, {
		background: 'rgba(238, 235, 230, .9)'
	})
	const pngData = resvg.render()
	const pngBuffer = pngData.asPng()

	setHeader(event, 'Content-Type', 'image/png')
	setHeader(
		event,
		'Cache-Control',
		'public, immutable, no-transform, max-age=31536000'
	)

	return pngBuffer
})
