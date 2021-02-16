<template>

    <h1>{{titulo}}</h1>

</template>

<script>

import axios from 'axios'

export default 
{
    data()
    {
        return{
            titulo : 'Hola'
        }
    },
    mounted()
    {
        this.getTodos();
    },
    methods:
    {
        getTodos()
        {
           
           //const token = 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1bmlxdWVfbmFtZSI6IlRFQ05PU0lTUkQiLCJUZWNub3Npc1ZhbHVlIjoiVmFsdWVUZWNub3NpcyIsImp0aSI6IjU2NWYwMjYxLWQzYzktNDFjYS05YWQ1LTdiY2NlOTRiYmFjZiIsImV4cCI6MTYxMzUyNzI4OX0.A4ux0lOpjpjNClGyek_gmbVQjA5Vl6Qs8A8TGc06tx0'

            axios.get('http://localhost:3000/GetAllBuyers', //http://localhost:53814/api/Zona/', 
            {
               //headers: 
               //{
               //    'Authorization': `Basic ${token}`
               //}
            })
            .then(response => 
            {
                this.titulo = response.status.toString()
                console.log(response)
            }).catch(e => console.log(e))
        }
    }
}

</script>

<style>

</style>