<script>

  export default {
    name: "SingleItem",
    props: {
      Field: String,
      Text: String
    }
  }
</script>

<template>
  <p>{{ Field }}: {{ Text }}</p>
</template>