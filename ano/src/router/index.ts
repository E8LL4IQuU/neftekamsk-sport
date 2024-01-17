import { createRouter, createWebHistory } from 'vue-router'
import LoginView from '@/views/LoginView.vue'
import AdminView from '@/views/AdminView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: '',
      // route level code-splitting
      // this generates a separate chunk (news.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import('@/views/Page.vue')
    },
    {
      path: '/news',
      name: 'news',
      component: () => import('@/views/NewsView.vue')
    },
    {
      path: '/schedule',
      name: 'schedule',
      component: () => import('@/views/ScheduleView.vue')
    },
    {
      path: '/sections',
      name: 'sections',
      component: () => import('@/views/SectionsView.vue')
    },
    {
      path: '/gallery',
      name: 'gallery',
      component: () => import('@/views/GalleryView.vue')
    },{
      path: '/events',
      name: 'events',
      component: () => import('@/views/EventsView.vue')
    },
    {
      path: '/documents',
      name: 'documents',
      component: () => import('@/views/DocumentsView.vue')
    },
    {
      path: '/people',
      name: 'people',
      component: () => import('@/views/PeoplePage.vue')
    }, 
    {
      path: '/login',
      name: 'login',
      // Not lazy loaded as it causes navbar/footer flickering while loading
      component: LoginView,
      meta: {
        hideNavbar: true,
      }
    },
    {
      path: '/admin',
      name: 'admin',
      component: AdminView,
      meta: {
      hideNavbar: true,
      hideFooter: true,
      }
    },
    {
      path: '/logout',
      name: 'logout',
      component: () => import ('@/views/LogoutView.vue'),
      meta: {
        hideNavbar: true,
      }
    },
    {
      path: '/manage',
      name: 'manage',
      component: () => import ('@/views/Manage/ManageView.vue'),
      meta: {
        hideNavbar: true,
        hideFooter: true,
      }
    },
    {
      path: '/manage/events',
      name: 'manage.events',
      component: () => import ('@/views/Manage/Events.vue'),
      meta: {
        hideNavbar: true,
        hideFooter: true,
        managementRoute: true,
      }
    },
    {
      path: '/manage/news',
      name: 'manage.news',
      component: () => import ('@/views/Manage/News.vue'),
      meta: {
        hideNavbar: true,
        hideFooter: true,
        managementRoute: true,
      }
    }
   
  ]
})

export default router
