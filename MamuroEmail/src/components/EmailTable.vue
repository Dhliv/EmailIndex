<script>
  import { MDBTable, MDBCol, MDBRow } from 'mdb-vue-ui-kit';
  import EmailTableRecord from './EmailTableRecord.vue'
  import EmailRecord from './EmailRecord.vue';
  import {getEmails} from '../requests/getEmails.js';
import Pagination from './Pagination.vue';
  export default {
    data() {
      return {
        emails: [],
        page: 1,
        showDetails: false,
        email: null,
        maxPage: 51743
      }
    },

    async beforeCreate() {
      this.emails = await getEmails(0)
    },

    methods: {
      onClickEmail: function(emailId) {
        this.showDetails = true;
        for(let i = 0; i < this.emails.length; i++){
          if(this.emails[i]._id == emailId){
            this.email = this.emails[i]._source;
            break
          }
        }
      },

      nextPage(delta) {
        this.page = Math.min(this.maxPage, this.page + delta)
      },

      previousPage(delta) {
        this.page = Math.max(1, this.page + delta)
      },

      async handlePageUpdate(delta) {
        if(delta < 0) this.previousPage(delta)
        else this.nextPage(delta)
        this.emails = await getEmails(this.page - 1)
      },

      async setPage(newPage) {
        this.page = newPage;
        this.emails = await getEmails(this.page - 1)
      }
    },

    components: {
    MDBTable,
    EmailTableRecord,
    EmailRecord,
    MDBCol,
    MDBRow,
    Pagination
},
  };
</script>

<template>
  <MDBCol md="7" class="container-size">
    <MDBRow class="container-table">
      <MDBTable hover striped class="border-table">
        <thead class="bg-light">
          <tr>
            <th>Subject</th>
            <th>From</th>
            <th>To</th>
          </tr>
        </thead>

        <tbody v-for="email in this.emails" :key="email._id">
          <EmailTableRecord :email=email._source @click="onClickEmail(email._id)"></EmailTableRecord>
        </tbody>
      </MDBTable>
    </MDBRow>
    <MDBRow center class="container-pagination">
      <Pagination @update-page="handlePageUpdate" @set-page="setPage" :currentPage="page" :maxPage="maxPage"/>
    </MDBRow>
  </MDBCol>

  <MDBCol md="5" class="container-size">
    <EmailRecord v-if="showDetails" :email="this.email"/>
  </MDBCol>
</template>

<style>
.container-size {
  overflow: auto;
  max-height: 85vh;
}

.container-table {
  height: 90%;
  overflow: auto;
}

.container-pagination {
  height: 10%;
}

.border-table {
  border: 2px solid black !important;
}
</style>