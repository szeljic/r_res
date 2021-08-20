<template>
	<v-form @submit.prevent="submit" ref="frm" :disabled="disabled" v-model="valid">
		<v-card>
			<v-card-title>Registracija</v-card-title>
			<v-divider></v-divider>
			<v-card-text>
				<v-row dense>
					<v-col xs="12" sm="12" md="6" lg="6" xl="6">
						<v-text-field
							outlined
							label="Ime"
							v-model="item.first_name"
							:rules="[$v.required, $v.minLength(4)]"
						></v-text-field>
					</v-col>
					<v-col xs="12" sm="12" md="6" lg="6" xl="6">
						<v-text-field
							outlined
							label="Prezime"
							v-model="item.last_name"
							:rules="[$v.required, $v.minLength(4)]"
						></v-text-field>
					</v-col>
				</v-row>
				<v-row dense>
					<v-col>
						<v-text-field
							outlined
							label="Datum rođenja"
							v-model="item.date_of_birth"
							:rules="[$v.required]"
						></v-text-field>
					</v-col>
				</v-row>
				<v-row dense>
					<v-col>
						<v-text-field
							outlined
							label="eMail"
							v-model="item.email"
							:rules="[$v.required, $v.email]"
						></v-text-field>
					</v-col>
				</v-row>
				<v-row dense>
					<v-col>
						<v-text-field
							outlined
							label="Korisničko ime"
							v-model="item.username"
							:rules="[$v.required, $v.minLength(5)]"
						></v-text-field>
					</v-col>
				</v-row>
				<v-row dense>
					<v-col>
						<v-text-field
							outlined
							label="Lozinka"
							v-model="item.password"
							type="password"
							:rules="[$v.required, $v.minLength(8), rulePassword]"
						></v-text-field>
					</v-col>
				</v-row>
				<v-row dense>
					<v-col>
						<v-text-field
							outlined
							label="Potvrda lozinke"
							v-model="password"
							type="password"
							:rules="[rulePasswordAgain]"
						></v-text-field>
					</v-col>
				</v-row>
			</v-card-text>
			<v-divider></v-divider>
			<v-card-actions>
				<v-spacer></v-spacer>
				<v-btn color="primary" :disabled="disabled || !valid" type="submit">Potvrdi</v-btn>
				<v-btn :disabled="disabled" @click.prevent="reset">Očisti</v-btn>
			</v-card-actions>
		</v-card>
	</v-form>
</template>

<script>
	export default {
		name: 'Registration',
		data()
		{
			return {
				disabled: false,
				item: {
					first_name: null,
					last_name: null,
					date_of_birth: null,
					email: null,
					username: null,
					password: null
				},
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
				if (!this.$refs.frm.validate())
				{
					return;
				}

				this.disabled = true;

				await this.$http({
					url: '/api/v1/auth/registration',
					data: this.item,
					method: 'POST'
				});

				this.disabled = false;
			},
			rulePassword(v)
			{
				return /\w/.test(v) && /\d/.test(v) && /\W/.test(v);
			},
			rulePasswordAgain()
			{
				return this.item.password === this.password;
			}
		}
	};
</script>
