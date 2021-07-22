import Vue from 'vue';

import router from './router';
import store from './store';
import vuetify from './plugins/vuetify';
import Layout from '@/layout/Layout';
import Validate from '@/plugins/validate';
import Http from '@/plugins/http';

Vue.use(new Validate());
Vue.use(new Http());

Vue.config.productionTip = false;

new Vue({
	router,
	store,
	vuetify,
	render: h => h(Layout)
}).$mount('#app');
