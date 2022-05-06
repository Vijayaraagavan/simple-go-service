<template>
    <form>
        <label for="fname">First name: </label>
            <input v-model="this.fname" type="text" id="fname" name="fname"><br>
        <label for="lname">Last name: </label>
            <input v-model="lname" type="text" id="lname" name="lname"><br>
        <button @click="pusher()">Send form</button>
    </form>
    <input type="text" v-model="add" >
        <p>the value of addition is : {{ add }}</p>
        <p>the value of multiplication is : {{ mul }}</p>
</template>

<script>
// import { defineComponent } from '@vue/composition-api'
import axios from 'axios';


export default{
    data() {
        return{
            title: "Naruto",
            fname:"",lname:"",add:"",mul:"",
        }
    },
    methods: {
    pusher() {
        alert(this.fname)
        var data = {
            "num": parseFloat(this.fname),
            "num2": this.lname
        }
        console.log(data)
        axios({method:"POST", url: "http://127.0.0.1:8090/calc", data: data, 
        headers:{"content-type":"text/plain"} }).then(result => {
            this.add = result.data['add']
            this.mul = result.data['mul']
            console.log(result.data)
        }).catch( error => {
            console.error(error);
        })
    },
    },
    // methods:
}
// console.log(this.fname,this.lname)
</script>
