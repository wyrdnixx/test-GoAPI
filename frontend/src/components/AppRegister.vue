<template>
<div id="main">
    <h1> Register me </h1>
    <button class="btn btn-info" v-on:click="register()"> Register</button>
    <br>
    <button class="btn btn-info" v-on:click="TestDelCookie()"> Delete cookie</button>
    <br>
    <button class="btn btn-info" v-on:click="testCookie()"> test cookie</button>
    <br>
    Cookie: {{this.UserCookie}}
</div>
</template>


<script>

  import { uuid } from 'vue-uuid'; 
import axios from 'axios';

export default{
    name: "AppRegister",
    props: {
        msg: String
    },
    data() {
        return {
            test :"",
            UserCookie:{}
            }
            
        
    },
    created() {
        this.testCookie()   
    },
    methods: {
        TestDelCookie() {
            this.$cookies.remove('MyvueAppCookie')
            this.UserCookie = this.$cookies.get('MyvueAppCookie')

        },
        register(){
            //alert("register")
            //this.$parent.authenicated = true

                var user = { id:uuid.v4(), name:'UserUUID' };
                this.$cookies.set('MyvueAppCookie',user)
                this.UserCookie = this.$cookies.get('MyvueAppCookie')
                this.testCookie()

        },
        testCookie() {
             
             // test the Alert
             //this.$parent.showAlert("test-Alert\n" )


                this.UserCookie = this.$cookies.get('MyvueAppCookie')
               
               if(this.UserCookie == null) {
                console.log("cookie not found")   
               } else {
                console.log("cookie found: " + this.UserCookie.id)
                console.log("testing against server...")
                this.checkCookiewithserver()
               }
               
        },
        async checkCookiewithserver() {
              let testResult =  await   axios.post(this.$parent.apiURL + "/checkUserCookie", {
           
                Id: this.UserCookie.id
              }, {
                headers: {
                    'Content-Type': 'application/json',
                }
              })              
              .catch((error)=> this.$parent.showAlert("Server returned an Error:\n" + error.response.data));  
             this.$parent.showAlert("cookie test result:\n" +  JSON.stringify(testResult))   
        }
    
    }
}
</script>

