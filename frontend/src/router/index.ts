import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'layout',
      component: () => import('../components/layouts/MainLayout.vue'),
      children: [
        {
          path: '',
          name: 'Process',
          component: () => import('../views/ProcessesView.vue')
        }
      ]
    }
  ]
})

export default router
