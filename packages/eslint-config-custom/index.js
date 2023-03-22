module.exports = {
	extends: ['next', 'turbo', 'prettier'],
	parser: '@typescript-eslint/parser',
	plugins: ['@typescript-eslint'],
	rules: {
		'@next/next/no-html-link-for-pages': 'off',
		'react/jsx-key': 'off',
		'@typescript-eslint/no-unused-vars': [
			'error',
			{ argsIgnorePattern: '^_' }
		],
		'@typescript-eslint/consistent-type-definitions': [
			'error',
			'interface'
		],
		camelcase: 'off',
		'@typescript-eslint/naming-convention': [
			'error',
			{
				selector: 'variable',
				format: [
					'camelCase',
					'PascalCase',
					'UPPER_CASE'
				]
			},
			{
				selector: 'parameter',
				format: ['camelCase'],
				leadingUnderscore: 'allow'
			},
			{
				selector: 'typeLike',
				format: ['PascalCase']
			}
		],
		'react/jsx-no-undef': 'off'
	}
}
