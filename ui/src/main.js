import Vue from 'vue';

import router from './router';
import store from './store';
import vuetify from './plugins/vuetify';
import Layout from '@/layout/Layout';

Vue.config.productionTip = false;

new Vue({
	router,
	store,
	vuetify,
	render: h => h(Layout)
}).$mount('#app');
