<template>
    <div class="title has-text-centered">
        各種書類認可
    </div>

    <div class="columns">
        
        <div class="column has-text-right">
            <div>未認可リスト</div>
        </div>
        
        <div class="column"></div>

        <div class="column">
            <div>ソート</div>
            <select class="select">
                <option>クラス略称</option>
                <option>名前</option>
                <option>分類</option>
            </select>
        </div>

    </div>

    <div class="container">
        
        <table class="table" v-if="users.length > 0">
            <thead>
                <tr>
                    <th>クラス略称</th>
                    <th>氏名</th>
                    <th>分類</th>
                    <th>    </th>
                </tr>
            </thead>
            <thead>
                <tr v-for="user in users" :key="user.id">
                    <th>{{ user.class }}</th>
                    <th>{{ user.name }}</th>
                    <th>{{ user.type }}</th>
                    <th><router-link to="/document_accept">詳細</router-link></th>
                </tr>
            </thead>
        </table>
        <div v-else>
            <div class="has-text-centered">
                <div class="title mb-5">
                    認可待ちの書類はありません
                </div>
            </div> 
        </div>
    </div>
</template>

<script setup>

    import { onMounted,ref } from 'vue'
    
    const users = ref([])

    onMounted(() => {

        // 認可待ちの書類一覧を取得
        fetch(new URL("ual" , import.meta.env.VITE_API_URL), {
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
        })
        .then((response) => {
            if (!response.ok) {
                throw new Error(`${response.status} ${response.statusText}`)
            }
            return response.json()
        })
        .then((data) => {
            
            // 取得したデータをusersに格納
            data['document'].forEach((user) => {
                users.value.push({
                    id: user['document_id'],
                    class: user['class_name'],
                    name: user['student_name'],
                    type: user['absence_category'],
                })
            })

        })
        .catch((error) => {
            console.log(error)
        })
    })

    
    
</script>