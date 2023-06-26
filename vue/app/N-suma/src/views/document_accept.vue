<template>
    <div class="title has-text-centered">
        書類受理
    </div>

    <router-link to="/document_auth" class="has-text-left" style="font-size: 1.5em">
        ◀戻る
    </router-link>

    <!--
        request_date: "2020-10-01T00:00:00Z"
        student_id: 20206001
        class_name: "CTA20"
        student_name: "斎藤勇"
        absence_start_date: "0001-01-01T00:00:00Z"
        start_flame: 0
        end_date: "0001-01-01T00:00:00Z"
        end_flame: 0
        location: "A棟"
        student_comment: "学生コメント"
        
        document_id: 1
        teacher_comment: ""
    -->

    <!--取得したデータを表示-->
    <!--documentがnullではない場合、画面に出す-->
    <!--
        このコードでは、何かしらのエラーでデータを取得できなかった場合の処理がないため、実装する必要がある。
        また、教員コメントが入力済みのデータが渡された場合の処理も考慮しなければならない
    -->
    <table class="table container" v-if="document != null">
        <thead>
            <tr>
                <th>申請日</th>
                <th>{{ document.request_date }}</th>
            </tr>
            <tr>
                <th>学籍番号</th>
                <th>{{ document.student_id }}</th>
            </tr>
            <tr>
                <th>クラス略称</th>
                <th>{{ document.class_name }}</th>
            </tr>
            <tr>
                <th>氏名</th>
                <th>{{ document.student_name }}</th>
            </tr>
            <tr>
                <th>開始日</th>
                <th>{{ document.absence_start_date }}</th>
            </tr>
            <tr>
                <th>開始時限</th>
                <th>{{ document.start_flame }}</th>
            </tr>
            <tr>
                <th>終了日</th>
                <th>{{ document.end_date }}</th>
            </tr>
            <tr>
                <th>終了時限</th>
                <th>{{ document.end_flame }}</th>
            </tr>
            <tr>
                <th>実施場所</th>
                <th>{{ document.location }}</th>
            </tr>
        </thead>
    </table>
    <div class="has-text-centered" v-else>
        <div class="title mb-5">
            読み込み中です
        </div>
    </div>

    <div class="container">
        <div class="is-grouped">
            <div class="is-flex mb-5">
                <input type="text" v-model="studentComment" disabled class="input disable" style="font-size: 1.25em">
            </div>
            <div class="is-flex mb-5">
                <input type="text" v-model="inputValue" class="input" style="font-size: 1.25em" placeholder="教員用コメント">
            </div>

            <div class="is-flex mb-5">
                <div class="ml-4 button is-medium" @click="OnClickAccept(true)">受理</div>
                <div class="ml-4 button is-medium" @click="OnClickAccept(false)">却下</div>
            </div>
        </div>
    </div>

</template>

<script setup>

    import store from '../store'
    import { useRouter } from 'vue-router'
    import { onMounted,ref } from 'vue'

    const router = useRouter();  

    const studentComment = ref('読み込み中')
    const inputValue = ref('')

    const DocId = store.state.DocId
    const document = ref()

    // 認可用のAPIを叩く
    const OnClickAccept = (isAcceptance) => {

        if (!inputValue.value){
            alert("コメントを入力してください。")
            return
        }     

        if(!isAcceptance){
            // 却下された時のAPI側の仕様が定まっていないため保留
            alert("却下されました")
            return;
        }


        fetch(new URL("ua" , import.meta.env.VITE_API_URL), {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'X-Requested-With': 'XMLHttpRequest',
            },
            body: JSON.stringify({
                "document_id": DocId,
	            "teacher_id": 1,
                "teacher_comment": inputValue.value,
            })
        })
        .then((response) => {
            if (!response.ok) {
                throw new Error(`${response.status} ${response.statusText}`)
            }
            return response.json()
        })
        .then((data) => {

            if(data.message != "200"){
                throw new Error(`${data.message}`)
            }

            alert("受理されました")
            router.push('/document_auth')

        })
        .catch((error) => {
            alert(error)
        })

    }

    onMounted(() => {

        // DocIdがnullの場合は、document_authに戻る
        if (!DocId) {
            router.push('/document_auth')
            return
        }
        
        fetch(new URL("rd" , import.meta.env.VITE_API_URL), {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'X-Requested-With': 'XMLHttpRequest',
            },
            body: JSON.stringify({
                "document_id": DocId,
            })
        })
        .then((response) => {
            if (!response.ok) {
                throw new Error(`${response.status} ${response.statusText}`)
            }
            return response.json()
        })
        .then((data) => {
            
            // 取得したデータを格納
            document.value = data.document;
            studentComment.value = data.document.student_comment;
            
            // データ確認用
            // console.log(document.value)

        })
        .catch((error) => {
            console.log(error)
        })
    })

</script>