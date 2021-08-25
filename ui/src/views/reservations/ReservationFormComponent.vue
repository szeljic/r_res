<template>
	<v-form ref="frm" @submit.prevent="submit" v-model="valid" :disabled="loading">
		<v-card elevation="3">
			<v-card-title>Dodaj</v-card-title>
			<v-divider></v-divider>
			<v-card-text>
				<v-row>
					<v-col sm="12" md="12" lg="6" xl="6">

						<v-row v-if="error !== false">
							<v-col>
								<v-alert type="error">{{ error }}</v-alert>
							</v-col>
						</v-row>

						<v-row>
							<v-col cols="12">
								<select-resource
									label="Resurs"
									v-model="item.resource"
									outlined
									hide-details
									@change="changeResource"
									:rules="[$v.required]"
									clearable
								></select-resource>
							</v-col>
						</v-row>
						<v-row>
							<v-col>
								<v-menu
									v-model="menuFrom"
									:close-on-content-click="false"
									:nudge-right="40"
									transition="scale-transition"
									offset-y
									min-width="auto"
								>
									<template v-slot:activator="{ on, attrs }">
										<v-text-field
											label="Rezervacija od"
											outlined
											prepend-icon="mdi-calendar"
											readonly
											v-bind="attrs"
											v-on="on"
											clearable
											:value="textDateFrom"
										></v-text-field>
									</template>
									<v-date-picker
										v-model="item.fromDate"
										@input="menuFrom = false"
										:rules="[$v.required]"
									></v-date-picker>
								</v-menu>
							</v-col>
							<v-col>
								<v-menu
									v-model="menuTo"
									:close-on-content-click="false"
									:nudge-right="40"
									transition="scale-transition"
									offset-y
									min-width="auto"
								>
									<template v-slot:activator="{ on, attrs }">
										<v-text-field
											label="Rezervacija do"
											outlined
											prepend-icon="mdi-calendar"
											readonly
											v-bind="attrs"
											v-on="on"
											clearable
											:value="textDateTo"
										></v-text-field>
									</template>
									<v-date-picker
										v-model="item.toDate"
										@input="menuTo = false"
										:rules="[$v.required]"
									></v-date-picker>
								</v-menu>
							</v-col>
						</v-row>
					</v-col>
					<v-divider vertical v-if="item.resource !== null"></v-divider>
					<v-col sm="12" md="12" lg="6" xl="6" v-if="item.resource !== null">
						<v-row>
							<v-col v-for="(x, idx) in aboutResource" :key="idx" cols="6">
								<v-text-field
									:label="x.text"
									readonly
									filled
									hide-details
									v-model="x.value"
									v-if="x.type === undefined || x.type === 'text' || x.type === 'integer' || x.type === 'float' || x.type === 'date'"
								></v-text-field>
								<v-switch
									:label="x.text"
									readonly
									hide-details
									v-model="x.value"
									v-if="x.type === 'boolean'"
								></v-switch>
							</v-col>
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
	import SelectResource from '@/components/SelectResource';

	export default {
		name: 'ReservationFormComponent',
		components: {
			SelectResource
		},
		data()
		{
			return {
				valid: null,
				loading: false,
				disabled: false,
				error: false,
				item: {
					name: null,
					resource: null,
					fromDate: null,
					toDate: null
				},
				resource: null,
				menuFrom: false,
				menuTo: false
			};
		},
		created()
		{
			window.op = this;
		},
		computed: {
			textDateFrom()
			{
				if (!this.item.fromDate)
				{
					return null;
				}

				return this.$dateFormatL18n(this.$dateParseISO(this.item.fromDate));
			},
			textDateTo()
			{
				if (!this.item.toDate)
				{
					return null;
				}

				return this.$dateFormatL18n(this.$dateParseISO(this.item.toDate));
			},
			aboutResource()
			{
				if (this.resource === null)
				{
					return null;
				}

				let obj = [{
					text: 'Naziv',
					value: this.resource.name
				}, {
					text: 'Kategorija',
					value: this.resource.category.name
				}];

				this.resource.category.specific_fields.forEach(f =>
				{
					obj.push({
						text: f.name,
						value: this.resource[f.sc_name],
						type: f.data_type
					});
				});

				return obj;
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
						url: '/api/v1/reservations',
						data: {
							name: this.item.name,
							resource_id: this.item.resource,
							from_date: this.item.fromDate + ' 00:00:00',
							to_date: this.item.toDate + ' 23:59:59'
						},
						method: 'POST'
					});

					me.$emit('success', data);
				} catch (e)
				{
					console.warn(e);

					this.error = e.response.data.message;

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
						url: `/api/v1/reservations/${this.id}`
					});

					this.item.resource = data.resource_id;
				} catch (e)
				{
					this.$emit('failed', e);
				} finally
				{
					this.loading = false;
				}
			},
			async changeResource()
			{
				if (!this.item.resource)
				{
					this.resource = null;

					return;
				}

				this.loading = true;

				try
				{
					const {data} = await this.$http({
						url: '/api/v1/resources/' + this.item.resource
					});

					this.resource = data;
				} catch (e)
				{
					console.warn(e);
				} finally
				{
					this.loading = false;
				}
			}
		}
	};
</script>
