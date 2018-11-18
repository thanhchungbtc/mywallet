<template>
  <div id="expenseCreate">
    <v-container grid-list-xl fluid>
      <v-layout row wrap>
        <v-flex lg12 sm12 xs12>
          <v-card>

            <v-card-title>
              <span class="headline">{{ formTitle }}</span>
            </v-card-title>

            <v-divider></v-divider>

            <v-card-text>
              <v-container grid-list-md>
                <v-layout wrap>
                  <v-flex>
                    <v-menu
                      :close-on-content-click="false"
                      v-model="menu"
                      :nudge-right="40"
                      lazy
                      transition="scale-transition"
                      offset-y
                      full-width
                      min-width="290px"
                    >
                      <v-text-field
                        slot="activator"
                        v-model="item.use_date"
                        label="Use date"
                        prepend-icon="event"
                        readonly
                      ></v-text-field>
                      <v-date-picker v-model="item.use_date" @input="menu = false"></v-date-picker>
                    </v-menu>
                  </v-flex>

                  <v-flex>
                    <v-text-field v-model="item.amount" label="Amount" type="number"></v-text-field>
                  </v-flex>
                  <v-flex>
                    <v-select
                      item-text="name"
                      item-value="id"
                      :items="categories"
                      v-model="item.category_id"
                      label="Category">
                    </v-select>
                  </v-flex>
                  <v-flex>
                    <v-select
                      item-text="name"
                      item-value="id"
                      :items="accounts"
                      v-model="item.account_id"
                      label="Account">
                    </v-select>
                  </v-flex>
                  <v-flex>
                    <v-text-field v-model="item.location" label="Location"></v-text-field>
                  </v-flex>
                  <v-flex>
                    <v-text-field v-model="item.memo" label="Memo"></v-text-field>
                  </v-flex>

                </v-layout>
              </v-container>
            </v-card-text>
            <v-card-actions>
              <v-spacer></v-spacer>
              <v-btn color="success" @click.native="handleSubmit">Save</v-btn>
            </v-card-actions>

          </v-card>
        </v-flex>
      </v-layout>
    </v-container>
  </div>

</template>

<script>
  import moment from 'moment'

  export default {
    data: () => ({
      menu: false,
      item: {
        use_date: moment().format('YYYY-MM-DD'),
        amount: 0,
        category_id: '',
        account_id: '',
        location: '',
        memo: ''
      },
      categories: [
        {id: 1, name: 'Food'},
      ],
      accounts: [
        {id: 1, name: 'Food'},
      ],
    }),
    computed: {
      formTitle() {
        return this.isUpdate === true ? 'Update Expense' : 'Add Expense'
      },
    },
    methods: {
      handleSubmit() {
        console.log(this.item)
      }
    }
  }

</script>