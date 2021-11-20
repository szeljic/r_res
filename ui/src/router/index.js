import Vue from 'vue';
import VueRouter from 'vue-router';

Vue.use(VueRouter);

const routes = [{
	path: '/',
	name: 'home'
},{
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
	name: 'resursi',
	component: () => import('@/views/resources/ResourcesTable')
}, {
	path: '/rezervacije',
	name: 'rezervacije',
	component: () => import('@/views/reservations/ReservationsTable')
}, {
	path: '/rezervacije/dodaj',
	name: 'rezervacije/dodaj',
	component: () => import('@/views/reservations/ReservationFormComponent')
}];

const router = new VueRouter({
	mode: 'history',
	routes
});

export default router;
