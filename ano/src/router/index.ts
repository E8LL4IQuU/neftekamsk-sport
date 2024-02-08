import { createRouter, createWebHistory } from 'vue-router'
import Page from '@/views/Page.vue'
import Events from '@/views/Manage/Events.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: '',
      component: Page
    },
    {
      path: '/events',
      name: 'events',
      // route level code-splitting
      // this generates a separate chunk (news.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      // routes that are the first to be loaded when visiting website are imported the usual way
      component: () => import('@/views/EventsView.vue')
    },
    {
      path: '/events/:id',
      component: () => import('@/views/EventView.vue')
    },
    {
      path: '/news',
      name: 'news',
      component: () => import('@/views/NewsBulkView.vue')
    },
    {
      path: '/news/:id',
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
      component: () => import('@/views/LoginView.vue'),
      meta: {
        hideNavbar: true,
        hideFooter: true,
      }
    },
    {
      path: '/logout',
      name: 'logout',
      component: () => import('@/views/LogoutView.vue'),
      meta: {
        hideNavbar: true,
      }
    },
    {
      path: '/manage',
      redirect: '/manage/events',
    },
    {
      path: '/manage/events',
      name: 'manage.events',
      component: Events,
      meta: {
        hideNavbar: true,
        hideFooter: true,
        // FIXME: meaning is unclear
        managementRoute: true,
      }
    },
    {
      path: '/manage/events/create',
      name: 'manage.events.create',
      component: () => import('@/views/Manage/CreateEvent.vue'),
      meta: {
        hideNavbar: true,
        hideFooter: true,
      }
    },
    {
      path: '/manage/news',
      name: 'manage.news',
      component: () => import('@/views/Manage/News.vue'),
      meta: {
        hideNavbar: true,
        hideFooter: true,
        managementRoute: true,
      }
    },
    {
      path: '/manage/news/create',
      name: 'manage.news.create',
      component: () => import('@/views/Manage/CreateNews.vue'),
      meta: {
        hideNavbar: true,
        hideFooter: true,
      }
    },
    {
      path: '/manage/news/:id',
      name: 'newsEdit',
      component: () => import('@/views/Manage/NewsEdit.vue'),
      meta: {
        hideNavbar: true,
        hideFooter: true,
      }
    },
    {
      path: '/manage/signups',
      name: 'manage.signups',
      component: () => import('@/views/Manage/Signups.vue'),
      meta: {
        hideNavbar: true,
        hideFooter: true,
        managementRoute: true,
      }
    },
    {
      path: '/:catchAll(.*)',
      redirect: '/',
      meta: {
        hideNavbar: true,
        hideFooter: true,
      }
    }

  ]
})

export default router
