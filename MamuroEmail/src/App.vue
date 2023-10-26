<script setup>
import EmailTable from './components/EmailTable.vue';
import { MDBCol, MDBRow, MDBContainer } from 'mdb-vue-ui-kit';
</script>

<template>
  <MDBContainer fluid class="border">
    <MDBRow>
      <h2>Mamuro Email</h2>
    </MDBRow>

    <MDBRow>
      <EmailTable/>
    </MDBRow>
  </MDBContainer>
</template>

<style scoped>
#app {
  font-family: Roboto, Helvetica, Arial, sans-serif;
}

.border {
  padding: 2rem;
}
</style>
