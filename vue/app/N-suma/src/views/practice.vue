<template>
    <div>
        <button @click="fetchData(false)">APIを呼び出す(POST)</button>
        <button @click="fetchData(true)">APIを呼び出す(GET)</button>
        <div>{{ responseData }}</div>
    </div>
</template>
  
<script setup>

    import { ref } from 'vue'

    const responseData = ref('')

    const GetParams = {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json',
            'X-Requested-With': 'XMLHttpRequest',
        },
    }

    const PostParams = {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'X-Requested-With': 'XMLHttpRequest',
        },
        body: JSON.stringify({
            "teacher_id": 1,
            "position": 1,
            "class_name": "CTA20"
        })
    }

    const fetchData = (isGet) => {

        fetch(new URL(isGet ? "ad/1" : "ual" , import.meta.env.VITE_API_URL), isGet ? GetParams : PostParams )
        .then((response) => {
            if (!response.ok) {
                throw new Error(`${response.status} ${response.statusText}`)
            }
            return response.json()
        })
        .then((data) => {
            responseData.value = data;
        })
        .catch((error) => {
            console.log(error)
        })

    }

</script>
  
