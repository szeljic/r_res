import axios from 'axios';
import Cookies from 'js-cookie';

const tokenKey = 'x-token';

class Http
{
	install(Vue)
	{
		Vue.prototype.$http = function (conf)
		{
			const token = Cookies.get(tokenKey);

			if (token)
			{
				axios.defaults.headers.common[tokenKey] = token;
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
