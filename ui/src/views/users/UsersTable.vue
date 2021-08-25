<template>
	<v-container>
		<v-row dense>
			<v-col>
				<v-toolbar dense elevation="2">
					<v-toolbar-title>Korisnici</v-toolbar-title>
					<v-spacer></v-spacer>
					<v-btn icon to="/korisnici/dodaj">
						<v-icon>mdi-plus</v-icon>
					</v-btn>
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
							<td>{{ item.date_of_birth }}</td>
							<td class="text-center">
								<status-icon v-model="item.status"></status-icon>
							</td>
							<td class="text-center">
								<v-btn small :to="`/korisnici/uredi/${item.id}`" icon>
									<v-icon>mdi-pencil</v-icon>
								</v-btn>
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
					value: 'statue',
					align: 'center',
					width: 90
				}, {
					text: '',
					value: null,
					sortable: false,
					filterable: false,
					width: 104,
					align: 'center'
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

				const {data} = await this.$http({
					url: '/api/v1/users'
				});

				this.total = data.total;
				this.items = data.items;

				this.loading = false;
			}
		}
	};
</script>
