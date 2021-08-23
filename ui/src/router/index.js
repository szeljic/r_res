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
	component: () => import('@/views/users/UsersTable')
}, {
	path: '/korisnici/uredi/:id',
	name: 'korisnici/uredi',
	component: () => import('@/views/users/UsersTable')
}, {
	path: '/kategorije',
	name: 'kategorije',
	component: () => import('@/views/categories/CategoriesTable')
}, {
	path: '/kategorije/dodaj',
	name: 'kategorije/dodaj',
	component: () => import('@/views/categories/CategoriesTable')
}, {
	path: '/kategorije/uredi/:id',
	name: 'kategorije/uredi',
	component: () => import('@/views/categories/CategoriesTable')
}, {
	path: '/resursi',
	name: '/resursi',
	component: () => import('@/views/resources/ResourcesTable')
}];

const router = new VueRouter({
	routes
});

export default router;
