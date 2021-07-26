module.exports = {
	transpileDependencies: ['vuetify'],
	productionSourceMap: false,
	configureWebpack: function (config)
	{
		if (!config.hasOwnProperty('devServer'))
		{
			config.devServer = {};
		}

		config.devServer.proxy = {
			'^/api/': {
				target: 'http://127.0.0.1:9000/'
			}
		};

		console.log(config);
	}
};
