import {createRouter, createWebHistory} from 'vue-router'
import {getAuth, onAuthStateChanged} from 'firebase/auth'

const router = createRouter({
    history: createWebHistory(),
    routes: [
        {  path: '/', component: () => import('../views/index.vue') },
        {  path: '/document_list', component: () => import('../views/document_list.vue') ,meta: { requiresAuth: true },},
        {  path: '/document_form', component: () => import('../views/document_form.vue') ,meta: { requiresAuth: true },},
        {  path: '/document_auth', component: () => import('../views/document_auth.vue') ,meta: { requiresAuth: true },},
        {  path: '/document_accept', component: () => import('../views/document_accept.vue') ,meta: { requiresAuth: true },},
        {  path: '/practice', component: () => import('../views/practice.vue') ,meta: {requiresAuth: true},},
        
    ],
})

const getCurrentUser = () => {
    return new Promise((resolve, reject) => {
        const removeListener = onAuthStateChanged(getAuth(), user => {
            removeListener()
            resolve(user)
        }, reject)
    })
}

router.beforeEach(async(to, from, next) => {
    if (to.matched.some(record => record.meta.requiresAuth)) {
        if (await getCurrentUser()) {
            next()
        } else {
            alert('ログインし直してください')
            next({ path: '/' })
        }
    } else {   
        next()
    }
})


export default router