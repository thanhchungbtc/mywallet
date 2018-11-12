<template>
  <div id="pageCategory">
    <v-container grid-list-xl fluid>
      <v-layout row wrap>
        <!--Category list-->
        <v-flex lg12 sm12 xs12>
          <category-table></category-table>
        </v-flex>

      </v-layout>
    </v-container>
  </div>
</template>

<script>
  import VWidget from '../components/VWidget'
  import CategoryTable from '../components/category/CategoryTable';

  export default {
    components: {
      VWidget,
      CategoryTable
    },
    data() {
      return {
        dialog: false,
      }
    },
    created() {


    }
  }
</script>