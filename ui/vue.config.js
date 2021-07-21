module.exports = {
	transpileDependencies: ['vuetify'],
	productionSourceMap: false,
	configureWebpack: {
		optimization: {
			splitChunks: {
				cacheGroups: {
					common: {
						chunks: 'all'
					}
				}
			}
		}
	}
};
