<template>
	<v-form ref="frm" @submit.prevent="submit" v-model="valid" :disabled="loading">
		<v-card elevation="3">
			<v-card-title>Dodaj</v-card-title>
			<v-divider></v-divider>
			<v-card-text>
				<v-row v-if="errorMessage">
					<v-col cols="12">
						<v-alert
							type="error"
							dismissible
							@input="errorMessage = null"
						>{{ errorMessage }}
						</v-alert>
					</v-col>
				</v-row>
				<v-row>
					<v-col sm="12" md="12" lg="6" xl="6">
						<v-text-field
							v-model="item.name"
							label="Naziv"
							outlined
							hide-details
							:rules="[$v.required]"
						></v-text-field>
					</v-col>
					<v-col sm="12" md="12" lg="6" xl="6">
						<select-category
							v-model="item.category"
							label="Kategorija"
							outlined
							hide-details
							@change="changeCategory"
							@loaded="loadedCategory"
							:rules="[$v.required]"
							clearable
						></select-category>
					</v-col>
				</v-row>
				<v-row v-if="item.category">
					<v-col>
						<v-divider></v-divider>
					</v-col>
				</v-row>
				<v-row>
					<v-col sm="12" md="12" lg="6" xl="6" v-for="(f, idx) in fields" :key="`field-${idx}`">
						<v-text-field
							:label="f.name"
							outlined
							hide-details
							:rules="f.required ? [$v.required] : []"
							v-if="f.data_type === 'integer'"
							v-model="item.fields[f.sc_name]"
						></v-text-field>

						<v-text-field
							:label="f.name"
							outlined
							hide-details
							:rules="f.required ? [$v.required] : []"
							v-if="f.data_type === 'float'"
							v-model="item.fields[f.sc_name]"
						></v-text-field>

						<v-text-field
							:label="f.name"
							outlined
							hide-details
							:rules="f.required ? [$v.required] : []"
							v-if="f.data_type === 'string'"
							v-model="item.fields[f.sc_name]"
						></v-text-field>

						<v-textarea
							:label="f.name"
							outlined
							hide-details
							:rules="f.required ? [$v.required] : []"
							v-if="f.data_type === 'text'"
							v-model="item.fields[f.sc_name]"
							counter
						></v-textarea>

						<v-checkbox
							:label="f.name"
							hide-details
							:rules="f.required ? [$v.required] : []"
							v-if="f.data_type === 'boolean'"
							v-model="item.fields[f.sc_name]"
						></v-checkbox>

						<v-menu
							v-model="menus[f.sc_name]"
							:close-on-content-click="false"
							:nudge-right="40"
							transition="scale-transition"
							offset-y
							min-width="auto"
							v-if="f.data_type === 'date'"
						>
							<template v-slot:activator="{ on, attrs }">
								<v-text-field
									:label="f.name"
									v-model="item.fields[f.sc_name]"
									outlined
									prepend-icon="mdi-calendar"
									readonly
									v-bind="attrs"
									v-on="on"
									clearable
								></v-text-field>
							</template>
							<v-date-picker
								v-model="item.fields[f.sc_name]"
								@input="menus[f.sc_name] = false"
							></v-date-picker>
						</v-menu>
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
	import SelectCategory from '@/components/SelectCategory';

	export default {
		name: 'ResourcesFormComponent',
		components: {
			SelectCategory
		},
		data()
		{
			return {
				valid: null,
				loading: false,
				disabled: false,
				errorMessage: null,
				item: {
					name: null,
					category: null,
					fields: {}
				},
				fields: [],
				menus: {}
			};
		},
		props: {
			preCategory: {
				type: Number,
				default: null
			},
			id: {
				type: Number,
				default: 0,
				required: false
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
					const submitData = {
						name: this.item.name,
						category_id: this.item.category
					};

					this.fields.forEach(f =>
					{
						let value = this.item.fields[f.sc_name];

						switch (f.data_type)
						{
							case 'integer':
							case 'float':
								if (!isNaN(parseFloat(value)) && isFinite(value))
								{
									value = Number(value);
								}
								break;
							case 'text':
							case 'string':
								value = String(value);
								break;
							case 'boolean':
								break;
							default:
								value = String(value);
						}

						submitData[f.sc_name] = value;
					});

					const {data} = await this.$http({
						url: '/api/v1/resources' + (this.id ? '/' + this.id : ''),
						data: submitData,
						method: this.id ? 'PATCH' : 'POST'
					});

					me.$emit('success', data);
				} catch (e)
				{
					console.warn(e);
					me.$emit('failed', e);

					me.errorMessage = e.response.data.message;
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
						url: `/api/v1/resources/${this.id}`
					});

					this.item.name = data.name;
					this.item.category = data.category_id;

					data.category.specific_fields.forEach(f => this.item.fields[f.sc_name] = data[f.sc_name]);
				} catch (e)
				{
					this.$emit('failed', e);
				} finally
				{
					this.loading = false;
				}
			},
			async changeCategory()
			{
				if (!this.item.category)
				{
					this.fields = [];

					return;
				}

				this.loading = true;

				try
				{
					const {data} = await this.$http({
						url: '/api/v1/categories/' + this.item.category
					});

					this.fields = data.specific_fields.map(f =>
					{
						f.value = null;

						return f;
					});
				} catch (e)
				{
					console.warn(e);
				} finally
				{
					this.loading = false;
				}
			},
			loadedCategory()
			{
				this.item.category = this.preCategory;

				this.changeCategory();
			}
		}
	};
</script>
