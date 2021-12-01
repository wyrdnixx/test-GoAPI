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


    <div  class="form-group formdiv" id="newFirma"> 
      
      <input type="text"  class="form-control" v-model="newFirma.Name" placeholder="FirmaName"> 
      <input type="checkbox"  class="form-control" v-model="newFirmaActive" placeholder="Enabled" name="chkEnabled" value="test">  <label for="chkEnabled">Aktiv</label>
      
      <br>
      <button class="btn btn-info" v-on:click="createFirma()">Eintragen</button> 

    </div>


    <div id="example-1">
      <button class="btn btn-info" v-on:click="getData()">Refresh</button>
      <button class="btn btn-info" v-on:click="clearData()">clear</button>
      <button class="btn btn-info" v-on:click="setCookie()">setCookie</button>
      <p>Info field: {{ this.info }} </p> <br> 
      <!-- <h2>Debug {{ this.info }} </h2>       --> 
      
      <h1 class="badge badge-warning" v-if="this.info.ErrorText != null">
         Debug error: {{this.info.ErrorText}} 
         
      </h1>
      
      <div id="Results" class="formdiv">
        <table class="table table-dark">
          <thead>
            <td>ID</td>
            <td>Name</td>
            <td>Enabled</td>
            <td>Delete entry</td>
            
          </thead>
       <tr  v-for="per in this.info.Firmen" v-bind:key="per.Id">
            <td>{{per.Id}}</td>
            <td>{{per.Name}}</td>
            <td>{{per.Enabled}}</td>
            <td><button class="btn btn-primary" @click='delEntry(per.Id)'>x</button></td>
       </tr>   
       </table>    
      </div>     
       
     </div>
  </div> 
  

  
</template>

<script>


import axios from 'axios';


const apiURL = window.location.protocol + "//"+ window.location.hostname +":8081/api"

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
    created() {
            this.getData()    
            this.$cookies.set('testcookie', 'this value')
            //$cookies.set('testcookie', 'this value')


  },
methods: {

  clearData() {
    this.info = "empty-string"
    
    
  },
  setCookie() {
    var user = { id:1, name:'Journal',session:'25j_7Sl6xDq2Kc3ym0fmrSSk2xV2XkUkX' };
        
    
    
    
    try {
      var user2 = this.$cookies.get('CB')
      this.info = user2.name + " - " + user2.session
      console.log("Cookie-B: " +JSON.stringify(user2))
       console.log("Cookie-B-config: " + JSON.stringify(this.$cookies.get('CB')))
    }catch {
      console.log("Cookie CB not existing... setting it...")
      this.$cookies.set('CB', user, '1Y' );
    }
    
  },
      

  async getData() {
      try {
        
        let response = await axios.get(apiURL  + "/getFirmen");
//        let response = await axios.get("http://localhost:8081/api/getFirmen");
        this.info = response.data;
      } catch (error) {
        console.log(error);
        this.showAlert("Server returned an Error:\n" + error); 
      }
    },
    async delEntry(Id) {
      console.log("delete: " + Id)
    
           await   axios.post(apiURL + "/delFirma", {
           //await   axios.post("http://localhost:8081/api/delFirma", {
                Id: Id
              }, {
                headers: {
                    'Content-Type': 'application/json',
                }
              }).catch((error)=> this.showAlert("Server returned an Error:\n" + error.response.data));  
    
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
        await   axios.post(apiURL + "/createFirma", {
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
