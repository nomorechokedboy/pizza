export const Size = {
	xs: 'xs',
	sm: 'sm',
	md: 'md',
	lg: 'lg',
	xl: 'xl'
}

export type UISize = keyof typeof Size

export const openColor = {
	dark: 'dark',
	gray: 'gray',
	red: 'red',
	pink: 'pink',
	grape: 'grape',
	violet: 'violet',
	indigo: 'indigo',
	blue: 'blue',
	cyan: 'cyan',
	teal: 'teal',
	green: 'green',
	lime: 'lime',
	yellow: 'yellow',
	orange: 'orange'
}

export type UIColor = keyof typeof openColor

export const Variant = {
	filled: 'filled',
	light: 'light',
	outline: 'outline',
	default: 'default',
	subtle: 'subtle'
}

export type UIVariant = keyof typeof Variant
