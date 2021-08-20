<template>
	<v-form ref="frm" @submit.prevent="submit" v-model="valid" :disabled="loading">
		<v-card elevation="3">
			<v-card-title>Dodaj</v-card-title>
			<v-divider></v-divider>
			<v-card-text>
				<v-row>
					<v-col sm="12" md="12" lg="6" xl="6">
						<v-text-field
							v-model="item.name"
							label="Naziv"
							outlined
							hide-details
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
							:rules="f.required ? [[$v.required]] : []"
							v-if="f.data_type === 'integer'"
							v-model="item.fields[f.sc_name]"
						></v-text-field>

						<v-text-field
							:label="f.name"
							outlined
							dense
							hide-details
							:rules="f.required ? [[$v.required]] : []"
							v-if="f.data_type === 'float'"
							v-model="item.fields[f.sc_name]"
						></v-text-field>
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
				item: {
					name: null,
					category: null,
					fields: {}
				},
				fields: []
			};
		},
		props: {
			preCategory: {
				type: Number,
				default: null
			}
		},
		methods: {
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

						console.log(f);

						switch (f.data_type)
						{
							case 'integer':
							case 'float':
								value = Number(value);
								break;
							case 'text':
							case 'string':
								value = String(value);
						}

						submitData[f.sc_name] = value;
					});

					const {data} = await this.$http({
						url: '/api/v1/resources',
						data: submitData,
						method: 'POST'
					});

					me.$emit('success', data);
				} catch (e)
				{
					console.log(e);
					me.$emit('failed', e);
				} finally
				{
					this.loading = false;
				}
			},
			async changeCategory()
			{
				if (!this.item.category)
				{
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
