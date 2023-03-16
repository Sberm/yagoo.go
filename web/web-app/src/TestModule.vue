<script setup>
defineProps({
    // parameter
    msg: { 
        type: String,
        required: true
    }
})
</script>
<template>
    <!-- <h1 :class="{green : true}">{{ msg }}</h1> -->
    <div class = "greetings">
        <h1 class = "red">{{ msg }}</h1>
    </div>
</template>

<style scoped>

.red {
    color: red;
}
h1 {
  font-weight: 500;
  font-size: 2.6rem;
  top: -10px;
}

h3 {
  font-size: 1.2rem;
}

.greetings h1,
.greetings h3 {
  text-align: center;
}

@media (min-width: 1024px) {
  .greetings h1,
  .greetings h3 {
    text-align: left;
  }
}

</style>   