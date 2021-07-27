import Vue from 'vue';

import vuetify from './plugins/vuetify';
import Layout from '@/layout/Layout';
import Validate from '@/plugins/validate';

import Bus from '@/plugins/bus';
import Http from '@/plugins/http';
import Store from '@/plugins/store';
import User from '@/plugins/user';
import Common from '@/plugins/common';

import router from './router';

Vue.use(new Bus());
Vue.use(new Store());
Vue.use(new Validate());
Vue.use(new Http());
Vue.use(new User());
Vue.use(new Common());

Vue.config.productionTip = false;

new Vue({
	router,
	vuetify,
	render: h => h(Layout)
}).$mount('#app');
