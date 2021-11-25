<template>
  <div class="hello">
    <h1>{{ msg }}</h1>
    <p>
      MY test app<br>
      
    </p>

<!-- Toast Message banner  -->
    <b-button @click="showAlert('Test-alarm')" variant="info" class="m-1">
      Show alert with count-down timer
    </b-button>
    <b-alert
      :show="dismissCountDown"
      dismissible
      variant="warning"
      @dismissed="dismissCountDown=0"
      @dismiss-count-down="countDownChanged"
    >
      <p> {{this.websiteAlertMessage}} {{ dismissCountDown }} </p>
      <b-progress
        variant="warning"
        :max="dismissSecs"
        :value="dismissCountDown"
        height="4px"
      ></b-progress>
    </b-alert>

<!-- Toast Message banner  -->

    <h3>Test here:</h3>

    <div id="newFirma"> 
      <input type="text" v-model="newFirma.Name" placeholder="FirmaName"> 
      <input type="checkbox" v-model="newFirmaActive" placeholder="Enabled" name="chkEnabled" value="test">  <label for="chkEnabled">Aktiv</label>
      
      <br>
      <button class="button" @click="createFirma">Eintragen</button> 

    </div>

    <div id="example-1">
      <button v-on:click="getData()">Refresh</button>
      <button v-on:click="clearData()">clear</button>
      <p>Info field: </p> <br> 
      <!-- <h2>Debug {{ this.info }} </h2>       --> 
      
      <h1 class="badge badge-warning" v-if="this.info.ErrorText != null">
         Debug error: {{this.info.ErrorText}} 
         
      </h1>
      
      <div id="Results">
        <table>
       <tr  v-for="per in this.info.Firmen" v-bind:key="per.Id">
            <td>{{per.Id}}</td>
            <td>{{per.Name}}</td>
            <td>{{per.Enabled}}</td>
            <td><button @click='delEntry(per.Id)'>x</button></td>
       </tr>   
       </table>    
      </div>     
       
     </div>
  </div> 
  

  
</template>

<script>


import axios from 'axios';

export default {
  name: 'HelloWorld',
  props: {
    msg: String
  },
    data () {
    return {
      info: "old-string",
      loading: true,
      errored: false,
      newFirma: {
        Name:"",
        Enabled:""
      },
      newFirmaActive:false,
        dismissSecs: 4,
        dismissCountDown: 0,
        showDismissibleAlert: false,
        websiteAlertMessage: ""
    }
  },
methods: {
  clearData() {
    this.info = "empty-string"
    
  },
  async getData() {
      try {
        let response = await fetch("http://localhost:8081/api/health");
        this.info = await response.json();
      } catch (error) {
        console.log(error);
      }
    },
    async delEntry(Id) {
      console.log("delete: " + Id)
    
           await   axios.post("http://localhost:8081/api/delFirma", {
                Id: Id
              }, {
                headers: {
                    'Content-Type': 'application/json',
                }
              }).catch((error)=> alert(("Server returned an Error:\n" + error.response.data)));  
    
            this.getData()    
    },
    async createFirma(){
      console.log("submiting new Firma", this.newFirma)
      if (this.newFirma.Name == "") {
        //alert("Kein Name eingetragen")        
        this.showAlert('Kein Name eingetragen')
      }else {
        if (this.newFirmaActive ) {
          this.newFirma.Enabled = '1'
        } else  {this.newFirma.Enabled = '0'}
        await   axios.post("http://localhost:8081/api/createFirma", {
                NewFirma:this.newFirma
                },
               {
                headers: {
                    'Content-Type': 'application/json',
                }
              }).catch((error)=> alert(("Server returned an Error:\n" + error.response.data)));  
    
            this.getData()    

      }
    
    },
    countDownChanged(dismissCountDown) {
      this.dismissCountDown = dismissCountDown
    },
    showAlert(msg) {
      this.websiteAlertMessage = msg
      this.dismissCountDown = this.dismissSecs
    }
}
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}


</style>
