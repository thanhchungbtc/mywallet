<template>
  <div id="pageCategory">
    <v-container grid-list-xl fluid>
      <v-layout row wrap>
        <!--Category list-->
        <v-flex lg12 sm12 xs12>
          <!--Category list-->
          <v-card>

            <v-toolbar card dense color="transparent">
              <v-toolbar-title><h4>Category list</h4></v-toolbar-title>
              <v-btn flat icon color="primary" @click.stop="refresh">
                <v-icon>refresh</v-icon>
              </v-btn>
              <v-spacer></v-spacer>
              <v-btn color="primary" flat icon @click="gotoCreate">
                <v-icon>add</v-icon>
              </v-btn>
            </v-toolbar>

            <v-divider></v-divider>

            <v-card-text class="pa-0">
              <template>
                <v-data-table
                  :headers="headers"
                  :items="categories"
                  hide-actions
                  class="elevation-0"
                >
                  <template slot="items" slot-scope="props">
                    <tr @click="selectRow(props.item)">
                    <td>
                      <v-avatar size="36px">
                        <img :src="props.item.avatar_url" :alt="props.item.name"/>
                      </v-avatar>
                    </td>
                    <td>{{ props.item.name }}</td>
                    <td>{{ props.item.memo }}</td>
                    <td>
                      <v-progress-linear
                        :value="props.item.total_expense * 100 /props.item.budget"
                        height="5"
                        color="red">

                      </v-progress-linear>
                      <div class="my-3 text-sm-center"><strong class="error--text text--accent-3">${{
                        props.item.total_expense.toLocaleString() }}</strong> /
                        <strong class="success--text text--darken-3">${{ props.item.budget.toLocaleString() }}</strong>
                      </div>
                    </td>
                    <td class="text-xs-right">
                      <v-btn flat icon color="grey" @click.stop="edit(props.item.id)">
                        <v-icon>edit</v-icon>
                      </v-btn>
                      <v-btn flat icon color="grey" @click.prevent="deleteItem(props.item.id)">
                        <v-icon>delete</v-icon>
                      </v-btn>
                    </td>
                    </tr>
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
  import Category from "../../models/Category";

  export default {
    components: {
      VWidget,
    },
    data() {
      return {
        headers: [
          {
            text: '',
            align: 'center',
            sortable: false,
            value: 'imageUrl'
          },
          {
            text: 'Name',
            align: 'left',
            sortable: true,
            value: 'name'
          },
          {
            text: 'Memo',
            align: 'left',
            sortable: true,
            value: 'memo'
          },
          {text: 'Progress', value: 'progress', sortable: true,},
          {text: 'Actions', value: 'action', align: 'right', sortable: false,},

        ],
        categories: [],
      }
    },
    async created() {
      await this.refresh()
    },

    methods: {
      edit(id) {
        this.$router.push({name: 'category_edit', params: {id: id}})
      },

      async deleteItem(id) {
        if (!confirm("Are you sure?")) {
          return
        }
        console.log('delete')
        try {
          await Category.delete(id)
          this.$message.success(`Category with id: ${id} deleted`)
          this.categories = this.categories.filter(item => item.id !== id)
        } catch (e) {
          const msg = e.error || e.toLocaleString()
          this.$message.error('Error ' + msg)
        }
      },

      async refresh() {
        try {
          this.categories = await Category.all()
        } catch (e) {
          const msg = e.error || e.toLocaleString()
          this.$message.error("Something when wrong " + msg)
        }
      },

      gotoCreate() {
        this.$router.push({name: 'category_create'})
      },

      selectRow(row) {
        console.log(row)
        this.$router.push({name: 'expense_list', query: {'category': row.id}})
      },
    },
  }
</script>