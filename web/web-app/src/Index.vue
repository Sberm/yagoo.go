<script setup> // setup is a must-have
import Author from './components/Author.vue'
import SearchResult from './components/SearchResult.vue'

</script>

<script>
export default  {

}
</script>

<template>
    <header class = "title">
        <h1 class = "word">Yagoo</h1>
        <Author msg = "Howard Chu"/>
    </header>

    <!-- <button>the button</button> -->

    <SearchResult />


</template>

<style>

button:hover{
    opacity:1;
}

.title {
  /* the two numbers in padding: has to be different */
  padding: 2rem; 
  place-items: center;
}

.word {
    /* the line-height in word is overwritten by word */
    line-height: 1.5;
    font-size: 5rem;
    margin: 1rem;
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
    header .something{
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