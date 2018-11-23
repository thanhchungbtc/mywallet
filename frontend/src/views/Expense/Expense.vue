<template>
  <div id="pageExpense">
    <v-container grid-list-xl fluid>
      <v-btn
        class="mb-3 mr-3"
        color="pink"
        dark
        fixed
        bottom
        right
        fab
        @click.stop="gotoCreate()"
      >
        <v-icon>add</v-icon>
      </v-btn>
      <v-layout row wrap>

        <v-flex lg12 sm12 xs12>
          <!--Expense list-->
          <v-card>

            <v-toolbar card dense color="transparent">
              <v-toolbar-title><h4>Expense list</h4></v-toolbar-title>
              <v-btn flat icon color="primary" @click.stop="refresh">
                <v-icon>refresh</v-icon>
              </v-btn>
              <v-spacer></v-spacer>

            </v-toolbar>

            <v-divider></v-divider>

            <!--table-->
            <v-card-text class="pa-0">
              <template>
                <v-data-table
                  :headers="headers"
                  :items="expenses"
                  hide-actions
                  class="elevation-0"
                >
                  <template slot="items" slot-scope="props">
                    <td>
                      <v-avatar size="36px">
                        <img :src="props.item.category.avatar_url" :alt="props.item.category.name"/>
                      </v-avatar>
                    </td>
                    <td>{{ props.item.use_date | formatDate }}</td>
                    <td>${{ props.item.amount.toLocaleString() }}</td>
                    <td>{{ props.item.location }}</td>
                    <td>{{ props.item.memo }}</td>
                    <td class="text-xs-right">
                      <v-btn flat icon color="grey" @click.stop="edit(props.item.id)">
                        <v-icon>edit</v-icon>
                      </v-btn>
                      <v-btn flat icon color="grey" @click.stop="deleteItem(props.item.id)">
                        <v-icon>delete</v-icon>
                      </v-btn>
                    </td>
                  </template>

                </v-data-table>
              </template>

              <v-divider></v-divider>

            </v-card-text>

          </v-card>
        </v-flex>

      </v-layout>
    </v-container>
  </div>
</template>

<script>
  import VWidget from '../../components/VWidget'
  import Expense from "../../models/Expense";

  export default {
    components: {
      VWidget,
    },
    data() {
      return {
        headers: [
          {
            text: 'Category',
            align: 'center',
            sortable: true,
            value: 'category.avatar_url'
          },
          {
            text: 'Use date',
            align: 'left',
            sortable: true,
            value: 'use_date'
          },
          {
            text: 'Amount',
            align: 'left',
            sortable: true,
            value: 'amount'
          },
          {
            text: 'Location',
            align: 'left',
            sortable: true,
            value: 'location'
          },
          {
            text: 'Memo',
            align: 'left',
            sortable: true,
            value: 'memo'
          },
          {text: 'Actions', value: 'action', align: 'right', sortable: false,},
        ],
        expenses: [],
      }
    },
    async created() {
      await this.refresh()
    },

    methods: {
      edit(id) {
        this.$router.push({name: 'expense_edit', params: {id: id}})
      },

      async deleteItem(id) {
        if (!confirm("Are you sure?")) {
          return
        }
        console.log('delete')
        try {
          await Expense.delete(id)
          this.$message.success(`Expense with id: ${id} deleted`)
          this.expenses = this.expenses.filter(item => item.id !== id)
        } catch (e) {
          const msg = e.error || e.toLocaleString()
          this.$message.error('Error ' + msg)
        }
      },

      async refresh() {
        try {
          this.expenses = await Expense.all(this.$route.query)
        } catch (e) {
          const msg = e.error || e.toLocaleString()
          this.$message.error("Something when wrong " + msg)
        }
      },

      gotoCreate() {
        this.$router.push({name: 'expense_create'})
      }
    },
  }
</script>