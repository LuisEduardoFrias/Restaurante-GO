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
           
           const token = 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1bmlxdWVfbmFtZSI6IlRFQ05PU0lTUkQiLCJUZWNub3Npc1ZhbHVlIjoiVmFsdWVUZWNub3NpcyIsImp0aSI6IjJjMWI0MmIzLTUyZDAtNGQyYi1iN2ZmLTAxNmEwODI5ZmNjNyIsImV4cCI6MTYxMzQ0NDczN30.YkotneSPJFNp5YyDRmdDnYMgmywkYRaQRo3IjfJxLqc'

            axios.get('http://localhost:53814/api/Zona/', 
            {
                headers: 
                {
                    'Authorization': `Basic ${token}`
                }
            })
            .then(response => 
            {
                this.titulo = response
                console.log(response)
            }).catch(e => console.log(e))
        }
    }
}

</script>

<style>

</style>