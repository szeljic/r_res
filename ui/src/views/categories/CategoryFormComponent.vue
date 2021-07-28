<template>
	<v-form ref="frm" @submit.prevent="submit" v-model="valid" :disabled="loading">
		<v-card elevation="3">
			<v-card-title>Dodaj</v-card-title>
			<v-divider></v-divider>
			<v-card-text>
				<v-text-field
					outlined
					v-model="item.name"
					label="Ime"
					:rules="[$v.required]"
				></v-text-field>
				<v-textarea
					outlined
					v-model="item.description"
					label="Opis"
					:rules="[$v.required]"
					counter
				></v-textarea>
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
					description: null
				}
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
							description: this.item.description
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
			}
		}
	};
</script>
