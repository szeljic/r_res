<template>
	<message-box
		:title="'Brisanje?'"
		:message="'Da li želite da nastavite brisanje?'"
		:show="internalShow"
		:loader="loader"
	>
		<template v-slot:actions>
			<v-spacer></v-spacer>
			<v-btn text color="default" :disabled="loader" @click="internalShow = false">Odustani</v-btn>
			<v-btn text color="primary" :disabled="loader" @click="perform">
				<v-icon left small>mdi-delete</v-icon>
				Briši
			</v-btn>
		</template>
	</message-box>
</template>

<script>

	export default {
		name: 'delete-dialog',
		data()
		{
			return {
				loader: false
			};
		},
		props: {
			url: {
				type: String,
				required: true
			},
			show: {
				type: Boolean,
				default: false
			},
			closeOnSuccess: {
				type: Boolean,
				default: true
			},
			closeOnFail: {
				type: Boolean,
				default: true
			}
		},
		computed: {
			internalShow: {
				get()
				{
					return this.show;
				},
				set(v)
				{
					if (!v)
					{
						this.$emit('update:show', false);
					}
				}
			}
		},
		methods: {
			perform()
			{
				const me = this;

				me.loader = true;

				me.$emit('update:show', false);

				this.$http({
					url: me.url,
					method: 'DELETE'
				})
					.then(() => me.$emit('success'))
					.catch((reason) => me.$emit('fail', reason))
					.finally(() => me.loader = false);
			}
		}
	};
</script>
