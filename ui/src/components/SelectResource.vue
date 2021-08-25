<template>
	<v-select
		:items="resources"
		v-bind="$attrs"
		:disabled="loading || $attrs.disabled"
		:value="value"
		@input="input"
	></v-select>
</template>

<script>
	export default {
		name: 'SelectResource',
		data()
		{
			return {
				resources: [],
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
						url: '/api/v1/resources'
					});

					this.resources = data.items.map(c => ({
						text: c.name,
						value: c.id
					}));

					this.$emit('update:resources', data.items);
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
