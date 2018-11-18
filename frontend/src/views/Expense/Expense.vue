<template>
  <div id="pageExpense">
    <v-container grid-list-xl fluid>
      <v-layout row wrap>

        <v-flex lg12 sm12 xs12>
          <!--Expense list-->
          <v-card>

            <v-toolbar card dense color="transparent">
              <v-toolbar-title><h4>Expense list</h4></v-toolbar-title>

              <v-spacer></v-spacer>
              <v-btn color="primary" flat icon @click.native="gotoCreate">
                <v-icon>add</v-icon>
              </v-btn>

            </v-toolbar>

            <v-divider></v-divider>

            <!--table-->
            <v-card-text class="pa-0">
              <template>
                <v-data-table
                  :headers="headers"
                  :items="items"
                  hide-actions
                  class="elevation-0"
                >
                  <template slot="items" slot-scope="props">
                    <td>
                      <v-avatar size="36px">
                        <img :src="props.item.category.image_url" :alt="props.item.username"/>
                      </v-avatar>
                    </td>
                    <td>{{ props.item.use_date | formatDate }}</td>
                    <td>${{ props.item.amount.toLocaleString() }}</td>
                    <td>{{ props.item.location }}</td>
                    <td>{{ props.item.memo }}</td>
                    <td class="text-xs-right">
                      <v-btn flat icon color="grey" @click.stop="edit(props.item)">
                        <v-icon>edit</v-icon>
                      </v-btn>
                      <v-btn flat icon color="grey">
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
  import moment from 'moment'
  import DatePicker from '../../components/DatePicker'

  export default {
    components: {
      DatePicker,
    },
    data() {
      return {
        isUpdate: false,
        defaultItem: {
          use_date: new Date(),
          amount: 0,
          account_id: null,
          category_id: null,
          location: '',
          memo: '',
        },
        item: {
          use_date: new Date(),
          amount: 0,
          account_id: null,
          category_id: null,
          location: '',
          memo: '',
        },

        dialog: false,

        headers: [
          {
            text: '',
            align: 'center',
            sortable: false,
            value: 'category.image_url'
          },

          {text: 'Use date', value: 'use_date', sortable: true,},
          {text: 'Amount', value: 'amount', sortable: true,},
          {text: 'Location', value: 'location', sortable: true,},
          {text: 'Memo', value: 'memo', sortable: true,},
          {text: 'Actions', value: 'action', align: 'right', sortable: false,},

        ],
      };
    },

    computed: {
      formTitle() {
        return this.isUpdate === true ? 'Update category' : 'Add category'
      },

      items() {
        return [
          {
            use_date: moment('2018-11-12T13:47:41.485Z'),
            amount: 10000,
            account: {
              id: 1,
              name: 'cash',
              image_url: '/images/category/food.png',
            },
            category: {
              id: 1,
              name: 'Food',
              image_url: '/images/category/food.png',
            },
            location: 'tsudanuma',
            memo: 'this is a memo',
            color: 'pink',
          },
        ]
      }
    },

    methods: {
      gotoCreate() {
        this.$router.push({name: 'expense_create'})
      }

    }
  }
</script>