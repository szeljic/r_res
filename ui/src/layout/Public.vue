<template>
	<v-app>
		<v-main>
			<v-container fluid>
				<v-row>
					<v-col class="flex-grow-1"></v-col>
					<v-col class="flex-shrink-1 flex-grow-0">
						<v-toolbar flat>
							<v-toolbar-items>
								<v-btn dark tile @click="login = true" :color="login ? 'green' : ''" class="mr-2">
									Prijava
								</v-btn>
								<v-btn dark tile @click="login = false" :color="!login ? 'green' : ''" class="ml-2">
									Registracija
								</v-btn>
							</v-toolbar-items>
						</v-toolbar>
					</v-col>
					<v-col class="flex-grow-1"></v-col>
				</v-row>
				<v-row>
					<v-col md="12" lg="4" offset-lg="4">
						<Login v-if="login" :pre-username="preUsername"></Login>
						<Registration v-if="!login" @registered="registered"></Registration>
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
		name: 'Public',
		data()
		{
			return {
				login: true,
				preUsername: null
			};
		},
		methods: {
			registered(username)
			{
				this.login = true;

				this.preUsername = username;
			}
		},
		watch: {
			login(v)
			{
				if (!v)
				{
					this.preUsername = null;
				}
			}
		}
	};
</script>
