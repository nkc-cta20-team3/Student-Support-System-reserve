import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import { getAuth,onAuthStateChanged } from 'firebase/auth'

const app = createApp(App)

// Firebaseのログイン状態を監視し、VuexストアにユーザーIDを保存する
onAuthStateChanged(getAuth(), (user) => {
  if (user) {
    store.dispatch('updateUserId', user.uid)
  }
})

app.use(router).use(store).mount('#app')