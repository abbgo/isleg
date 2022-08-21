import tm from './locales/tm'
import ru from './locales/ru'
require('dotenv').config()
export default {
  // Global page headers: https://go.nuxtjs.dev/config-head
  head: {
    title: 'Isleg',
    htmlAttrs: {
      lang: 'en',
    },
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      { hid: 'description', name: 'description', content: '' },
      { name: 'format-detection', content: 'telephone=no' },
    ],
    link: [{ rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }],
  },
  server: {
    port: 8000,
    host: '0.0.0.0',
  },
  // Global CSS: https://go.nuxtjs.dev/config-css
  css: ['@/assets/css/style.css'],

  loading: {
    height: '1px',
    color: '#fd5e29',
  },
  // Plugins to run before rendering page: https://go.nuxtjs.dev/config-plugins
  plugins: [
    '@/plugins/vue-validate',
    '@/plugins/toast.js',
    { src: '@/plugins/vue-awesome-swiper', mode: 'client' },
  ],

  // Auto import components: https://go.nuxtjs.dev/config-components
  components: true,

  // Modules for dev and build (recommended): https://go.nuxtjs.dev/config-modules
  buildModules: [],

  // Modules: https://go.nuxtjs.dev/config-modules
  ssr: true,
  target: 'server',

  modules: [
    // https://go.nuxtjs.dev/axios
    '@nuxtjs/axios',
    '@nuxtjs/auth-next',
    '@nuxtjs/i18n',
    'cookie-universal-nuxt',
  ],
  env: {
    BASE_API: process.env.BASE_API,
    SITE_URL: process.env.SITE_URL,
    IMAGE_URL: process.env.IMAGE_URL,
  },
  // Axios module configuration: https://go.nuxtjs.dev/config-axios
  axios: {
    // Workaround to avoid enforcing hard-coded localhost:3000: https://github.com/nuxt-community/axios-module/issues/308
    baseURL: process.env.BASE_API,
  },
  auth: {
    strategies: {
      userRegister: {
        scheme: 'refresh',
        token: {
          property: 'access_token',
          global: true,
          type: 'JWT',
        },
        refreshToken: {
          property: 'refresh_token',
          data: 'refresh_token',
          maxAge: 60,
        },
        user: {
          property: 'user',
          // autoFetch: true
        },
        endpoints: {
          login: { url: '/auth/register', method: 'post' },
          refresh: { url: '/auth/refresh', method: 'post' },
          user: false,
          logout: false,
        },
        // autoLogout: false
      },
      userLogin: {
        scheme: 'refresh',
        token: {
          property: 'access_token',
          global: true,
          type: 'JWT',
        },
        refreshToken: {
          property: 'refresh_token',
          data: 'refresh_token',
          maxAge: 60,
        },
        user: {
          property: 'user',
          // autoFetch: true
        },
        endpoints: {
          login: { url: '/auth/login', method: 'post' },
          refresh: { url: '/auth/refresh', method: 'post' },
          user: false,
          logout: false,
        },
        // autoLogout: false
      },
      admin: {
        scheme: 'refresh',
        token: {
          property: 'access_token',
          global: true,
        },
        refreshToken: {
          property: 'refresh_token',
          data: 'refresh_token',
          // maxAge: 60 * 60 * 24 * 30,
        },
        user: {
          property: false,
          // autoFetch: true
        },
        endpoints: {
          login: { url: '/api/auth/login', method: 'post' },
          refresh: { url: '/api/auth/refresh', method: 'post' },
          user: { url: '/api/auth/user', method: 'get' },
          logout: false,
        },
        // autoLogout: false
      },
    },
  },
  i18n: {
    baseUrl: process.env.SITE_URL,
    locales: [
      {
        code: 'tm',
        lang: 'TM',
        name: 'TM',
        iso: 'tm-TM',
        file: 'tm',
        isCatchallLocale: true,
      },
      { code: 'ru', lang: 'RU', name: 'RU', iso: 'ru-RU', file: 'ru' },
    ],
    defaultLocale: 'tm',
    seo: true,
    vueI18n: {
      fallbackLocale: 'tm',
      messages: {
        tm,
        ru,
      },
    },
  },

  router: {
    linkExactActiveClass: '_active',
    linkActiveClass: '_active',
  },

  // Build Configuration: https://go.nuxtjs.dev/config-build
  build: {},
}
