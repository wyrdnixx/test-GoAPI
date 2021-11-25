<template>
  <div class="hello">
    <h1>{{ msg }}</h1>
    <p>
      MY test app<br>
      
    </p>
    <h3>Test here:</h3>

    <div id="newFirma"> 
      <input type="text" v-model="newFirma.Name" placeholder="FirmaName"> 
      <input type="checkbox" v-model="newFirma.Enabled" placeholder="Enabled" name="chkEnabled">  <label for="chkEnabled">Aktiv</label>
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
            <td>{{per.id}}</td>
            <td>{{per.name}}</td>
            <td>{{per.enabled}}</td>
            <td><button @click='delEntry(per.id)'>x</button></td>
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
      }
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
    async delEntry(id) {
      console.log("delete: " + id)
    
           await   axios.post("http://localhost:8081/api/delFirma", {
                id: id
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
        alert("Kein Name eingetragen")        
      }else {

 await   axios.post("http://localhost:8081/api/createFirma", {
                newFirma: this.newFirma},
               {
                headers: {
                    'Content-Type': 'application/json',
                }
              }).catch((error)=> alert(("Server returned an Error:\n" + error.response.data)));  
    
            this.getData()    

      }
    
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
