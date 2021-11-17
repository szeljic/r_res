module.exports = {
	root: true,
	env: {
		node: true
	},
	extends: ['plugin:vue/essential', 'eslint:recommended'],
	parserOptions: {
		parser: '@babel/eslint-parser'
	},
	rules: {
		'no-console': process.env.NODE_ENV === 'production' ? 'warn' : 'off',
		'no-debugger': process.env.NODE_ENV === 'production' ? 'error' : 'off',

		'indent': [
			'error',
			'tab',
			{
				'SwitchCase': 1
			}
		],
		'quotes': [
			'warn',
			'single'
		],
		'comma-dangle': [
			'error',
			'never'
		],
		'semi': 'error',
		'no-unused-vars': 'error',
		'compat/compat': 'warn',
		'no-trailing-spaces': 'error',
		'no-dupe-keys': 'error',
		'no-prototype-builtins': 'off',
		'no-mixed-spaces-and-tabs': ['warn'],
		'eqeqeq': 'warn',
		'vue/valid-v-slot': ['error', {
			'allowModifiers': true
		}]
	},
	overrides: [
		{
			'files': ['*.vue'],
			'rules': {
				'indent': 'off',
				'vue/script-indent': [
					'error',
					'tab',
					{
						'baseIndent': 1,
						'switchCase': 1
					}
				]
			}
		}
	],
	plugins: ['compat']
};
