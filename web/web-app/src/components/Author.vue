<script setup>
defineProps({
  msg: {
    type: String,
    required: true
  }
})
</script>

<template>
    <div class = "author">
        <h3>{{ msg }}</h3>
    </div>
</template>

<style scoped>
.author{
    /* if an element has a padding it can be select on F12 */
    padding: 0.3rem;
    margin: 0.2rem;
}

/* class name and tag name: child tag in the parent class */
.author h3{
    text-align: center;
    font-size: 1.4rem;
}


</style>
