<template>
	<v-form @submit.prevent="submit" ref="frm" :disabled="disabled" v-model="valid">
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
				<v-btn color="primary" type="submit">Login</v-btn>
				<v-btn @click.prevent="reset">Očisti</v-btn>
			</v-card-actions>
		</v-card>
	</v-form>
</template>

<script>
	import Cookies from 'js-cookie';

	export default {
		name: 'Login',
		data()
		{
			return {
				disabled: false,
				item: {
					username: null,
					password: null
				},
				valid: null
			};
		},
		methods: {
			reset()
			{
				this.$refs.frm.reset();
				this.$refs.frm.resetValidation();
			},
			async submit()
			{
				this.disabled = true;

				const response = await this.$http({
					url: '/login',
					data: this.item,
					method: 'POST'
				});

				if (response.data.code === 200)
				{
					Cookies.set('x-token', response.data.access_token);
				} else if (response.data.code === 403)
				{
					Cookies.remove('x-token');
				}

				this.disabled = false;
			}
		}
	};
</script>
