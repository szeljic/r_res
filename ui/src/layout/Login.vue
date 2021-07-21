<template>
	<v-form @submit="submit" :disabled="disabled">
		<v-card>
			<v-card-title>Login</v-card-title>
			<v-divider></v-divider>
			<v-card-text>
				<v-row dense>
					<v-col>
						<v-text-field
							outlined
							label="Korisničko ime"
							v-model="item.username"
						></v-text-field>
					</v-col>
				</v-row>
				<v-row>
					<v-col>
						<v-text-field
							outlined
							label="Lozinka"
							type="password"
							v-model="item.password"
						></v-text-field>
					</v-col>
				</v-row>
			</v-card-text>
			<v-divider></v-divider>
			<v-card-actions>
				<v-spacer></v-spacer>
				<v-btn color="primary">Login</v-btn>
				<v-btn @click.prevent="reset">Očisti</v-btn>
			</v-card-actions>
		</v-card>
	</v-form>
</template>

<script>
	import axios from 'axios';

	export default {
		name: 'Login',
		data()
		{
			return {
				disabled: false,
				item: {
					username: null,
					password: null
				}
			};
		},
		methods: {
			async submit()
			{
				const response = await axios({});

				console.log(response);
			},
			reset()
			{
				this.item.username = null;
				this.item.password = null;
			}
		}
	};
</script>
