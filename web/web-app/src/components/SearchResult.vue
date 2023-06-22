<script>

export default {
    data() {
        return {
            foo : "",
            items : [
                {"text" : "hello", "message" : "hello"},
                {"text" : "world", "message" : "world"}
            ],
            data: null,
            is_loading : false
        }
    },
    methods : {
        async fetchData(q) {
            // GET request using fetch with async/await
            console.log("fetching", q)
            var request = {
                query : q 
            }

            var requestOptions = {
                method: "POST",
                headers: { "Content-Type": "application/json;charset=UTF-8"},
                body: JSON.stringify(request)
            }

            console.log("body: ", requestOptions.body)

            this.is_loading = true 

            var temp = await fetch("http://43.153.212.240:8080/search", requestOptions)
            this.data = await temp.json()

            this.is_loading = false 
        },
    }
}
</script>

<template>
    <div class = "wrapper">
        <div v-if = "!is_loading">
            <div v-if = "!is_loading" class = "item" v-for = "{id, dir,text, token, url} in data" :key = "id">
                <div v-if = "!is_loading">{{ text }}</div>
                <div v-if = "!is_loading" class = "imgCont">
                    <img :src = "url" />
                </div>
            </div>
        </div>
        <div v-else class = "loading">
            <img src = "./icons/loading.gif" />
        </div>
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
    margin: 2rem 10rem;
    /* border: solid 1px black; */
    border:none;
}

.item div{
    font-size: 20px;
    margin: 0 4px;
}

.imgCont {
    padding: 1rem;
    display: flex;
    justify-content: center;
}

.item img {
    max-width: 100%;
    max-height: 20rem;
    display: inline-block;
}

.loading {
    display: flex;
    /* justify-content is the king of centering */
    justify-content: center; 
}

.loading img{
    max-width: 100%;
    max-height: 100px;
}
</style>
