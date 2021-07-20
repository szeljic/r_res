<template>
	<v-app>
		<v-navigation-drawer
			v-model="drawer"
			app
		></v-navigation-drawer>

		<v-app-bar app>
			<v-app-bar-nav-icon @click="drawer = !drawer"></v-app-bar-nav-icon>

			<v-toolbar-title>Reservator</v-toolbar-title>
		</v-app-bar>

		<v-main>
			<router-view></router-view>
		</v-main>
	</v-app>
</template>

<script>

	export default {
		name: 'App',

		data: () => ({drawer: null})
	};
</script>
