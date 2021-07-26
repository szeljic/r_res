import Vue from 'vue';
import VueRouter from 'vue-router';

Vue.use(VueRouter);

const routes = [{
	path: '/korisnici',
	name: 'korisnici',
	component: () => import('@/views/users/UsersTable')
}];

const router = new VueRouter({
	routes
});

export default router;
