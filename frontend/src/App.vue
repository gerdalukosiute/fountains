<script setup>
import Fountains from './components/Fountains.vue'
</script>

<template>
  <div id="app">
    <header class="bg-blue-600 text-white p-4">
      <h1 id="title" class="text-2xl font-bold">These are the fountains of Rome!</h1>
    </header>
    <main class="container mx-auto p-4">
      <Fountains />
    </main>
  </div>
</template>

<style>
@import 'tailwindcss';

body {
  font-family: Arial, sans-serif;
  line-height: 1.6;
  color: #333;
}

#title{
  color: rgb(114, 159, 237);
  font-family: 'Press Start 2P';
  font-weight: bold;
  font-size: 25px;
}
</style>


