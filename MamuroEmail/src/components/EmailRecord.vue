<script>
import { MDBRow, MDBContainer } from 'mdb-vue-ui-kit';
import EMAIL_FIELD_ORDER from '../email/fieldOrder';
import SingleItem from './SingleItem.vue'
export default {
  name: "EmailRecord",
  data() {
    return {
      email_field_order: EMAIL_FIELD_ORDER
    }
  },
  props: {
      email: JSON
  },
  components: {
    SingleItem,
    MDBContainer,
    MDBRow
  }
}
</script>

<template>
  <MDBContainer class="email-overflow">
    <MDBRow v-for="emailField in this.email_field_order" :key="emailField">
      <SingleItem :Field="emailField" :Text="email[emailField]"></SingleItem>
    </MDBRow>
  </MDBContainer>
</template>

<style>
.email-overflow {
  overflow: scroll;
  height: 80vh;
}
</style>