<script lang="ts" setup>
import { reactive, ref } from "vue";
import { G_获取游戏名称 } from "../../wailsjs/go/main/App";

// const firstName = ref("");

const rules = [
  (value: any) => {
    // console.log(value);
    if (value) return true;
    return "请输入英文游戏名称";
  },
];

const data = reactive({
  搜索关键词: "",
  返回的数据: {
    标题: "",
    下载链接: "",
  },
});

function 搜索游戏() {
  console.log(data.搜索关键词);
  if (data.搜索关键词 === "") {
    return;
  } else {
    G_获取游戏名称(data.搜索关键词).then((result: any) => {
      data.返回的数据.标题 = result.标题;
      data.返回的数据.下载链接 = result.下载链接;
      console.log(result);
    });
  }
}
</script>

<template>
  <div class="home">
    <v-sheet class="mx-auto" width="300">
      <v-form @submit.prevent>
        <v-text-field
          v-model="data.搜索关键词"
          :rules="rules"
          label="游戏名称"
        ></v-text-field>
        <v-btn class="mt-2" type="submit" block @click="搜索游戏()">搜索</v-btn>
      </v-form>
    </v-sheet>
    <!-- 搜索的内容 -->
    <v-card v-show="data.返回的数据.标题" class="search-content" title="搜索结果"
      >{{ data.返回的数据.标题 }}:<a :href="data.返回的数据.下载链接">{{
        data.返回的数据.下载链接
      }}</a></v-card
    >
  </div>
</template>

<style lang="css" scoped>
/* 设置文本颜色 */
.mx-auto {
  color: #000;
  background-color: antiquewhite;
  .mt-2 {
    margin-top: 2rem;
    background-color: antiquewhite;
    color: rgb(79, 122, 122);
  }
}
.search-content {
  margin-top: 2rem;
  background-color: antiquewhite;
  color: rgb(79, 122, 122);
  a {
    color: rgb(79, 122, 122);
    text-decoration: none;
    &:hover {
      text-decoration: underline;
    }
  }
}
</style>
