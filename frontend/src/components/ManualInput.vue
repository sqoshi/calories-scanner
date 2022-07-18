<template>
  <div class="hello">
    <h1>Manual Input</h1>
    <button @click="count++">
      Count is: {{ count }}
    </button>
  </div>
</template>

<script>
export default {
  name: 'ManualInput',
  data(){
    return {
      count: 0
    }
  }
  // props: {
  //   msg: String
  // }
}
</script>

<style scoped>
</style>
