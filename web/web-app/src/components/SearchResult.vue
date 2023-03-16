<script>
export default {
    data() {
        return {
            foo : "bar",
            items : [
                {"text" : "hello", "message" : "hello"},
                {"text" : "world", "message" : "world"}
            ],
            data: null
        }
    },
    methods : {
        async fetchData() {
            // GET request using fetch with async/await
            console.log("fetching")
            this.data = await (await fetch("http://localhost:8080/albums")).json();
        }
    }
}
</script>

<template>
    <div class = "wrapper">
        <!-- <ol>
            <li class = "item" v-for = "item in items" :key = "item.message">{{ item.text }}</li>
        </ol> -->
        <div class = "buttonD">
            <button @click="fetchData">fetch</button>
        </div>
        <ol>
            <li class = "item" v-for = "{id, title, artist, price} in data" :key = "id">
                <span>{{ id }}</span>
                <span>{{ title }}</span>
                <span>{{ artist }}</span>
                <span>{{ price }}</span>
            </li>
        </ol>
    </div>
</template>

<style>

button {
    margin: 0 auto;
}
.buttonD {
    /* align-items: center; */
    display: flex;
    justify-content: center;
    align-items: center;
    height: 3rem;
}

.wrapper {
    margin: 2rem;
}
.item {
    margin: 0 15rem;
    border: solid 1px black;
}

.item > span{
    margin: 0 4px;
}
</style>