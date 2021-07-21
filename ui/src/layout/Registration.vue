<template>
	<v-form @submit="submit" ref="frm" :disabled="disabled">
		<v-card>
			<v-card-title>Registracija</v-card-title>
			<v-divider></v-divider>
			<v-card-text>
				<v-row dense>
					<v-col>
						<v-text-field
							outlined
							label="Ime"
							v-model="item.first_name"
						></v-text-field>
					</v-col>
					<v-col>
						<v-text-field
							outlined
							label="Prezime"
							v-model="item.last_name"
						></v-text-field>
					</v-col>
				</v-row>
				<v-row dense>
					<v-col>
						<v-text-field
							outlined
							label="Datum rođenja"
							v-model="item.date_of_birth"
						></v-text-field>
					</v-col>
				</v-row>
				<v-row dense>
					<v-col>
						<v-text-field
							outlined
							label="eMail"
							v-model="item.email"
						></v-text-field>
					</v-col>
				</v-row>
				<v-row dense>
					<v-col>
						<v-text-field
							outlined
							label="Korisničko ime"
							v-model="item.username"
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
						></v-text-field>
					</v-col>
				</v-row>
			</v-card-text>
			<v-divider></v-divider>
			<v-card-actions>
				<v-spacer></v-spacer>
				<v-btn color="primary" :disabled="disabled" type="submit">Potvrdi</v-btn>
				<v-btn :disabled="disabled" @click.prevent="reset">Očisti</v-btn>
			</v-card-actions>
		</v-card>
	</v-form>
</template>

<script>
	import axios from 'axios';

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
				itemClone: null
			};
		},
		created()
		{
			this.itemClone = Object.assign({}, this.item);
		},
		methods: {
			reset()
			{
				this.item = this.itemClone;
				this.password = null;
			},
			async submit()
			{
				const formData = new FormData();

				Object.entries(this.item).forEach(([k, v]) => formData.append(k, v));

				const response = await axios.post('/registration', formData);

				console.log(response);
			}
		}
	};
</script>
