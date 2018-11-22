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
                        v-model="expense.use_date"
                        label="Use date"
                        prepend-icon="event"
                        readonly
                      ></v-text-field>
                      <v-date-picker v-model="expense.use_date" @input="menu=false"></v-date-picker>
                    </v-menu>
                  </v-flex>

                  <v-flex>
                    <v-text-field v-model.number="expense.amount" label="Amount" type="number"></v-text-field>
                  </v-flex>
                  <v-flex>
                    <v-select
                      item-text="name"
                      item-value="id"
                      :items="categories"
                      v-model="expense.category_id"
                      label="Category">
                    </v-select>
                  </v-flex>
                  <v-flex>
                    <v-select
                      item-text="name"
                      item-value="id"
                      :items="accounts"
                      v-model="expense.account_id"
                      label="Account">
                    </v-select>
                  </v-flex>
                  <v-flex>
                    <v-text-field v-model="expense.location" label="Location"></v-text-field>
                  </v-flex>
                  <v-flex>
                    <v-text-field v-model="expense.memo" label="Memo"></v-text-field>
                  </v-flex>

                </v-layout>
              </v-container>
            </v-card-text>
            <v-card-actions>
              <v-spacer></v-spacer>
              <v-btn color="default" @click.native="back">Cancel</v-btn>
              <v-btn color="success" @click.native="save">Save</v-btn>
            </v-card-actions>

          </v-card>
        </v-flex>
      </v-layout>
    </v-container>
  </div>

</template>

<script>
  import Expense from "../../models/Expense";
  import Category from "../../models/Category";
  import Account from "../../models/Account";

  export default {
    data: () => ({
      menu: false,
      loading: false,
      categories: [],

      accounts: [],
      expense: {
        use_date: '',
        budget: 0,
        memo: '',
      }
    }),
    computed: {
      formTitle() {
        return !!this.$route.params.id ? 'Update expense' : 'Create expense'
      }
    },
    async created() {
      const id = this.$route.params.id
      this.categories = await Category.all()
      this.accounts = await Account.all()
      if (id) {
        this.loading = true
        try {
          this.expense = await Expense.findByID(id)

        } catch (e) {
        }
        this.loading = false
      }
    },

    methods: {
      back() {
        this.$router.push({name: 'expense_list'})
      },
      async save() {
        const id = this.$route.params.id
        this.loading = true
        if (!id) {
          try {
            await Expense.create(this.expense)
            this.$store.dispatch('app/success', `Expense ${this.expense.name} created`)
            this.$router.push({name: 'expense_list'})
          } catch (e) {
            const msg = e.error || e.toLocaleString()
            this.$store.dispatch('app/error', 'Error: ' + msg)
          }
        } else {
          try {
            await Expense.update(id, this.expense)
            this.$store.dispatch('app/success', `Expense ${this.expense.name} updated`)
            this.$router.push({name: 'expense_list'})
          } catch (e) {
            const msg = e.error || e.toLocaleString()
            this.$store.dispatch('app/error', 'Error: ' + msg)
          }
        }
        this.loading = false
      }

    }
  }
</script>
