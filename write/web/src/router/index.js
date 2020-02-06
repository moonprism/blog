import Vue from 'vue'
import Router from 'vue-router'
import Layout from '@/views'
import store from '@/store'

Vue.use(Router)

const router = new Router({
    routes: [
        {
            path: '/login',
            component: () => import('@/views/login/index'),
        },
        {
            path: '/article/edit/:id',
            name: 'article_edit',
            component: () => import('@/views/article/edit'),
        },
        {
            path: '/',
            component: Layout,
            redirect: '/article',
            children: [
                {
                    path: 'article',
                    name: 'article',
                    component: () => import('@/views/article/list'),
                    children: [
                        {
                            path: 'list/:page',
                            name: 'article_list',
                        }
                    ]
                },
                {
                    path: 'article/create',
                    name: 'article_create',
                    component: () => import('@/views/article/create')
                },
                {
                    path: 'tag',
                    name: 'tag',
                    component: () => import('@/views/tag/list')
                },
                {
                    path: 'code',
                    name: 'code',
                    component: () => import('@/views/code/list'),
                    children: [
                        {
                            path: 'list/:page',
                            name: 'code_list'
                        }
                    ]
                },
                {
                    path: 'code/create',
                    name: 'code_create',
                    component: () => import('@/views/code/create')
                },
                {
                    path: 'code/edit/:id',
                    name: 'code_edit',
                    component: () => import('@/views/code/edit'),
                },
                {
                    path: 'file',
                    name: 'file',
                    component: () => import('@/views/file/index'),
                }
            ]
        }
    ]
})

router.beforeEach((to, from, next) => {
    let token = store.getters['auth/token']
    if (!token && to.path != '/login') {
        next('/login')
    } else if (token && to.path == '/login') {
        next('/')
    } else {
        next()
    }
})

export default router