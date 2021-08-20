<template>
	<v-form @submit.prevent="submit" ref="frm" :disabled="disabled" v-model="valid">
		<v-card>
			<v-card-title>Prijava</v-card-title>
			<v-divider></v-divider>
			<v-card-text>
				<v-row dense>
					<v-col>
						<v-text-field
							outlined
							label="Korisničko ime"
							v-model="username"
						></v-text-field>
					</v-col>
				</v-row>
				<v-row>
					<v-col>
						<v-text-field
							outlined
							label="Lozinka"
							type="password"
							v-model="password"
						></v-text-field>
					</v-col>
				</v-row>
			</v-card-text>
			<v-divider></v-divider>
			<v-card-actions>
				<v-spacer></v-spacer>
				<v-btn color="primary" type="submit" :disabled="disabled">Login</v-btn>
				<v-btn @click.prevent="reset" :disabled="disabled">Očisti</v-btn>
			</v-card-actions>
		</v-card>
	</v-form>
</template>

<script>
	export default {
		name: 'Login',
		data()
		{
			return {
				disabled: false,
				username: null,
				password: null,
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

				await this.$user.login(this.username, this.password);

				this.disabled = false;
			}
		}
	};
</script>
