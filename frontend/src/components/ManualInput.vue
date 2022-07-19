<template>
  <div class="hello">
    <h1>Manual Input</h1>
    <div class="wrapper">
      <div class="single-item" v-for="item in items" :key="item.count">
        <input type="text"/>
        <input type="number"/>
        <button type="reset"> -</button>
      </div>
      <button @click="items.push(count);count++;">
        Count is: {{ count }}
      </button>
    </div>
    <button type="submit" @click="created()">
      Compute calories
    </button>
  </div>
</template>

<script>
export default {
  name: 'ManualInput',
  data() {
    return {
      count: 0,
      items: []
    }
  }, methods: {
    created: function () {
      // POST request using fetch with error handling
      const requestOptions = {
        method: 'POST',
        header: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify([{"Name": "Avocado", "Weight": 100}, {"Name": "Apple", "Weight": 150}])
      };
      fetch('http://localhost:8080/compute-calories', requestOptions)
          .then(async response => {
            const data = await response.json();
            console.log(data)
            // check for error response
            if (!response.ok) {
              // get error message from body or default to response status
              const error = (data && data.message) || response.status;
              return Promise.reject(error);
            }

            this.postId = data.id;
          })
          .catch(error => {
            this.errorMessage = error;
            console.error('There was an error!', error);
          });
    }
  }
  // props: {
  //   msg: String
  // }
}
</script>

<style scoped>
.wrapper {
  display: flex;
  flex-direction: column;
  padding: 10px;
  margin: auto;
}

.single-item {
  justify-content: center;
  display: flex;
  flex-direction: row;
}
</style>
