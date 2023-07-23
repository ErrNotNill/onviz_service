import { createRouter, createWebHistory } from 'vue-router'
import Home from '../src/views/Home.vue'

const router = createRouter({
    history: createWebHistory(),
    routes: [
        {
            path: '/',
            component: Home
        },
        {
            path: '/about',
            component: () => import('../src/views/About.vue')
        },
    ],
})

export default router