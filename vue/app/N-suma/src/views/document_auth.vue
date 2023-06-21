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
        
        <table class="table" v-if="documents.length > 0">
            <thead>
                <tr>
                    <th>クラス略称</th>
                    <th>氏名</th>
                    <th>分類</th>
                    <th>    </th>
                </tr>
            </thead>
            <thead>
                <tr v-for="document in documents" :key="document.id">
                    <th>{{ document.class }}</th>
                    <th>{{ document.name }}</th>
                    <th>{{ document.type }}</th>
                    <th><button @click="showAccept(document.id)">詳細</button></th>
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

    import store from '../store'
    import { useRouter } from 'vue-router'
    import { onMounted,ref } from 'vue'

    const router = useRouter();    
    const documents = ref([])

    onMounted(() => {

        // 認可待ちの書類一覧を取得
        // 本来はログインユーザーのIDを取得し、そのユーザーの認可待ちの書類一覧を取得する
        // 現在は、ユーザーIDをハードコーディングしているため、修正する必要があります
        fetch(new URL("ral" , import.meta.env.VITE_API_URL), {
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
            data.document.forEach((document) => {
                documents.value.push({
                    id: document.document_id,
                    class: document.class_name,
                    name: document.student_name,
                    type: document.absence_category,
                })
            })

        })
        .catch((error) => {
            console.log(error)
        })
    })

    // ドキュメントIDをStoreに格納し、認可画面に遷移
    const showAccept = (id) => {
        store.commit('setDocId', id)
        router.push('/document_accept')
    }
    
</script>