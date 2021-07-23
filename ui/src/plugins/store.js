import Vuex from 'vuex';
import store from '../store/store';

class Store {
	install(Vue)
	{
		Vue.use(Vuex);

		Vue.prototype.$store = new Vuex.Store(store);
	}
}

export default Store;
