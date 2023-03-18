<script>
export default {
    data() {
        return {
            search_text : null
        }
    },
    methods : {
        query() {
            this.$emit('returnQuery', this.search_text)
        }
    }
}
</script>

<template>
    <div class = "container">
        <!-- <form class = "search-bar" method = "post"> -->
        <form class = "search-bar" @submit.prevent="query">
            <input v-model = "search_text" placeholder="Search anything" type = "text">
            <button type = 'button' @click="query"><img src = "./icons/search_icon.png"></button>
        </form>
    </div>
</template>


<style>
.container {
    margin: 2rem;
    min-height: 1rem;
    padding: 5%;
    display: flex;
    align-items: center;
    justify-content: center;
}
.search-bar{
    width: 100%;
    max-width: 700px;
    background: rgba(109, 109, 109, 0.2);
    display: flex;
    align-items: center;
    border-radius: 60px;
    padding: 10px 20px;
}

.search-bar input {
    background: transparent;
    flex : 1;
    border: 0;
    outline: none;
    font-size: 20px;
    color: rgb(30, 30, 30);
}

::placeholder { /* Chrome, Firefox, Opera, Safari 10.1+ */
    color: rgb(30, 30, 30);
    opacity: 1; /* Firefox */
}

.search-bar button img{
    width: 30px;
}

.search-bar button {
    border: 0;
    border-radius: 50%;
    background: transparent;
    outline: none;
    cursor: pointer;
}

</style>