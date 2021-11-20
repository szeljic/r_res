<template>
	<v-app>
		<v-navigation-drawer
			v-model="drawer"
			app
			clipped
		>
			<v-list nav dense>
				<v-list-item to="/korisnici" v-if="ut === 'admin'">
					<v-list-item-icon>
						<v-icon>mdi-account</v-icon>
					</v-list-item-icon>
					<v-list-item-content>
						<v-list-item-title>Korisnici</v-list-item-title>
					</v-list-item-content>
				</v-list-item>

				<v-list-item to="/kategorije" v-if="ut === 'admin'">
					<v-list-item-icon>
						<v-icon>mdi-shape</v-icon>
					</v-list-item-icon>
					<v-list-item-content>
						<v-list-item-title>Kategorije</v-list-item-title>
					</v-list-item-content>
				</v-list-item>

				<v-list-item to="/resursi" v-if="ut === 'admin' || ut === 'user'">
					<v-list-item-icon>
						<v-icon>mdi-inbox-multiple</v-icon>
					</v-list-item-icon>
					<v-list-item-content>
						<v-list-item-title>Resursi</v-list-item-title>
					</v-list-item-content>
				</v-list-item>

				<v-list-item to="/rezervacije" v-if="ut === 'admin' || ut === 'user'">
					<v-list-item-icon>
						<v-icon>mdi-newspaper</v-icon>
					</v-list-item-icon>
					<v-list-item-content>
						<v-list-item-title>Rezervacije</v-list-item-title>
					</v-list-item-content>
				</v-list-item>
			</v-list>
		</v-navigation-drawer>

		<v-app-bar app clipped-left>
			<v-app-bar-nav-icon @click="drawer = !drawer"></v-app-bar-nav-icon>
			<v-toolbar-title>Rezervator</v-toolbar-title>
			<v-spacer></v-spacer>
			<v-btn text small>
				<v-icon left>{{ ut === 'admin' ? 'mdi mdi-account-plus' : 'mdi mdi-account' }}</v-icon>
				{{ name }}
			</v-btn>
			<v-tooltip left>
				<template v-slot:activator="{ on, attr }">
					<v-btn icon v-on="on" v-bind="attr" @click.prevent="logout">
						<v-icon>mdi-logout</v-icon>
					</v-btn>
				</template>
				<span>Izloguj se</span>
			</v-tooltip>
		</v-app-bar>

		<v-main>
			<v-container>
				<router-view></router-view>
			</v-container>
		</v-main>
	</v-app>
</template>

<script>
	import {mapGetters} from 'vuex';

	export default {
		name: 'Application',
		data: () =>
		{
			return {
				drawer: true
			};
		},
		created()
		{
			if (this.ut === 'admin' && this.$route.name !== 'korisnici')
			{
				this.$router.push({name: 'korisnici'});
			} else if (this.ut === 'user' && this.$route.name !== 'rezervacije')
			{
				this.$router.push({name: 'rezervacije'});
			}
		},
		computed: {
			...mapGetters({
				whoami: 'user/whoami'
			}),
			name()
			{
				return this.whoami.firstName + ' ' + this.whoami.lastName;
			},
			ut()
			{
				return this.whoami.userType;
			}
		},
		methods: {
			logout()
			{
				this.$user.logout();
			}
		}
	};
</script>
