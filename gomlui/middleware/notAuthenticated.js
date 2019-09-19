export default function ({ store, redirect }) {
  console.log("middleware:notauth:" + store.state.auth)
  // If the user is authenticated redirect to home page
  if (store.state.auth) {
    return redirect('/dashboard')
  }
}
