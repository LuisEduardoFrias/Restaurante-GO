<template>

    <div>
        
        <h1>{{ Token }}</h1>

        <h1>{{ titulo }}</h1>

        <h2>{{ Id }}</h2>

        <h2>{{ Name }}</h2>

        <h2>{{ Age }}</h2>

    </div>

</template>

<script>

import axios from 'axios'

export default 
{
    data()
    {
        return{
            titulo : "Status Code",
            Id : "0",
            Name : "",
            Age : "0",
            Token : "token"
        }
    },
    mounted()
    {
        this.PostTodos();
        this.getTodos();
    },
    methods:
    {
        PostTodos()
        {
            axios.post('http://localhost:53814/api/Credenciales',
            {
                userName: 'TECNOSISRD',
                passwordHash: 'TECNOSISRD'
            })
            .then(response => 
            {
                this.Token = response.data.token.token.toString()
                console.log(response.data.token.token);
            })
            .catch(e => console.log(e));
        },

        
        getTodos()
        {
            axios.get('http://localhost:53814/api/Zona/', //http://localhost:53814/api/Zona/',http://localhost:3000/GetBuyerPerId/95195187 
            {
               headers: 
               {
                   "Authorization": `Bearer ${this.Token}`
               }
            })
            .then(response => 
            {
                this.titulo = "Status Code: " + response.status.toString()
                this.Id = " ID: " + response.data.Id
                this.Name = " NAME: " + response.data.Name 
                this.Age = " AGE: " + response.data.Age

                console.log(response)
            })
            .catch(e => {
                 this.titulo = e.toString()
                console.log(e)})
        }

    }
}

</script>

<style>

</style>