<template>
  <div id="accountForm">
    <v-container grid-list-xl fluid>
      <v-layout wrap>
        <v-flex lg12 sm12 xs12>
          <v-card>
            <v-progress-linear v-show="loading" :indeterminate="true"></v-progress-linear>

            <v-toolbar card dense color="transparent">
              <v-toolbar-title><h4>{{formTitle}}</h4></v-toolbar-title>
              <v-spacer></v-spacer>
            </v-toolbar>
            <v-divider></v-divider>

            <v-card-text class="pa-3">
              <template>
                <v-form>

                  <v-layout row wrap>
                    <v-flex xs12>
                      <v-text-field v-model="account.name" label="Name" placeholder="Account name"></v-text-field>
                    </v-flex>

                    <v-flex xs12>
                      <v-text-field v-model.number="account.budget" type="number" label="Budget"
                                    placeholder="Budget for account"></v-text-field>
                    </v-flex>

                    <v-flex xs12>
                      <v-textarea v-model="account.memo" label="Memo"
                                  placeholder="Description about account"></v-textarea>
                    </v-flex>

                    <v-btn color="default" :disabled="loading" @click.prevent="back">Cancel</v-btn>
                    <v-btn color="info" :disabled="loading" @click.prevent="save">Save</v-btn>
                  </v-layout>
                </v-form>
              </template>
            </v-card-text>
          </v-card>
        </v-flex>
      </v-layout>
    </v-container>
  </div>
</template>
<script>
  import Account from "../../models/Account";

  export default {
    data: () => ({
      loading: false,
      account: {
        name: '',
        budget: 0,
        memo: '',
      }
    }),
    computed: {
     formTitle() {
       return !!this.$route.params.id ? 'Update account' : 'Create account'
     }
    },
    async created() {
      const id = this.$route.params.id
      if (id) {
        this.loading = true
        try {
          this.account = await Account.findByID(id)
        } catch (e) {

        }
        this.loading = false
      }
    },

    methods: {
      back() {
        this.$router.push({name: 'account_list'})
      },
      async save() {
        const id = this.$route.params.id
        this.loading = true
        if (!id) {
          try {
            await Account.create(this.account)
            this.$store.dispatch('app/success', `Account ${this.account.name} created`)
            this.$router.push({name: 'account_list'})
          } catch (e) {
            const msg = e.error || e.toLocaleString()
            this.$store.dispatch('app/error', 'Error: ' + msg)
          }
        } else {
          try {
            await Account.update(id, this.account)
            this.$store.dispatch('app/success', `Account ${this.account.name} updated`)
            this.$router.push({name: 'account_list'})
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