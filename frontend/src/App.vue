<template>
  <MainPanel msg="Welcome to Your Vue.js App"/>
</template>

<script>
import MainPanel from './components/MainPanel.vue'

export default {
  name: 'App',
  components: {
    MainPanel
  }
}
</script>

<style>
body {
  color: #42b983;
  background-color: #343232;
}

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  margin-top: 60px;
}
</style>
