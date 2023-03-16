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
    /* display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    text-align: center; */
    /* dont use text-align in parent class */
    /* text-align: center; */

    /* if an element has a padding it can be select on F12 */
    padding: 1rem;
    margin: 1rem;
}

/* class name and tag name: child tag in the parent class */
.author h3{
    text-align: center;
    font-size: 1.4rem;
}


</style>