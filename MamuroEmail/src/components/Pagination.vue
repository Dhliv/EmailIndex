<script>
import { MDBPagination, MDBPageNav, MDBPageItem } from 'mdb-vue-ui-kit';
export default {
  props: {
    currentPage: 1,
    maxPage: 1
  },

  methods: {
    shouldShowAsActive(pageIndex) {
      return pageIndex == this.currentPage
    },

    realPageIndex(index) {
      if(this.currentPage < 5) return index;
      if(this.currentPage > this.maxPage - 5) return this.maxPage - 9 + index;
      return this.currentPage + index - 5;
    }
  },

  components: {
    MDBPagination,
    MDBPageNav,
    MDBPageItem
  }
}
</script>

<template>
  <nav>
    <MDBPagination circle class="center-all">
      <MDBPageItem @click="$emit('updatePage', -10000000)">
        <i class="fas fa-angle-double-left"></i>
      </MDBPageItem>

      <MDBPageItem @click="$emit('updatePage', -1)">
        <i class="fas fa-angle-left"></i>
      </MDBPageItem>

      <div v-for="i in 9" :key="i">
        <MDBPageItem @click="$emit('setPage', realPageIndex(i))" icon active v-if="shouldShowAsActive(realPageIndex(i))">{{realPageIndex(i)}}</MDBPageItem>
        <MDBPageItem @click="$emit('setPage', realPageIndex(i))" icon v-else="shouldShowAsActive(realPageIndex(i))">{{realPageIndex(i)}}</MDBPageItem>
      </div>

      <MDBPageItem @click="$emit('updatePage', 1)">
        <i class="fas fa-angle-right"></i>
      </MDBPageItem>

      <MDBPageItem @click="$emit('updatePage', 1000000)">
        <i class="fas fa-angle-double-right"></i>
      </MDBPageItem>
    </MDBPagination>
  </nav>
</template>

<style>
.center-all {
  justify-content: center !important;
  padding-top: 10px;
}
</style>