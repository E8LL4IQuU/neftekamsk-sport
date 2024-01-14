<template>
  <div class="wrapper">
    <Hero/>
    <Slider :SliderData="SliderData"/>
    <div>
      <h3 class="text-6xl pt-4 pb-3 text-center text-black font-normal">Новости</h3>
      <div class="flex flex-wrap justify-center">
        <News :NewsData="NewsData"/>
      </div>
      <h3 class="text-6xl pt-4 pb-3 text-center text-black font-normal">Галерея</h3>
      <div>
        <Gallery :GalleryData="GalleryData"/>
      </div>
      <h3 class="text-6xl pt-4 pb-3 text-center text-black font-normal">Истории людей</h3>
      <div>
        <StoriesPeople :PeopleData="PeopleData"/>
      </div>
      <div>
        <Connect/>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import Hero from "@/components/Hero.vue";
import Slider from "@/components/Slider.vue";
import News from "@/components/News.vue";
import Gallery from "@/components/Gallery.vue";
import StoriesPeople from "@/components/StoriesPeople.vue";
import Connect from "@/components/Connect.vue";
// mock data
const SliderData  = [{
  "title": "Мероприятие1",
  "description": "Это описание просто имба",
  "links": "events",
  "img": "https://white-rainbow.ru/upload/iblock/b26/7sf66d03tagu4tqpxywzr9e18325c05v/BSA_3747_Pano_copy.750.jpg"
},
  {
    "title": "Мероприятие2",
    "description": "Это описание просто имба",
    "links": "events",
    "img": "https://www.rgo.ru/sites/default/files/styles/head_image_article/public/node/26417/fucking-finger-array-usvinskie-pillars-sergei-afanasev-0.jpg?itok=StYRWb9k"
  },
  {
    "title": "Мероприятие3",
    "description": "Это описание просто имба",
    "links": "events",
    "img": "https://rusclimbing.ru/upload/resize_cache/webp/iblock/9b3/IMG_20230611_000706.webp"
  }
]
const NewsData =[{
  "title": "Это заголовок новости",
  "img": "https://thumb.tildacdn.com/tild6538-3765-4639-b764-656637656238/-/format/webp/_1.JPG",
  "description": "Это краткое описание новости , которое поможет, понять нравится ли вам наш сайтец",
  "links": "news",
  "date" : "21.09.2023"
},
  {
    "title": "Это заголовок новости",
    "img": "https://thumb.tildacdn.com/tild6538-3765-4639-b764-656637656238/-/format/webp/_1.JPG",
    "description": "Это краткое описание новости , которое поможет, понять нравится ли вам наш сайтец",
    "links": "news",
    "date" : "21.09.2023"
  },
  {
    "title": "Это заголовок новости",
    "img": "https://thumb.tildacdn.com/tild6538-3765-4639-b764-656637656238/-/format/webp/_1.JPG",
    "description": "Это краткое описание новости , которое поможет, понять нравится ли вам наш сайтец",
    "links": "news",
    "date" : "21.09.2023"
  }
]
const GalleryData = {
  "linkImgOne": "https://thumb.tildacdn.com/tild6538-3765-4639-b764-656637656238/-/format/webp/_1.JPG",
  "linkImgTwo": "https://thumb.tildacdn.com/tild6538-3765-4639-b764-656637656238/-/format/webp/_1.JPG",
  "linkImgThree": "https://thumb.tildacdn.com/tild6538-3765-4639-b764-656637656238/-/format/webp/_1.JPG",
    "videos": "https://res.cloudinary.com/demo/video/upload/q_auto,f_auto/dog.mp4"
}
const PeopleData = [
  {
    "img": "https://static.tildacdn.com/tild3235-3065-4462-a662-666363353834/i.jpg",
    "name": "Dmitry Shishkin",
    "activity": "Предприниматель",
    "description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Phasellus nisi metus, cursus et diam in, feugiat porttitor sapien. Mauris vel fringilla quam",
  },
  {
    "img": "https://static.tildacdn.com/tild3235-3065-4462-a662-666363353834/i.jpg",
    "name": "Dmitry Shishkin",
    "activity": "Предприниматель",
    "description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Phasellus nisi metus, cursus et diam in, feugiat porttitor sapien. Mauris vel fringilla quam",
  },
  {
    "img": "https://static.tildacdn.com/tild3235-3065-4462-a662-666363353834/i.jpg",
    "name": "Dmitry Shishkin",
    "activity": "Предприниматель",
    "description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Phasellus nisi metus, cursus et diam in, feugiat porttitor sapien. Mauris vel fringilla quam",
  },
]
</script>
<style scoped>
.wrapper{
  max-width: 1380px;
  padding: 18px 0px;
  margin: 0 auto;
}
</style>