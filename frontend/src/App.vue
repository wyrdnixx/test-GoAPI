

<template>
  <div id="app" >
    <img alt="Vue logo" src="./assets/logo.png">
    <div v-if="authenicated">
      <HelloWorld  msg="Welcome asdf Your Vue.js App"/>
    </div>
    <div v-else>
      <AppRegister/>
    </div>
  </div>
</template>

<script>
import HelloWorld from './components/HelloWorld.vue'
import AppRegister from './components/AppRegister.vue'




export default {
  name: 'App',
  components: {
    HelloWorld,
    AppRegister
  },
  data() {
    return {
      authenicated: false
    }
    
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}

table, tr, td {
  margin-left: auto;
  margin-right: auto;
    
  border: 1px solid black;
  text-align: left;  
   }


.formdiv {
 padding-top: 8px; 
padding-right: 90px; 
padding-bottom: 8px; 
padding-left: 90px; 
}
</style>
