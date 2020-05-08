import Vue from 'vue'
import Router from 'vue-router'
import Layout from '@/views'
import store from '@/store'

Vue.use(Router)

const router = new Router({
    routes: [
        {
            path: "/login",
            component: (resolve) => require(["@/views/login/index"], resolve),
        },
        {
            path: "/article/edit/:id",
            name: "article_edit",
            component: (resolve) => require(["@/views/article/edit"], resolve),
        },
        {
            path: "/",
            component: Layout,
            redirect: "/setting",
            children: [
                {
                    path: "/Setting",
                    component: (resolve) =>
                        require(["@/views/setting/index"], resolve),
                },
                {
                    path: "article",
                    name: "article",
                    component: (resolve) =>
                        require(["@/views/article/list"], resolve),
                    children: [
                        {
                            path: "list/:page",
                            name: "article_list",
                        },
                    ],
                },
                {
                    path: "article/create",
                    name: "article_create",
                    component: (resolve) =>
                        require(["@/views/article/create"], resolve),
                },
                {
                    path: "tag",
                    name: "tag",
                    component: (resolve) =>
                        require(["@/views/tag/list"], resolve),
                },
                {
                    path: "code",
                    name: "code",
                    component: (resolve) =>
                        require(["@/views/code/list"], resolve),
                    children: [
                        {
                            path: "list/:page",
                            name: "code_list",
                        },
                    ],
                },
                {
                    path: "code/create",
                    name: "code_create",
                    component: (resolve) =>
                        require(["@/views/code/create"], resolve),
                },
                {
                    path: "code/edit/:id",
                    name: "code_edit",
                    component: (resolve) =>
                        require(["@/views/code/edit"], resolve),
                },
                {
                    path: "file",
                    name: "file",
                    component: (resolve) =>
                        require(["@/views/file/index"], resolve),
                },
            ],
        },
    ],
});

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