<template>
	<v-container>
		<v-row dense>
			<v-col>
				<v-toolbar dense elevation="2">
					<v-toolbar-title>Korisnici</v-toolbar-title>
					<v-spacer></v-spacer>
					<v-btn icon @click.prevent="fetch">
						<v-icon>mdi-refresh</v-icon>
					</v-btn>
				</v-toolbar>
			</v-col>
		</v-row>
		<v-row dense>
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
					:footer-props="{
						itemsPerPageText: 'Redova po stranici',
						pageText: '{0}-{1} od {2}'
					}"
				>
					<template v-slot:item="{item}">
						<tr>
							<td class="text-center">{{ item.id }}</td>
							<td>{{ item.first_name }} {{ item.last_name }}</td>
							<td>{{ item.username }}</td>
							<td>{{ item.email }}</td>
							<td>{{ $dateFormatL18n(new Date(item.date_of_birth)) }}</td>
							<td class="text-center">
								<v-checkbox
									hide-details
									class="pa-0 ma-0 ml-2"
									:disabled="loading"
									v-model="item.status"
									@change="updateStatus(item)"
									:color="item.status ? 'green': 'grey'"
								></v-checkbox>
							</td>
						</tr>
					</template>
				</v-data-table>
			</v-col>
		</v-row>
	</v-container>
</template>

<script>
	export default {
		name: 'UsersTable',
		data()
		{
			return {
				headers: [{
					text: '#',
					value: 'ID',
					width: 100,
					align: 'center'
				}, {
					text: 'Ime i prezime',
					value: 'first_name'
				}, {
					text: 'Korisnicko ime',
					value: 'username',
					width: 240
				}, {
					text: 'Email',
					value: 'email',
					width: 300
				}, {
					text: 'Datum rodjenja',
					value: 'date_of_birth',
					width: 140
				}, {
					text: 'Status',
					value: 'status',
					align: 'center',
					width: 90
				}],
				items: [],
				total: null,
				loading: false
			};
		},
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
						url: '/api/v1/users'
					});

					this.total = data.total;
					this.items = data.items;
				} catch (e)
				{
					console.warn(e);
				}

				this.loading = false;
			},
			async updateStatus(item)
			{
				this.loading = true;

				try
				{
					await this.$http({
						method: 'PATCH',
						url: '/api/v1/users/' + item.id,
						data: {
							status: item.status ? '1' : '0'
						}
					});
				} catch (e)
				{
					console.warn(e);
				}

				this.loading = false;
			}
		}
	};
</script>
