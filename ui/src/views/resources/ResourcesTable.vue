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
					:categories.sync="categories"
				></select-category>
			</v-col>
			<v-col md="3">
				<v-text-field
					label="Trazi"
					solo
					hide-details
					v-model="search"
					@keyup="fetch()"
				></v-text-field>
			</v-col>
		</v-row>

		<v-row dense v-if="!form.show && items.length > 0">
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
							<td>{{ item.user.first_name + ' ' + item.user.last_name }}</td>
							<td>{{ $format(new Date(item.created_at * 1000), 'dd.MM.yyyy. HH:mm:ss') }}</td>

							<td v-for="f in selectedSpecificFields" :key="f.sc_name">
								{{ item[f.sc_name] || '' }}
							</td>

							<td class="text-center">
								<table-menu-btn>
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
				staticHeaders: [{
					text: '#',
					value: 'id',
					width: 60,
					align: 'center',
					sortable: false
				}, {
					text: 'Naziv',
					value: 'name',
					width: 180
				}, {
					text: 'Napravio',
					value: 'user',
					width: 180
				}, {
					text: 'Datum',
					value: 'created_at',
					width: 180
				}],
				items: [],
				total: null,
				category: null,
				categories: [],
				search: null,
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
				let headers = [...this.staticHeaders];

				if (this.category !== null)
				{
					if (this.selectedCategory)
					{
						this.selectedSpecificFields.forEach(f =>
							headers.push({
								text: f.name,
								value: f.sc_name
							}));
					}
				}

				headers.push({
					text: '',
					value: 'action',
					width: 100
				});

				return headers;
			},
			selectedCategory()
			{
				return this.categories.find(item => item.id === this.category);
			},
			selectedSpecificFields()
			{
				const c = this.selectedCategory;

				return c ? c.specific_fields : [];
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
						url: '/api/v1/resources/?category_id=' + this.category + (this.search ? '&q=' + this.search : '')
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
