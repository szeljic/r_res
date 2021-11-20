<template>
	<v-form @submit.prevent="submit" ref="frm" :disabled="disabled" v-model="valid">
		<v-card>
			<v-card-title>Registracija</v-card-title>
			<v-divider></v-divider>
			<v-card-text>

				<v-row v-if="error !== false">
					<v-col>
						<v-alert
							type="error"
							dismissible
							@input="error = false"
						>{{ error }}
						</v-alert>
					</v-col>
				</v-row>

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
						<v-menu
							v-model="menuDOB"
							:close-on-content-click="false"
							:nudge-right="40"
							:nudge-top="20"
							transition="scale-transition"
							offset-y
							min-width="auto"
						>
							<template v-slot:activator="{ on, attrs }">
								<v-text-field
									label="Datum rođenja"
									outlined
									prepend-inner-icon="mdi-calendar"
									readonly
									v-bind="attrs"
									v-on="on"
									clearable
									:value="textDateDOB"
									:rules="[$v.required]"
								></v-text-field>
							</template>
							<v-date-picker
								v-model="item.date_of_birth"
								@input="menuDOB = false"
								:rules="[$v.required]"
							></v-date-picker>
						</v-menu>
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
							:rules="[$v.required, rulePasswordAgain]"
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
				error: false,
				item: {
					first_name: null,
					last_name: null,
					date_of_birth: null,
					email: null,
					username: null,
					password: null
				},
				password: null,
				valid: null,

				menuDOB: false
			};
		},
		computed: {
			textDateDOB()
			{
				if (!this.item.date_of_birth)
				{
					return null;
				}

				return this.$dateFormatL18n(this.$dateParseISO(this.item.date_of_birth));
			}
		},
		methods: {
			reset()
			{
				this.$refs.frm.reset();
				this.$refs.frm.resetValidation();

				this.error = false;
			},
			async submit()
			{
				if (!this.$refs.frm.validate())
				{
					return;
				}

				this.disabled = true;

				try
				{
					const {data} = await this.$http({
						url: '/api/v1/auth/registration',
						data: this.item,
						method: 'POST'
					});

					if (data.code === -1)
					{
						this.error = data.message;
					} else if (data.code === 200)
					{
						this.$emit('registered', this.item.username);
					}
				} catch (e)
				{
					console.warn(e);

					this.error = e.response.data.message || 'Internal server error';
				}

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
