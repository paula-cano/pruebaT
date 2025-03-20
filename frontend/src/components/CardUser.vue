<template>
    <div class="flex flex-wrap justify-center gap-6 p-6">
      <div
        v-for="(item, index) in marketData"
        :key="index"
        class="bg-gray-800 text-white p-6 rounded-lg shadow-lg w-80"
      >
        <h3 class="text-lg font-bold mb-2">{{ item.name }}</h3>
        <p><span class="font-semibold">Ticker:</span> {{ item.ticker }}</p>
        <p><span class="font-semibold">Action:</span> {{ item.action }}</p>
        <p><span class="font-semibold">Brokerage:</span> {{ item.brokerage }}</p>
        <p><span class="font-semibold">Rating:</span> {{ item.rating }}</p>
        <p><span class="font-semibold">Target:</span> {{ item.target_from }}</p>
        <p><span class="font-semibold">Time:</span> {{ item.time }}</p>
      </div>
    </div>
  </template>
  
  <script setup>
  import { ref, onMounted } from "vue";
  import { fetchData } from "@/connectionDb/api";
  
  const marketData = ref([]);
  
  onMounted(async () => {
    marketData.value = await fetchData();
  });
  </script>
  
  