<template>
  <div id="categoryForm">
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
                <v-form :diabled="loading">

                  <v-layout row wrap>
                    <v-flex xs12>
                      <v-text-field v-model="category.name" label="Name" placeholder="Category name"></v-text-field>
                    </v-flex>

                    <v-flex xs12>
                      <v-text-field v-model="category.budget" type="number" label="Budget"
                                    placeholder="Budget for category"></v-text-field>
                    </v-flex>

                    <v-flex xs12>
                      <v-textarea v-model="category.memo" label="Memo"
                                  placeholder="Description about category"></v-textarea>
                    </v-flex>

                    <v-btn color="default" @click.prevent="back">Cancel</v-btn>
                    <v-btn color="info" @click.prevent="save">Save</v-btn>
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
  import Category from "../../models/Category";

  export default {
    data: () => ({
      loading: false,
      category: {
        name: '',
        budget: 0,
        memo: '',
      }
    }),
    computed: {
     formTitle() {
       return !!this.$route.params.id ? 'Update category' : 'Create category'
     }
    },
    async created() {
      const id = this.$route.params.id
      if (id) {
        this.loading = true
        try {
          this.category = await Category.findByID(id)
        } catch (e) {

        }
        this.loading = false
      }
    },

    methods: {
      back() {
        this.$router.push({name: 'category_list'})
      },
      async save() {
        const id = this.$route.params.id
        this.loading = true
        if (!id) {
          try {
            await Category.create(this.category)
          } catch (e) {

          }
        } else {
          try {
            await Category.update(id, this.category)
          } catch (e) {

          }
        }
        this.loading = false
      }

    }
  }
</script>