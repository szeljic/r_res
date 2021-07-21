<template>
	<v-app>
		<v-main class="d-flex align-center">
			<v-container fluid>
				<v-row>
					<v-col offset="1" cols="4">
						<Login></Login>
					</v-col>
					<v-col offset="2" cols="4">
						<Registration></Registration>
					</v-col>
				</v-row>
			</v-container>
		</v-main>
	</v-app>
</template>

<script>
	import Login from '@/layout/Login';
	import Registration from '@/layout/Registration';

	export default {
		components: {Login, Registration},
		name: 'Public'
	};
</script>
