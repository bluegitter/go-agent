import Vue from 'vue'
import Router from 'vue-router'
import HelloWorld from '@/components/HelloWorld'
import SysInfo from '@/components/SysInfo'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/hello',
      name: 'HelloWorld',
      component: HelloWorld,
    },
    {
      path: '/info',
      name: 'SysInfo',
      component: SysInfo,
    },
    {
      path: '/',
      name: 'SysInfo',
      component: SysInfo,
    },
  ],
})
