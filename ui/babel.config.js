module.exports = {
	'presets': [
		['@babel/preset-env', {useBuiltIns: 'entry', corejs: '3'}]
	],
	plugins: [
		'@babel/plugin-proposal-nullish-coalescing-operator',
		'@babel/plugin-proposal-class-static-block',
		'@babel/plugin-proposal-class-properties'
	]
};
