/** @type {import('tailwindcss').Config} */
module.exports = {
	content: ['./components/**/*.vue'],
	theme: {
		extend: {
			borderRadius: {
				'4xl': '32px'
			},
			padding: {
				5.5: '22px',
				6.5: '26px'
			}
		}
	},
	plugins: []
}
