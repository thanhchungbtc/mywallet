<template>
  <div id="pageCategory">
    <v-container grid-list-xl fluid>
      <v-layout row wrap>
        <!--Account list-->
        <v-flex lg12 sm12 xs12>
          <!--Account list-->
          <v-card>

            <v-toolbar card dense color="transparent">
              <v-toolbar-title><h4>Account list</h4></v-toolbar-title>
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
                  :items="accounts"
                  hide-actions
                  class="elevation-0"
                >
                  <template slot="items" slot-scope="props">
                    <td>
                      <v-avatar size="36px">
                        <img :src="props.item.imageUrl" :alt="props.item.username"/>
                      </v-avatar>
                    </td>
                    <td>{{ props.item.name }}</td>
                    <td>{{ props.item.memo }}</td>
                    <td class="text-xs-left">
                      <v-progress-linear :value="props.item.progress" height="5"
                                         :color="props.item.color"></v-progress-linear>
                      <div class="my-3 text-sm-center"><strong class="error--text text--accent-3">$1,000</strong> /
                        <strong class="success--text text--darken-3">$15,000</strong></div>
                    </td>
                    <td class="text-xs-right">
                      <v-btn flat icon color="grey" @click.stop="edit(props.item.id)">
                        <v-icon>edit</v-icon>
                      </v-btn>
                      <v-btn flat icon color="grey" @click.prevent="deleteItem(props.item.id)">
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
  import Account from "../../models/Account";

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
        accounts: [],
      }
    },
    async created() {
      await this.refresh()
    },

    methods: {
      edit(id) {
        this.$router.push({name: 'account_edit', params: {id: id}})
      },

      async deleteItem(id) {
        if (!confirm("Are you sure?")) {
          return
        }
        console.log('delete')
        try {
          await Account.delete(id)
          this.$message.success(`Account with id: ${id} deleted`)
          this.accounts = this.accounts.filter(item => item.id !== id)
        } catch (e) {
          const msg = e.error || e.toLocaleString()
          this.$message.error('Error ' + msg)
        }
      },

      async refresh() {
        try {
          this.accounts = await Account.all()
        } catch (e) {
          const msg = e.error || e.toLocaleString()
          this.$message.error("Something when wrong " + msg)
        }
      },

      gotoCreate() {
        this.$router.push({name: 'account_create'})
      }
    },
  }
</script>