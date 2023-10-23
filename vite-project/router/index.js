import { createRouter, createWebHistory } from 'vue-router'
import Home from '@/views/Leads.vue'
import About from '../src/views/About.vue'
import Chat from '@/views/Chat.vue'
import Team from '@/views/Team.vue'
import Login from '@/views/Login.vue'
import Start from "@/views/Start.vue";
import Leads from "@/views/Leads.vue";

const router = createRouter({
    history: createWebHistory(),
    routes: [
        {
            path: '/',
            component: Start
        },
        {
            path: '/leads',
            component: Leads
        },
        {
            path: '/login',
            component: Login
        },

        {
            path: '/about',
            component: About
        },
        {
            path: '/team',
            component: Team
        },
        {
            path: '/chat',
            component: Chat,
            options: {
                headers: {
                    'X-Requested-With': 'XMLHttpRequest',
                    'Accept': 'application/json',
                    'Content-Type': 'application/json',
                    'Access-Control-Allow-Origin': '*',
                    'Access-Control-Allow-Methods': 'GET, POST, PUT, DELETE, PATCH, OPTIONS',
                    'Access-Control-Allow-Headers': 'X-Requested-With, content-type, Authorization',
                    'WebSocket-Accept': 'application/json',
                }
            }
        },
    ],
})

export default router