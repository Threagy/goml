export const state = () => {
  return {
    drawer: null,
    color: 'general',
    image: '' // change if you want a picture but if you are using this internally get rid of the picture links here and in filter.vue
  }
}

const set = property => (state, payload) => (state[property] = payload)
const toggle = property => state => (state[property] = !state[property])

export const mutations = {
  setDrawer: set('drawer'),
  setImage: set('image'),
  setColor: set('color'),
  toggleDrawer: toggle('drawer')
}

