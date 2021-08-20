<template>
	<v-select
		:items="categories"
		v-bind="$attrs"
		:disabled="loading || $attrs.disabled"
		:value="value"
		@input="input"
	></v-select>
</template>

<script>
	export default {
		name: 'SelectCategory',
		data()
		{
			return {
				categories: [],
				loading: false
			};
		},
		props: {value: null},
		created()
		{
			this.fetch();
		},
		methods: {
			async fetch()
			{
				this.loading = true;

				try
				{
					const {data} = await this.$http({
						url: '/api/v1/categories'
					});

					this.categories = data.items.map(c => ({
						text: c.name,
						value: c.id
					}));
				} catch (e)
				{
					console.warn(e);
				} finally
				{
					this.loading = false;

					this.$emit('loaded');
				}
			},
			input(v)
			{
				this.$emit('input', v);
				this.$emit('change', v);
			}
		}
	};
</script>
