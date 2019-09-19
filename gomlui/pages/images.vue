<template>
  <div>
    <h1>Images page.</h1>
    <NuxtLink to="/">
      Back home
    </NuxtLink>
  </div>
</template>

<script>
export default {
  layout: 'dashboard',
  middleware: 'authenticated'
}
</script>
