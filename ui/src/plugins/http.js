import axios from 'axios';

class Http
{
	install(Vue)
	{
		Vue.prototype.$http = function (conf)
		{
			const vue = new Vue(),
				store = vue.$store;

			const token = store.getters['user/token'];

			if (token)
			{
				axios.defaults.headers.common['x-token'] = token;
			}

			const promise = axios(conf);

			promise.catch(async reason =>
			{
				console.warn(reason);
			});

			return promise;
		};
	}
}

export default Http;
