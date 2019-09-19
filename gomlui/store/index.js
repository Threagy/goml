const cookieparser = process.server ? require('cookieparser') : undefined

export const state = () => {
  return {
    auth: '',
  }
}
export const mutations = {
  setAuth (state, auth) {
    state.auth = auth
  }
}

export const actions = {
  nuxtServerInit ({ app, commit }, { req }) {
    let auth = null

    if (req.headers.cookie) {
      const parsed = cookieparser.parse(req.headers.cookie)
      try {
        auth = parsed.auth
      } catch (err) {
        console.error(err)
      }
    }
    commit('setAuth', auth)
  },
}
