<template>
	<v-form ref="frm" @submit.prevent="submit" v-model="valid">
		<v-card>
			<v-card-title>Dodaj</v-card-title>
			<v-card-text>
				<v-text-field
					outlined
					v-model="item.name"
					label="Ime"
				></v-text-field>
				<v-textarea
					outlined
					v-model="item.description"
					label="Opis"
				></v-textarea>
				<v-divider></v-divider>
			</v-card-text>
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
		name: 'UsersForm',
		data()
		{
			return {
				disabled: false,
				valid: false,
				item: {
					name: null
				}
			};
		},
		props: {
			id: {
				type: Number
			}
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

				const response = await this.$http({
					url: '/api/v1/categories',
					data: this.item,
					method: 'POST'
				});

				console.log(response);
			}
		}
	};
</script>
