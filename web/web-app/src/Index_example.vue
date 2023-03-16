<script setup> // setup is a must-have
import TestModule from './TestModule.vue'

</script>

<script>
export default  {
    methods :{
        print() {
            console.log("Hello World")
        }
    }
}
</script>

<template>
    <header>
        <div class = "wrapper">Hello World</div>
        <img alt="Vue logo" class="logo" src="./assets/logo.svg" width="125" height="125" />
        <div class = "wrapper">
            <TestModule msg="Foo"/>
        </div>
        <TestModule msg="hey" />
    </header>

    <div class = "word">
        <button @click = "print">click me!</button>
    </div>

    <div class = "word">
        does it /r?
    </div>
</template>

<style>

header {
  line-height: 1.5;
}

.word {
    text-align: center; 
}

.noclass {
    display: flex;
    justify-content: center;
    align-items: center;
    font-size: xx-large;
}

/* given by example */
@media (min-width: 1024px){
    header {
        display: flex;
        place-items: center;
        padding-right: calc(var(--section-gap) / 2);
    }
    header .wrapper {
    display: flex;
    place-items: flex-start;
    flex-wrap: wrap;
  }
}
</style>