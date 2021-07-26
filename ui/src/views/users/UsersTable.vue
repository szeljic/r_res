<template>
	<v-container>
		<v-row>
			<v-col>
				<v-toolbar dense elevation="2">
					<v-toolbar-title>Korisnici</v-toolbar-title>
				</v-toolbar>
			</v-col>
		</v-row>
		<v-row>
			<v-col>
				<v-data-table
					:headers="headers"
					:items="items"
					:server-items-length="total"
					:items-per-page="10"
					class="elevation-2"
					no-data-text="Nema podataka"
					:loading="loading"
				>
					<template v-slot:item="{item}">
						<tr>
							<td>{{item.ID}}</td>
							<td>{{item.FirstName}} {{item.LastName}}</td>
							<td>{{item.Username}}</td>
							<td>{{item.Email}}</td>
							<td class="text-center">
								<v-btn :to="`/korisnici/uredi/${item.ID}`" icon>
									<v-icon>mdi-menu</v-icon>
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
		data() {
			return {
				headers: [{
					text:'#',
					value: 'ID',
					width: 100
				}, {
					text: 'Ime i prezime',
					value: 'FirstName'
				}, {
					text: 'Korisnicko ime',
					value: 'Username',
					width: 240
				}, {
					text: 'Email',
					value: 'Email',
					width: 300
				}, {
					text: '',
					value: null,
					sortable: false,
					filterable: false,
					width: 100,
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
				this.items = data.users;

				this.loading = false;
			}
		}
	};
</script>
