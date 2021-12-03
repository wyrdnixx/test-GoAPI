

<template>
  <div id="app" >
    <img alt="Vue logo" src="./assets/logo.png">

    <!-- Toast Message banner  -->

    <b-alert
      :show="AlertDismissCountDown"
      dismissible
      variant="warning"
      @dismissed="AlertDismissCountDown=0"
      @dismiss-count-down="countDownChanged"
    >
      <p> {{this.AlertWebsiteAlertMessage}} {{ AlertDismissCountDown }} </p>
      <b-progress
        variant="warning"
        :max="AlertDismissSecs"
        :value="AlertDismissCountDown"
        height="4px"
      ></b-progress>
    </b-alert>

<!-- Toast Message banner  -->


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
      authenicated: false,
      apiURL :  window.location.protocol + "//"+ window.location.hostname +":8081/api",
      AlertDismissSecs: 4,
      AlertDismissCountDown: 0,
      AlertShowDismissibleAlert: false,
      AlertWebsiteAlertMessage: ""
    }
    
  },
  methods: {
        countDownChanged(AlertDismissCountDown) {
      this.AlertDismissCountDown = AlertDismissCountDown
    },
    showAlert(msg) {
      this.AlertWebsiteAlertMessage = msg
      this.AlertDismissCountDown = this.AlertDismissSecs
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
