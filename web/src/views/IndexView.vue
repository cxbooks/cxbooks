<template>
    <div class="index">
        <v-row>
            <v-col cols=12>
                <p class="ma-0 title">随便推荐</p>
            </v-col>
            <v-col cols=6 xs=6 sm=4 md=2 lg=2 v-for="(book,idx) in render_books(books.random)" :key="'rec'+idx+book.id"
                class="book-card">
                <v-card :to="book.href" class='rounded-0'>
                    <v-img :src="book.cover_url" :aspect-ratio="11/15"> </v-img>
                </v-card>
            </v-col>
        </v-row>
        <v-row>
            <v-col cols=12>
                <v-divider class="new-legend"></v-divider>
                <p class="ma-0 title">新书推荐</p>
            </v-col>
            <v-col cols=12>
                <book-cards :books="books.recent"></book-cards>
            </v-col>
        </v-row>
        <v-row>
            <v-col cols=12>
                <v-divider class="new-legend"></v-divider>
                <p class="ma-0 title">分类浏览</p>
            </v-col>
            <v-col cols=12 sm=6 md=4 v-for="nav in navs" :key="nav.text">
                <v-card outlined>
                    <v-list>
                        <v-list-item :to="nav.href">
                            <v-avatar large color='primary'>
                                <v-icon dark>{{nav.icon}}</v-icon>
                            </v-avatar>
 
                                <v-list-item-title>{{nav.text}} </v-list-item-title>
                                <v-list-item-subtitle>{{nav.subtitle}}</v-list-item-subtitle>

                            <v-list-item-action>
                                <v-icon>mdi-arrow-right</v-icon>
                            </v-list-item-action>
                        </v-list-item>
                    </v-list>
                </v-card>
            </v-col>
        </v-row>
    </div>
</template>

<script setup lang='ts'>

import BookCards from "@/components/BookCards.vue";
import {  onMounted, ref } from 'vue';
import {bookIndex} from '@/services/book';
import type { RespData, Book,Nav } from '@/types';


interface IndexBook {
    random: Book[],
    recent: Book[],
}

const books = ref <IndexBook> ({random:[],recent:[]})


const render_books = (books: Book[] ) => {
     books.map(b => {
        if (b['href'] == undefined) {
            b['href'] = "/book/" + b.id;
        }
        return b;
    });
    return books
}

onMounted(()=>{

    bookIndex()
        .then((response) => {
            // this.todo.id = response.data.id;
            if (response.code != 0) { //状态码异常
                console.log(response.data);
            }

            // console.log(response.data);
            books.value = response.data
         

        }).catch((e) => {
            console.log(e);
        })


}
)

const navs: Nav[] = [
    { icon: 'widgets', href: '/nav', text: '分类导览', count: 0 },
    { icon: 'mdi-human-greeting', href: '/author', text: '作者', count: 0 },
    { icon: 'mdi-home-group', href: '/publisher', text: '出版社', count: 0 },
    { icon: 'mdi-tag-heart', href: '/tag', text: '标签', count: 0 },
    { icon: 'mdi-history', href: '/recent', text: '所有书籍', count:0},
    { icon: 'mdi-trending-up', href: '/hot', text: '热度榜单', count: 0 },
]   



// const head: () => ({
//     titleTemplate: "%s",
// })


</script>

<style>
.new-legend {
    margin-top: 30px;
    margin-bottom: 20px;
}
</style>
