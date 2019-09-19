export default function ({ store, redirect }) {
  // If the user is not authenticated
  console.log("middleware:auth")
  if (!store.state.auth) {
    return redirect('/login')
  }
}
