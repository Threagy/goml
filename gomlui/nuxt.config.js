import VuetifyLoaderPlugin from 'vuetify-loader/lib/plugin'
const webpack = require('webpack')

export default {
  /*
  ** Head elements
  ** Add Roboto font and Material Icons
  */
  head: {
    link: [
      {
        rel: 'stylesheet',
        type: 'text/css',
        href: 'https://fonts.googleapis.com/css?family=Roboto:300,400,500,700,300italic,400italic'
      }
    ],
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      { hid: 'description', name: 'description', content: 'Meta description' }
    ]
  },
  css: [
    'vuetify/dist/vuetify.min.css',
    'xterm/dist/xterm.css',
    // 'xterm/dist/addons/fullscreen/fullscreen.css'
  ],

  build: {
    plugins: [
      new VuetifyLoaderPlugin(),
      new webpack.ProvidePlugin({
        '_': 'lodash'
      })
    ],
    extractCSS: true,
    transpile: ['vuetify/lib']
  },
  devModules: [
    '@nuxtjs/vuetify',
  ],
  plugins: [
    '~/plugins/chartist.js',
    '~/plugins/register.js',
    { src: '~/plugins/avatar.js', injectAs: 'avatar', ssr: false},
    { src: '~/plugins/prism.js', injectAs: 'prism', ssr: false},
    { src: '~/plugins/gomlapi.js', injectAs: 'gomlapi', ssr: false},
  ],
  vuetify: {
    customVariables: ["~/assets/index.scss"],
    optionsPath: "./vuetify.options.js"
  }
}
