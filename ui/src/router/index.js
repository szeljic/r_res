import Vue from 'vue';
import VueRouter from 'vue-router';

Vue.use(VueRouter);

const routes = [{
	path: '/korisnici',
	name: 'korisnici',
	component: () => import('@/views/users/UsersTable')
}, {
	path: '/korisnici/dodaj',
	name: 'korisnici/dodaj',
	component: () => import('@/views/users/UserForm')
}, {
	path: '/korisnici/uredi/:id',
	name: 'korisnici/uredi',
	component: () => import('@/views/users/UserForm')
}, {
	path: '/kategorije',
	name: 'kategorije',
	component: () => import('@/views/categories/CategoriesTable')
}];

const router = new VueRouter({
	routes
});

export default router;
