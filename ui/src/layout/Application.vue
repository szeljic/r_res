<template>
	<v-app>
		<v-navigation-drawer
			v-model="drawer"
			app
			clipped
		>
			<v-list nav dense>
				<v-list-item to="/korisnici">
					<v-list-item-icon>
						<v-icon>mdi-account</v-icon>
					</v-list-item-icon>
					<v-list-item-content>
						<v-list-item-title>Korisnici</v-list-item-title>
					</v-list-item-content>
				</v-list-item>

				<v-list-item to="/kategorije">
					<v-list-item-icon>
						<v-icon>mdi-shape</v-icon>
					</v-list-item-icon>
					<v-list-item-content>
						<v-list-item-title>Kategorije</v-list-item-title>
					</v-list-item-content>
				</v-list-item>

				<v-list-item to="/resursi">
					<v-list-item-icon>
						<v-icon>mdi-inbox-multiple</v-icon>
					</v-list-item-icon>
					<v-list-item-content>
						<v-list-item-title>Resursi</v-list-item-title>
					</v-list-item-content>
				</v-list-item>

				<v-list-item to="/rezervacije">
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
		</v-app-bar>

		<v-main>
			<v-container>
				<router-view></router-view>
			</v-container>
		</v-main>
	</v-app>
</template>

<script>
	export default {
		name: 'Application',
		data: () =>
		{
			return {
				drawer: true
			};
		}
	};
</script>
