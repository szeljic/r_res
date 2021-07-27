<template>
	<v-dialog persistent v-model="internalShow" :width="width">
		<v-card>
			<v-card-title class="title mb-0 word-break">{{ title }}</v-card-title>
			<v-card-text>
				{{ message }}
			</v-card-text>
			<v-card-actions>
				<slot name="actions"></slot>
			</v-card-actions>
		</v-card>
	</v-dialog>
</template>

<script>
	export default {
		name: 'MessageBox',
		data()
		{
			return {};
		},
		props: {
			show: {
				type: Boolean,
				default: false
			},
			title: String,
			message: String,
			width: {
				type: Number,
				default: 480
			},
			loader: {
				type: Boolean,
				default: false
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
		}
	};
</script>
