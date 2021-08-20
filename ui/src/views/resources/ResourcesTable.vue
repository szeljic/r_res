<template>
	<v-container>
		<v-row dense>
			<v-col>
				<v-toolbar dense elevation="2">
					<v-toolbar-title>Resursi</v-toolbar-title>
					<v-spacer></v-spacer>
					<v-btn icon @click="showForm()" v-if="!form.show">
						<v-icon>mdi-plus</v-icon>
					</v-btn>
					<v-btn icon @click.prevent="fetch" v-if="!form.show">
						<v-icon>mdi-refresh</v-icon>
					</v-btn>
				</v-toolbar>
			</v-col>
		</v-row>

		<v-row dense v-if="form.show">
			<v-col>
				<resource-form-component
					:id="form.id"
					@success="success"
					@close="form.show = false"
					:pre-category="category"
				></resource-form-component>
			</v-col>
		</v-row>

		<v-row dense v-if="!form.show">
			<v-col md="3">
				<select-category
					label="Kategorija"
					v-model="category"
					:disabled="loading"
					@change="fetch()"
					hide-details
					solo
				></select-category>
			</v-col>
		</v-row>

		<v-row dense v-if="!form.show">
			<v-col>
				<v-data-table
					:headers="headers"
					:items="items"
					:server-items-length="total"
					:items-per-page="10"
					class="elevation-2"
					no-data-text="Nema podataka"
					no-results-text="Nema rezultata"
					:loading="loading"
				>
					<template v-slot:item="{item}">
						<tr>
							<td class="text-center">{{ item.id }}</td>
							<td>{{ item.name }}</td>
							<td>{{ item.description }}</td>
							<td>{{ item.created_by }}</td>
							<td>{{ item.created_at }}</td>
							<td class="text-center">
								<table-menu-btn :disabled="editing">
									<v-list dense>
										<v-list-item-group>
											<v-list-item @click.prevent="showForm(item)">
												<v-list-item-icon>
													<v-icon>mdi-pencil</v-icon>
												</v-list-item-icon>
												<v-list-item-content>
													<v-list-item-title>Uredi</v-list-item-title>
												</v-list-item-content>
											</v-list-item>

											<v-list-item>
												<v-list-item-icon>
													<v-icon>mdi-delete-forever-outline</v-icon>
												</v-list-item-icon>
												<v-list-item-content>
													<v-list-item-title>Briši</v-list-item-title>
												</v-list-item-content>
											</v-list-item>
										</v-list-item-group>
									</v-list>
								</table-menu-btn>
							</td>
						</tr>
					</template>
				</v-data-table>
			</v-col>
		</v-row>
	</v-container>
</template>

<script>
	import ResourceFormComponent from '@/views/resources/ResourceFormComponent';
	import SelectCategory from '@/components/SelectCategory';

	export default {
		name: 'ResourcesTable',
		components: {
			ResourceFormComponent,
			SelectCategory
		},
		data()
		{
			return {
				staticHeaders: [],
				items: [],
				total: null,
				category: null,
				categories: [],
				loading: false,
				form: {
					show: false,
					id: null
				}
			};
		},
		created()
		{
			this.fetch();
		},
		computed: {
			headers()
			{
				return [...this.staticHeaders];
			}
		},
		methods: {
			async fetch()
			{
				if (this.category === null)
				{
					return;
				}

				this.loading = true;

				try
				{
					const {data} = await this.$http({
						url: '/api/v1/resources'
					});

					this.total = data.total;

					this.items = data.items || [];
				} catch (e)
				{
					console.warn(e);
				} finally
				{
					this.loading = false;
				}
			},
			showForm(item)
			{
				this.form.id = item ? item.id : null;
				this.form.show = true;
			},
			success()
			{
				this.fetch();

				this.form.show = false;
			}
		}
	};
</script>
