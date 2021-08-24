<template>
	<v-form ref="frm" @submit.prevent="submit" v-model="valid" :disabled="loading">
		<v-card elevation="3">
			<v-card-title>Dodaj</v-card-title>
			<v-divider></v-divider>
			<v-card-text>
				<v-row>
					<v-col>
						<v-text-field
							outlined
							v-model="item.name"
							label="Naziv"
							:rules="[$v.required]"
						></v-text-field>
						<v-textarea
							outlined
							v-model="item.description"
							label="Opis"
							:rules="[$v.required]"
							counter
						></v-textarea>
					</v-col>
					<v-divider vertical></v-divider>
					<v-col>
						<v-row dense v-for="(s, idx) in item.specific_fields" :key="`sf-${idx}`">
							<v-col>
								<v-sheet outlined class="pa-3">
									<v-row dense>
										<v-col class="flex-md-grow-1">
											<v-text-field
												v-model="s.name"
												label="Naziv"
												outlined
												dense
												hide-details
											></v-text-field>
										</v-col>

										<v-col class="flex-md-grow-1">
											<v-select
												v-model="s.data_type"
												label="Tip podatka"
												:items="typeItems"
												outlined
												dense
												hide-details
											></v-select>
										</v-col>

										<v-col class="flex-shrink-1 flex-grow-0">
											<v-checkbox
												v-model="s.required"
												label="Obavezno"
												dense
												hide-details
											></v-checkbox>
										</v-col>

										<v-col class="flex-shrink-1 flex-grow-0">
											<v-btn icon @click.prevent="removeTypeItem(s)">
												<v-icon dense>mdi-delete</v-icon>
											</v-btn>
										</v-col>

									</v-row>
								</v-sheet>
							</v-col>
						</v-row>
						<v-row>
							<v-spacer></v-spacer>
							<v-btn small @click.prevent="addTypeItem" icon>
								<v-icon>mdi-plus</v-icon>
							</v-btn>
						</v-row>
					</v-col>
				</v-row>
			</v-card-text>
			<v-divider></v-divider>
			<v-card-actions>
				<v-spacer></v-spacer>
				<v-btn color="primary" :disabled="disabled || !valid" type="submit">Potvrdi</v-btn>
				<v-btn :disabled="disabled" @click.prevent="reset">Očisti</v-btn>
				<v-btn :disabled="disabled" @click.prevent="$emit('close')">Zatvori</v-btn>
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
				loading: false,
				valid: false,
				item: {
					name: null,
					description: null,
					specific_fields: []
				},
				typeItems: [
					{text: 'Cijeli broj', value: 'integer'},
					{text: 'Razlomljeni broj', value: 'float'},
					{text: 'Datum', value: 'date'},
					{text: 'Kratak tekst', value: 'string'},
					{text: 'Duži tekst', value: 'text'},
					{text: 'Istina/Laž', value: 'boolean'}
				]
			};
		},
		props: {
			id: {
				type: Number,
				default: null
			}
		},
		created()
		{
			if (this.id !== null)
			{
				this.fetch();
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
				const me = this;

				if (!this.$refs.frm.validate())
				{
					return;
				}

				this.loading = true;

				try
				{
					const {data} = await this.$http({
						url: '/api/v1/categories' + (this.id ? '/' + this.id : ''),
						data: {
							name: this.item.name,
							description: this.item.description,
							specific_fields: this.item.specific_fields.filter(item => item.name && item.data_type).map(item =>
							{
								delete item.sc_name;

								return item;
							})
						},
						method: this.id ? 'PATCH' : 'POST'
					});

					me.$emit('success', data);
				} catch (e)
				{
					me.$emit('failed', e);
				} finally
				{
					this.loading = false;
				}
			},
			async fetch()
			{
				this.loading = true;

				try
				{
					const {data} = await this.$http({
						url: `/api/v1/categories/${this.id}`
					});

					Object.keys(this.item).forEach(k => this.item[k] = data[k]);
				} catch (e)
				{
					this.$emit('failed', e);
				} finally
				{
					this.loading = false;
				}
			},
			addTypeItem()
			{
				this.item.specific_fields.push({
					name: null,
					data_type: null,
					required: false
				});
			},
			removeTypeItem(s)
			{
				this.item.specific_fields = this.item.specific_fields.filter(item => item !== s);
			}
		}
	};
</script>
