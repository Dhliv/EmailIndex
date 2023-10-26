<script>
  export default {
    name: "EmailTableRecord",
    props: {
      email: JSON
    }
  }
</script>

<template>
  <tr>
    <td>
      {{ email.Subject }}
    </td>

    <td>
      {{ email.From }}
    </td>

    <td>
      {{ email.To }}
    </td>
  </tr>
</template>