<template>
  <div class="tabs-br  flex justify-between items-center bg-[#E4E4E4] text-center tablet:block tablet:bg-transparent">
    <input name="tab-btn" type="radio" value="tab-br1" id="tab-br1" v-model="selectedTab"/>
    <label class="text-[24px] mac:text-[18px] pc:text-[14px] p-[10px] tablet:p-[4px] tablet:m-[5px] tablet:text-[14px]" for="tab-br1">Физкультурно-оздоровительные
      <br class="tablet:hidden">
      услуги</label>
    <input name="tab-btn" type="radio" value="tab-br2" id="tab-br2" v-model="selectedTab"/>
    <label class="text-[24px] mac:text-[18px] pc:text-[14px] p-[10px] tablet:p-[4px] tablet:m-[5px] tablet:text-[14px]" for="tab-br2">Туристические и спортивные <br class="tablet:hidden">
      мероприятия</label>
    <input name="tab-btn" type="radio" value="tab-br3" id="tab-br3" v-model="selectedTab"/>
    <label class="text-[24px] mac:text-[18px] pc:text-[14px] p-[10px] tablet:p-[4px] tablet:m-[5px] tablet:text-[14px]" for="tab-br3">Организация <br class="tablet:hidden">
      походов и сплавов </label>
    <input name="tab-btn" type="radio" value="tab-br4" id="tab-br4" v-model="selectedTab"/>
    <label class="text-[24px] mac:text-[18px] pc:text-[14px] p-[10px] tablet:p-[4px] tablet:m-[5px] tablet:text-[14px]" for="tab-br4">Поддержка <br class="tablet:hidden">
      ветеранов спорта</label>
    <input name="tab-btn" type="radio" value="tab-br5" id="tab-br5" v-model="selectedTab"/>
    <label class="text-[24px] mac:text-[18px] pc:text-[14px] p-[10px] tablet:p-[4px] tablet:m-[5px] tablet:text-[14px]" for="tab-br5">Спортивный <br class="tablet:hidden">
      клуб “Наш ковчег”</label>
  </div>
  <div class="tab-br" :id="selectedTab + '-content'">
    <slot :name="selectedTab"/>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';

const selectedTab = ref('tab-br1');
</script>


<style scoped>

.tabs-br input{
  display: none;
}
.tabs-br label{
  cursor: pointer;
  color: black;
  font-style: normal;
  font-weight: 300;
  line-height: normal;
}
.tabs-br input:checked + label {
  background-color: #333;
  color: white;
}
</style>
