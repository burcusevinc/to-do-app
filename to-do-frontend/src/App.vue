<template>
  <div id="app">
    <input type="text" class="text-box" placeholder="Add your tasks"
    v-model="newTask">
    <div v-for="task in tasks" :key="task.id" class="title">
      {{ task.title }}
    </div>
    <button class="add-button" @click="addTask">Add</button>
  </div>
</template>

<script>
import API from "@/api"
export default {
  name: 'App',
  data() {
    return {
      //task array
      tasks: [],
      //text to be written
      newTask: "",
      //task id
      taskId: null,
    }
  },
  methods: {
    async addTask() {
      // Increment the task id
      this.taskId = this.tasks.length + 1
      // If text is empty stop the execute
      if (this.newTask.trim().length === 0) {
        return
      }
      // Task object
      const task = {id: this.taskId, title: this.newTask}
      // API post request method called with task object
      await API.createTasks(task)
      // Push the task object to tasks array
      this.tasks.push(task)
      // Empty text
      this.newTask = ""
    }
  },
  async created() {
    // Assign API get request method to task array
    this.tasks = await API.getTasks();
  }
}
</script>

<style scoped>
body {
  font-family: Verdana, monospace, sans-serif;
  font-size: 15px;
  height: 100%;
}
#app {
  width: 600px;
  height: 300px;
  padding: 20px;
  margin: auto;
}
</style>